


package migrate

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

type AutomaticResolutionPropertiesResponse struct {
	MoveResourceId *string `pulumi:"moveResourceId"`
}





type AutomaticResolutionPropertiesResponseInput interface {
	pulumi.Input

	ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput
	ToAutomaticResolutionPropertiesResponseOutputWithContext(context.Context) AutomaticResolutionPropertiesResponseOutput
}

type AutomaticResolutionPropertiesResponseArgs struct {
	MoveResourceId pulumi.StringPtrInput `pulumi:"moveResourceId"`
}

func (AutomaticResolutionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput {
	return i.ToAutomaticResolutionPropertiesResponseOutputWithContext(context.Background())
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponseOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponseOutput)
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return i.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AutomaticResolutionPropertiesResponseArgs) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponseOutput).ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx)
}









type AutomaticResolutionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput
	ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Context) AutomaticResolutionPropertiesResponsePtrOutput
}

type automaticResolutionPropertiesResponsePtrType AutomaticResolutionPropertiesResponseArgs

func AutomaticResolutionPropertiesResponsePtr(v *AutomaticResolutionPropertiesResponseArgs) AutomaticResolutionPropertiesResponsePtrInput {
	return (*automaticResolutionPropertiesResponsePtrType)(v)
}

func (*automaticResolutionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (i *automaticResolutionPropertiesResponsePtrType) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return i.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *automaticResolutionPropertiesResponsePtrType) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomaticResolutionPropertiesResponsePtrOutput)
}

type AutomaticResolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AutomaticResolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponseOutput() AutomaticResolutionPropertiesResponseOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponseOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponseOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AutomaticResolutionPropertiesResponseOutput) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomaticResolutionPropertiesResponse) *AutomaticResolutionPropertiesResponse {
		return &v
	}).(AutomaticResolutionPropertiesResponsePtrOutput)
}

func (o AutomaticResolutionPropertiesResponseOutput) MoveResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomaticResolutionPropertiesResponse) *string { return v.MoveResourceId }).(pulumi.StringPtrOutput)
}

type AutomaticResolutionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutomaticResolutionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomaticResolutionPropertiesResponse)(nil)).Elem()
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) ToAutomaticResolutionPropertiesResponsePtrOutput() AutomaticResolutionPropertiesResponsePtrOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) ToAutomaticResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) AutomaticResolutionPropertiesResponsePtrOutput {
	return o
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) Elem() AutomaticResolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *AutomaticResolutionPropertiesResponse) AutomaticResolutionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AutomaticResolutionPropertiesResponse
		return ret
	}).(AutomaticResolutionPropertiesResponseOutput)
}

func (o AutomaticResolutionPropertiesResponsePtrOutput) MoveResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomaticResolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.MoveResourceId
	}).(pulumi.StringPtrOutput)
}

type AvailabilitySetResourceSettings struct {
	FaultDomain        *int   `pulumi:"faultDomain"`
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
	UpdateDomain       *int   `pulumi:"updateDomain"`
}





type AvailabilitySetResourceSettingsInput interface {
	pulumi.Input

	ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput
	ToAvailabilitySetResourceSettingsOutputWithContext(context.Context) AvailabilitySetResourceSettingsOutput
}

type AvailabilitySetResourceSettingsArgs struct {
	FaultDomain        pulumi.IntPtrInput `pulumi:"faultDomain"`
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
	UpdateDomain       pulumi.IntPtrInput `pulumi:"updateDomain"`
}

func (AvailabilitySetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettings)(nil)).Elem()
}

func (i AvailabilitySetResourceSettingsArgs) ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput {
	return i.ToAvailabilitySetResourceSettingsOutputWithContext(context.Background())
}

func (i AvailabilitySetResourceSettingsArgs) ToAvailabilitySetResourceSettingsOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetResourceSettingsOutput)
}

type AvailabilitySetResourceSettingsOutput struct{ *pulumi.OutputState }

func (AvailabilitySetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettings)(nil)).Elem()
}

func (o AvailabilitySetResourceSettingsOutput) ToAvailabilitySetResourceSettingsOutput() AvailabilitySetResourceSettingsOutput {
	return o
}

func (o AvailabilitySetResourceSettingsOutput) ToAvailabilitySetResourceSettingsOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsOutput {
	return o
}

func (o AvailabilitySetResourceSettingsOutput) FaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) *int { return v.FaultDomain }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsOutput) UpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettings) *int { return v.UpdateDomain }).(pulumi.IntPtrOutput)
}

type AvailabilitySetResourceSettingsResponse struct {
	FaultDomain        *int   `pulumi:"faultDomain"`
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
	UpdateDomain       *int   `pulumi:"updateDomain"`
}





type AvailabilitySetResourceSettingsResponseInput interface {
	pulumi.Input

	ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput
	ToAvailabilitySetResourceSettingsResponseOutputWithContext(context.Context) AvailabilitySetResourceSettingsResponseOutput
}

type AvailabilitySetResourceSettingsResponseArgs struct {
	FaultDomain        pulumi.IntPtrInput `pulumi:"faultDomain"`
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
	UpdateDomain       pulumi.IntPtrInput `pulumi:"updateDomain"`
}

func (AvailabilitySetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettingsResponse)(nil)).Elem()
}

func (i AvailabilitySetResourceSettingsResponseArgs) ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput {
	return i.ToAvailabilitySetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i AvailabilitySetResourceSettingsResponseArgs) ToAvailabilitySetResourceSettingsResponseOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetResourceSettingsResponseOutput)
}

type AvailabilitySetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (AvailabilitySetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilitySetResourceSettingsResponse)(nil)).Elem()
}

func (o AvailabilitySetResourceSettingsResponseOutput) ToAvailabilitySetResourceSettingsResponseOutput() AvailabilitySetResourceSettingsResponseOutput {
	return o
}

func (o AvailabilitySetResourceSettingsResponseOutput) ToAvailabilitySetResourceSettingsResponseOutputWithContext(ctx context.Context) AvailabilitySetResourceSettingsResponseOutput {
	return o
}

func (o AvailabilitySetResourceSettingsResponseOutput) FaultDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) *int { return v.FaultDomain }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o AvailabilitySetResourceSettingsResponseOutput) UpdateDomain() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilitySetResourceSettingsResponse) *int { return v.UpdateDomain }).(pulumi.IntPtrOutput)
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

type DatabaseProjectSummaryResponse struct {
	ExtendedSummary          map[string]string `pulumi:"extendedSummary"`
	InstanceType             string            `pulumi:"instanceType"`
	LastSummaryRefreshedTime *string           `pulumi:"lastSummaryRefreshedTime"`
	RefreshSummaryState      *string           `pulumi:"refreshSummaryState"`
}





type DatabaseProjectSummaryResponseInput interface {
	pulumi.Input

	ToDatabaseProjectSummaryResponseOutput() DatabaseProjectSummaryResponseOutput
	ToDatabaseProjectSummaryResponseOutputWithContext(context.Context) DatabaseProjectSummaryResponseOutput
}

type DatabaseProjectSummaryResponseArgs struct {
	ExtendedSummary          pulumi.StringMapInput `pulumi:"extendedSummary"`
	InstanceType             pulumi.StringInput    `pulumi:"instanceType"`
	LastSummaryRefreshedTime pulumi.StringPtrInput `pulumi:"lastSummaryRefreshedTime"`
	RefreshSummaryState      pulumi.StringPtrInput `pulumi:"refreshSummaryState"`
}

func (DatabaseProjectSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseProjectSummaryResponse)(nil)).Elem()
}

func (i DatabaseProjectSummaryResponseArgs) ToDatabaseProjectSummaryResponseOutput() DatabaseProjectSummaryResponseOutput {
	return i.ToDatabaseProjectSummaryResponseOutputWithContext(context.Background())
}

func (i DatabaseProjectSummaryResponseArgs) ToDatabaseProjectSummaryResponseOutputWithContext(ctx context.Context) DatabaseProjectSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseProjectSummaryResponseOutput)
}

type DatabaseProjectSummaryResponseOutput struct{ *pulumi.OutputState }

func (DatabaseProjectSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseProjectSummaryResponse)(nil)).Elem()
}

func (o DatabaseProjectSummaryResponseOutput) ToDatabaseProjectSummaryResponseOutput() DatabaseProjectSummaryResponseOutput {
	return o
}

func (o DatabaseProjectSummaryResponseOutput) ToDatabaseProjectSummaryResponseOutputWithContext(ctx context.Context) DatabaseProjectSummaryResponseOutput {
	return o
}

func (o DatabaseProjectSummaryResponseOutput) ExtendedSummary() pulumi.StringMapOutput {
	return o.ApplyT(func(v DatabaseProjectSummaryResponse) map[string]string { return v.ExtendedSummary }).(pulumi.StringMapOutput)
}

func (o DatabaseProjectSummaryResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseProjectSummaryResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o DatabaseProjectSummaryResponseOutput) LastSummaryRefreshedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProjectSummaryResponse) *string { return v.LastSummaryRefreshedTime }).(pulumi.StringPtrOutput)
}

func (o DatabaseProjectSummaryResponseOutput) RefreshSummaryState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseProjectSummaryResponse) *string { return v.RefreshSummaryState }).(pulumi.StringPtrOutput)
}

type DatabasesSolutionSummaryResponse struct {
	DatabaseInstancesAssessedCount *int   `pulumi:"databaseInstancesAssessedCount"`
	DatabasesAssessedCount         *int   `pulumi:"databasesAssessedCount"`
	InstanceType                   string `pulumi:"instanceType"`
	MigrationReadyCount            *int   `pulumi:"migrationReadyCount"`
}





type DatabasesSolutionSummaryResponseInput interface {
	pulumi.Input

	ToDatabasesSolutionSummaryResponseOutput() DatabasesSolutionSummaryResponseOutput
	ToDatabasesSolutionSummaryResponseOutputWithContext(context.Context) DatabasesSolutionSummaryResponseOutput
}

type DatabasesSolutionSummaryResponseArgs struct {
	DatabaseInstancesAssessedCount pulumi.IntPtrInput `pulumi:"databaseInstancesAssessedCount"`
	DatabasesAssessedCount         pulumi.IntPtrInput `pulumi:"databasesAssessedCount"`
	InstanceType                   pulumi.StringInput `pulumi:"instanceType"`
	MigrationReadyCount            pulumi.IntPtrInput `pulumi:"migrationReadyCount"`
}

func (DatabasesSolutionSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasesSolutionSummaryResponse)(nil)).Elem()
}

func (i DatabasesSolutionSummaryResponseArgs) ToDatabasesSolutionSummaryResponseOutput() DatabasesSolutionSummaryResponseOutput {
	return i.ToDatabasesSolutionSummaryResponseOutputWithContext(context.Background())
}

func (i DatabasesSolutionSummaryResponseArgs) ToDatabasesSolutionSummaryResponseOutputWithContext(ctx context.Context) DatabasesSolutionSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasesSolutionSummaryResponseOutput)
}

type DatabasesSolutionSummaryResponseOutput struct{ *pulumi.OutputState }

func (DatabasesSolutionSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasesSolutionSummaryResponse)(nil)).Elem()
}

func (o DatabasesSolutionSummaryResponseOutput) ToDatabasesSolutionSummaryResponseOutput() DatabasesSolutionSummaryResponseOutput {
	return o
}

func (o DatabasesSolutionSummaryResponseOutput) ToDatabasesSolutionSummaryResponseOutputWithContext(ctx context.Context) DatabasesSolutionSummaryResponseOutput {
	return o
}

func (o DatabasesSolutionSummaryResponseOutput) DatabaseInstancesAssessedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatabasesSolutionSummaryResponse) *int { return v.DatabaseInstancesAssessedCount }).(pulumi.IntPtrOutput)
}

func (o DatabasesSolutionSummaryResponseOutput) DatabasesAssessedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatabasesSolutionSummaryResponse) *int { return v.DatabasesAssessedCount }).(pulumi.IntPtrOutput)
}

func (o DatabasesSolutionSummaryResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasesSolutionSummaryResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o DatabasesSolutionSummaryResponseOutput) MigrationReadyCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatabasesSolutionSummaryResponse) *int { return v.MigrationReadyCount }).(pulumi.IntPtrOutput)
}

type DiskEncryptionSetResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type DiskEncryptionSetResourceSettingsInput interface {
	pulumi.Input

	ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput
	ToDiskEncryptionSetResourceSettingsOutputWithContext(context.Context) DiskEncryptionSetResourceSettingsOutput
}

type DiskEncryptionSetResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (DiskEncryptionSetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettings)(nil)).Elem()
}

func (i DiskEncryptionSetResourceSettingsArgs) ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput {
	return i.ToDiskEncryptionSetResourceSettingsOutputWithContext(context.Background())
}

func (i DiskEncryptionSetResourceSettingsArgs) ToDiskEncryptionSetResourceSettingsOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetResourceSettingsOutput)
}

type DiskEncryptionSetResourceSettingsOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettings)(nil)).Elem()
}

func (o DiskEncryptionSetResourceSettingsOutput) ToDiskEncryptionSetResourceSettingsOutput() DiskEncryptionSetResourceSettingsOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsOutput) ToDiskEncryptionSetResourceSettingsOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type DiskEncryptionSetResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type DiskEncryptionSetResourceSettingsResponseInput interface {
	pulumi.Input

	ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput
	ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(context.Context) DiskEncryptionSetResourceSettingsResponseOutput
}

type DiskEncryptionSetResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (DiskEncryptionSetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettingsResponse)(nil)).Elem()
}

func (i DiskEncryptionSetResourceSettingsResponseArgs) ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput {
	return i.ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i DiskEncryptionSetResourceSettingsResponseArgs) ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionSetResourceSettingsResponseOutput)
}

type DiskEncryptionSetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (DiskEncryptionSetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionSetResourceSettingsResponse)(nil)).Elem()
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ToDiskEncryptionSetResourceSettingsResponseOutput() DiskEncryptionSetResourceSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ToDiskEncryptionSetResourceSettingsResponseOutputWithContext(ctx context.Context) DiskEncryptionSetResourceSettingsResponseOutput {
	return o
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o DiskEncryptionSetResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v DiskEncryptionSetResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
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

type Identity struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId *string `pulumi:"principalId"`
	TenantId    *string `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
	TenantId    pulumi.StringPtrInput `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
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

type JobStatusResponse struct {
	JobName     string `pulumi:"jobName"`
	JobProgress string `pulumi:"jobProgress"`
}





type JobStatusResponseInput interface {
	pulumi.Input

	ToJobStatusResponseOutput() JobStatusResponseOutput
	ToJobStatusResponseOutputWithContext(context.Context) JobStatusResponseOutput
}

type JobStatusResponseArgs struct {
	JobName     pulumi.StringInput `pulumi:"jobName"`
	JobProgress pulumi.StringInput `pulumi:"jobProgress"`
}

func (JobStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return i.ToJobStatusResponseOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput)
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i JobStatusResponseArgs) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponseOutput).ToJobStatusResponsePtrOutputWithContext(ctx)
}









type JobStatusResponsePtrInput interface {
	pulumi.Input

	ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput
	ToJobStatusResponsePtrOutputWithContext(context.Context) JobStatusResponsePtrOutput
}

type jobStatusResponsePtrType JobStatusResponseArgs

func JobStatusResponsePtr(v *JobStatusResponseArgs) JobStatusResponsePtrInput {
	return (*jobStatusResponsePtrType)(v)
}

func (*jobStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return i.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (i *jobStatusResponsePtrType) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStatusResponsePtrOutput)
}

type JobStatusResponseOutput struct{ *pulumi.OutputState }

func (JobStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutput() JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) ToJobStatusResponseOutputWithContext(ctx context.Context) JobStatusResponseOutput {
	return o
}

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o.ToJobStatusResponsePtrOutputWithContext(context.Background())
}

func (o JobStatusResponseOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStatusResponse) *JobStatusResponse {
		return &v
	}).(JobStatusResponsePtrOutput)
}

func (o JobStatusResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o JobStatusResponseOutput) JobProgress() pulumi.StringOutput {
	return o.ApplyT(func(v JobStatusResponse) string { return v.JobProgress }).(pulumi.StringOutput)
}

type JobStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStatusResponse)(nil)).Elem()
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutput() JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) ToJobStatusResponsePtrOutputWithContext(ctx context.Context) JobStatusResponsePtrOutput {
	return o
}

func (o JobStatusResponsePtrOutput) Elem() JobStatusResponseOutput {
	return o.ApplyT(func(v *JobStatusResponse) JobStatusResponse {
		if v != nil {
			return *v
		}
		var ret JobStatusResponse
		return ret
	}).(JobStatusResponseOutput)
}

func (o JobStatusResponsePtrOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.JobName
	}).(pulumi.StringPtrOutput)
}

func (o JobStatusResponsePtrOutput) JobProgress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.JobProgress
	}).(pulumi.StringPtrOutput)
}

type KeyVaultResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type KeyVaultResourceSettingsInput interface {
	pulumi.Input

	ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput
	ToKeyVaultResourceSettingsOutputWithContext(context.Context) KeyVaultResourceSettingsOutput
}

type KeyVaultResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (KeyVaultResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettings)(nil)).Elem()
}

func (i KeyVaultResourceSettingsArgs) ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput {
	return i.ToKeyVaultResourceSettingsOutputWithContext(context.Background())
}

func (i KeyVaultResourceSettingsArgs) ToKeyVaultResourceSettingsOutputWithContext(ctx context.Context) KeyVaultResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultResourceSettingsOutput)
}

type KeyVaultResourceSettingsOutput struct{ *pulumi.OutputState }

func (KeyVaultResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettings)(nil)).Elem()
}

func (o KeyVaultResourceSettingsOutput) ToKeyVaultResourceSettingsOutput() KeyVaultResourceSettingsOutput {
	return o
}

func (o KeyVaultResourceSettingsOutput) ToKeyVaultResourceSettingsOutputWithContext(ctx context.Context) KeyVaultResourceSettingsOutput {
	return o
}

func (o KeyVaultResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o KeyVaultResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type KeyVaultResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type KeyVaultResourceSettingsResponseInput interface {
	pulumi.Input

	ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput
	ToKeyVaultResourceSettingsResponseOutputWithContext(context.Context) KeyVaultResourceSettingsResponseOutput
}

type KeyVaultResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (KeyVaultResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettingsResponse)(nil)).Elem()
}

func (i KeyVaultResourceSettingsResponseArgs) ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput {
	return i.ToKeyVaultResourceSettingsResponseOutputWithContext(context.Background())
}

func (i KeyVaultResourceSettingsResponseArgs) ToKeyVaultResourceSettingsResponseOutputWithContext(ctx context.Context) KeyVaultResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultResourceSettingsResponseOutput)
}

type KeyVaultResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultResourceSettingsResponse)(nil)).Elem()
}

func (o KeyVaultResourceSettingsResponseOutput) ToKeyVaultResourceSettingsResponseOutput() KeyVaultResourceSettingsResponseOutput {
	return o
}

func (o KeyVaultResourceSettingsResponseOutput) ToKeyVaultResourceSettingsResponseOutputWithContext(ctx context.Context) KeyVaultResourceSettingsResponseOutput {
	return o
}

func (o KeyVaultResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o KeyVaultResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type LBBackendAddressPoolResourceSettings struct {
	Name *string `pulumi:"name"`
}





type LBBackendAddressPoolResourceSettingsInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput
	ToLBBackendAddressPoolResourceSettingsOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsOutput
}

type LBBackendAddressPoolResourceSettingsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LBBackendAddressPoolResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsArgs) ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput {
	return i.ToLBBackendAddressPoolResourceSettingsOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsArgs) ToLBBackendAddressPoolResourceSettingsOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsOutput)
}





type LBBackendAddressPoolResourceSettingsArrayInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput
	ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsArrayOutput
}

type LBBackendAddressPoolResourceSettingsArray []LBBackendAddressPoolResourceSettingsInput

func (LBBackendAddressPoolResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsArray) ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput {
	return i.ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsArray) ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsArrayOutput)
}

type LBBackendAddressPoolResourceSettingsOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsOutput) ToLBBackendAddressPoolResourceSettingsOutput() LBBackendAddressPoolResourceSettingsOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsOutput) ToLBBackendAddressPoolResourceSettingsOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBBackendAddressPoolResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LBBackendAddressPoolResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettings)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) ToLBBackendAddressPoolResourceSettingsArrayOutput() LBBackendAddressPoolResourceSettingsArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) ToLBBackendAddressPoolResourceSettingsArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsArrayOutput) Index(i pulumi.IntInput) LBBackendAddressPoolResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBBackendAddressPoolResourceSettings {
		return vs[0].([]LBBackendAddressPoolResourceSettings)[vs[1].(int)]
	}).(LBBackendAddressPoolResourceSettingsOutput)
}

type LBBackendAddressPoolResourceSettingsResponse struct {
	Name *string `pulumi:"name"`
}





type LBBackendAddressPoolResourceSettingsResponseInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput
	ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsResponseOutput
}

type LBBackendAddressPoolResourceSettingsResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LBBackendAddressPoolResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsResponseArgs) ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput {
	return i.ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsResponseArgs) ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsResponseOutput)
}





type LBBackendAddressPoolResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput
	ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput
}

type LBBackendAddressPoolResourceSettingsResponseArray []LBBackendAddressPoolResourceSettingsResponseInput

func (LBBackendAddressPoolResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (i LBBackendAddressPoolResourceSettingsResponseArray) ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return i.ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i LBBackendAddressPoolResourceSettingsResponseArray) ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBBackendAddressPoolResourceSettingsResponseArrayOutput)
}

type LBBackendAddressPoolResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) ToLBBackendAddressPoolResourceSettingsResponseOutput() LBBackendAddressPoolResourceSettingsResponseOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) ToLBBackendAddressPoolResourceSettingsResponseOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBBackendAddressPoolResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type LBBackendAddressPoolResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LBBackendAddressPoolResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBBackendAddressPoolResourceSettingsResponse)(nil)).Elem()
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) ToLBBackendAddressPoolResourceSettingsResponseArrayOutput() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) ToLBBackendAddressPoolResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o
}

func (o LBBackendAddressPoolResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) LBBackendAddressPoolResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBBackendAddressPoolResourceSettingsResponse {
		return vs[0].([]LBBackendAddressPoolResourceSettingsResponse)[vs[1].(int)]
	}).(LBBackendAddressPoolResourceSettingsResponseOutput)
}

type LBFrontendIPConfigurationResourceSettings struct {
	Name                      *string          `pulumi:"name"`
	PrivateIpAddress          *string          `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string          `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReference `pulumi:"subnet"`
	Zones                     *string          `pulumi:"zones"`
}





type LBFrontendIPConfigurationResourceSettingsInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput
	ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsOutput
}

type LBFrontendIPConfigurationResourceSettingsArgs struct {
	Name                      pulumi.StringPtrInput   `pulumi:"name"`
	PrivateIpAddress          pulumi.StringPtrInput   `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod pulumi.StringPtrInput   `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubnetReferencePtrInput `pulumi:"subnet"`
	Zones                     pulumi.StringPtrInput   `pulumi:"zones"`
}

func (LBFrontendIPConfigurationResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsArgs) ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsArgs) ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsOutput)
}





type LBFrontendIPConfigurationResourceSettingsArrayInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput
	ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput
}

type LBFrontendIPConfigurationResourceSettingsArray []LBFrontendIPConfigurationResourceSettingsInput

func (LBFrontendIPConfigurationResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsArray) ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsArray) ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsArrayOutput)
}

type LBFrontendIPConfigurationResourceSettingsOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) ToLBFrontendIPConfigurationResourceSettingsOutput() LBFrontendIPConfigurationResourceSettingsOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) ToLBFrontendIPConfigurationResourceSettingsOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Subnet() SubnetReferencePtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *SubnetReference { return v.Subnet }).(SubnetReferencePtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type LBFrontendIPConfigurationResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettings)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) ToLBFrontendIPConfigurationResourceSettingsArrayOutput() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) ToLBFrontendIPConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsArrayOutput) Index(i pulumi.IntInput) LBFrontendIPConfigurationResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBFrontendIPConfigurationResourceSettings {
		return vs[0].([]LBFrontendIPConfigurationResourceSettings)[vs[1].(int)]
	}).(LBFrontendIPConfigurationResourceSettingsOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponse struct {
	Name                      *string                  `pulumi:"name"`
	PrivateIpAddress          *string                  `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod *string                  `pulumi:"privateIpAllocationMethod"`
	Subnet                    *SubnetReferenceResponse `pulumi:"subnet"`
	Zones                     *string                  `pulumi:"zones"`
}





type LBFrontendIPConfigurationResourceSettingsResponseInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput
	ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput
}

type LBFrontendIPConfigurationResourceSettingsResponseArgs struct {
	Name                      pulumi.StringPtrInput           `pulumi:"name"`
	PrivateIpAddress          pulumi.StringPtrInput           `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod pulumi.StringPtrInput           `pulumi:"privateIpAllocationMethod"`
	Subnet                    SubnetReferenceResponsePtrInput `pulumi:"subnet"`
	Zones                     pulumi.StringPtrInput           `pulumi:"zones"`
}

func (LBFrontendIPConfigurationResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArgs) ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArgs) ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsResponseOutput)
}





type LBFrontendIPConfigurationResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput
	ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput
}

type LBFrontendIPConfigurationResourceSettingsResponseArray []LBFrontendIPConfigurationResourceSettingsResponseInput

func (LBFrontendIPConfigurationResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArray) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return i.ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i LBFrontendIPConfigurationResourceSettingsResponseArray) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) ToLBFrontendIPConfigurationResourceSettingsResponseOutput() LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) ToLBFrontendIPConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Subnet() SubnetReferenceResponsePtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *SubnetReferenceResponse { return v.Subnet }).(SubnetReferenceResponsePtrOutput)
}

func (o LBFrontendIPConfigurationResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LBFrontendIPConfigurationResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type LBFrontendIPConfigurationResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LBFrontendIPConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutput() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) ToLBFrontendIPConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o LBFrontendIPConfigurationResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) LBFrontendIPConfigurationResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LBFrontendIPConfigurationResourceSettingsResponse {
		return vs[0].([]LBFrontendIPConfigurationResourceSettingsResponse)[vs[1].(int)]
	}).(LBFrontendIPConfigurationResourceSettingsResponseOutput)
}

type LoadBalancerBackendAddressPoolReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerBackendAddressPoolReferenceInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput
	ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceOutput
}

type LoadBalancerBackendAddressPoolReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerBackendAddressPoolReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceArgs) ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceArgs) ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceOutput)
}





type LoadBalancerBackendAddressPoolReferenceArrayInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput
	ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput
}

type LoadBalancerBackendAddressPoolReferenceArray []LoadBalancerBackendAddressPoolReferenceInput

func (LoadBalancerBackendAddressPoolReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceArray) ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceArray) ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceArrayOutput)
}

type LoadBalancerBackendAddressPoolReferenceOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) ToLoadBalancerBackendAddressPoolReferenceOutput() LoadBalancerBackendAddressPoolReferenceOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) ToLoadBalancerBackendAddressPoolReferenceOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerBackendAddressPoolReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerBackendAddressPoolReferenceArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReference)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) ToLoadBalancerBackendAddressPoolReferenceArrayOutput() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) ToLoadBalancerBackendAddressPoolReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceArrayOutput) Index(i pulumi.IntInput) LoadBalancerBackendAddressPoolReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerBackendAddressPoolReference {
		return vs[0].([]LoadBalancerBackendAddressPoolReference)[vs[1].(int)]
	}).(LoadBalancerBackendAddressPoolReferenceOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerBackendAddressPoolReferenceResponseInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput
	ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput
}

type LoadBalancerBackendAddressPoolReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerBackendAddressPoolReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArgs) ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArgs) ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceResponseOutput)
}





type LoadBalancerBackendAddressPoolReferenceResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput
	ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput
}

type LoadBalancerBackendAddressPoolReferenceResponseArray []LoadBalancerBackendAddressPoolReferenceResponseInput

func (LoadBalancerBackendAddressPoolReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArray) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return i.ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancerBackendAddressPoolReferenceResponseArray) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) ToLoadBalancerBackendAddressPoolReferenceResponseOutput() LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) ToLoadBalancerBackendAddressPoolReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerBackendAddressPoolReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerBackendAddressPoolReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerBackendAddressPoolReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerBackendAddressPoolReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutput() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) ToLoadBalancerBackendAddressPoolReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerBackendAddressPoolReferenceResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerBackendAddressPoolReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerBackendAddressPoolReferenceResponse {
		return vs[0].([]LoadBalancerBackendAddressPoolReferenceResponse)[vs[1].(int)]
	}).(LoadBalancerBackendAddressPoolReferenceResponseOutput)
}

type LoadBalancerNatRuleReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerNatRuleReferenceInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput
	ToLoadBalancerNatRuleReferenceOutputWithContext(context.Context) LoadBalancerNatRuleReferenceOutput
}

type LoadBalancerNatRuleReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerNatRuleReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReference)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceArgs) ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput {
	return i.ToLoadBalancerNatRuleReferenceOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceArgs) ToLoadBalancerNatRuleReferenceOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceOutput)
}





type LoadBalancerNatRuleReferenceArrayInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput
	ToLoadBalancerNatRuleReferenceArrayOutputWithContext(context.Context) LoadBalancerNatRuleReferenceArrayOutput
}

type LoadBalancerNatRuleReferenceArray []LoadBalancerNatRuleReferenceInput

func (LoadBalancerNatRuleReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReference)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceArray) ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput {
	return i.ToLoadBalancerNatRuleReferenceArrayOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceArray) ToLoadBalancerNatRuleReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceArrayOutput)
}

type LoadBalancerNatRuleReferenceOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReference)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceOutput) ToLoadBalancerNatRuleReferenceOutput() LoadBalancerNatRuleReferenceOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceOutput) ToLoadBalancerNatRuleReferenceOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerNatRuleReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerNatRuleReferenceArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReference)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceArrayOutput) ToLoadBalancerNatRuleReferenceArrayOutput() LoadBalancerNatRuleReferenceArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceArrayOutput) ToLoadBalancerNatRuleReferenceArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceArrayOutput) Index(i pulumi.IntInput) LoadBalancerNatRuleReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerNatRuleReference {
		return vs[0].([]LoadBalancerNatRuleReference)[vs[1].(int)]
	}).(LoadBalancerNatRuleReferenceOutput)
}

type LoadBalancerNatRuleReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type LoadBalancerNatRuleReferenceResponseInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput
	ToLoadBalancerNatRuleReferenceResponseOutputWithContext(context.Context) LoadBalancerNatRuleReferenceResponseOutput
}

type LoadBalancerNatRuleReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (LoadBalancerNatRuleReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceResponseArgs) ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput {
	return i.ToLoadBalancerNatRuleReferenceResponseOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceResponseArgs) ToLoadBalancerNatRuleReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceResponseOutput)
}





type LoadBalancerNatRuleReferenceResponseArrayInput interface {
	pulumi.Input

	ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput
	ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput
}

type LoadBalancerNatRuleReferenceResponseArray []LoadBalancerNatRuleReferenceResponseInput

func (LoadBalancerNatRuleReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (i LoadBalancerNatRuleReferenceResponseArray) ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return i.ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(context.Background())
}

func (i LoadBalancerNatRuleReferenceResponseArray) ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerNatRuleReferenceResponseArrayOutput)
}

type LoadBalancerNatRuleReferenceResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceResponseOutput) ToLoadBalancerNatRuleReferenceResponseOutput() LoadBalancerNatRuleReferenceResponseOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseOutput) ToLoadBalancerNatRuleReferenceResponseOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerNatRuleReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerNatRuleReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type LoadBalancerNatRuleReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerNatRuleReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerNatRuleReferenceResponse)(nil)).Elem()
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) ToLoadBalancerNatRuleReferenceResponseArrayOutput() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) ToLoadBalancerNatRuleReferenceResponseArrayOutputWithContext(ctx context.Context) LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o
}

func (o LoadBalancerNatRuleReferenceResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerNatRuleReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerNatRuleReferenceResponse {
		return vs[0].([]LoadBalancerNatRuleReferenceResponse)[vs[1].(int)]
	}).(LoadBalancerNatRuleReferenceResponseOutput)
}

type LoadBalancerResourceSettings struct {
	BackendAddressPools      []LBBackendAddressPoolResourceSettings      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations []LBFrontendIPConfigurationResourceSettings `pulumi:"frontendIPConfigurations"`
	ResourceType             string                                      `pulumi:"resourceType"`
	Sku                      *string                                     `pulumi:"sku"`
	TargetResourceName       string                                      `pulumi:"targetResourceName"`
	Zones                    *string                                     `pulumi:"zones"`
}





type LoadBalancerResourceSettingsInput interface {
	pulumi.Input

	ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput
	ToLoadBalancerResourceSettingsOutputWithContext(context.Context) LoadBalancerResourceSettingsOutput
}

type LoadBalancerResourceSettingsArgs struct {
	BackendAddressPools      LBBackendAddressPoolResourceSettingsArrayInput      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations LBFrontendIPConfigurationResourceSettingsArrayInput `pulumi:"frontendIPConfigurations"`
	ResourceType             pulumi.StringInput                                  `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput                               `pulumi:"sku"`
	TargetResourceName       pulumi.StringInput                                  `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput                               `pulumi:"zones"`
}

func (LoadBalancerResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettings)(nil)).Elem()
}

func (i LoadBalancerResourceSettingsArgs) ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput {
	return i.ToLoadBalancerResourceSettingsOutputWithContext(context.Background())
}

func (i LoadBalancerResourceSettingsArgs) ToLoadBalancerResourceSettingsOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerResourceSettingsOutput)
}

type LoadBalancerResourceSettingsOutput struct{ *pulumi.OutputState }

func (LoadBalancerResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettings)(nil)).Elem()
}

func (o LoadBalancerResourceSettingsOutput) ToLoadBalancerResourceSettingsOutput() LoadBalancerResourceSettingsOutput {
	return o
}

func (o LoadBalancerResourceSettingsOutput) ToLoadBalancerResourceSettingsOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsOutput {
	return o
}

func (o LoadBalancerResourceSettingsOutput) BackendAddressPools() LBBackendAddressPoolResourceSettingsArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) []LBBackendAddressPoolResourceSettings {
		return v.BackendAddressPools
	}).(LBBackendAddressPoolResourceSettingsArrayOutput)
}

func (o LoadBalancerResourceSettingsOutput) FrontendIPConfigurations() LBFrontendIPConfigurationResourceSettingsArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) []LBFrontendIPConfigurationResourceSettings {
		return v.FrontendIPConfigurations
	}).(LBFrontendIPConfigurationResourceSettingsArrayOutput)
}

func (o LoadBalancerResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type LoadBalancerResourceSettingsResponse struct {
	BackendAddressPools      []LBBackendAddressPoolResourceSettingsResponse      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations []LBFrontendIPConfigurationResourceSettingsResponse `pulumi:"frontendIPConfigurations"`
	ResourceType             string                                              `pulumi:"resourceType"`
	Sku                      *string                                             `pulumi:"sku"`
	TargetResourceName       string                                              `pulumi:"targetResourceName"`
	Zones                    *string                                             `pulumi:"zones"`
}





type LoadBalancerResourceSettingsResponseInput interface {
	pulumi.Input

	ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput
	ToLoadBalancerResourceSettingsResponseOutputWithContext(context.Context) LoadBalancerResourceSettingsResponseOutput
}

type LoadBalancerResourceSettingsResponseArgs struct {
	BackendAddressPools      LBBackendAddressPoolResourceSettingsResponseArrayInput      `pulumi:"backendAddressPools"`
	FrontendIPConfigurations LBFrontendIPConfigurationResourceSettingsResponseArrayInput `pulumi:"frontendIPConfigurations"`
	ResourceType             pulumi.StringInput                                          `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput                                       `pulumi:"sku"`
	TargetResourceName       pulumi.StringInput                                          `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput                                       `pulumi:"zones"`
}

func (LoadBalancerResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettingsResponse)(nil)).Elem()
}

func (i LoadBalancerResourceSettingsResponseArgs) ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput {
	return i.ToLoadBalancerResourceSettingsResponseOutputWithContext(context.Background())
}

func (i LoadBalancerResourceSettingsResponseArgs) ToLoadBalancerResourceSettingsResponseOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerResourceSettingsResponseOutput)
}

type LoadBalancerResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerResourceSettingsResponse)(nil)).Elem()
}

func (o LoadBalancerResourceSettingsResponseOutput) ToLoadBalancerResourceSettingsResponseOutput() LoadBalancerResourceSettingsResponseOutput {
	return o
}

func (o LoadBalancerResourceSettingsResponseOutput) ToLoadBalancerResourceSettingsResponseOutputWithContext(ctx context.Context) LoadBalancerResourceSettingsResponseOutput {
	return o
}

func (o LoadBalancerResourceSettingsResponseOutput) BackendAddressPools() LBBackendAddressPoolResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) []LBBackendAddressPoolResourceSettingsResponse {
		return v.BackendAddressPools
	}).(LBBackendAddressPoolResourceSettingsResponseArrayOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) FrontendIPConfigurations() LBFrontendIPConfigurationResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) []LBFrontendIPConfigurationResourceSettingsResponse {
		return v.FrontendIPConfigurations
	}).(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o LoadBalancerResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type ManualResolutionPropertiesResponse struct {
	TargetId *string `pulumi:"targetId"`
}





type ManualResolutionPropertiesResponseInput interface {
	pulumi.Input

	ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput
	ToManualResolutionPropertiesResponseOutputWithContext(context.Context) ManualResolutionPropertiesResponseOutput
}

type ManualResolutionPropertiesResponseArgs struct {
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (ManualResolutionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput {
	return i.ToManualResolutionPropertiesResponseOutputWithContext(context.Background())
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponseOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponseOutput)
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return i.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ManualResolutionPropertiesResponseArgs) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponseOutput).ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx)
}









type ManualResolutionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput
	ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Context) ManualResolutionPropertiesResponsePtrOutput
}

type manualResolutionPropertiesResponsePtrType ManualResolutionPropertiesResponseArgs

func ManualResolutionPropertiesResponsePtr(v *ManualResolutionPropertiesResponseArgs) ManualResolutionPropertiesResponsePtrInput {
	return (*manualResolutionPropertiesResponsePtrType)(v)
}

func (*manualResolutionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (i *manualResolutionPropertiesResponsePtrType) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return i.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *manualResolutionPropertiesResponsePtrType) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManualResolutionPropertiesResponsePtrOutput)
}

type ManualResolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ManualResolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponseOutput() ManualResolutionPropertiesResponseOutput {
	return o
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponseOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponseOutput {
	return o
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return o.ToManualResolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ManualResolutionPropertiesResponseOutput) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManualResolutionPropertiesResponse) *ManualResolutionPropertiesResponse {
		return &v
	}).(ManualResolutionPropertiesResponsePtrOutput)
}

func (o ManualResolutionPropertiesResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManualResolutionPropertiesResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type ManualResolutionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ManualResolutionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManualResolutionPropertiesResponse)(nil)).Elem()
}

func (o ManualResolutionPropertiesResponsePtrOutput) ToManualResolutionPropertiesResponsePtrOutput() ManualResolutionPropertiesResponsePtrOutput {
	return o
}

func (o ManualResolutionPropertiesResponsePtrOutput) ToManualResolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) ManualResolutionPropertiesResponsePtrOutput {
	return o
}

func (o ManualResolutionPropertiesResponsePtrOutput) Elem() ManualResolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *ManualResolutionPropertiesResponse) ManualResolutionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ManualResolutionPropertiesResponse
		return ret
	}).(ManualResolutionPropertiesResponseOutput)
}

func (o ManualResolutionPropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManualResolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetId
	}).(pulumi.StringPtrOutput)
}

type MigrateProjectProperties struct {
	ProvisioningState *string  `pulumi:"provisioningState"`
	RegisteredTools   []string `pulumi:"registeredTools"`
}





type MigrateProjectPropertiesInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput
	ToMigrateProjectPropertiesOutputWithContext(context.Context) MigrateProjectPropertiesOutput
}

type MigrateProjectPropertiesArgs struct {
	ProvisioningState pulumi.StringPtrInput   `pulumi:"provisioningState"`
	RegisteredTools   pulumi.StringArrayInput `pulumi:"registeredTools"`
}

func (MigrateProjectPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return i.ToMigrateProjectPropertiesOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput)
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesArgs) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesOutput).ToMigrateProjectPropertiesPtrOutputWithContext(ctx)
}









type MigrateProjectPropertiesPtrInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput
	ToMigrateProjectPropertiesPtrOutputWithContext(context.Context) MigrateProjectPropertiesPtrOutput
}

type migrateProjectPropertiesPtrType MigrateProjectPropertiesArgs

func MigrateProjectPropertiesPtr(v *MigrateProjectPropertiesArgs) MigrateProjectPropertiesPtrInput {
	return (*migrateProjectPropertiesPtrType)(v)
}

func (*migrateProjectPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return i.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i *migrateProjectPropertiesPtrType) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesPtrOutput)
}

type MigrateProjectPropertiesOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutput() MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesOutputWithContext(ctx context.Context) MigrateProjectPropertiesOutput {
	return o
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o.ToMigrateProjectPropertiesPtrOutputWithContext(context.Background())
}

func (o MigrateProjectPropertiesOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectProperties) *MigrateProjectProperties {
		return &v
	}).(MigrateProjectPropertiesPtrOutput)
}

func (o MigrateProjectPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrateProjectProperties) []string { return v.RegisteredTools }).(pulumi.StringArrayOutput)
}

type MigrateProjectPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectProperties)(nil)).Elem()
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutput() MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) ToMigrateProjectPropertiesPtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesPtrOutput {
	return o
}

func (o MigrateProjectPropertiesPtrOutput) Elem() MigrateProjectPropertiesOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) MigrateProjectProperties {
		if v != nil {
			return *v
		}
		var ret MigrateProjectProperties
		return ret
	}).(MigrateProjectPropertiesOutput)
}

func (o MigrateProjectPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesPtrOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MigrateProjectProperties) []string {
		if v == nil {
			return nil
		}
		return v.RegisteredTools
	}).(pulumi.StringArrayOutput)
}

type MigrateProjectPropertiesResponse struct {
	LastSummaryRefreshedTime string                 `pulumi:"lastSummaryRefreshedTime"`
	ProvisioningState        *string                `pulumi:"provisioningState"`
	RefreshSummaryState      string                 `pulumi:"refreshSummaryState"`
	RegisteredTools          []string               `pulumi:"registeredTools"`
	Summary                  map[string]interface{} `pulumi:"summary"`
}





type MigrateProjectPropertiesResponseInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesResponseOutput() MigrateProjectPropertiesResponseOutput
	ToMigrateProjectPropertiesResponseOutputWithContext(context.Context) MigrateProjectPropertiesResponseOutput
}

type MigrateProjectPropertiesResponseArgs struct {
	LastSummaryRefreshedTime pulumi.StringInput      `pulumi:"lastSummaryRefreshedTime"`
	ProvisioningState        pulumi.StringPtrInput   `pulumi:"provisioningState"`
	RefreshSummaryState      pulumi.StringInput      `pulumi:"refreshSummaryState"`
	RegisteredTools          pulumi.StringArrayInput `pulumi:"registeredTools"`
	Summary                  pulumi.MapInput         `pulumi:"summary"`
}

func (MigrateProjectPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (i MigrateProjectPropertiesResponseArgs) ToMigrateProjectPropertiesResponseOutput() MigrateProjectPropertiesResponseOutput {
	return i.ToMigrateProjectPropertiesResponseOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesResponseArgs) ToMigrateProjectPropertiesResponseOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesResponseOutput)
}

func (i MigrateProjectPropertiesResponseArgs) ToMigrateProjectPropertiesResponsePtrOutput() MigrateProjectPropertiesResponsePtrOutput {
	return i.ToMigrateProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MigrateProjectPropertiesResponseArgs) ToMigrateProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesResponseOutput).ToMigrateProjectPropertiesResponsePtrOutputWithContext(ctx)
}









type MigrateProjectPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMigrateProjectPropertiesResponsePtrOutput() MigrateProjectPropertiesResponsePtrOutput
	ToMigrateProjectPropertiesResponsePtrOutputWithContext(context.Context) MigrateProjectPropertiesResponsePtrOutput
}

type migrateProjectPropertiesResponsePtrType MigrateProjectPropertiesResponseArgs

func MigrateProjectPropertiesResponsePtr(v *MigrateProjectPropertiesResponseArgs) MigrateProjectPropertiesResponsePtrInput {
	return (*migrateProjectPropertiesResponsePtrType)(v)
}

func (*migrateProjectPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (i *migrateProjectPropertiesResponsePtrType) ToMigrateProjectPropertiesResponsePtrOutput() MigrateProjectPropertiesResponsePtrOutput {
	return i.ToMigrateProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *migrateProjectPropertiesResponsePtrType) ToMigrateProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectPropertiesResponsePtrOutput)
}

type MigrateProjectPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutput() MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponseOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponseOutput {
	return o
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponsePtrOutput() MigrateProjectPropertiesResponsePtrOutput {
	return o.ToMigrateProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MigrateProjectPropertiesResponseOutput) ToMigrateProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectPropertiesResponse) *MigrateProjectPropertiesResponse {
		return &v
	}).(MigrateProjectPropertiesResponsePtrOutput)
}

func (o MigrateProjectPropertiesResponseOutput) LastSummaryRefreshedTime() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.LastSummaryRefreshedTime }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RefreshSummaryState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) string { return v.RefreshSummaryState }).(pulumi.StringOutput)
}

func (o MigrateProjectPropertiesResponseOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) []string { return v.RegisteredTools }).(pulumi.StringArrayOutput)
}

func (o MigrateProjectPropertiesResponseOutput) Summary() pulumi.MapOutput {
	return o.ApplyT(func(v MigrateProjectPropertiesResponse) map[string]interface{} { return v.Summary }).(pulumi.MapOutput)
}

type MigrateProjectPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectPropertiesResponse)(nil)).Elem()
}

func (o MigrateProjectPropertiesResponsePtrOutput) ToMigrateProjectPropertiesResponsePtrOutput() MigrateProjectPropertiesResponsePtrOutput {
	return o
}

func (o MigrateProjectPropertiesResponsePtrOutput) ToMigrateProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrateProjectPropertiesResponsePtrOutput {
	return o
}

func (o MigrateProjectPropertiesResponsePtrOutput) Elem() MigrateProjectPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) MigrateProjectPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MigrateProjectPropertiesResponse
		return ret
	}).(MigrateProjectPropertiesResponseOutput)
}

func (o MigrateProjectPropertiesResponsePtrOutput) LastSummaryRefreshedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSummaryRefreshedTime
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponsePtrOutput) RefreshSummaryState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RefreshSummaryState
	}).(pulumi.StringPtrOutput)
}

func (o MigrateProjectPropertiesResponsePtrOutput) RegisteredTools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.RegisteredTools
	}).(pulumi.StringArrayOutput)
}

func (o MigrateProjectPropertiesResponsePtrOutput) Summary() pulumi.MapOutput {
	return o.ApplyT(func(v *MigrateProjectPropertiesResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Summary
	}).(pulumi.MapOutput)
}

type MigrateProjectResponseTags struct {
	AdditionalProperties *string `pulumi:"additionalProperties"`
}





type MigrateProjectResponseTagsInput interface {
	pulumi.Input

	ToMigrateProjectResponseTagsOutput() MigrateProjectResponseTagsOutput
	ToMigrateProjectResponseTagsOutputWithContext(context.Context) MigrateProjectResponseTagsOutput
}

type MigrateProjectResponseTagsArgs struct {
	AdditionalProperties pulumi.StringPtrInput `pulumi:"additionalProperties"`
}

func (MigrateProjectResponseTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectResponseTags)(nil)).Elem()
}

func (i MigrateProjectResponseTagsArgs) ToMigrateProjectResponseTagsOutput() MigrateProjectResponseTagsOutput {
	return i.ToMigrateProjectResponseTagsOutputWithContext(context.Background())
}

func (i MigrateProjectResponseTagsArgs) ToMigrateProjectResponseTagsOutputWithContext(ctx context.Context) MigrateProjectResponseTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectResponseTagsOutput)
}

func (i MigrateProjectResponseTagsArgs) ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput {
	return i.ToMigrateProjectResponseTagsPtrOutputWithContext(context.Background())
}

func (i MigrateProjectResponseTagsArgs) ToMigrateProjectResponseTagsPtrOutputWithContext(ctx context.Context) MigrateProjectResponseTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectResponseTagsOutput).ToMigrateProjectResponseTagsPtrOutputWithContext(ctx)
}









type MigrateProjectResponseTagsPtrInput interface {
	pulumi.Input

	ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput
	ToMigrateProjectResponseTagsPtrOutputWithContext(context.Context) MigrateProjectResponseTagsPtrOutput
}

type migrateProjectResponseTagsPtrType MigrateProjectResponseTagsArgs

func MigrateProjectResponseTagsPtr(v *MigrateProjectResponseTagsArgs) MigrateProjectResponseTagsPtrInput {
	return (*migrateProjectResponseTagsPtrType)(v)
}

func (*migrateProjectResponseTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectResponseTags)(nil)).Elem()
}

func (i *migrateProjectResponseTagsPtrType) ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput {
	return i.ToMigrateProjectResponseTagsPtrOutputWithContext(context.Background())
}

func (i *migrateProjectResponseTagsPtrType) ToMigrateProjectResponseTagsPtrOutputWithContext(ctx context.Context) MigrateProjectResponseTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectResponseTagsPtrOutput)
}

type MigrateProjectResponseTagsOutput struct{ *pulumi.OutputState }

func (MigrateProjectResponseTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectResponseTags)(nil)).Elem()
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsOutput() MigrateProjectResponseTagsOutput {
	return o
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsOutputWithContext(ctx context.Context) MigrateProjectResponseTagsOutput {
	return o
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput {
	return o.ToMigrateProjectResponseTagsPtrOutputWithContext(context.Background())
}

func (o MigrateProjectResponseTagsOutput) ToMigrateProjectResponseTagsPtrOutputWithContext(ctx context.Context) MigrateProjectResponseTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectResponseTags) *MigrateProjectResponseTags {
		return &v
	}).(MigrateProjectResponseTagsPtrOutput)
}

func (o MigrateProjectResponseTagsOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectResponseTags) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

type MigrateProjectResponseTagsPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectResponseTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectResponseTags)(nil)).Elem()
}

func (o MigrateProjectResponseTagsPtrOutput) ToMigrateProjectResponseTagsPtrOutput() MigrateProjectResponseTagsPtrOutput {
	return o
}

func (o MigrateProjectResponseTagsPtrOutput) ToMigrateProjectResponseTagsPtrOutputWithContext(ctx context.Context) MigrateProjectResponseTagsPtrOutput {
	return o
}

func (o MigrateProjectResponseTagsPtrOutput) Elem() MigrateProjectResponseTagsOutput {
	return o.ApplyT(func(v *MigrateProjectResponseTags) MigrateProjectResponseTags {
		if v != nil {
			return *v
		}
		var ret MigrateProjectResponseTags
		return ret
	}).(MigrateProjectResponseTagsOutput)
}

func (o MigrateProjectResponseTagsPtrOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectResponseTags) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(pulumi.StringPtrOutput)
}

type MigrateProjectTags struct {
	AdditionalProperties *string `pulumi:"additionalProperties"`
}





type MigrateProjectTagsInput interface {
	pulumi.Input

	ToMigrateProjectTagsOutput() MigrateProjectTagsOutput
	ToMigrateProjectTagsOutputWithContext(context.Context) MigrateProjectTagsOutput
}

type MigrateProjectTagsArgs struct {
	AdditionalProperties pulumi.StringPtrInput `pulumi:"additionalProperties"`
}

func (MigrateProjectTagsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectTags)(nil)).Elem()
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsOutput() MigrateProjectTagsOutput {
	return i.ToMigrateProjectTagsOutputWithContext(context.Background())
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsOutputWithContext(ctx context.Context) MigrateProjectTagsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsOutput)
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return i.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (i MigrateProjectTagsArgs) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsOutput).ToMigrateProjectTagsPtrOutputWithContext(ctx)
}









type MigrateProjectTagsPtrInput interface {
	pulumi.Input

	ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput
	ToMigrateProjectTagsPtrOutputWithContext(context.Context) MigrateProjectTagsPtrOutput
}

type migrateProjectTagsPtrType MigrateProjectTagsArgs

func MigrateProjectTagsPtr(v *MigrateProjectTagsArgs) MigrateProjectTagsPtrInput {
	return (*migrateProjectTagsPtrType)(v)
}

func (*migrateProjectTagsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectTags)(nil)).Elem()
}

func (i *migrateProjectTagsPtrType) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return i.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (i *migrateProjectTagsPtrType) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrateProjectTagsPtrOutput)
}

type MigrateProjectTagsOutput struct{ *pulumi.OutputState }

func (MigrateProjectTagsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrateProjectTags)(nil)).Elem()
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsOutput() MigrateProjectTagsOutput {
	return o
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsOutputWithContext(ctx context.Context) MigrateProjectTagsOutput {
	return o
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return o.ToMigrateProjectTagsPtrOutputWithContext(context.Background())
}

func (o MigrateProjectTagsOutput) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrateProjectTags) *MigrateProjectTags {
		return &v
	}).(MigrateProjectTagsPtrOutput)
}

func (o MigrateProjectTagsOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MigrateProjectTags) *string { return v.AdditionalProperties }).(pulumi.StringPtrOutput)
}

type MigrateProjectTagsPtrOutput struct{ *pulumi.OutputState }

func (MigrateProjectTagsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrateProjectTags)(nil)).Elem()
}

func (o MigrateProjectTagsPtrOutput) ToMigrateProjectTagsPtrOutput() MigrateProjectTagsPtrOutput {
	return o
}

func (o MigrateProjectTagsPtrOutput) ToMigrateProjectTagsPtrOutputWithContext(ctx context.Context) MigrateProjectTagsPtrOutput {
	return o
}

func (o MigrateProjectTagsPtrOutput) Elem() MigrateProjectTagsOutput {
	return o.ApplyT(func(v *MigrateProjectTags) MigrateProjectTags {
		if v != nil {
			return *v
		}
		var ret MigrateProjectTags
		return ret
	}).(MigrateProjectTagsOutput)
}

func (o MigrateProjectTagsPtrOutput) AdditionalProperties() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrateProjectTags) *string {
		if v == nil {
			return nil
		}
		return v.AdditionalProperties
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionProperties struct {
	SourceRegion string `pulumi:"sourceRegion"`
	TargetRegion string `pulumi:"targetRegion"`
}





type MoveCollectionPropertiesInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput
	ToMoveCollectionPropertiesOutputWithContext(context.Context) MoveCollectionPropertiesOutput
}

type MoveCollectionPropertiesArgs struct {
	SourceRegion pulumi.StringInput `pulumi:"sourceRegion"`
	TargetRegion pulumi.StringInput `pulumi:"targetRegion"`
}

func (MoveCollectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionProperties)(nil)).Elem()
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput {
	return i.ToMoveCollectionPropertiesOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesOutputWithContext(ctx context.Context) MoveCollectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesOutput)
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return i.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesArgs) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesOutput).ToMoveCollectionPropertiesPtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesPtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput
	ToMoveCollectionPropertiesPtrOutputWithContext(context.Context) MoveCollectionPropertiesPtrOutput
}

type moveCollectionPropertiesPtrType MoveCollectionPropertiesArgs

func MoveCollectionPropertiesPtr(v *MoveCollectionPropertiesArgs) MoveCollectionPropertiesPtrInput {
	return (*moveCollectionPropertiesPtrType)(v)
}

func (*moveCollectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionProperties)(nil)).Elem()
}

func (i *moveCollectionPropertiesPtrType) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return i.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesPtrType) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesPtrOutput)
}

type MoveCollectionPropertiesOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionProperties)(nil)).Elem()
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesOutput() MoveCollectionPropertiesOutput {
	return o
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesOutputWithContext(ctx context.Context) MoveCollectionPropertiesOutput {
	return o
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return o.ToMoveCollectionPropertiesPtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesOutput) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionProperties) *MoveCollectionProperties {
		return &v
	}).(MoveCollectionPropertiesPtrOutput)
}

func (o MoveCollectionPropertiesOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionProperties) string { return v.SourceRegion }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesOutput) TargetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionProperties) string { return v.TargetRegion }).(pulumi.StringOutput)
}

type MoveCollectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionProperties)(nil)).Elem()
}

func (o MoveCollectionPropertiesPtrOutput) ToMoveCollectionPropertiesPtrOutput() MoveCollectionPropertiesPtrOutput {
	return o
}

func (o MoveCollectionPropertiesPtrOutput) ToMoveCollectionPropertiesPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesPtrOutput {
	return o
}

func (o MoveCollectionPropertiesPtrOutput) Elem() MoveCollectionPropertiesOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) MoveCollectionProperties {
		if v != nil {
			return *v
		}
		var ret MoveCollectionProperties
		return ret
	}).(MoveCollectionPropertiesOutput)
}

func (o MoveCollectionPropertiesPtrOutput) SourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesPtrOutput) TargetRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TargetRegion
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionPropertiesResponse struct {
	Errors            MoveCollectionPropertiesResponseErrors `pulumi:"errors"`
	ProvisioningState string                                 `pulumi:"provisioningState"`
	SourceRegion      string                                 `pulumi:"sourceRegion"`
	TargetRegion      string                                 `pulumi:"targetRegion"`
}





type MoveCollectionPropertiesResponseInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput
	ToMoveCollectionPropertiesResponseOutputWithContext(context.Context) MoveCollectionPropertiesResponseOutput
}

type MoveCollectionPropertiesResponseArgs struct {
	Errors            MoveCollectionPropertiesResponseErrorsInput `pulumi:"errors"`
	ProvisioningState pulumi.StringInput                          `pulumi:"provisioningState"`
	SourceRegion      pulumi.StringInput                          `pulumi:"sourceRegion"`
	TargetRegion      pulumi.StringInput                          `pulumi:"targetRegion"`
}

func (MoveCollectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput {
	return i.ToMoveCollectionPropertiesResponseOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponseOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseOutput)
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return i.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseArgs) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseOutput).ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput
	ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Context) MoveCollectionPropertiesResponsePtrOutput
}

type moveCollectionPropertiesResponsePtrType MoveCollectionPropertiesResponseArgs

func MoveCollectionPropertiesResponsePtr(v *MoveCollectionPropertiesResponseArgs) MoveCollectionPropertiesResponsePtrInput {
	return (*moveCollectionPropertiesResponsePtrType)(v)
}

func (*moveCollectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (i *moveCollectionPropertiesResponsePtrType) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return i.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesResponsePtrType) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponsePtrOutput)
}

type MoveCollectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponseOutput() MoveCollectionPropertiesResponseOutput {
	return o
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponseOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseOutput {
	return o
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return o.ToMoveCollectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesResponseOutput) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionPropertiesResponse) *MoveCollectionPropertiesResponse {
		return &v
	}).(MoveCollectionPropertiesResponsePtrOutput)
}

func (o MoveCollectionPropertiesResponseOutput) Errors() MoveCollectionPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) MoveCollectionPropertiesResponseErrors { return v.Errors }).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (o MoveCollectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesResponseOutput) SourceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.SourceRegion }).(pulumi.StringOutput)
}

func (o MoveCollectionPropertiesResponseOutput) TargetRegion() pulumi.StringOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponse) string { return v.TargetRegion }).(pulumi.StringOutput)
}

type MoveCollectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponse)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponsePtrOutput) ToMoveCollectionPropertiesResponsePtrOutput() MoveCollectionPropertiesResponsePtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponsePtrOutput) ToMoveCollectionPropertiesResponsePtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponsePtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponsePtrOutput) Elem() MoveCollectionPropertiesResponseOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) MoveCollectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MoveCollectionPropertiesResponse
		return ret
	}).(MoveCollectionPropertiesResponseOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) Errors() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *MoveCollectionPropertiesResponseErrors {
		if v == nil {
			return nil
		}
		return &v.Errors
	}).(MoveCollectionPropertiesResponseErrorsPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) SourceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceRegion
	}).(pulumi.StringPtrOutput)
}

func (o MoveCollectionPropertiesResponsePtrOutput) TargetRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetRegion
	}).(pulumi.StringPtrOutput)
}

type MoveCollectionPropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}





type MoveCollectionPropertiesResponseErrorsInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput
	ToMoveCollectionPropertiesResponseErrorsOutputWithContext(context.Context) MoveCollectionPropertiesResponseErrorsOutput
}

type MoveCollectionPropertiesResponseErrorsArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveCollectionPropertiesResponseErrorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i MoveCollectionPropertiesResponseErrorsArgs) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsOutput).ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx)
}









type MoveCollectionPropertiesResponseErrorsPtrInput interface {
	pulumi.Input

	ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput
	ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput
}

type moveCollectionPropertiesResponseErrorsPtrType MoveCollectionPropertiesResponseErrorsArgs

func MoveCollectionPropertiesResponseErrorsPtr(v *MoveCollectionPropertiesResponseErrorsArgs) MoveCollectionPropertiesResponseErrorsPtrInput {
	return (*moveCollectionPropertiesResponseErrorsPtrType)(v)
}

func (*moveCollectionPropertiesResponseErrorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (i *moveCollectionPropertiesResponseErrorsPtrType) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return i.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i *moveCollectionPropertiesResponseErrorsPtrType) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveCollectionPropertiesResponseErrorsPtrOutput)
}

type MoveCollectionPropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsOutput() MoveCollectionPropertiesResponseErrorsOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (o MoveCollectionPropertiesResponseErrorsOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCollectionPropertiesResponseErrors) *MoveCollectionPropertiesResponseErrors {
		return &v
	}).(MoveCollectionPropertiesResponseErrorsPtrOutput)
}

func (o MoveCollectionPropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveCollectionPropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveCollectionPropertiesResponseErrorsPtrOutput struct{ *pulumi.OutputState }

func (MoveCollectionPropertiesResponseErrorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCollectionPropertiesResponseErrors)(nil)).Elem()
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutput() MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) ToMoveCollectionPropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveCollectionPropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) Elem() MoveCollectionPropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponseErrors) MoveCollectionPropertiesResponseErrors {
		if v != nil {
			return *v
		}
		var ret MoveCollectionPropertiesResponseErrors
		return ret
	}).(MoveCollectionPropertiesResponseErrorsOutput)
}

func (o MoveCollectionPropertiesResponseErrorsPtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveCollectionPropertiesResponseErrors) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceDependencyOverride struct {
	Id       *string `pulumi:"id"`
	TargetId *string `pulumi:"targetId"`
}





type MoveResourceDependencyOverrideInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput
	ToMoveResourceDependencyOverrideOutputWithContext(context.Context) MoveResourceDependencyOverrideOutput
}

type MoveResourceDependencyOverrideArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (MoveResourceDependencyOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverride)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideArgs) ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput {
	return i.ToMoveResourceDependencyOverrideOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideArgs) ToMoveResourceDependencyOverrideOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideOutput)
}





type MoveResourceDependencyOverrideArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput
	ToMoveResourceDependencyOverrideArrayOutputWithContext(context.Context) MoveResourceDependencyOverrideArrayOutput
}

type MoveResourceDependencyOverrideArray []MoveResourceDependencyOverrideInput

func (MoveResourceDependencyOverrideArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverride)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideArray) ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput {
	return i.ToMoveResourceDependencyOverrideArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideArray) ToMoveResourceDependencyOverrideArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideArrayOutput)
}

type MoveResourceDependencyOverrideOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverride)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideOutput) ToMoveResourceDependencyOverrideOutput() MoveResourceDependencyOverrideOutput {
	return o
}

func (o MoveResourceDependencyOverrideOutput) ToMoveResourceDependencyOverrideOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideOutput {
	return o
}

func (o MoveResourceDependencyOverrideOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverride) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyOverrideOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverride) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyOverrideArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverride)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideArrayOutput) ToMoveResourceDependencyOverrideArrayOutput() MoveResourceDependencyOverrideArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideArrayOutput) ToMoveResourceDependencyOverrideArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyOverrideOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyOverride {
		return vs[0].([]MoveResourceDependencyOverride)[vs[1].(int)]
	}).(MoveResourceDependencyOverrideOutput)
}

type MoveResourceDependencyOverrideResponse struct {
	Id       *string `pulumi:"id"`
	TargetId *string `pulumi:"targetId"`
}





type MoveResourceDependencyOverrideResponseInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput
	ToMoveResourceDependencyOverrideResponseOutputWithContext(context.Context) MoveResourceDependencyOverrideResponseOutput
}

type MoveResourceDependencyOverrideResponseArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	TargetId pulumi.StringPtrInput `pulumi:"targetId"`
}

func (MoveResourceDependencyOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideResponseArgs) ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput {
	return i.ToMoveResourceDependencyOverrideResponseOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideResponseArgs) ToMoveResourceDependencyOverrideResponseOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideResponseOutput)
}





type MoveResourceDependencyOverrideResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput
	ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(context.Context) MoveResourceDependencyOverrideResponseArrayOutput
}

type MoveResourceDependencyOverrideResponseArray []MoveResourceDependencyOverrideResponseInput

func (MoveResourceDependencyOverrideResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (i MoveResourceDependencyOverrideResponseArray) ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput {
	return i.ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyOverrideResponseArray) ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyOverrideResponseArrayOutput)
}

type MoveResourceDependencyOverrideResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideResponseOutput) ToMoveResourceDependencyOverrideResponseOutput() MoveResourceDependencyOverrideResponseOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseOutput) ToMoveResourceDependencyOverrideResponseOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverrideResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyOverrideResponseOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyOverrideResponse) *string { return v.TargetId }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyOverrideResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyOverrideResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyOverrideResponse)(nil)).Elem()
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) ToMoveResourceDependencyOverrideResponseArrayOutput() MoveResourceDependencyOverrideResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) ToMoveResourceDependencyOverrideResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyOverrideResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyOverrideResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyOverrideResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyOverrideResponse {
		return vs[0].([]MoveResourceDependencyOverrideResponse)[vs[1].(int)]
	}).(MoveResourceDependencyOverrideResponseOutput)
}

type MoveResourceDependencyResponse struct {
	AutomaticResolution *AutomaticResolutionPropertiesResponse `pulumi:"automaticResolution"`
	DependencyType      *string                                `pulumi:"dependencyType"`
	Id                  *string                                `pulumi:"id"`
	IsOptional          *string                                `pulumi:"isOptional"`
	ManualResolution    *ManualResolutionPropertiesResponse    `pulumi:"manualResolution"`
	ResolutionStatus    *string                                `pulumi:"resolutionStatus"`
	ResolutionType      *string                                `pulumi:"resolutionType"`
}





type MoveResourceDependencyResponseInput interface {
	pulumi.Input

	ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput
	ToMoveResourceDependencyResponseOutputWithContext(context.Context) MoveResourceDependencyResponseOutput
}

type MoveResourceDependencyResponseArgs struct {
	AutomaticResolution AutomaticResolutionPropertiesResponsePtrInput `pulumi:"automaticResolution"`
	DependencyType      pulumi.StringPtrInput                         `pulumi:"dependencyType"`
	Id                  pulumi.StringPtrInput                         `pulumi:"id"`
	IsOptional          pulumi.StringPtrInput                         `pulumi:"isOptional"`
	ManualResolution    ManualResolutionPropertiesResponsePtrInput    `pulumi:"manualResolution"`
	ResolutionStatus    pulumi.StringPtrInput                         `pulumi:"resolutionStatus"`
	ResolutionType      pulumi.StringPtrInput                         `pulumi:"resolutionType"`
}

func (MoveResourceDependencyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyResponse)(nil)).Elem()
}

func (i MoveResourceDependencyResponseArgs) ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput {
	return i.ToMoveResourceDependencyResponseOutputWithContext(context.Background())
}

func (i MoveResourceDependencyResponseArgs) ToMoveResourceDependencyResponseOutputWithContext(ctx context.Context) MoveResourceDependencyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyResponseOutput)
}





type MoveResourceDependencyResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput
	ToMoveResourceDependencyResponseArrayOutputWithContext(context.Context) MoveResourceDependencyResponseArrayOutput
}

type MoveResourceDependencyResponseArray []MoveResourceDependencyResponseInput

func (MoveResourceDependencyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyResponse)(nil)).Elem()
}

func (i MoveResourceDependencyResponseArray) ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput {
	return i.ToMoveResourceDependencyResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceDependencyResponseArray) ToMoveResourceDependencyResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceDependencyResponseArrayOutput)
}

type MoveResourceDependencyResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceDependencyResponse)(nil)).Elem()
}

func (o MoveResourceDependencyResponseOutput) ToMoveResourceDependencyResponseOutput() MoveResourceDependencyResponseOutput {
	return o
}

func (o MoveResourceDependencyResponseOutput) ToMoveResourceDependencyResponseOutputWithContext(ctx context.Context) MoveResourceDependencyResponseOutput {
	return o
}

func (o MoveResourceDependencyResponseOutput) AutomaticResolution() AutomaticResolutionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *AutomaticResolutionPropertiesResponse {
		return v.AutomaticResolution
	}).(AutomaticResolutionPropertiesResponsePtrOutput)
}

func (o MoveResourceDependencyResponseOutput) DependencyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.DependencyType }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) IsOptional() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.IsOptional }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ManualResolution() ManualResolutionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *ManualResolutionPropertiesResponse { return v.ManualResolution }).(ManualResolutionPropertiesResponsePtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ResolutionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.ResolutionStatus }).(pulumi.StringPtrOutput)
}

func (o MoveResourceDependencyResponseOutput) ResolutionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceDependencyResponse) *string { return v.ResolutionType }).(pulumi.StringPtrOutput)
}

type MoveResourceDependencyResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceDependencyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceDependencyResponse)(nil)).Elem()
}

func (o MoveResourceDependencyResponseArrayOutput) ToMoveResourceDependencyResponseArrayOutput() MoveResourceDependencyResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyResponseArrayOutput) ToMoveResourceDependencyResponseArrayOutputWithContext(ctx context.Context) MoveResourceDependencyResponseArrayOutput {
	return o
}

func (o MoveResourceDependencyResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceDependencyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceDependencyResponse {
		return vs[0].([]MoveResourceDependencyResponse)[vs[1].(int)]
	}).(MoveResourceDependencyResponseOutput)
}

type MoveResourceErrorBodyResponse struct {
	Code    string                          `pulumi:"code"`
	Details []MoveResourceErrorBodyResponse `pulumi:"details"`
	Message string                          `pulumi:"message"`
	Target  string                          `pulumi:"target"`
}





type MoveResourceErrorBodyResponseInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput
	ToMoveResourceErrorBodyResponseOutputWithContext(context.Context) MoveResourceErrorBodyResponseOutput
}

type MoveResourceErrorBodyResponseArgs struct {
	Code    pulumi.StringInput                      `pulumi:"code"`
	Details MoveResourceErrorBodyResponseArrayInput `pulumi:"details"`
	Message pulumi.StringInput                      `pulumi:"message"`
	Target  pulumi.StringInput                      `pulumi:"target"`
}

func (MoveResourceErrorBodyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput {
	return i.ToMoveResourceErrorBodyResponseOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponseOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseOutput)
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return i.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArgs) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseOutput).ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx)
}









type MoveResourceErrorBodyResponsePtrInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput
	ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Context) MoveResourceErrorBodyResponsePtrOutput
}

type moveResourceErrorBodyResponsePtrType MoveResourceErrorBodyResponseArgs

func MoveResourceErrorBodyResponsePtr(v *MoveResourceErrorBodyResponseArgs) MoveResourceErrorBodyResponsePtrInput {
	return (*moveResourceErrorBodyResponsePtrType)(v)
}

func (*moveResourceErrorBodyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i *moveResourceErrorBodyResponsePtrType) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return i.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourceErrorBodyResponsePtrType) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponsePtrOutput)
}





type MoveResourceErrorBodyResponseArrayInput interface {
	pulumi.Input

	ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput
	ToMoveResourceErrorBodyResponseArrayOutputWithContext(context.Context) MoveResourceErrorBodyResponseArrayOutput
}

type MoveResourceErrorBodyResponseArray []MoveResourceErrorBodyResponseInput

func (MoveResourceErrorBodyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (i MoveResourceErrorBodyResponseArray) ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput {
	return i.ToMoveResourceErrorBodyResponseArrayOutputWithContext(context.Background())
}

func (i MoveResourceErrorBodyResponseArray) ToMoveResourceErrorBodyResponseArrayOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorBodyResponseArrayOutput)
}

type MoveResourceErrorBodyResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponseOutput() MoveResourceErrorBodyResponseOutput {
	return o
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponseOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseOutput {
	return o
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return o.ToMoveResourceErrorBodyResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourceErrorBodyResponseOutput) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceErrorBodyResponse) *MoveResourceErrorBodyResponse {
		return &v
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Details() MoveResourceErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) []MoveResourceErrorBodyResponse { return v.Details }).(MoveResourceErrorBodyResponseArrayOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o MoveResourceErrorBodyResponseOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceErrorBodyResponse) string { return v.Target }).(pulumi.StringOutput)
}

type MoveResourceErrorBodyResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponsePtrOutput) ToMoveResourceErrorBodyResponsePtrOutput() MoveResourceErrorBodyResponsePtrOutput {
	return o
}

func (o MoveResourceErrorBodyResponsePtrOutput) ToMoveResourceErrorBodyResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponsePtrOutput {
	return o
}

func (o MoveResourceErrorBodyResponsePtrOutput) Elem() MoveResourceErrorBodyResponseOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) MoveResourceErrorBodyResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourceErrorBodyResponse
		return ret
	}).(MoveResourceErrorBodyResponseOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Details() MoveResourceErrorBodyResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) []MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(MoveResourceErrorBodyResponseArrayOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourceErrorBodyResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorBodyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Target
	}).(pulumi.StringPtrOutput)
}

type MoveResourceErrorBodyResponseArrayOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorBodyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MoveResourceErrorBodyResponse)(nil)).Elem()
}

func (o MoveResourceErrorBodyResponseArrayOutput) ToMoveResourceErrorBodyResponseArrayOutput() MoveResourceErrorBodyResponseArrayOutput {
	return o
}

func (o MoveResourceErrorBodyResponseArrayOutput) ToMoveResourceErrorBodyResponseArrayOutputWithContext(ctx context.Context) MoveResourceErrorBodyResponseArrayOutput {
	return o
}

func (o MoveResourceErrorBodyResponseArrayOutput) Index(i pulumi.IntInput) MoveResourceErrorBodyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MoveResourceErrorBodyResponse {
		return vs[0].([]MoveResourceErrorBodyResponse)[vs[1].(int)]
	}).(MoveResourceErrorBodyResponseOutput)
}

type MoveResourceErrorResponse struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}





type MoveResourceErrorResponseInput interface {
	pulumi.Input

	ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput
	ToMoveResourceErrorResponseOutputWithContext(context.Context) MoveResourceErrorResponseOutput
}

type MoveResourceErrorResponseArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveResourceErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorResponse)(nil)).Elem()
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput {
	return i.ToMoveResourceErrorResponseOutputWithContext(context.Background())
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponseOutputWithContext(ctx context.Context) MoveResourceErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponseOutput)
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return i.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourceErrorResponseArgs) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponseOutput).ToMoveResourceErrorResponsePtrOutputWithContext(ctx)
}









type MoveResourceErrorResponsePtrInput interface {
	pulumi.Input

	ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput
	ToMoveResourceErrorResponsePtrOutputWithContext(context.Context) MoveResourceErrorResponsePtrOutput
}

type moveResourceErrorResponsePtrType MoveResourceErrorResponseArgs

func MoveResourceErrorResponsePtr(v *MoveResourceErrorResponseArgs) MoveResourceErrorResponsePtrInput {
	return (*moveResourceErrorResponsePtrType)(v)
}

func (*moveResourceErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorResponse)(nil)).Elem()
}

func (i *moveResourceErrorResponsePtrType) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return i.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourceErrorResponsePtrType) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourceErrorResponsePtrOutput)
}

type MoveResourceErrorResponseOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceErrorResponse)(nil)).Elem()
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponseOutput() MoveResourceErrorResponseOutput {
	return o
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponseOutputWithContext(ctx context.Context) MoveResourceErrorResponseOutput {
	return o
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return o.ToMoveResourceErrorResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourceErrorResponseOutput) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceErrorResponse) *MoveResourceErrorResponse {
		return &v
	}).(MoveResourceErrorResponsePtrOutput)
}

func (o MoveResourceErrorResponseOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveResourceErrorResponse) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourceErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceErrorResponse)(nil)).Elem()
}

func (o MoveResourceErrorResponsePtrOutput) ToMoveResourceErrorResponsePtrOutput() MoveResourceErrorResponsePtrOutput {
	return o
}

func (o MoveResourceErrorResponsePtrOutput) ToMoveResourceErrorResponsePtrOutputWithContext(ctx context.Context) MoveResourceErrorResponsePtrOutput {
	return o
}

func (o MoveResourceErrorResponsePtrOutput) Elem() MoveResourceErrorResponseOutput {
	return o.ApplyT(func(v *MoveResourceErrorResponse) MoveResourceErrorResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourceErrorResponse
		return ret
	}).(MoveResourceErrorResponseOutput)
}

func (o MoveResourceErrorResponsePtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourceErrorResponse) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourceProperties struct {
	DependsOnOverrides []MoveResourceDependencyOverride `pulumi:"dependsOnOverrides"`
	ExistingTargetId   *string                          `pulumi:"existingTargetId"`
	ResourceSettings   interface{}                      `pulumi:"resourceSettings"`
	SourceId           string                           `pulumi:"sourceId"`
}





type MoveResourcePropertiesInput interface {
	pulumi.Input

	ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput
	ToMoveResourcePropertiesOutputWithContext(context.Context) MoveResourcePropertiesOutput
}

type MoveResourcePropertiesArgs struct {
	DependsOnOverrides MoveResourceDependencyOverrideArrayInput `pulumi:"dependsOnOverrides"`
	ExistingTargetId   pulumi.StringPtrInput                    `pulumi:"existingTargetId"`
	ResourceSettings   pulumi.Input                             `pulumi:"resourceSettings"`
	SourceId           pulumi.StringInput                       `pulumi:"sourceId"`
}

func (MoveResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceProperties)(nil)).Elem()
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput {
	return i.ToMoveResourcePropertiesOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesOutputWithContext(ctx context.Context) MoveResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesOutput)
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return i.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesArgs) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesOutput).ToMoveResourcePropertiesPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput
	ToMoveResourcePropertiesPtrOutputWithContext(context.Context) MoveResourcePropertiesPtrOutput
}

type moveResourcePropertiesPtrType MoveResourcePropertiesArgs

func MoveResourcePropertiesPtr(v *MoveResourcePropertiesArgs) MoveResourcePropertiesPtrInput {
	return (*moveResourcePropertiesPtrType)(v)
}

func (*moveResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceProperties)(nil)).Elem()
}

func (i *moveResourcePropertiesPtrType) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return i.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesPtrType) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesPtrOutput)
}

type MoveResourcePropertiesOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourceProperties)(nil)).Elem()
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesOutput() MoveResourcePropertiesOutput {
	return o
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesOutputWithContext(ctx context.Context) MoveResourcePropertiesOutput {
	return o
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return o.ToMoveResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesOutput) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourceProperties) *MoveResourceProperties {
		return &v
	}).(MoveResourcePropertiesPtrOutput)
}

func (o MoveResourcePropertiesOutput) DependsOnOverrides() MoveResourceDependencyOverrideArrayOutput {
	return o.ApplyT(func(v MoveResourceProperties) []MoveResourceDependencyOverride { return v.DependsOnOverrides }).(MoveResourceDependencyOverrideArrayOutput)
}

func (o MoveResourcePropertiesOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourceProperties) *string { return v.ExistingTargetId }).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourceProperties) interface{} { return v.ResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourceProperties) string { return v.SourceId }).(pulumi.StringOutput)
}

type MoveResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourceProperties)(nil)).Elem()
}

func (o MoveResourcePropertiesPtrOutput) ToMoveResourcePropertiesPtrOutput() MoveResourcePropertiesPtrOutput {
	return o
}

func (o MoveResourcePropertiesPtrOutput) ToMoveResourcePropertiesPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesPtrOutput {
	return o
}

func (o MoveResourcePropertiesPtrOutput) Elem() MoveResourcePropertiesOutput {
	return o.ApplyT(func(v *MoveResourceProperties) MoveResourceProperties {
		if v != nil {
			return *v
		}
		var ret MoveResourceProperties
		return ret
	}).(MoveResourcePropertiesOutput)
}

func (o MoveResourcePropertiesPtrOutput) DependsOnOverrides() MoveResourceDependencyOverrideArrayOutput {
	return o.ApplyT(func(v *MoveResourceProperties) []MoveResourceDependencyOverride {
		if v == nil {
			return nil
		}
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideArrayOutput)
}

func (o MoveResourcePropertiesPtrOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ExistingTargetId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesPtrOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourceProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesPtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SourceId
	}).(pulumi.StringPtrOutput)
}

type MoveResourcePropertiesResponse struct {
	DependsOn              []MoveResourceDependencyResponse         `pulumi:"dependsOn"`
	DependsOnOverrides     []MoveResourceDependencyOverrideResponse `pulumi:"dependsOnOverrides"`
	Errors                 MoveResourcePropertiesResponseErrors     `pulumi:"errors"`
	ExistingTargetId       *string                                  `pulumi:"existingTargetId"`
	IsResolveRequired      bool                                     `pulumi:"isResolveRequired"`
	MoveStatus             MoveResourcePropertiesResponseMoveStatus `pulumi:"moveStatus"`
	ProvisioningState      string                                   `pulumi:"provisioningState"`
	ResourceSettings       interface{}                              `pulumi:"resourceSettings"`
	SourceId               string                                   `pulumi:"sourceId"`
	SourceResourceSettings interface{}                              `pulumi:"sourceResourceSettings"`
	TargetId               string                                   `pulumi:"targetId"`
}





type MoveResourcePropertiesResponseInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput
	ToMoveResourcePropertiesResponseOutputWithContext(context.Context) MoveResourcePropertiesResponseOutput
}

type MoveResourcePropertiesResponseArgs struct {
	DependsOn              MoveResourceDependencyResponseArrayInput         `pulumi:"dependsOn"`
	DependsOnOverrides     MoveResourceDependencyOverrideResponseArrayInput `pulumi:"dependsOnOverrides"`
	Errors                 MoveResourcePropertiesResponseErrorsInput        `pulumi:"errors"`
	ExistingTargetId       pulumi.StringPtrInput                            `pulumi:"existingTargetId"`
	IsResolveRequired      pulumi.BoolInput                                 `pulumi:"isResolveRequired"`
	MoveStatus             MoveResourcePropertiesResponseMoveStatusInput    `pulumi:"moveStatus"`
	ProvisioningState      pulumi.StringInput                               `pulumi:"provisioningState"`
	ResourceSettings       pulumi.Input                                     `pulumi:"resourceSettings"`
	SourceId               pulumi.StringInput                               `pulumi:"sourceId"`
	SourceResourceSettings pulumi.Input                                     `pulumi:"sourceResourceSettings"`
	TargetId               pulumi.StringInput                               `pulumi:"targetId"`
}

func (MoveResourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponse)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput {
	return i.ToMoveResourcePropertiesResponseOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponseOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseOutput)
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return i.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseArgs) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseOutput).ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput
	ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Context) MoveResourcePropertiesResponsePtrOutput
}

type moveResourcePropertiesResponsePtrType MoveResourcePropertiesResponseArgs

func MoveResourcePropertiesResponsePtr(v *MoveResourcePropertiesResponseArgs) MoveResourcePropertiesResponsePtrInput {
	return (*moveResourcePropertiesResponsePtrType)(v)
}

func (*moveResourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponse)(nil)).Elem()
}

func (i *moveResourcePropertiesResponsePtrType) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return i.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponsePtrType) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponsePtrOutput)
}

type MoveResourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponse)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponseOutput() MoveResourcePropertiesResponseOutput {
	return o
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponseOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseOutput {
	return o
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return o.ToMoveResourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseOutput) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponse) *MoveResourcePropertiesResponse {
		return &v
	}).(MoveResourcePropertiesResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseOutput) DependsOn() MoveResourceDependencyResponseArrayOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) []MoveResourceDependencyResponse { return v.DependsOn }).(MoveResourceDependencyResponseArrayOutput)
}

func (o MoveResourcePropertiesResponseOutput) DependsOnOverrides() MoveResourceDependencyOverrideResponseArrayOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) []MoveResourceDependencyOverrideResponse {
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideResponseArrayOutput)
}

func (o MoveResourcePropertiesResponseOutput) Errors() MoveResourcePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) MoveResourcePropertiesResponseErrors { return v.Errors }).(MoveResourcePropertiesResponseErrorsOutput)
}

func (o MoveResourcePropertiesResponseOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) *string { return v.ExistingTargetId }).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponseOutput) IsResolveRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) bool { return v.IsResolveRequired }).(pulumi.BoolOutput)
}

func (o MoveResourcePropertiesResponseOutput) MoveStatus() MoveResourcePropertiesResponseMoveStatusOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) MoveResourcePropertiesResponseMoveStatus { return v.MoveStatus }).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (o MoveResourcePropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MoveResourcePropertiesResponseOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) interface{} { return v.ResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponseOutput) SourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.SourceId }).(pulumi.StringOutput)
}

func (o MoveResourcePropertiesResponseOutput) SourceResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) interface{} { return v.SourceResourceSettings }).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponseOutput) TargetId() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponse) string { return v.TargetId }).(pulumi.StringOutput)
}

type MoveResourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponse)(nil)).Elem()
}

func (o MoveResourcePropertiesResponsePtrOutput) ToMoveResourcePropertiesResponsePtrOutput() MoveResourcePropertiesResponsePtrOutput {
	return o
}

func (o MoveResourcePropertiesResponsePtrOutput) ToMoveResourcePropertiesResponsePtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponsePtrOutput {
	return o
}

func (o MoveResourcePropertiesResponsePtrOutput) Elem() MoveResourcePropertiesResponseOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) MoveResourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponse
		return ret
	}).(MoveResourcePropertiesResponseOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) DependsOn() MoveResourceDependencyResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) []MoveResourceDependencyResponse {
		if v == nil {
			return nil
		}
		return v.DependsOn
	}).(MoveResourceDependencyResponseArrayOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) DependsOnOverrides() MoveResourceDependencyOverrideResponseArrayOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) []MoveResourceDependencyOverrideResponse {
		if v == nil {
			return nil
		}
		return v.DependsOnOverrides
	}).(MoveResourceDependencyOverrideResponseArrayOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) Errors() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *MoveResourcePropertiesResponseErrors {
		if v == nil {
			return nil
		}
		return &v.Errors
	}).(MoveResourcePropertiesResponseErrorsPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ExistingTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExistingTargetId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) IsResolveRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsResolveRequired
	}).(pulumi.BoolPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) MoveStatus() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *MoveResourcePropertiesResponseMoveStatus {
		if v == nil {
			return nil
		}
		return &v.MoveStatus
	}).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) ResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceId
	}).(pulumi.StringPtrOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) SourceResourceSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.SourceResourceSettings
	}).(pulumi.AnyOutput)
}

func (o MoveResourcePropertiesResponsePtrOutput) TargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetId
	}).(pulumi.StringPtrOutput)
}

type MoveResourcePropertiesResponseErrors struct {
	Properties *MoveResourceErrorBodyResponse `pulumi:"properties"`
}





type MoveResourcePropertiesResponseErrorsInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput
	ToMoveResourcePropertiesResponseErrorsOutputWithContext(context.Context) MoveResourcePropertiesResponseErrorsOutput
}

type MoveResourcePropertiesResponseErrorsArgs struct {
	Properties MoveResourceErrorBodyResponsePtrInput `pulumi:"properties"`
}

func (MoveResourcePropertiesResponseErrorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput {
	return i.ToMoveResourcePropertiesResponseErrorsOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsOutput)
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return i.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseErrorsArgs) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsOutput).ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponseErrorsPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput
	ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Context) MoveResourcePropertiesResponseErrorsPtrOutput
}

type moveResourcePropertiesResponseErrorsPtrType MoveResourcePropertiesResponseErrorsArgs

func MoveResourcePropertiesResponseErrorsPtr(v *MoveResourcePropertiesResponseErrorsArgs) MoveResourcePropertiesResponseErrorsPtrInput {
	return (*moveResourcePropertiesResponseErrorsPtrType)(v)
}

func (*moveResourcePropertiesResponseErrorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (i *moveResourcePropertiesResponseErrorsPtrType) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return i.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponseErrorsPtrType) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseErrorsPtrOutput)
}

type MoveResourcePropertiesResponseErrorsOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseErrorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsOutput() MoveResourcePropertiesResponseErrorsOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseErrorsOutput) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponseErrors) *MoveResourcePropertiesResponseErrors {
		return &v
	}).(MoveResourcePropertiesResponseErrorsPtrOutput)
}

func (o MoveResourcePropertiesResponseErrorsOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseErrors) *MoveResourceErrorBodyResponse { return v.Properties }).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourcePropertiesResponseErrorsPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseErrorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseErrors)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) ToMoveResourcePropertiesResponseErrorsPtrOutput() MoveResourcePropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) ToMoveResourcePropertiesResponseErrorsPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseErrorsPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) Elem() MoveResourcePropertiesResponseErrorsOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseErrors) MoveResourcePropertiesResponseErrors {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponseErrors
		return ret
	}).(MoveResourcePropertiesResponseErrorsOutput)
}

func (o MoveResourcePropertiesResponseErrorsPtrOutput) Properties() MoveResourceErrorBodyResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseErrors) *MoveResourceErrorBodyResponse {
		if v == nil {
			return nil
		}
		return v.Properties
	}).(MoveResourceErrorBodyResponsePtrOutput)
}

type MoveResourcePropertiesResponseMoveStatus struct {
	Errors    *MoveResourceErrorResponse `pulumi:"errors"`
	JobStatus *JobStatusResponse         `pulumi:"jobStatus"`
	MoveState string                     `pulumi:"moveState"`
}





type MoveResourcePropertiesResponseMoveStatusInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput
	ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(context.Context) MoveResourcePropertiesResponseMoveStatusOutput
}

type MoveResourcePropertiesResponseMoveStatusArgs struct {
	Errors    MoveResourceErrorResponsePtrInput `pulumi:"errors"`
	JobStatus JobStatusResponsePtrInput         `pulumi:"jobStatus"`
	MoveState pulumi.StringInput                `pulumi:"moveState"`
}

func (MoveResourcePropertiesResponseMoveStatusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (i MoveResourcePropertiesResponseMoveStatusArgs) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusOutput).ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx)
}









type MoveResourcePropertiesResponseMoveStatusPtrInput interface {
	pulumi.Input

	ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput
	ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput
}

type moveResourcePropertiesResponseMoveStatusPtrType MoveResourcePropertiesResponseMoveStatusArgs

func MoveResourcePropertiesResponseMoveStatusPtr(v *MoveResourcePropertiesResponseMoveStatusArgs) MoveResourcePropertiesResponseMoveStatusPtrInput {
	return (*moveResourcePropertiesResponseMoveStatusPtrType)(v)
}

func (*moveResourcePropertiesResponseMoveStatusPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (i *moveResourcePropertiesResponseMoveStatusPtrType) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return i.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (i *moveResourcePropertiesResponseMoveStatusPtrType) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
}

type MoveResourcePropertiesResponseMoveStatusOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseMoveStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusOutput() MoveResourcePropertiesResponseMoveStatusOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(context.Background())
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveResourcePropertiesResponseMoveStatus) *MoveResourcePropertiesResponseMoveStatus {
		return &v
	}).(MoveResourcePropertiesResponseMoveStatusPtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) Errors() MoveResourceErrorResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) *MoveResourceErrorResponse { return v.Errors }).(MoveResourceErrorResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) JobStatus() JobStatusResponsePtrOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) *JobStatusResponse { return v.JobStatus }).(JobStatusResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusOutput) MoveState() pulumi.StringOutput {
	return o.ApplyT(func(v MoveResourcePropertiesResponseMoveStatus) string { return v.MoveState }).(pulumi.StringOutput)
}

type MoveResourcePropertiesResponseMoveStatusPtrOutput struct{ *pulumi.OutputState }

func (MoveResourcePropertiesResponseMoveStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveResourcePropertiesResponseMoveStatus)(nil)).Elem()
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutput() MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) ToMoveResourcePropertiesResponseMoveStatusPtrOutputWithContext(ctx context.Context) MoveResourcePropertiesResponseMoveStatusPtrOutput {
	return o
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) Elem() MoveResourcePropertiesResponseMoveStatusOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) MoveResourcePropertiesResponseMoveStatus {
		if v != nil {
			return *v
		}
		var ret MoveResourcePropertiesResponseMoveStatus
		return ret
	}).(MoveResourcePropertiesResponseMoveStatusOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) Errors() MoveResourceErrorResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *MoveResourceErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(MoveResourceErrorResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) JobStatus() JobStatusResponsePtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *JobStatusResponse {
		if v == nil {
			return nil
		}
		return v.JobStatus
	}).(JobStatusResponsePtrOutput)
}

func (o MoveResourcePropertiesResponseMoveStatusPtrOutput) MoveState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MoveResourcePropertiesResponseMoveStatus) *string {
		if v == nil {
			return nil
		}
		return &v.MoveState
	}).(pulumi.StringPtrOutput)
}

type NetworkInterfaceResourceSettings struct {
	EnableAcceleratedNetworking *bool                                `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettings `pulumi:"ipConfigurations"`
	ResourceType                string                               `pulumi:"resourceType"`
	TargetResourceName          string                               `pulumi:"targetResourceName"`
}





type NetworkInterfaceResourceSettingsInput interface {
	pulumi.Input

	ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput
	ToNetworkInterfaceResourceSettingsOutputWithContext(context.Context) NetworkInterfaceResourceSettingsOutput
}

type NetworkInterfaceResourceSettingsArgs struct {
	EnableAcceleratedNetworking pulumi.BoolPtrInput                          `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            NicIpConfigurationResourceSettingsArrayInput `pulumi:"ipConfigurations"`
	ResourceType                pulumi.StringInput                           `pulumi:"resourceType"`
	TargetResourceName          pulumi.StringInput                           `pulumi:"targetResourceName"`
}

func (NetworkInterfaceResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettings)(nil)).Elem()
}

func (i NetworkInterfaceResourceSettingsArgs) ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput {
	return i.ToNetworkInterfaceResourceSettingsOutputWithContext(context.Background())
}

func (i NetworkInterfaceResourceSettingsArgs) ToNetworkInterfaceResourceSettingsOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResourceSettingsOutput)
}

type NetworkInterfaceResourceSettingsOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettings)(nil)).Elem()
}

func (o NetworkInterfaceResourceSettingsOutput) ToNetworkInterfaceResourceSettingsOutput() NetworkInterfaceResourceSettingsOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsOutput) ToNetworkInterfaceResourceSettingsOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) IpConfigurations() NicIpConfigurationResourceSettingsArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) []NicIpConfigurationResourceSettings {
		return v.IpConfigurations
	}).(NicIpConfigurationResourceSettingsArrayOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkInterfaceResourceSettingsResponse struct {
	EnableAcceleratedNetworking *bool                                        `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            []NicIpConfigurationResourceSettingsResponse `pulumi:"ipConfigurations"`
	ResourceType                string                                       `pulumi:"resourceType"`
	TargetResourceName          string                                       `pulumi:"targetResourceName"`
}





type NetworkInterfaceResourceSettingsResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput
	ToNetworkInterfaceResourceSettingsResponseOutputWithContext(context.Context) NetworkInterfaceResourceSettingsResponseOutput
}

type NetworkInterfaceResourceSettingsResponseArgs struct {
	EnableAcceleratedNetworking pulumi.BoolPtrInput                                  `pulumi:"enableAcceleratedNetworking"`
	IpConfigurations            NicIpConfigurationResourceSettingsResponseArrayInput `pulumi:"ipConfigurations"`
	ResourceType                pulumi.StringInput                                   `pulumi:"resourceType"`
	TargetResourceName          pulumi.StringInput                                   `pulumi:"targetResourceName"`
}

func (NetworkInterfaceResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettingsResponse)(nil)).Elem()
}

func (i NetworkInterfaceResourceSettingsResponseArgs) ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput {
	return i.ToNetworkInterfaceResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceResourceSettingsResponseArgs) ToNetworkInterfaceResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResourceSettingsResponseOutput)
}

type NetworkInterfaceResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResourceSettingsResponse)(nil)).Elem()
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ToNetworkInterfaceResourceSettingsResponseOutput() NetworkInterfaceResourceSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ToNetworkInterfaceResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkInterfaceResourceSettingsResponseOutput {
	return o
}

func (o NetworkInterfaceResourceSettingsResponseOutput) EnableAcceleratedNetworking() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) *bool { return v.EnableAcceleratedNetworking }).(pulumi.BoolPtrOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) IpConfigurations() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) []NicIpConfigurationResourceSettingsResponse {
		return v.IpConfigurations
	}).(NicIpConfigurationResourceSettingsResponseArrayOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResourceSettings struct {
	ResourceType       string            `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRule `pulumi:"securityRules"`
	TargetResourceName string            `pulumi:"targetResourceName"`
}





type NetworkSecurityGroupResourceSettingsInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput
	ToNetworkSecurityGroupResourceSettingsOutputWithContext(context.Context) NetworkSecurityGroupResourceSettingsOutput
}

type NetworkSecurityGroupResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput        `pulumi:"resourceType"`
	SecurityRules      NsgSecurityRuleArrayInput `pulumi:"securityRules"`
	TargetResourceName pulumi.StringInput        `pulumi:"targetResourceName"`
}

func (NetworkSecurityGroupResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettings)(nil)).Elem()
}

func (i NetworkSecurityGroupResourceSettingsArgs) ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput {
	return i.ToNetworkSecurityGroupResourceSettingsOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResourceSettingsArgs) ToNetworkSecurityGroupResourceSettingsOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResourceSettingsOutput)
}

type NetworkSecurityGroupResourceSettingsOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettings)(nil)).Elem()
}

func (o NetworkSecurityGroupResourceSettingsOutput) ToNetworkSecurityGroupResourceSettingsOutput() NetworkSecurityGroupResourceSettingsOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsOutput) ToNetworkSecurityGroupResourceSettingsOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResourceSettingsOutput) SecurityRules() NsgSecurityRuleArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) []NsgSecurityRule { return v.SecurityRules }).(NsgSecurityRuleArrayOutput)
}

func (o NetworkSecurityGroupResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NetworkSecurityGroupResourceSettingsResponse struct {
	ResourceType       string                    `pulumi:"resourceType"`
	SecurityRules      []NsgSecurityRuleResponse `pulumi:"securityRules"`
	TargetResourceName string                    `pulumi:"targetResourceName"`
}





type NetworkSecurityGroupResourceSettingsResponseInput interface {
	pulumi.Input

	ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput
	ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(context.Context) NetworkSecurityGroupResourceSettingsResponseOutput
}

type NetworkSecurityGroupResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput                `pulumi:"resourceType"`
	SecurityRules      NsgSecurityRuleResponseArrayInput `pulumi:"securityRules"`
	TargetResourceName pulumi.StringInput                `pulumi:"targetResourceName"`
}

func (NetworkSecurityGroupResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettingsResponse)(nil)).Elem()
}

func (i NetworkSecurityGroupResourceSettingsResponseArgs) ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput {
	return i.ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NetworkSecurityGroupResourceSettingsResponseArgs) ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityGroupResourceSettingsResponseOutput)
}

type NetworkSecurityGroupResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NetworkSecurityGroupResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkSecurityGroupResourceSettingsResponse)(nil)).Elem()
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ToNetworkSecurityGroupResourceSettingsResponseOutput() NetworkSecurityGroupResourceSettingsResponseOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ToNetworkSecurityGroupResourceSettingsResponseOutputWithContext(ctx context.Context) NetworkSecurityGroupResourceSettingsResponseOutput {
	return o
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) SecurityRules() NsgSecurityRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) []NsgSecurityRuleResponse { return v.SecurityRules }).(NsgSecurityRuleResponseArrayOutput)
}

func (o NetworkSecurityGroupResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkSecurityGroupResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type NicIpConfigurationResourceSettings struct {
	LoadBalancerBackendAddressPools []LoadBalancerBackendAddressPoolReference `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            []LoadBalancerNatRuleReference            `pulumi:"loadBalancerNatRules"`
	Name                            *string                                   `pulumi:"name"`
	Primary                         *bool                                     `pulumi:"primary"`
	PrivateIpAddress                *string                                   `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       *string                                   `pulumi:"privateIpAllocationMethod"`
	PublicIp                        *PublicIpReference                        `pulumi:"publicIp"`
	Subnet                          *SubnetReference                          `pulumi:"subnet"`
}





type NicIpConfigurationResourceSettingsInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput
	ToNicIpConfigurationResourceSettingsOutputWithContext(context.Context) NicIpConfigurationResourceSettingsOutput
}

type NicIpConfigurationResourceSettingsArgs struct {
	LoadBalancerBackendAddressPools LoadBalancerBackendAddressPoolReferenceArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            LoadBalancerNatRuleReferenceArrayInput            `pulumi:"loadBalancerNatRules"`
	Name                            pulumi.StringPtrInput                             `pulumi:"name"`
	Primary                         pulumi.BoolPtrInput                               `pulumi:"primary"`
	PrivateIpAddress                pulumi.StringPtrInput                             `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       pulumi.StringPtrInput                             `pulumi:"privateIpAllocationMethod"`
	PublicIp                        PublicIpReferencePtrInput                         `pulumi:"publicIp"`
	Subnet                          SubnetReferencePtrInput                           `pulumi:"subnet"`
}

func (NicIpConfigurationResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsArgs) ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput {
	return i.ToNicIpConfigurationResourceSettingsOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsArgs) ToNicIpConfigurationResourceSettingsOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsOutput)
}





type NicIpConfigurationResourceSettingsArrayInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput
	ToNicIpConfigurationResourceSettingsArrayOutputWithContext(context.Context) NicIpConfigurationResourceSettingsArrayOutput
}

type NicIpConfigurationResourceSettingsArray []NicIpConfigurationResourceSettingsInput

func (NicIpConfigurationResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsArray) ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput {
	return i.ToNicIpConfigurationResourceSettingsArrayOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsArray) ToNicIpConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsArrayOutput)
}

type NicIpConfigurationResourceSettingsOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsOutput) ToNicIpConfigurationResourceSettingsOutput() NicIpConfigurationResourceSettingsOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsOutput) ToNicIpConfigurationResourceSettingsOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsOutput) LoadBalancerBackendAddressPools() LoadBalancerBackendAddressPoolReferenceArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) []LoadBalancerBackendAddressPoolReference {
		return v.LoadBalancerBackendAddressPools
	}).(LoadBalancerBackendAddressPoolReferenceArrayOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) LoadBalancerNatRules() LoadBalancerNatRuleReferenceArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) []LoadBalancerNatRuleReference {
		return v.LoadBalancerNatRules
	}).(LoadBalancerNatRuleReferenceArrayOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) PublicIp() PublicIpReferencePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *PublicIpReference { return v.PublicIp }).(PublicIpReferencePtrOutput)
}

func (o NicIpConfigurationResourceSettingsOutput) Subnet() SubnetReferencePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettings) *SubnetReference { return v.Subnet }).(SubnetReferencePtrOutput)
}

type NicIpConfigurationResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettings)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsArrayOutput) ToNicIpConfigurationResourceSettingsArrayOutput() NicIpConfigurationResourceSettingsArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsArrayOutput) ToNicIpConfigurationResourceSettingsArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsArrayOutput) Index(i pulumi.IntInput) NicIpConfigurationResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NicIpConfigurationResourceSettings {
		return vs[0].([]NicIpConfigurationResourceSettings)[vs[1].(int)]
	}).(NicIpConfigurationResourceSettingsOutput)
}

type NicIpConfigurationResourceSettingsResponse struct {
	LoadBalancerBackendAddressPools []LoadBalancerBackendAddressPoolReferenceResponse `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            []LoadBalancerNatRuleReferenceResponse            `pulumi:"loadBalancerNatRules"`
	Name                            *string                                           `pulumi:"name"`
	Primary                         *bool                                             `pulumi:"primary"`
	PrivateIpAddress                *string                                           `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       *string                                           `pulumi:"privateIpAllocationMethod"`
	PublicIp                        *PublicIpReferenceResponse                        `pulumi:"publicIp"`
	Subnet                          *SubnetReferenceResponse                          `pulumi:"subnet"`
}





type NicIpConfigurationResourceSettingsResponseInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput
	ToNicIpConfigurationResourceSettingsResponseOutputWithContext(context.Context) NicIpConfigurationResourceSettingsResponseOutput
}

type NicIpConfigurationResourceSettingsResponseArgs struct {
	LoadBalancerBackendAddressPools LoadBalancerBackendAddressPoolReferenceResponseArrayInput `pulumi:"loadBalancerBackendAddressPools"`
	LoadBalancerNatRules            LoadBalancerNatRuleReferenceResponseArrayInput            `pulumi:"loadBalancerNatRules"`
	Name                            pulumi.StringPtrInput                                     `pulumi:"name"`
	Primary                         pulumi.BoolPtrInput                                       `pulumi:"primary"`
	PrivateIpAddress                pulumi.StringPtrInput                                     `pulumi:"privateIpAddress"`
	PrivateIpAllocationMethod       pulumi.StringPtrInput                                     `pulumi:"privateIpAllocationMethod"`
	PublicIp                        PublicIpReferenceResponsePtrInput                         `pulumi:"publicIp"`
	Subnet                          SubnetReferenceResponsePtrInput                           `pulumi:"subnet"`
}

func (NicIpConfigurationResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsResponseArgs) ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput {
	return i.ToNicIpConfigurationResourceSettingsResponseOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsResponseArgs) ToNicIpConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsResponseOutput)
}





type NicIpConfigurationResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput
	ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput
}

type NicIpConfigurationResourceSettingsResponseArray []NicIpConfigurationResourceSettingsResponseInput

func (NicIpConfigurationResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (i NicIpConfigurationResourceSettingsResponseArray) ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return i.ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i NicIpConfigurationResourceSettingsResponseArray) ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicIpConfigurationResourceSettingsResponseArrayOutput)
}

type NicIpConfigurationResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsResponseOutput) ToNicIpConfigurationResourceSettingsResponseOutput() NicIpConfigurationResourceSettingsResponseOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseOutput) ToNicIpConfigurationResourceSettingsResponseOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseOutput) LoadBalancerBackendAddressPools() LoadBalancerBackendAddressPoolReferenceResponseArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) []LoadBalancerBackendAddressPoolReferenceResponse {
		return v.LoadBalancerBackendAddressPools
	}).(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) LoadBalancerNatRules() LoadBalancerNatRuleReferenceResponseArrayOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) []LoadBalancerNatRuleReferenceResponse {
		return v.LoadBalancerNatRules
	}).(LoadBalancerNatRuleReferenceResponseArrayOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Primary() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *bool { return v.Primary }).(pulumi.BoolPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PrivateIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *string { return v.PrivateIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) PublicIp() PublicIpReferenceResponsePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *PublicIpReferenceResponse { return v.PublicIp }).(PublicIpReferenceResponsePtrOutput)
}

func (o NicIpConfigurationResourceSettingsResponseOutput) Subnet() SubnetReferenceResponsePtrOutput {
	return o.ApplyT(func(v NicIpConfigurationResourceSettingsResponse) *SubnetReferenceResponse { return v.Subnet }).(SubnetReferenceResponsePtrOutput)
}

type NicIpConfigurationResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (NicIpConfigurationResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NicIpConfigurationResourceSettingsResponse)(nil)).Elem()
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) ToNicIpConfigurationResourceSettingsResponseArrayOutput() NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) ToNicIpConfigurationResourceSettingsResponseArrayOutputWithContext(ctx context.Context) NicIpConfigurationResourceSettingsResponseArrayOutput {
	return o
}

func (o NicIpConfigurationResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) NicIpConfigurationResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NicIpConfigurationResourceSettingsResponse {
		return vs[0].([]NicIpConfigurationResourceSettingsResponse)[vs[1].(int)]
	}).(NicIpConfigurationResourceSettingsResponseOutput)
}

type NsgReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type NsgReferenceInput interface {
	pulumi.Input

	ToNsgReferenceOutput() NsgReferenceOutput
	ToNsgReferenceOutputWithContext(context.Context) NsgReferenceOutput
}

type NsgReferenceArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (NsgReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReference)(nil)).Elem()
}

func (i NsgReferenceArgs) ToNsgReferenceOutput() NsgReferenceOutput {
	return i.ToNsgReferenceOutputWithContext(context.Background())
}

func (i NsgReferenceArgs) ToNsgReferenceOutputWithContext(ctx context.Context) NsgReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceOutput)
}

func (i NsgReferenceArgs) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return i.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (i NsgReferenceArgs) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceOutput).ToNsgReferencePtrOutputWithContext(ctx)
}









type NsgReferencePtrInput interface {
	pulumi.Input

	ToNsgReferencePtrOutput() NsgReferencePtrOutput
	ToNsgReferencePtrOutputWithContext(context.Context) NsgReferencePtrOutput
}

type nsgReferencePtrType NsgReferenceArgs

func NsgReferencePtr(v *NsgReferenceArgs) NsgReferencePtrInput {
	return (*nsgReferencePtrType)(v)
}

func (*nsgReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReference)(nil)).Elem()
}

func (i *nsgReferencePtrType) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return i.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (i *nsgReferencePtrType) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferencePtrOutput)
}

type NsgReferenceOutput struct{ *pulumi.OutputState }

func (NsgReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReference)(nil)).Elem()
}

func (o NsgReferenceOutput) ToNsgReferenceOutput() NsgReferenceOutput {
	return o
}

func (o NsgReferenceOutput) ToNsgReferenceOutputWithContext(ctx context.Context) NsgReferenceOutput {
	return o
}

func (o NsgReferenceOutput) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return o.ToNsgReferencePtrOutputWithContext(context.Background())
}

func (o NsgReferenceOutput) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NsgReference) *NsgReference {
		return &v
	}).(NsgReferencePtrOutput)
}

func (o NsgReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NsgReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type NsgReferencePtrOutput struct{ *pulumi.OutputState }

func (NsgReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReference)(nil)).Elem()
}

func (o NsgReferencePtrOutput) ToNsgReferencePtrOutput() NsgReferencePtrOutput {
	return o
}

func (o NsgReferencePtrOutput) ToNsgReferencePtrOutputWithContext(ctx context.Context) NsgReferencePtrOutput {
	return o
}

func (o NsgReferencePtrOutput) Elem() NsgReferenceOutput {
	return o.ApplyT(func(v *NsgReference) NsgReference {
		if v != nil {
			return *v
		}
		var ret NsgReference
		return ret
	}).(NsgReferenceOutput)
}

func (o NsgReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsgReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type NsgReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type NsgReferenceResponseInput interface {
	pulumi.Input

	ToNsgReferenceResponseOutput() NsgReferenceResponseOutput
	ToNsgReferenceResponseOutputWithContext(context.Context) NsgReferenceResponseOutput
}

type NsgReferenceResponseArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (NsgReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReferenceResponse)(nil)).Elem()
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponseOutput() NsgReferenceResponseOutput {
	return i.ToNsgReferenceResponseOutputWithContext(context.Background())
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponseOutputWithContext(ctx context.Context) NsgReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponseOutput)
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return i.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (i NsgReferenceResponseArgs) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponseOutput).ToNsgReferenceResponsePtrOutputWithContext(ctx)
}









type NsgReferenceResponsePtrInput interface {
	pulumi.Input

	ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput
	ToNsgReferenceResponsePtrOutputWithContext(context.Context) NsgReferenceResponsePtrOutput
}

type nsgReferenceResponsePtrType NsgReferenceResponseArgs

func NsgReferenceResponsePtr(v *NsgReferenceResponseArgs) NsgReferenceResponsePtrInput {
	return (*nsgReferenceResponsePtrType)(v)
}

func (*nsgReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReferenceResponse)(nil)).Elem()
}

func (i *nsgReferenceResponsePtrType) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return i.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *nsgReferenceResponsePtrType) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgReferenceResponsePtrOutput)
}

type NsgReferenceResponseOutput struct{ *pulumi.OutputState }

func (NsgReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgReferenceResponse)(nil)).Elem()
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponseOutput() NsgReferenceResponseOutput {
	return o
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponseOutputWithContext(ctx context.Context) NsgReferenceResponseOutput {
	return o
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return o.ToNsgReferenceResponsePtrOutputWithContext(context.Background())
}

func (o NsgReferenceResponseOutput) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NsgReferenceResponse) *NsgReferenceResponse {
		return &v
	}).(NsgReferenceResponsePtrOutput)
}

func (o NsgReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v NsgReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type NsgReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (NsgReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgReferenceResponse)(nil)).Elem()
}

func (o NsgReferenceResponsePtrOutput) ToNsgReferenceResponsePtrOutput() NsgReferenceResponsePtrOutput {
	return o
}

func (o NsgReferenceResponsePtrOutput) ToNsgReferenceResponsePtrOutputWithContext(ctx context.Context) NsgReferenceResponsePtrOutput {
	return o
}

func (o NsgReferenceResponsePtrOutput) Elem() NsgReferenceResponseOutput {
	return o.ApplyT(func(v *NsgReferenceResponse) NsgReferenceResponse {
		if v != nil {
			return *v
		}
		var ret NsgReferenceResponse
		return ret
	}).(NsgReferenceResponseOutput)
}

func (o NsgReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsgReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type NsgSecurityRule struct {
	Access                   *string `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                *string `pulumi:"direction"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 *string `pulumi:"protocol"`
	SourceAddressPrefix      *string `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}





type NsgSecurityRuleInput interface {
	pulumi.Input

	ToNsgSecurityRuleOutput() NsgSecurityRuleOutput
	ToNsgSecurityRuleOutputWithContext(context.Context) NsgSecurityRuleOutput
}

type NsgSecurityRuleArgs struct {
	Access                   pulumi.StringPtrInput `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringPtrInput `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringPtrInput `pulumi:"direction"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringPtrInput `pulumi:"protocol"`
	SourceAddressPrefix      pulumi.StringPtrInput `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (NsgSecurityRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRule)(nil)).Elem()
}

func (i NsgSecurityRuleArgs) ToNsgSecurityRuleOutput() NsgSecurityRuleOutput {
	return i.ToNsgSecurityRuleOutputWithContext(context.Background())
}

func (i NsgSecurityRuleArgs) ToNsgSecurityRuleOutputWithContext(ctx context.Context) NsgSecurityRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleOutput)
}





type NsgSecurityRuleArrayInput interface {
	pulumi.Input

	ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput
	ToNsgSecurityRuleArrayOutputWithContext(context.Context) NsgSecurityRuleArrayOutput
}

type NsgSecurityRuleArray []NsgSecurityRuleInput

func (NsgSecurityRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRule)(nil)).Elem()
}

func (i NsgSecurityRuleArray) ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput {
	return i.ToNsgSecurityRuleArrayOutputWithContext(context.Background())
}

func (i NsgSecurityRuleArray) ToNsgSecurityRuleArrayOutputWithContext(ctx context.Context) NsgSecurityRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleArrayOutput)
}

type NsgSecurityRuleOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRule)(nil)).Elem()
}

func (o NsgSecurityRuleOutput) ToNsgSecurityRuleOutput() NsgSecurityRuleOutput {
	return o
}

func (o NsgSecurityRuleOutput) ToNsgSecurityRuleOutputWithContext(ctx context.Context) NsgSecurityRuleOutput {
	return o
}

func (o NsgSecurityRuleOutput) Access() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Access }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o NsgSecurityRuleOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRule) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type NsgSecurityRuleArrayOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRule)(nil)).Elem()
}

func (o NsgSecurityRuleArrayOutput) ToNsgSecurityRuleArrayOutput() NsgSecurityRuleArrayOutput {
	return o
}

func (o NsgSecurityRuleArrayOutput) ToNsgSecurityRuleArrayOutputWithContext(ctx context.Context) NsgSecurityRuleArrayOutput {
	return o
}

func (o NsgSecurityRuleArrayOutput) Index(i pulumi.IntInput) NsgSecurityRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsgSecurityRule {
		return vs[0].([]NsgSecurityRule)[vs[1].(int)]
	}).(NsgSecurityRuleOutput)
}

type NsgSecurityRuleResponse struct {
	Access                   *string `pulumi:"access"`
	Description              *string `pulumi:"description"`
	DestinationAddressPrefix *string `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     *string `pulumi:"destinationPortRange"`
	Direction                *string `pulumi:"direction"`
	Name                     *string `pulumi:"name"`
	Priority                 *int    `pulumi:"priority"`
	Protocol                 *string `pulumi:"protocol"`
	SourceAddressPrefix      *string `pulumi:"sourceAddressPrefix"`
	SourcePortRange          *string `pulumi:"sourcePortRange"`
}





type NsgSecurityRuleResponseInput interface {
	pulumi.Input

	ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput
	ToNsgSecurityRuleResponseOutputWithContext(context.Context) NsgSecurityRuleResponseOutput
}

type NsgSecurityRuleResponseArgs struct {
	Access                   pulumi.StringPtrInput `pulumi:"access"`
	Description              pulumi.StringPtrInput `pulumi:"description"`
	DestinationAddressPrefix pulumi.StringPtrInput `pulumi:"destinationAddressPrefix"`
	DestinationPortRange     pulumi.StringPtrInput `pulumi:"destinationPortRange"`
	Direction                pulumi.StringPtrInput `pulumi:"direction"`
	Name                     pulumi.StringPtrInput `pulumi:"name"`
	Priority                 pulumi.IntPtrInput    `pulumi:"priority"`
	Protocol                 pulumi.StringPtrInput `pulumi:"protocol"`
	SourceAddressPrefix      pulumi.StringPtrInput `pulumi:"sourceAddressPrefix"`
	SourcePortRange          pulumi.StringPtrInput `pulumi:"sourcePortRange"`
}

func (NsgSecurityRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRuleResponse)(nil)).Elem()
}

func (i NsgSecurityRuleResponseArgs) ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput {
	return i.ToNsgSecurityRuleResponseOutputWithContext(context.Background())
}

func (i NsgSecurityRuleResponseArgs) ToNsgSecurityRuleResponseOutputWithContext(ctx context.Context) NsgSecurityRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleResponseOutput)
}





type NsgSecurityRuleResponseArrayInput interface {
	pulumi.Input

	ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput
	ToNsgSecurityRuleResponseArrayOutputWithContext(context.Context) NsgSecurityRuleResponseArrayOutput
}

type NsgSecurityRuleResponseArray []NsgSecurityRuleResponseInput

func (NsgSecurityRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRuleResponse)(nil)).Elem()
}

func (i NsgSecurityRuleResponseArray) ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput {
	return i.ToNsgSecurityRuleResponseArrayOutputWithContext(context.Background())
}

func (i NsgSecurityRuleResponseArray) ToNsgSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NsgSecurityRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsgSecurityRuleResponseArrayOutput)
}

type NsgSecurityRuleResponseOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgSecurityRuleResponse)(nil)).Elem()
}

func (o NsgSecurityRuleResponseOutput) ToNsgSecurityRuleResponseOutput() NsgSecurityRuleResponseOutput {
	return o
}

func (o NsgSecurityRuleResponseOutput) ToNsgSecurityRuleResponseOutputWithContext(ctx context.Context) NsgSecurityRuleResponseOutput {
	return o
}

func (o NsgSecurityRuleResponseOutput) Access() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Access }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) DestinationAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.DestinationAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) DestinationPortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.DestinationPortRange }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) SourceAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.SourceAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o NsgSecurityRuleResponseOutput) SourcePortRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NsgSecurityRuleResponse) *string { return v.SourcePortRange }).(pulumi.StringPtrOutput)
}

type NsgSecurityRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NsgSecurityRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NsgSecurityRuleResponse)(nil)).Elem()
}

func (o NsgSecurityRuleResponseArrayOutput) ToNsgSecurityRuleResponseArrayOutput() NsgSecurityRuleResponseArrayOutput {
	return o
}

func (o NsgSecurityRuleResponseArrayOutput) ToNsgSecurityRuleResponseArrayOutputWithContext(ctx context.Context) NsgSecurityRuleResponseArrayOutput {
	return o
}

func (o NsgSecurityRuleResponseArrayOutput) Index(i pulumi.IntInput) NsgSecurityRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NsgSecurityRuleResponse {
		return vs[0].([]NsgSecurityRuleResponse)[vs[1].(int)]
	}).(NsgSecurityRuleResponseOutput)
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

type PublicIPAddressResourceSettings struct {
	DomainNameLabel          *string `pulumi:"domainNameLabel"`
	Fqdn                     *string `pulumi:"fqdn"`
	PublicIpAllocationMethod *string `pulumi:"publicIpAllocationMethod"`
	ResourceType             string  `pulumi:"resourceType"`
	Sku                      *string `pulumi:"sku"`
	TargetResourceName       string  `pulumi:"targetResourceName"`
	Zones                    *string `pulumi:"zones"`
}





type PublicIPAddressResourceSettingsInput interface {
	pulumi.Input

	ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput
	ToPublicIPAddressResourceSettingsOutputWithContext(context.Context) PublicIPAddressResourceSettingsOutput
}

type PublicIPAddressResourceSettingsArgs struct {
	DomainNameLabel          pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn                     pulumi.StringPtrInput `pulumi:"fqdn"`
	PublicIpAllocationMethod pulumi.StringPtrInput `pulumi:"publicIpAllocationMethod"`
	ResourceType             pulumi.StringInput    `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput `pulumi:"sku"`
	TargetResourceName       pulumi.StringInput    `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput `pulumi:"zones"`
}

func (PublicIPAddressResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettings)(nil)).Elem()
}

func (i PublicIPAddressResourceSettingsArgs) ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput {
	return i.ToPublicIPAddressResourceSettingsOutputWithContext(context.Background())
}

func (i PublicIPAddressResourceSettingsArgs) ToPublicIPAddressResourceSettingsOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResourceSettingsOutput)
}

type PublicIPAddressResourceSettingsOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettings)(nil)).Elem()
}

func (o PublicIPAddressResourceSettingsOutput) ToPublicIPAddressResourceSettingsOutput() PublicIPAddressResourceSettingsOutput {
	return o
}

func (o PublicIPAddressResourceSettingsOutput) ToPublicIPAddressResourceSettingsOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsOutput {
	return o
}

func (o PublicIPAddressResourceSettingsOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) PublicIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.PublicIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettings) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type PublicIPAddressResourceSettingsResponse struct {
	DomainNameLabel          *string `pulumi:"domainNameLabel"`
	Fqdn                     *string `pulumi:"fqdn"`
	PublicIpAllocationMethod *string `pulumi:"publicIpAllocationMethod"`
	ResourceType             string  `pulumi:"resourceType"`
	Sku                      *string `pulumi:"sku"`
	TargetResourceName       string  `pulumi:"targetResourceName"`
	Zones                    *string `pulumi:"zones"`
}





type PublicIPAddressResourceSettingsResponseInput interface {
	pulumi.Input

	ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput
	ToPublicIPAddressResourceSettingsResponseOutputWithContext(context.Context) PublicIPAddressResourceSettingsResponseOutput
}

type PublicIPAddressResourceSettingsResponseArgs struct {
	DomainNameLabel          pulumi.StringPtrInput `pulumi:"domainNameLabel"`
	Fqdn                     pulumi.StringPtrInput `pulumi:"fqdn"`
	PublicIpAllocationMethod pulumi.StringPtrInput `pulumi:"publicIpAllocationMethod"`
	ResourceType             pulumi.StringInput    `pulumi:"resourceType"`
	Sku                      pulumi.StringPtrInput `pulumi:"sku"`
	TargetResourceName       pulumi.StringInput    `pulumi:"targetResourceName"`
	Zones                    pulumi.StringPtrInput `pulumi:"zones"`
}

func (PublicIPAddressResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettingsResponse)(nil)).Elem()
}

func (i PublicIPAddressResourceSettingsResponseArgs) ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput {
	return i.ToPublicIPAddressResourceSettingsResponseOutputWithContext(context.Background())
}

func (i PublicIPAddressResourceSettingsResponseArgs) ToPublicIPAddressResourceSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIPAddressResourceSettingsResponseOutput)
}

type PublicIPAddressResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (PublicIPAddressResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIPAddressResourceSettingsResponse)(nil)).Elem()
}

func (o PublicIPAddressResourceSettingsResponseOutput) ToPublicIPAddressResourceSettingsResponseOutput() PublicIPAddressResourceSettingsResponseOutput {
	return o
}

func (o PublicIPAddressResourceSettingsResponseOutput) ToPublicIPAddressResourceSettingsResponseOutputWithContext(ctx context.Context) PublicIPAddressResourceSettingsResponseOutput {
	return o
}

func (o PublicIPAddressResourceSettingsResponseOutput) DomainNameLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.DomainNameLabel }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Fqdn }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) PublicIpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.PublicIpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o PublicIPAddressResourceSettingsResponseOutput) Zones() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PublicIPAddressResourceSettingsResponse) *string { return v.Zones }).(pulumi.StringPtrOutput)
}

type PublicIpReference struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type PublicIpReferenceInput interface {
	pulumi.Input

	ToPublicIpReferenceOutput() PublicIpReferenceOutput
	ToPublicIpReferenceOutputWithContext(context.Context) PublicIpReferenceOutput
}

type PublicIpReferenceArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (PublicIpReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReference)(nil)).Elem()
}

func (i PublicIpReferenceArgs) ToPublicIpReferenceOutput() PublicIpReferenceOutput {
	return i.ToPublicIpReferenceOutputWithContext(context.Background())
}

func (i PublicIpReferenceArgs) ToPublicIpReferenceOutputWithContext(ctx context.Context) PublicIpReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceOutput)
}

func (i PublicIpReferenceArgs) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return i.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (i PublicIpReferenceArgs) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceOutput).ToPublicIpReferencePtrOutputWithContext(ctx)
}









type PublicIpReferencePtrInput interface {
	pulumi.Input

	ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput
	ToPublicIpReferencePtrOutputWithContext(context.Context) PublicIpReferencePtrOutput
}

type publicIpReferencePtrType PublicIpReferenceArgs

func PublicIpReferencePtr(v *PublicIpReferenceArgs) PublicIpReferencePtrInput {
	return (*publicIpReferencePtrType)(v)
}

func (*publicIpReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReference)(nil)).Elem()
}

func (i *publicIpReferencePtrType) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return i.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (i *publicIpReferencePtrType) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferencePtrOutput)
}

type PublicIpReferenceOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReference)(nil)).Elem()
}

func (o PublicIpReferenceOutput) ToPublicIpReferenceOutput() PublicIpReferenceOutput {
	return o
}

func (o PublicIpReferenceOutput) ToPublicIpReferenceOutputWithContext(ctx context.Context) PublicIpReferenceOutput {
	return o
}

func (o PublicIpReferenceOutput) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return o.ToPublicIpReferencePtrOutputWithContext(context.Background())
}

func (o PublicIpReferenceOutput) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpReference) *PublicIpReference {
		return &v
	}).(PublicIpReferencePtrOutput)
}

func (o PublicIpReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIpReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type PublicIpReferencePtrOutput struct{ *pulumi.OutputState }

func (PublicIpReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReference)(nil)).Elem()
}

func (o PublicIpReferencePtrOutput) ToPublicIpReferencePtrOutput() PublicIpReferencePtrOutput {
	return o
}

func (o PublicIpReferencePtrOutput) ToPublicIpReferencePtrOutputWithContext(ctx context.Context) PublicIpReferencePtrOutput {
	return o
}

func (o PublicIpReferencePtrOutput) Elem() PublicIpReferenceOutput {
	return o.ApplyT(func(v *PublicIpReference) PublicIpReference {
		if v != nil {
			return *v
		}
		var ret PublicIpReference
		return ret
	}).(PublicIpReferenceOutput)
}

func (o PublicIpReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type PublicIpReferenceResponse struct {
	SourceArmResourceId string `pulumi:"sourceArmResourceId"`
}





type PublicIpReferenceResponseInput interface {
	pulumi.Input

	ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput
	ToPublicIpReferenceResponseOutputWithContext(context.Context) PublicIpReferenceResponseOutput
}

type PublicIpReferenceResponseArgs struct {
	SourceArmResourceId pulumi.StringInput `pulumi:"sourceArmResourceId"`
}

func (PublicIpReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReferenceResponse)(nil)).Elem()
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput {
	return i.ToPublicIpReferenceResponseOutputWithContext(context.Background())
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponseOutputWithContext(ctx context.Context) PublicIpReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponseOutput)
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return i.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (i PublicIpReferenceResponseArgs) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponseOutput).ToPublicIpReferenceResponsePtrOutputWithContext(ctx)
}









type PublicIpReferenceResponsePtrInput interface {
	pulumi.Input

	ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput
	ToPublicIpReferenceResponsePtrOutputWithContext(context.Context) PublicIpReferenceResponsePtrOutput
}

type publicIpReferenceResponsePtrType PublicIpReferenceResponseArgs

func PublicIpReferenceResponsePtr(v *PublicIpReferenceResponseArgs) PublicIpReferenceResponsePtrInput {
	return (*publicIpReferenceResponsePtrType)(v)
}

func (*publicIpReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReferenceResponse)(nil)).Elem()
}

func (i *publicIpReferenceResponsePtrType) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return i.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *publicIpReferenceResponsePtrType) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicIpReferenceResponsePtrOutput)
}

type PublicIpReferenceResponseOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicIpReferenceResponse)(nil)).Elem()
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponseOutput() PublicIpReferenceResponseOutput {
	return o
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponseOutputWithContext(ctx context.Context) PublicIpReferenceResponseOutput {
	return o
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return o.ToPublicIpReferenceResponsePtrOutputWithContext(context.Background())
}

func (o PublicIpReferenceResponseOutput) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicIpReferenceResponse) *PublicIpReferenceResponse {
		return &v
	}).(PublicIpReferenceResponsePtrOutput)
}

func (o PublicIpReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v PublicIpReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type PublicIpReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (PublicIpReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicIpReferenceResponse)(nil)).Elem()
}

func (o PublicIpReferenceResponsePtrOutput) ToPublicIpReferenceResponsePtrOutput() PublicIpReferenceResponsePtrOutput {
	return o
}

func (o PublicIpReferenceResponsePtrOutput) ToPublicIpReferenceResponsePtrOutputWithContext(ctx context.Context) PublicIpReferenceResponsePtrOutput {
	return o
}

func (o PublicIpReferenceResponsePtrOutput) Elem() PublicIpReferenceResponseOutput {
	return o.ApplyT(func(v *PublicIpReferenceResponse) PublicIpReferenceResponse {
		if v != nil {
			return *v
		}
		var ret PublicIpReferenceResponse
		return ret
	}).(PublicIpReferenceResponseOutput)
}

func (o PublicIpReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PublicIpReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceGroupResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type ResourceGroupResourceSettingsInput interface {
	pulumi.Input

	ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput
	ToResourceGroupResourceSettingsOutputWithContext(context.Context) ResourceGroupResourceSettingsOutput
}

type ResourceGroupResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (ResourceGroupResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettings)(nil)).Elem()
}

func (i ResourceGroupResourceSettingsArgs) ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput {
	return i.ToResourceGroupResourceSettingsOutputWithContext(context.Background())
}

func (i ResourceGroupResourceSettingsArgs) ToResourceGroupResourceSettingsOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupResourceSettingsOutput)
}

type ResourceGroupResourceSettingsOutput struct{ *pulumi.OutputState }

func (ResourceGroupResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettings)(nil)).Elem()
}

func (o ResourceGroupResourceSettingsOutput) ToResourceGroupResourceSettingsOutput() ResourceGroupResourceSettingsOutput {
	return o
}

func (o ResourceGroupResourceSettingsOutput) ToResourceGroupResourceSettingsOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsOutput {
	return o
}

func (o ResourceGroupResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o ResourceGroupResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type ResourceGroupResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type ResourceGroupResourceSettingsResponseInput interface {
	pulumi.Input

	ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput
	ToResourceGroupResourceSettingsResponseOutputWithContext(context.Context) ResourceGroupResourceSettingsResponseOutput
}

type ResourceGroupResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (ResourceGroupResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettingsResponse)(nil)).Elem()
}

func (i ResourceGroupResourceSettingsResponseArgs) ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput {
	return i.ToResourceGroupResourceSettingsResponseOutputWithContext(context.Background())
}

func (i ResourceGroupResourceSettingsResponseArgs) ToResourceGroupResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGroupResourceSettingsResponseOutput)
}

type ResourceGroupResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (ResourceGroupResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGroupResourceSettingsResponse)(nil)).Elem()
}

func (o ResourceGroupResourceSettingsResponseOutput) ToResourceGroupResourceSettingsResponseOutput() ResourceGroupResourceSettingsResponseOutput {
	return o
}

func (o ResourceGroupResourceSettingsResponseOutput) ToResourceGroupResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceGroupResourceSettingsResponseOutput {
	return o
}

func (o ResourceGroupResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o ResourceGroupResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceGroupResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
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

type ServersProjectSummaryResponse struct {
	AssessedCount            *int              `pulumi:"assessedCount"`
	DiscoveredCount          *int              `pulumi:"discoveredCount"`
	ExtendedSummary          map[string]string `pulumi:"extendedSummary"`
	InstanceType             string            `pulumi:"instanceType"`
	LastSummaryRefreshedTime *string           `pulumi:"lastSummaryRefreshedTime"`
	MigratedCount            *int              `pulumi:"migratedCount"`
	RefreshSummaryState      *string           `pulumi:"refreshSummaryState"`
	ReplicatingCount         *int              `pulumi:"replicatingCount"`
	TestMigratedCount        *int              `pulumi:"testMigratedCount"`
}





type ServersProjectSummaryResponseInput interface {
	pulumi.Input

	ToServersProjectSummaryResponseOutput() ServersProjectSummaryResponseOutput
	ToServersProjectSummaryResponseOutputWithContext(context.Context) ServersProjectSummaryResponseOutput
}

type ServersProjectSummaryResponseArgs struct {
	AssessedCount            pulumi.IntPtrInput    `pulumi:"assessedCount"`
	DiscoveredCount          pulumi.IntPtrInput    `pulumi:"discoveredCount"`
	ExtendedSummary          pulumi.StringMapInput `pulumi:"extendedSummary"`
	InstanceType             pulumi.StringInput    `pulumi:"instanceType"`
	LastSummaryRefreshedTime pulumi.StringPtrInput `pulumi:"lastSummaryRefreshedTime"`
	MigratedCount            pulumi.IntPtrInput    `pulumi:"migratedCount"`
	RefreshSummaryState      pulumi.StringPtrInput `pulumi:"refreshSummaryState"`
	ReplicatingCount         pulumi.IntPtrInput    `pulumi:"replicatingCount"`
	TestMigratedCount        pulumi.IntPtrInput    `pulumi:"testMigratedCount"`
}

func (ServersProjectSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServersProjectSummaryResponse)(nil)).Elem()
}

func (i ServersProjectSummaryResponseArgs) ToServersProjectSummaryResponseOutput() ServersProjectSummaryResponseOutput {
	return i.ToServersProjectSummaryResponseOutputWithContext(context.Background())
}

func (i ServersProjectSummaryResponseArgs) ToServersProjectSummaryResponseOutputWithContext(ctx context.Context) ServersProjectSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServersProjectSummaryResponseOutput)
}

type ServersProjectSummaryResponseOutput struct{ *pulumi.OutputState }

func (ServersProjectSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServersProjectSummaryResponse)(nil)).Elem()
}

func (o ServersProjectSummaryResponseOutput) ToServersProjectSummaryResponseOutput() ServersProjectSummaryResponseOutput {
	return o
}

func (o ServersProjectSummaryResponseOutput) ToServersProjectSummaryResponseOutputWithContext(ctx context.Context) ServersProjectSummaryResponseOutput {
	return o
}

func (o ServersProjectSummaryResponseOutput) AssessedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *int { return v.AssessedCount }).(pulumi.IntPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) DiscoveredCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *int { return v.DiscoveredCount }).(pulumi.IntPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) ExtendedSummary() pulumi.StringMapOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) map[string]string { return v.ExtendedSummary }).(pulumi.StringMapOutput)
}

func (o ServersProjectSummaryResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o ServersProjectSummaryResponseOutput) LastSummaryRefreshedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *string { return v.LastSummaryRefreshedTime }).(pulumi.StringPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) MigratedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *int { return v.MigratedCount }).(pulumi.IntPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) RefreshSummaryState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *string { return v.RefreshSummaryState }).(pulumi.StringPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) ReplicatingCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *int { return v.ReplicatingCount }).(pulumi.IntPtrOutput)
}

func (o ServersProjectSummaryResponseOutput) TestMigratedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersProjectSummaryResponse) *int { return v.TestMigratedCount }).(pulumi.IntPtrOutput)
}

type ServersSolutionSummaryResponse struct {
	AssessedCount     *int   `pulumi:"assessedCount"`
	DiscoveredCount   *int   `pulumi:"discoveredCount"`
	InstanceType      string `pulumi:"instanceType"`
	MigratedCount     *int   `pulumi:"migratedCount"`
	ReplicatingCount  *int   `pulumi:"replicatingCount"`
	TestMigratedCount *int   `pulumi:"testMigratedCount"`
}





type ServersSolutionSummaryResponseInput interface {
	pulumi.Input

	ToServersSolutionSummaryResponseOutput() ServersSolutionSummaryResponseOutput
	ToServersSolutionSummaryResponseOutputWithContext(context.Context) ServersSolutionSummaryResponseOutput
}

type ServersSolutionSummaryResponseArgs struct {
	AssessedCount     pulumi.IntPtrInput `pulumi:"assessedCount"`
	DiscoveredCount   pulumi.IntPtrInput `pulumi:"discoveredCount"`
	InstanceType      pulumi.StringInput `pulumi:"instanceType"`
	MigratedCount     pulumi.IntPtrInput `pulumi:"migratedCount"`
	ReplicatingCount  pulumi.IntPtrInput `pulumi:"replicatingCount"`
	TestMigratedCount pulumi.IntPtrInput `pulumi:"testMigratedCount"`
}

func (ServersSolutionSummaryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServersSolutionSummaryResponse)(nil)).Elem()
}

func (i ServersSolutionSummaryResponseArgs) ToServersSolutionSummaryResponseOutput() ServersSolutionSummaryResponseOutput {
	return i.ToServersSolutionSummaryResponseOutputWithContext(context.Background())
}

func (i ServersSolutionSummaryResponseArgs) ToServersSolutionSummaryResponseOutputWithContext(ctx context.Context) ServersSolutionSummaryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServersSolutionSummaryResponseOutput)
}

type ServersSolutionSummaryResponseOutput struct{ *pulumi.OutputState }

func (ServersSolutionSummaryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServersSolutionSummaryResponse)(nil)).Elem()
}

func (o ServersSolutionSummaryResponseOutput) ToServersSolutionSummaryResponseOutput() ServersSolutionSummaryResponseOutput {
	return o
}

func (o ServersSolutionSummaryResponseOutput) ToServersSolutionSummaryResponseOutputWithContext(ctx context.Context) ServersSolutionSummaryResponseOutput {
	return o
}

func (o ServersSolutionSummaryResponseOutput) AssessedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) *int { return v.AssessedCount }).(pulumi.IntPtrOutput)
}

func (o ServersSolutionSummaryResponseOutput) DiscoveredCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) *int { return v.DiscoveredCount }).(pulumi.IntPtrOutput)
}

func (o ServersSolutionSummaryResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o ServersSolutionSummaryResponseOutput) MigratedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) *int { return v.MigratedCount }).(pulumi.IntPtrOutput)
}

func (o ServersSolutionSummaryResponseOutput) ReplicatingCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) *int { return v.ReplicatingCount }).(pulumi.IntPtrOutput)
}

func (o ServersSolutionSummaryResponseOutput) TestMigratedCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServersSolutionSummaryResponse) *int { return v.TestMigratedCount }).(pulumi.IntPtrOutput)
}

type SolutionDetails struct {
	AssessmentCount *int              `pulumi:"assessmentCount"`
	ExtendedDetails map[string]string `pulumi:"extendedDetails"`
	GroupCount      *int              `pulumi:"groupCount"`
}





type SolutionDetailsInput interface {
	pulumi.Input

	ToSolutionDetailsOutput() SolutionDetailsOutput
	ToSolutionDetailsOutputWithContext(context.Context) SolutionDetailsOutput
}

type SolutionDetailsArgs struct {
	AssessmentCount pulumi.IntPtrInput    `pulumi:"assessmentCount"`
	ExtendedDetails pulumi.StringMapInput `pulumi:"extendedDetails"`
	GroupCount      pulumi.IntPtrInput    `pulumi:"groupCount"`
}

func (SolutionDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetails)(nil)).Elem()
}

func (i SolutionDetailsArgs) ToSolutionDetailsOutput() SolutionDetailsOutput {
	return i.ToSolutionDetailsOutputWithContext(context.Background())
}

func (i SolutionDetailsArgs) ToSolutionDetailsOutputWithContext(ctx context.Context) SolutionDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsOutput)
}

func (i SolutionDetailsArgs) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return i.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (i SolutionDetailsArgs) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsOutput).ToSolutionDetailsPtrOutputWithContext(ctx)
}









type SolutionDetailsPtrInput interface {
	pulumi.Input

	ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput
	ToSolutionDetailsPtrOutputWithContext(context.Context) SolutionDetailsPtrOutput
}

type solutionDetailsPtrType SolutionDetailsArgs

func SolutionDetailsPtr(v *SolutionDetailsArgs) SolutionDetailsPtrInput {
	return (*solutionDetailsPtrType)(v)
}

func (*solutionDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetails)(nil)).Elem()
}

func (i *solutionDetailsPtrType) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return i.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (i *solutionDetailsPtrType) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsPtrOutput)
}

type SolutionDetailsOutput struct{ *pulumi.OutputState }

func (SolutionDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetails)(nil)).Elem()
}

func (o SolutionDetailsOutput) ToSolutionDetailsOutput() SolutionDetailsOutput {
	return o
}

func (o SolutionDetailsOutput) ToSolutionDetailsOutputWithContext(ctx context.Context) SolutionDetailsOutput {
	return o
}

func (o SolutionDetailsOutput) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return o.ToSolutionDetailsPtrOutputWithContext(context.Background())
}

func (o SolutionDetailsOutput) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionDetails) *SolutionDetails {
		return &v
	}).(SolutionDetailsPtrOutput)
}

func (o SolutionDetailsOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetails) *int { return v.AssessmentCount }).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v SolutionDetails) map[string]string { return v.ExtendedDetails }).(pulumi.StringMapOutput)
}

func (o SolutionDetailsOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetails) *int { return v.GroupCount }).(pulumi.IntPtrOutput)
}

type SolutionDetailsPtrOutput struct{ *pulumi.OutputState }

func (SolutionDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetails)(nil)).Elem()
}

func (o SolutionDetailsPtrOutput) ToSolutionDetailsPtrOutput() SolutionDetailsPtrOutput {
	return o
}

func (o SolutionDetailsPtrOutput) ToSolutionDetailsPtrOutputWithContext(ctx context.Context) SolutionDetailsPtrOutput {
	return o
}

func (o SolutionDetailsPtrOutput) Elem() SolutionDetailsOutput {
	return o.ApplyT(func(v *SolutionDetails) SolutionDetails {
		if v != nil {
			return *v
		}
		var ret SolutionDetails
		return ret
	}).(SolutionDetailsOutput)
}

func (o SolutionDetailsPtrOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetails) *int {
		if v == nil {
			return nil
		}
		return v.AssessmentCount
	}).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsPtrOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SolutionDetails) map[string]string {
		if v == nil {
			return nil
		}
		return v.ExtendedDetails
	}).(pulumi.StringMapOutput)
}

func (o SolutionDetailsPtrOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetails) *int {
		if v == nil {
			return nil
		}
		return v.GroupCount
	}).(pulumi.IntPtrOutput)
}

type SolutionDetailsResponse struct {
	AssessmentCount *int              `pulumi:"assessmentCount"`
	ExtendedDetails map[string]string `pulumi:"extendedDetails"`
	GroupCount      *int              `pulumi:"groupCount"`
}





type SolutionDetailsResponseInput interface {
	pulumi.Input

	ToSolutionDetailsResponseOutput() SolutionDetailsResponseOutput
	ToSolutionDetailsResponseOutputWithContext(context.Context) SolutionDetailsResponseOutput
}

type SolutionDetailsResponseArgs struct {
	AssessmentCount pulumi.IntPtrInput    `pulumi:"assessmentCount"`
	ExtendedDetails pulumi.StringMapInput `pulumi:"extendedDetails"`
	GroupCount      pulumi.IntPtrInput    `pulumi:"groupCount"`
}

func (SolutionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetailsResponse)(nil)).Elem()
}

func (i SolutionDetailsResponseArgs) ToSolutionDetailsResponseOutput() SolutionDetailsResponseOutput {
	return i.ToSolutionDetailsResponseOutputWithContext(context.Background())
}

func (i SolutionDetailsResponseArgs) ToSolutionDetailsResponseOutputWithContext(ctx context.Context) SolutionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsResponseOutput)
}

func (i SolutionDetailsResponseArgs) ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput {
	return i.ToSolutionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i SolutionDetailsResponseArgs) ToSolutionDetailsResponsePtrOutputWithContext(ctx context.Context) SolutionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsResponseOutput).ToSolutionDetailsResponsePtrOutputWithContext(ctx)
}









type SolutionDetailsResponsePtrInput interface {
	pulumi.Input

	ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput
	ToSolutionDetailsResponsePtrOutputWithContext(context.Context) SolutionDetailsResponsePtrOutput
}

type solutionDetailsResponsePtrType SolutionDetailsResponseArgs

func SolutionDetailsResponsePtr(v *SolutionDetailsResponseArgs) SolutionDetailsResponsePtrInput {
	return (*solutionDetailsResponsePtrType)(v)
}

func (*solutionDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetailsResponse)(nil)).Elem()
}

func (i *solutionDetailsResponsePtrType) ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput {
	return i.ToSolutionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *solutionDetailsResponsePtrType) ToSolutionDetailsResponsePtrOutputWithContext(ctx context.Context) SolutionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionDetailsResponsePtrOutput)
}

type SolutionDetailsResponseOutput struct{ *pulumi.OutputState }

func (SolutionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionDetailsResponse)(nil)).Elem()
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponseOutput() SolutionDetailsResponseOutput {
	return o
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponseOutputWithContext(ctx context.Context) SolutionDetailsResponseOutput {
	return o
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput {
	return o.ToSolutionDetailsResponsePtrOutputWithContext(context.Background())
}

func (o SolutionDetailsResponseOutput) ToSolutionDetailsResponsePtrOutputWithContext(ctx context.Context) SolutionDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionDetailsResponse) *SolutionDetailsResponse {
		return &v
	}).(SolutionDetailsResponsePtrOutput)
}

func (o SolutionDetailsResponseOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) *int { return v.AssessmentCount }).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsResponseOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) map[string]string { return v.ExtendedDetails }).(pulumi.StringMapOutput)
}

func (o SolutionDetailsResponseOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SolutionDetailsResponse) *int { return v.GroupCount }).(pulumi.IntPtrOutput)
}

type SolutionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (SolutionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionDetailsResponse)(nil)).Elem()
}

func (o SolutionDetailsResponsePtrOutput) ToSolutionDetailsResponsePtrOutput() SolutionDetailsResponsePtrOutput {
	return o
}

func (o SolutionDetailsResponsePtrOutput) ToSolutionDetailsResponsePtrOutputWithContext(ctx context.Context) SolutionDetailsResponsePtrOutput {
	return o
}

func (o SolutionDetailsResponsePtrOutput) Elem() SolutionDetailsResponseOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) SolutionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret SolutionDetailsResponse
		return ret
	}).(SolutionDetailsResponseOutput)
}

func (o SolutionDetailsResponsePtrOutput) AssessmentCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.AssessmentCount
	}).(pulumi.IntPtrOutput)
}

func (o SolutionDetailsResponsePtrOutput) ExtendedDetails() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.ExtendedDetails
	}).(pulumi.StringMapOutput)
}

func (o SolutionDetailsResponsePtrOutput) GroupCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SolutionDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return v.GroupCount
	}).(pulumi.IntPtrOutput)
}

type SolutionProperties struct {
	CleanupState *string          `pulumi:"cleanupState"`
	Details      *SolutionDetails `pulumi:"details"`
	Goal         *string          `pulumi:"goal"`
	Purpose      *string          `pulumi:"purpose"`
	Status       *string          `pulumi:"status"`
	Tool         *string          `pulumi:"tool"`
}





type SolutionPropertiesInput interface {
	pulumi.Input

	ToSolutionPropertiesOutput() SolutionPropertiesOutput
	ToSolutionPropertiesOutputWithContext(context.Context) SolutionPropertiesOutput
}

type SolutionPropertiesArgs struct {
	CleanupState pulumi.StringPtrInput   `pulumi:"cleanupState"`
	Details      SolutionDetailsPtrInput `pulumi:"details"`
	Goal         pulumi.StringPtrInput   `pulumi:"goal"`
	Purpose      pulumi.StringPtrInput   `pulumi:"purpose"`
	Status       pulumi.StringPtrInput   `pulumi:"status"`
	Tool         pulumi.StringPtrInput   `pulumi:"tool"`
}

func (SolutionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return i.ToSolutionPropertiesOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput)
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i SolutionPropertiesArgs) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesOutput).ToSolutionPropertiesPtrOutputWithContext(ctx)
}









type SolutionPropertiesPtrInput interface {
	pulumi.Input

	ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput
	ToSolutionPropertiesPtrOutputWithContext(context.Context) SolutionPropertiesPtrOutput
}

type solutionPropertiesPtrType SolutionPropertiesArgs

func SolutionPropertiesPtr(v *SolutionPropertiesArgs) SolutionPropertiesPtrInput {
	return (*solutionPropertiesPtrType)(v)
}

func (*solutionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return i.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (i *solutionPropertiesPtrType) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesPtrOutput)
}

type SolutionPropertiesOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutput() SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesOutputWithContext(ctx context.Context) SolutionPropertiesOutput {
	return o
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o.ToSolutionPropertiesPtrOutputWithContext(context.Background())
}

func (o SolutionPropertiesOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionProperties) *SolutionProperties {
		return &v
	}).(SolutionPropertiesPtrOutput)
}

func (o SolutionPropertiesOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.CleanupState }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Details() SolutionDetailsPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *SolutionDetails { return v.Details }).(SolutionDetailsPtrOutput)
}

func (o SolutionPropertiesOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Goal }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionProperties) *string { return v.Tool }).(pulumi.StringPtrOutput)
}

type SolutionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionProperties)(nil)).Elem()
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutput() SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) ToSolutionPropertiesPtrOutputWithContext(ctx context.Context) SolutionPropertiesPtrOutput {
	return o
}

func (o SolutionPropertiesPtrOutput) Elem() SolutionPropertiesOutput {
	return o.ApplyT(func(v *SolutionProperties) SolutionProperties {
		if v != nil {
			return *v
		}
		var ret SolutionProperties
		return ret
	}).(SolutionPropertiesOutput)
}

func (o SolutionPropertiesPtrOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.CleanupState
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Details() SolutionDetailsPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *SolutionDetails {
		if v == nil {
			return nil
		}
		return v.Details
	}).(SolutionDetailsPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Goal
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Purpose
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesPtrOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionProperties) *string {
		if v == nil {
			return nil
		}
		return v.Tool
	}).(pulumi.StringPtrOutput)
}

type SolutionPropertiesResponse struct {
	CleanupState *string                  `pulumi:"cleanupState"`
	Details      *SolutionDetailsResponse `pulumi:"details"`
	Goal         *string                  `pulumi:"goal"`
	Purpose      *string                  `pulumi:"purpose"`
	Status       *string                  `pulumi:"status"`
	Summary      interface{}              `pulumi:"summary"`
	Tool         *string                  `pulumi:"tool"`
}





type SolutionPropertiesResponseInput interface {
	pulumi.Input

	ToSolutionPropertiesResponseOutput() SolutionPropertiesResponseOutput
	ToSolutionPropertiesResponseOutputWithContext(context.Context) SolutionPropertiesResponseOutput
}

type SolutionPropertiesResponseArgs struct {
	CleanupState pulumi.StringPtrInput           `pulumi:"cleanupState"`
	Details      SolutionDetailsResponsePtrInput `pulumi:"details"`
	Goal         pulumi.StringPtrInput           `pulumi:"goal"`
	Purpose      pulumi.StringPtrInput           `pulumi:"purpose"`
	Status       pulumi.StringPtrInput           `pulumi:"status"`
	Summary      pulumi.Input                    `pulumi:"summary"`
	Tool         pulumi.StringPtrInput           `pulumi:"tool"`
}

func (SolutionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPropertiesResponse)(nil)).Elem()
}

func (i SolutionPropertiesResponseArgs) ToSolutionPropertiesResponseOutput() SolutionPropertiesResponseOutput {
	return i.ToSolutionPropertiesResponseOutputWithContext(context.Background())
}

func (i SolutionPropertiesResponseArgs) ToSolutionPropertiesResponseOutputWithContext(ctx context.Context) SolutionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesResponseOutput)
}

func (i SolutionPropertiesResponseArgs) ToSolutionPropertiesResponsePtrOutput() SolutionPropertiesResponsePtrOutput {
	return i.ToSolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SolutionPropertiesResponseArgs) ToSolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) SolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesResponseOutput).ToSolutionPropertiesResponsePtrOutputWithContext(ctx)
}









type SolutionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSolutionPropertiesResponsePtrOutput() SolutionPropertiesResponsePtrOutput
	ToSolutionPropertiesResponsePtrOutputWithContext(context.Context) SolutionPropertiesResponsePtrOutput
}

type solutionPropertiesResponsePtrType SolutionPropertiesResponseArgs

func SolutionPropertiesResponsePtr(v *SolutionPropertiesResponseArgs) SolutionPropertiesResponsePtrInput {
	return (*solutionPropertiesResponsePtrType)(v)
}

func (*solutionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionPropertiesResponse)(nil)).Elem()
}

func (i *solutionPropertiesResponsePtrType) ToSolutionPropertiesResponsePtrOutput() SolutionPropertiesResponsePtrOutput {
	return i.ToSolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *solutionPropertiesResponsePtrType) ToSolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) SolutionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SolutionPropertiesResponsePtrOutput)
}

type SolutionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SolutionPropertiesResponse)(nil)).Elem()
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutput() SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponseOutputWithContext(ctx context.Context) SolutionPropertiesResponseOutput {
	return o
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponsePtrOutput() SolutionPropertiesResponsePtrOutput {
	return o.ToSolutionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SolutionPropertiesResponseOutput) ToSolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) SolutionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SolutionPropertiesResponse) *SolutionPropertiesResponse {
		return &v
	}).(SolutionPropertiesResponsePtrOutput)
}

func (o SolutionPropertiesResponseOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.CleanupState }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Details() SolutionDetailsResponsePtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *SolutionDetailsResponse { return v.Details }).(SolutionDetailsResponsePtrOutput)
}

func (o SolutionPropertiesResponseOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Goal }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Purpose }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponseOutput) Summary() pulumi.AnyOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) interface{} { return v.Summary }).(pulumi.AnyOutput)
}

func (o SolutionPropertiesResponseOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SolutionPropertiesResponse) *string { return v.Tool }).(pulumi.StringPtrOutput)
}

type SolutionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SolutionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SolutionPropertiesResponse)(nil)).Elem()
}

func (o SolutionPropertiesResponsePtrOutput) ToSolutionPropertiesResponsePtrOutput() SolutionPropertiesResponsePtrOutput {
	return o
}

func (o SolutionPropertiesResponsePtrOutput) ToSolutionPropertiesResponsePtrOutputWithContext(ctx context.Context) SolutionPropertiesResponsePtrOutput {
	return o
}

func (o SolutionPropertiesResponsePtrOutput) Elem() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) SolutionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SolutionPropertiesResponse
		return ret
	}).(SolutionPropertiesResponseOutput)
}

func (o SolutionPropertiesResponsePtrOutput) CleanupState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CleanupState
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Details() SolutionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *SolutionDetailsResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(SolutionDetailsResponsePtrOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Goal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Goal
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Purpose() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Purpose
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Summary() pulumi.AnyOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Summary
	}).(pulumi.AnyOutput)
}

func (o SolutionPropertiesResponsePtrOutput) Tool() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SolutionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tool
	}).(pulumi.StringPtrOutput)
}

type SqlDatabaseResourceSettings struct {
	ResourceType       string  `pulumi:"resourceType"`
	TargetResourceName string  `pulumi:"targetResourceName"`
	ZoneRedundant      *string `pulumi:"zoneRedundant"`
}





type SqlDatabaseResourceSettingsInput interface {
	pulumi.Input

	ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput
	ToSqlDatabaseResourceSettingsOutputWithContext(context.Context) SqlDatabaseResourceSettingsOutput
}

type SqlDatabaseResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlDatabaseResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettings)(nil)).Elem()
}

func (i SqlDatabaseResourceSettingsArgs) ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput {
	return i.ToSqlDatabaseResourceSettingsOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceSettingsArgs) ToSqlDatabaseResourceSettingsOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceSettingsOutput)
}

type SqlDatabaseResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettings)(nil)).Elem()
}

func (o SqlDatabaseResourceSettingsOutput) ToSqlDatabaseResourceSettingsOutput() SqlDatabaseResourceSettingsOutput {
	return o
}

func (o SqlDatabaseResourceSettingsOutput) ToSqlDatabaseResourceSettingsOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsOutput {
	return o
}

func (o SqlDatabaseResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettings) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlDatabaseResourceSettingsResponse struct {
	ResourceType       string  `pulumi:"resourceType"`
	TargetResourceName string  `pulumi:"targetResourceName"`
	ZoneRedundant      *string `pulumi:"zoneRedundant"`
}





type SqlDatabaseResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput
	ToSqlDatabaseResourceSettingsResponseOutputWithContext(context.Context) SqlDatabaseResourceSettingsResponseOutput
}

type SqlDatabaseResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlDatabaseResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettingsResponse)(nil)).Elem()
}

func (i SqlDatabaseResourceSettingsResponseArgs) ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput {
	return i.ToSqlDatabaseResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceSettingsResponseArgs) ToSqlDatabaseResourceSettingsResponseOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceSettingsResponseOutput)
}

type SqlDatabaseResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResourceSettingsResponse)(nil)).Elem()
}

func (o SqlDatabaseResourceSettingsResponseOutput) ToSqlDatabaseResourceSettingsResponseOutput() SqlDatabaseResourceSettingsResponseOutput {
	return o
}

func (o SqlDatabaseResourceSettingsResponseOutput) ToSqlDatabaseResourceSettingsResponseOutputWithContext(ctx context.Context) SqlDatabaseResourceSettingsResponseOutput {
	return o
}

func (o SqlDatabaseResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlDatabaseResourceSettingsResponseOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseResourceSettingsResponse) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlElasticPoolResourceSettings struct {
	ResourceType       string  `pulumi:"resourceType"`
	TargetResourceName string  `pulumi:"targetResourceName"`
	ZoneRedundant      *string `pulumi:"zoneRedundant"`
}





type SqlElasticPoolResourceSettingsInput interface {
	pulumi.Input

	ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput
	ToSqlElasticPoolResourceSettingsOutputWithContext(context.Context) SqlElasticPoolResourceSettingsOutput
}

type SqlElasticPoolResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlElasticPoolResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettings)(nil)).Elem()
}

func (i SqlElasticPoolResourceSettingsArgs) ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput {
	return i.ToSqlElasticPoolResourceSettingsOutputWithContext(context.Background())
}

func (i SqlElasticPoolResourceSettingsArgs) ToSqlElasticPoolResourceSettingsOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlElasticPoolResourceSettingsOutput)
}

type SqlElasticPoolResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlElasticPoolResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettings)(nil)).Elem()
}

func (o SqlElasticPoolResourceSettingsOutput) ToSqlElasticPoolResourceSettingsOutput() SqlElasticPoolResourceSettingsOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsOutput) ToSqlElasticPoolResourceSettingsOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettings) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlElasticPoolResourceSettingsResponse struct {
	ResourceType       string  `pulumi:"resourceType"`
	TargetResourceName string  `pulumi:"targetResourceName"`
	ZoneRedundant      *string `pulumi:"zoneRedundant"`
}





type SqlElasticPoolResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput
	ToSqlElasticPoolResourceSettingsResponseOutputWithContext(context.Context) SqlElasticPoolResourceSettingsResponseOutput
}

type SqlElasticPoolResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput    `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput    `pulumi:"targetResourceName"`
	ZoneRedundant      pulumi.StringPtrInput `pulumi:"zoneRedundant"`
}

func (SqlElasticPoolResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettingsResponse)(nil)).Elem()
}

func (i SqlElasticPoolResourceSettingsResponseArgs) ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput {
	return i.ToSqlElasticPoolResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlElasticPoolResourceSettingsResponseArgs) ToSqlElasticPoolResourceSettingsResponseOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlElasticPoolResourceSettingsResponseOutput)
}

type SqlElasticPoolResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlElasticPoolResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlElasticPoolResourceSettingsResponse)(nil)).Elem()
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ToSqlElasticPoolResourceSettingsResponseOutput() SqlElasticPoolResourceSettingsResponseOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ToSqlElasticPoolResourceSettingsResponseOutputWithContext(ctx context.Context) SqlElasticPoolResourceSettingsResponseOutput {
	return o
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o SqlElasticPoolResourceSettingsResponseOutput) ZoneRedundant() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlElasticPoolResourceSettingsResponse) *string { return v.ZoneRedundant }).(pulumi.StringPtrOutput)
}

type SqlServerResourceSettings struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type SqlServerResourceSettingsInput interface {
	pulumi.Input

	ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput
	ToSqlServerResourceSettingsOutputWithContext(context.Context) SqlServerResourceSettingsOutput
}

type SqlServerResourceSettingsArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (SqlServerResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettings)(nil)).Elem()
}

func (i SqlServerResourceSettingsArgs) ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput {
	return i.ToSqlServerResourceSettingsOutputWithContext(context.Background())
}

func (i SqlServerResourceSettingsArgs) ToSqlServerResourceSettingsOutputWithContext(ctx context.Context) SqlServerResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerResourceSettingsOutput)
}

type SqlServerResourceSettingsOutput struct{ *pulumi.OutputState }

func (SqlServerResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettings)(nil)).Elem()
}

func (o SqlServerResourceSettingsOutput) ToSqlServerResourceSettingsOutput() SqlServerResourceSettingsOutput {
	return o
}

func (o SqlServerResourceSettingsOutput) ToSqlServerResourceSettingsOutputWithContext(ctx context.Context) SqlServerResourceSettingsOutput {
	return o
}

func (o SqlServerResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlServerResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type SqlServerResourceSettingsResponse struct {
	ResourceType       string `pulumi:"resourceType"`
	TargetResourceName string `pulumi:"targetResourceName"`
}





type SqlServerResourceSettingsResponseInput interface {
	pulumi.Input

	ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput
	ToSqlServerResourceSettingsResponseOutputWithContext(context.Context) SqlServerResourceSettingsResponseOutput
}

type SqlServerResourceSettingsResponseArgs struct {
	ResourceType       pulumi.StringInput `pulumi:"resourceType"`
	TargetResourceName pulumi.StringInput `pulumi:"targetResourceName"`
}

func (SqlServerResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettingsResponse)(nil)).Elem()
}

func (i SqlServerResourceSettingsResponseArgs) ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput {
	return i.ToSqlServerResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SqlServerResourceSettingsResponseArgs) ToSqlServerResourceSettingsResponseOutputWithContext(ctx context.Context) SqlServerResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlServerResourceSettingsResponseOutput)
}

type SqlServerResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlServerResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlServerResourceSettingsResponse)(nil)).Elem()
}

func (o SqlServerResourceSettingsResponseOutput) ToSqlServerResourceSettingsResponseOutput() SqlServerResourceSettingsResponseOutput {
	return o
}

func (o SqlServerResourceSettingsResponseOutput) ToSqlServerResourceSettingsResponseOutputWithContext(ctx context.Context) SqlServerResourceSettingsResponseOutput {
	return o
}

func (o SqlServerResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SqlServerResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v SqlServerResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type SubnetReference struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type SubnetReferenceInput interface {
	pulumi.Input

	ToSubnetReferenceOutput() SubnetReferenceOutput
	ToSubnetReferenceOutputWithContext(context.Context) SubnetReferenceOutput
}

type SubnetReferenceArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (SubnetReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReference)(nil)).Elem()
}

func (i SubnetReferenceArgs) ToSubnetReferenceOutput() SubnetReferenceOutput {
	return i.ToSubnetReferenceOutputWithContext(context.Background())
}

func (i SubnetReferenceArgs) ToSubnetReferenceOutputWithContext(ctx context.Context) SubnetReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceOutput)
}

func (i SubnetReferenceArgs) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return i.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (i SubnetReferenceArgs) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceOutput).ToSubnetReferencePtrOutputWithContext(ctx)
}









type SubnetReferencePtrInput interface {
	pulumi.Input

	ToSubnetReferencePtrOutput() SubnetReferencePtrOutput
	ToSubnetReferencePtrOutputWithContext(context.Context) SubnetReferencePtrOutput
}

type subnetReferencePtrType SubnetReferenceArgs

func SubnetReferencePtr(v *SubnetReferenceArgs) SubnetReferencePtrInput {
	return (*subnetReferencePtrType)(v)
}

func (*subnetReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReference)(nil)).Elem()
}

func (i *subnetReferencePtrType) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return i.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (i *subnetReferencePtrType) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferencePtrOutput)
}

type SubnetReferenceOutput struct{ *pulumi.OutputState }

func (SubnetReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReference)(nil)).Elem()
}

func (o SubnetReferenceOutput) ToSubnetReferenceOutput() SubnetReferenceOutput {
	return o
}

func (o SubnetReferenceOutput) ToSubnetReferenceOutputWithContext(ctx context.Context) SubnetReferenceOutput {
	return o
}

func (o SubnetReferenceOutput) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return o.ToSubnetReferencePtrOutputWithContext(context.Background())
}

func (o SubnetReferenceOutput) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetReference) *SubnetReference {
		return &v
	}).(SubnetReferencePtrOutput)
}

func (o SubnetReferenceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetReference) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetReference) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type SubnetReferencePtrOutput struct{ *pulumi.OutputState }

func (SubnetReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReference)(nil)).Elem()
}

func (o SubnetReferencePtrOutput) ToSubnetReferencePtrOutput() SubnetReferencePtrOutput {
	return o
}

func (o SubnetReferencePtrOutput) ToSubnetReferencePtrOutputWithContext(ctx context.Context) SubnetReferencePtrOutput {
	return o
}

func (o SubnetReferencePtrOutput) Elem() SubnetReferenceOutput {
	return o.ApplyT(func(v *SubnetReference) SubnetReference {
		if v != nil {
			return *v
		}
		var ret SubnetReference
		return ret
	}).(SubnetReferenceOutput)
}

func (o SubnetReferencePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReference) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetReferencePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReference) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type SubnetReferenceResponse struct {
	Name                *string `pulumi:"name"`
	SourceArmResourceId string  `pulumi:"sourceArmResourceId"`
}





type SubnetReferenceResponseInput interface {
	pulumi.Input

	ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput
	ToSubnetReferenceResponseOutputWithContext(context.Context) SubnetReferenceResponseOutput
}

type SubnetReferenceResponseArgs struct {
	Name                pulumi.StringPtrInput `pulumi:"name"`
	SourceArmResourceId pulumi.StringInput    `pulumi:"sourceArmResourceId"`
}

func (SubnetReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReferenceResponse)(nil)).Elem()
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput {
	return i.ToSubnetReferenceResponseOutputWithContext(context.Background())
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponseOutputWithContext(ctx context.Context) SubnetReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponseOutput)
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return i.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (i SubnetReferenceResponseArgs) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponseOutput).ToSubnetReferenceResponsePtrOutputWithContext(ctx)
}









type SubnetReferenceResponsePtrInput interface {
	pulumi.Input

	ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput
	ToSubnetReferenceResponsePtrOutputWithContext(context.Context) SubnetReferenceResponsePtrOutput
}

type subnetReferenceResponsePtrType SubnetReferenceResponseArgs

func SubnetReferenceResponsePtr(v *SubnetReferenceResponseArgs) SubnetReferenceResponsePtrInput {
	return (*subnetReferenceResponsePtrType)(v)
}

func (*subnetReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReferenceResponse)(nil)).Elem()
}

func (i *subnetReferenceResponsePtrType) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return i.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *subnetReferenceResponsePtrType) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetReferenceResponsePtrOutput)
}

type SubnetReferenceResponseOutput struct{ *pulumi.OutputState }

func (SubnetReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetReferenceResponse)(nil)).Elem()
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponseOutput() SubnetReferenceResponseOutput {
	return o
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponseOutputWithContext(ctx context.Context) SubnetReferenceResponseOutput {
	return o
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return o.ToSubnetReferenceResponsePtrOutputWithContext(context.Background())
}

func (o SubnetReferenceResponseOutput) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubnetReferenceResponse) *SubnetReferenceResponse {
		return &v
	}).(SubnetReferenceResponsePtrOutput)
}

func (o SubnetReferenceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetReferenceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceResponseOutput) SourceArmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SubnetReferenceResponse) string { return v.SourceArmResourceId }).(pulumi.StringOutput)
}

type SubnetReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubnetReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubnetReferenceResponse)(nil)).Elem()
}

func (o SubnetReferenceResponsePtrOutput) ToSubnetReferenceResponsePtrOutput() SubnetReferenceResponsePtrOutput {
	return o
}

func (o SubnetReferenceResponsePtrOutput) ToSubnetReferenceResponsePtrOutputWithContext(ctx context.Context) SubnetReferenceResponsePtrOutput {
	return o
}

func (o SubnetReferenceResponsePtrOutput) Elem() SubnetReferenceResponseOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) SubnetReferenceResponse {
		if v != nil {
			return *v
		}
		var ret SubnetReferenceResponse
		return ret
	}).(SubnetReferenceResponseOutput)
}

func (o SubnetReferenceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SubnetReferenceResponsePtrOutput) SourceArmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubnetReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceArmResourceId
	}).(pulumi.StringPtrOutput)
}

type SubnetResourceSettings struct {
	AddressPrefix        *string       `pulumi:"addressPrefix"`
	Name                 *string       `pulumi:"name"`
	NetworkSecurityGroup *NsgReference `pulumi:"networkSecurityGroup"`
}





type SubnetResourceSettingsInput interface {
	pulumi.Input

	ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput
	ToSubnetResourceSettingsOutputWithContext(context.Context) SubnetResourceSettingsOutput
}

type SubnetResourceSettingsArgs struct {
	AddressPrefix        pulumi.StringPtrInput `pulumi:"addressPrefix"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	NetworkSecurityGroup NsgReferencePtrInput  `pulumi:"networkSecurityGroup"`
}

func (SubnetResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettings)(nil)).Elem()
}

func (i SubnetResourceSettingsArgs) ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput {
	return i.ToSubnetResourceSettingsOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsArgs) ToSubnetResourceSettingsOutputWithContext(ctx context.Context) SubnetResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsOutput)
}





type SubnetResourceSettingsArrayInput interface {
	pulumi.Input

	ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput
	ToSubnetResourceSettingsArrayOutputWithContext(context.Context) SubnetResourceSettingsArrayOutput
}

type SubnetResourceSettingsArray []SubnetResourceSettingsInput

func (SubnetResourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettings)(nil)).Elem()
}

func (i SubnetResourceSettingsArray) ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput {
	return i.ToSubnetResourceSettingsArrayOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsArray) ToSubnetResourceSettingsArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsArrayOutput)
}

type SubnetResourceSettingsOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettings)(nil)).Elem()
}

func (o SubnetResourceSettingsOutput) ToSubnetResourceSettingsOutput() SubnetResourceSettingsOutput {
	return o
}

func (o SubnetResourceSettingsOutput) ToSubnetResourceSettingsOutputWithContext(ctx context.Context) SubnetResourceSettingsOutput {
	return o
}

func (o SubnetResourceSettingsOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsOutput) NetworkSecurityGroup() NsgReferencePtrOutput {
	return o.ApplyT(func(v SubnetResourceSettings) *NsgReference { return v.NetworkSecurityGroup }).(NsgReferencePtrOutput)
}

type SubnetResourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettings)(nil)).Elem()
}

func (o SubnetResourceSettingsArrayOutput) ToSubnetResourceSettingsArrayOutput() SubnetResourceSettingsArrayOutput {
	return o
}

func (o SubnetResourceSettingsArrayOutput) ToSubnetResourceSettingsArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsArrayOutput {
	return o
}

func (o SubnetResourceSettingsArrayOutput) Index(i pulumi.IntInput) SubnetResourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResourceSettings {
		return vs[0].([]SubnetResourceSettings)[vs[1].(int)]
	}).(SubnetResourceSettingsOutput)
}

type SubnetResourceSettingsResponse struct {
	AddressPrefix        *string               `pulumi:"addressPrefix"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *NsgReferenceResponse `pulumi:"networkSecurityGroup"`
}





type SubnetResourceSettingsResponseInput interface {
	pulumi.Input

	ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput
	ToSubnetResourceSettingsResponseOutputWithContext(context.Context) SubnetResourceSettingsResponseOutput
}

type SubnetResourceSettingsResponseArgs struct {
	AddressPrefix        pulumi.StringPtrInput        `pulumi:"addressPrefix"`
	Name                 pulumi.StringPtrInput        `pulumi:"name"`
	NetworkSecurityGroup NsgReferenceResponsePtrInput `pulumi:"networkSecurityGroup"`
}

func (SubnetResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettingsResponse)(nil)).Elem()
}

func (i SubnetResourceSettingsResponseArgs) ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput {
	return i.ToSubnetResourceSettingsResponseOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsResponseArgs) ToSubnetResourceSettingsResponseOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsResponseOutput)
}





type SubnetResourceSettingsResponseArrayInput interface {
	pulumi.Input

	ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput
	ToSubnetResourceSettingsResponseArrayOutputWithContext(context.Context) SubnetResourceSettingsResponseArrayOutput
}

type SubnetResourceSettingsResponseArray []SubnetResourceSettingsResponseInput

func (SubnetResourceSettingsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettingsResponse)(nil)).Elem()
}

func (i SubnetResourceSettingsResponseArray) ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput {
	return i.ToSubnetResourceSettingsResponseArrayOutputWithContext(context.Background())
}

func (i SubnetResourceSettingsResponseArray) ToSubnetResourceSettingsResponseArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetResourceSettingsResponseArrayOutput)
}

type SubnetResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetResourceSettingsResponse)(nil)).Elem()
}

func (o SubnetResourceSettingsResponseOutput) ToSubnetResourceSettingsResponseOutput() SubnetResourceSettingsResponseOutput {
	return o
}

func (o SubnetResourceSettingsResponseOutput) ToSubnetResourceSettingsResponseOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseOutput {
	return o
}

func (o SubnetResourceSettingsResponseOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetResourceSettingsResponseOutput) NetworkSecurityGroup() NsgReferenceResponsePtrOutput {
	return o.ApplyT(func(v SubnetResourceSettingsResponse) *NsgReferenceResponse { return v.NetworkSecurityGroup }).(NsgReferenceResponsePtrOutput)
}

type SubnetResourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (SubnetResourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubnetResourceSettingsResponse)(nil)).Elem()
}

func (o SubnetResourceSettingsResponseArrayOutput) ToSubnetResourceSettingsResponseArrayOutput() SubnetResourceSettingsResponseArrayOutput {
	return o
}

func (o SubnetResourceSettingsResponseArrayOutput) ToSubnetResourceSettingsResponseArrayOutputWithContext(ctx context.Context) SubnetResourceSettingsResponseArrayOutput {
	return o
}

func (o SubnetResourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) SubnetResourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubnetResourceSettingsResponse {
		return vs[0].([]SubnetResourceSettingsResponse)[vs[1].(int)]
	}).(SubnetResourceSettingsResponseOutput)
}

type VirtualMachineResourceSettings struct {
	ResourceType            string  `pulumi:"resourceType"`
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  *string `pulumi:"targetAvailabilityZone"`
	TargetResourceName      string  `pulumi:"targetResourceName"`
	TargetVmSize            *string `pulumi:"targetVmSize"`
}





type VirtualMachineResourceSettingsInput interface {
	pulumi.Input

	ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput
	ToVirtualMachineResourceSettingsOutputWithContext(context.Context) VirtualMachineResourceSettingsOutput
}

type VirtualMachineResourceSettingsArgs struct {
	ResourceType            pulumi.StringInput    `pulumi:"resourceType"`
	TargetAvailabilitySetId pulumi.StringPtrInput `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  pulumi.StringPtrInput `pulumi:"targetAvailabilityZone"`
	TargetResourceName      pulumi.StringInput    `pulumi:"targetResourceName"`
	TargetVmSize            pulumi.StringPtrInput `pulumi:"targetVmSize"`
}

func (VirtualMachineResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettings)(nil)).Elem()
}

func (i VirtualMachineResourceSettingsArgs) ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput {
	return i.ToVirtualMachineResourceSettingsOutputWithContext(context.Background())
}

func (i VirtualMachineResourceSettingsArgs) ToVirtualMachineResourceSettingsOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResourceSettingsOutput)
}

type VirtualMachineResourceSettingsOutput struct{ *pulumi.OutputState }

func (VirtualMachineResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettings)(nil)).Elem()
}

func (o VirtualMachineResourceSettingsOutput) ToVirtualMachineResourceSettingsOutput() VirtualMachineResourceSettingsOutput {
	return o
}

func (o VirtualMachineResourceSettingsOutput) ToVirtualMachineResourceSettingsOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsOutput {
	return o
}

func (o VirtualMachineResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettings) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

type VirtualMachineResourceSettingsResponse struct {
	ResourceType            string  `pulumi:"resourceType"`
	TargetAvailabilitySetId *string `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  *string `pulumi:"targetAvailabilityZone"`
	TargetResourceName      string  `pulumi:"targetResourceName"`
	TargetVmSize            *string `pulumi:"targetVmSize"`
}





type VirtualMachineResourceSettingsResponseInput interface {
	pulumi.Input

	ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput
	ToVirtualMachineResourceSettingsResponseOutputWithContext(context.Context) VirtualMachineResourceSettingsResponseOutput
}

type VirtualMachineResourceSettingsResponseArgs struct {
	ResourceType            pulumi.StringInput    `pulumi:"resourceType"`
	TargetAvailabilitySetId pulumi.StringPtrInput `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone  pulumi.StringPtrInput `pulumi:"targetAvailabilityZone"`
	TargetResourceName      pulumi.StringInput    `pulumi:"targetResourceName"`
	TargetVmSize            pulumi.StringPtrInput `pulumi:"targetVmSize"`
}

func (VirtualMachineResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettingsResponse)(nil)).Elem()
}

func (i VirtualMachineResourceSettingsResponseArgs) ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput {
	return i.ToVirtualMachineResourceSettingsResponseOutputWithContext(context.Background())
}

func (i VirtualMachineResourceSettingsResponseArgs) ToVirtualMachineResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineResourceSettingsResponseOutput)
}

type VirtualMachineResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualMachineResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineResourceSettingsResponse)(nil)).Elem()
}

func (o VirtualMachineResourceSettingsResponseOutput) ToVirtualMachineResourceSettingsResponseOutput() VirtualMachineResourceSettingsResponseOutput {
	return o
}

func (o VirtualMachineResourceSettingsResponseOutput) ToVirtualMachineResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualMachineResourceSettingsResponseOutput {
	return o
}

func (o VirtualMachineResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineResourceSettingsResponseOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachineResourceSettingsResponse) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

type VirtualNetworkResourceSettings struct {
	AddressSpace         []string                 `pulumi:"addressSpace"`
	DnsServers           []string                 `pulumi:"dnsServers"`
	EnableDdosProtection *bool                    `pulumi:"enableDdosProtection"`
	ResourceType         string                   `pulumi:"resourceType"`
	Subnets              []SubnetResourceSettings `pulumi:"subnets"`
	TargetResourceName   string                   `pulumi:"targetResourceName"`
}





type VirtualNetworkResourceSettingsInput interface {
	pulumi.Input

	ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput
	ToVirtualNetworkResourceSettingsOutputWithContext(context.Context) VirtualNetworkResourceSettingsOutput
}

type VirtualNetworkResourceSettingsArgs struct {
	AddressSpace         pulumi.StringArrayInput          `pulumi:"addressSpace"`
	DnsServers           pulumi.StringArrayInput          `pulumi:"dnsServers"`
	EnableDdosProtection pulumi.BoolPtrInput              `pulumi:"enableDdosProtection"`
	ResourceType         pulumi.StringInput               `pulumi:"resourceType"`
	Subnets              SubnetResourceSettingsArrayInput `pulumi:"subnets"`
	TargetResourceName   pulumi.StringInput               `pulumi:"targetResourceName"`
}

func (VirtualNetworkResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettings)(nil)).Elem()
}

func (i VirtualNetworkResourceSettingsArgs) ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput {
	return i.ToVirtualNetworkResourceSettingsOutputWithContext(context.Background())
}

func (i VirtualNetworkResourceSettingsArgs) ToVirtualNetworkResourceSettingsOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceSettingsOutput)
}

type VirtualNetworkResourceSettingsOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettings)(nil)).Elem()
}

func (o VirtualNetworkResourceSettingsOutput) ToVirtualNetworkResourceSettingsOutput() VirtualNetworkResourceSettingsOutput {
	return o
}

func (o VirtualNetworkResourceSettingsOutput) ToVirtualNetworkResourceSettingsOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsOutput {
	return o
}

func (o VirtualNetworkResourceSettingsOutput) AddressSpace() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []string { return v.AddressSpace }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkResourceSettingsOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualNetworkResourceSettingsOutput) Subnets() SubnetResourceSettingsArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) []SubnetResourceSettings { return v.Subnets }).(SubnetResourceSettingsArrayOutput)
}

func (o VirtualNetworkResourceSettingsOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettings) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

type VirtualNetworkResourceSettingsResponse struct {
	AddressSpace         []string                         `pulumi:"addressSpace"`
	DnsServers           []string                         `pulumi:"dnsServers"`
	EnableDdosProtection *bool                            `pulumi:"enableDdosProtection"`
	ResourceType         string                           `pulumi:"resourceType"`
	Subnets              []SubnetResourceSettingsResponse `pulumi:"subnets"`
	TargetResourceName   string                           `pulumi:"targetResourceName"`
}





type VirtualNetworkResourceSettingsResponseInput interface {
	pulumi.Input

	ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput
	ToVirtualNetworkResourceSettingsResponseOutputWithContext(context.Context) VirtualNetworkResourceSettingsResponseOutput
}

type VirtualNetworkResourceSettingsResponseArgs struct {
	AddressSpace         pulumi.StringArrayInput                  `pulumi:"addressSpace"`
	DnsServers           pulumi.StringArrayInput                  `pulumi:"dnsServers"`
	EnableDdosProtection pulumi.BoolPtrInput                      `pulumi:"enableDdosProtection"`
	ResourceType         pulumi.StringInput                       `pulumi:"resourceType"`
	Subnets              SubnetResourceSettingsResponseArrayInput `pulumi:"subnets"`
	TargetResourceName   pulumi.StringInput                       `pulumi:"targetResourceName"`
}

func (VirtualNetworkResourceSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettingsResponse)(nil)).Elem()
}

func (i VirtualNetworkResourceSettingsResponseArgs) ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput {
	return i.ToVirtualNetworkResourceSettingsResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkResourceSettingsResponseArgs) ToVirtualNetworkResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkResourceSettingsResponseOutput)
}

type VirtualNetworkResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkResourceSettingsResponse)(nil)).Elem()
}

func (o VirtualNetworkResourceSettingsResponseOutput) ToVirtualNetworkResourceSettingsResponseOutput() VirtualNetworkResourceSettingsResponseOutput {
	return o
}

func (o VirtualNetworkResourceSettingsResponseOutput) ToVirtualNetworkResourceSettingsResponseOutputWithContext(ctx context.Context) VirtualNetworkResourceSettingsResponseOutput {
	return o
}

func (o VirtualNetworkResourceSettingsResponseOutput) AddressSpace() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []string { return v.AddressSpace }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) EnableDdosProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) *bool { return v.EnableDdosProtection }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) Subnets() SubnetResourceSettingsResponseArrayOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) []SubnetResourceSettingsResponse { return v.Subnets }).(SubnetResourceSettingsResponseArrayOutput)
}

func (o VirtualNetworkResourceSettingsResponseOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkResourceSettingsResponse) string { return v.TargetResourceName }).(pulumi.StringOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentPropertiesInput)(nil)).Elem(), AssessmentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentPropertiesPtrInput)(nil)).Elem(), AssessmentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentPropertiesResponseInput)(nil)).Elem(), AssessmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssessmentPropertiesResponsePtrInput)(nil)).Elem(), AssessmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomaticResolutionPropertiesResponseInput)(nil)).Elem(), AutomaticResolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomaticResolutionPropertiesResponsePtrInput)(nil)).Elem(), AutomaticResolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AvailabilitySetResourceSettingsInput)(nil)).Elem(), AvailabilitySetResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AvailabilitySetResourceSettingsResponseInput)(nil)).Elem(), AvailabilitySetResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorAgentPropertiesInput)(nil)).Elem(), CollectorAgentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorAgentPropertiesPtrInput)(nil)).Elem(), CollectorAgentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorAgentPropertiesResponseInput)(nil)).Elem(), CollectorAgentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorAgentPropertiesResponsePtrInput)(nil)).Elem(), CollectorAgentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorBodyAgentSpnPropertiesInput)(nil)).Elem(), CollectorBodyAgentSpnPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorBodyAgentSpnPropertiesPtrInput)(nil)).Elem(), CollectorBodyAgentSpnPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorBodyAgentSpnPropertiesResponseInput)(nil)).Elem(), CollectorBodyAgentSpnPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorBodyAgentSpnPropertiesResponsePtrInput)(nil)).Elem(), CollectorBodyAgentSpnPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorPropertiesInput)(nil)).Elem(), CollectorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorPropertiesPtrInput)(nil)).Elem(), CollectorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorPropertiesResponseInput)(nil)).Elem(), CollectorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CollectorPropertiesResponsePtrInput)(nil)).Elem(), CollectorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseProjectSummaryResponseInput)(nil)).Elem(), DatabaseProjectSummaryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabasesSolutionSummaryResponseInput)(nil)).Elem(), DatabasesSolutionSummaryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetResourceSettingsInput)(nil)).Elem(), DiskEncryptionSetResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionSetResourceSettingsResponseInput)(nil)).Elem(), DiskEncryptionSetResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPropertiesInput)(nil)).Elem(), GroupPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPropertiesPtrInput)(nil)).Elem(), GroupPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPropertiesResponseInput)(nil)).Elem(), GroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupPropertiesResponsePtrInput)(nil)).Elem(), GroupPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportCollectorPropertiesInput)(nil)).Elem(), ImportCollectorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportCollectorPropertiesPtrInput)(nil)).Elem(), ImportCollectorPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportCollectorPropertiesResponseInput)(nil)).Elem(), ImportCollectorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImportCollectorPropertiesResponsePtrInput)(nil)).Elem(), ImportCollectorPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStatusResponseInput)(nil)).Elem(), JobStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStatusResponsePtrInput)(nil)).Elem(), JobStatusResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultResourceSettingsInput)(nil)).Elem(), KeyVaultResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultResourceSettingsResponseInput)(nil)).Elem(), KeyVaultResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBBackendAddressPoolResourceSettingsInput)(nil)).Elem(), LBBackendAddressPoolResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBBackendAddressPoolResourceSettingsArrayInput)(nil)).Elem(), LBBackendAddressPoolResourceSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponseInput)(nil)).Elem(), LBBackendAddressPoolResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBBackendAddressPoolResourceSettingsResponseArrayInput)(nil)).Elem(), LBBackendAddressPoolResourceSettingsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsInput)(nil)).Elem(), LBFrontendIPConfigurationResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsArrayInput)(nil)).Elem(), LBFrontendIPConfigurationResourceSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponseInput)(nil)).Elem(), LBFrontendIPConfigurationResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LBFrontendIPConfigurationResourceSettingsResponseArrayInput)(nil)).Elem(), LBFrontendIPConfigurationResourceSettingsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceInput)(nil)).Elem(), LoadBalancerBackendAddressPoolReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceArrayInput)(nil)).Elem(), LoadBalancerBackendAddressPoolReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponseInput)(nil)).Elem(), LoadBalancerBackendAddressPoolReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerBackendAddressPoolReferenceResponseArrayInput)(nil)).Elem(), LoadBalancerBackendAddressPoolReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerNatRuleReferenceInput)(nil)).Elem(), LoadBalancerNatRuleReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerNatRuleReferenceArrayInput)(nil)).Elem(), LoadBalancerNatRuleReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerNatRuleReferenceResponseInput)(nil)).Elem(), LoadBalancerNatRuleReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerNatRuleReferenceResponseArrayInput)(nil)).Elem(), LoadBalancerNatRuleReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerResourceSettingsInput)(nil)).Elem(), LoadBalancerResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerResourceSettingsResponseInput)(nil)).Elem(), LoadBalancerResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManualResolutionPropertiesResponseInput)(nil)).Elem(), ManualResolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ManualResolutionPropertiesResponsePtrInput)(nil)).Elem(), ManualResolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectPropertiesInput)(nil)).Elem(), MigrateProjectPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectPropertiesPtrInput)(nil)).Elem(), MigrateProjectPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectPropertiesResponseInput)(nil)).Elem(), MigrateProjectPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectPropertiesResponsePtrInput)(nil)).Elem(), MigrateProjectPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectResponseTagsInput)(nil)).Elem(), MigrateProjectResponseTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectResponseTagsPtrInput)(nil)).Elem(), MigrateProjectResponseTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectTagsInput)(nil)).Elem(), MigrateProjectTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrateProjectTagsPtrInput)(nil)).Elem(), MigrateProjectTagsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesInput)(nil)).Elem(), MoveCollectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesPtrInput)(nil)).Elem(), MoveCollectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesResponseInput)(nil)).Elem(), MoveCollectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesResponsePtrInput)(nil)).Elem(), MoveCollectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesResponseErrorsInput)(nil)).Elem(), MoveCollectionPropertiesResponseErrorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveCollectionPropertiesResponseErrorsPtrInput)(nil)).Elem(), MoveCollectionPropertiesResponseErrorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyOverrideInput)(nil)).Elem(), MoveResourceDependencyOverrideArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyOverrideArrayInput)(nil)).Elem(), MoveResourceDependencyOverrideArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyOverrideResponseInput)(nil)).Elem(), MoveResourceDependencyOverrideResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyOverrideResponseArrayInput)(nil)).Elem(), MoveResourceDependencyOverrideResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyResponseInput)(nil)).Elem(), MoveResourceDependencyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceDependencyResponseArrayInput)(nil)).Elem(), MoveResourceDependencyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceErrorBodyResponseInput)(nil)).Elem(), MoveResourceErrorBodyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceErrorBodyResponsePtrInput)(nil)).Elem(), MoveResourceErrorBodyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceErrorBodyResponseArrayInput)(nil)).Elem(), MoveResourceErrorBodyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceErrorResponseInput)(nil)).Elem(), MoveResourceErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourceErrorResponsePtrInput)(nil)).Elem(), MoveResourceErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesInput)(nil)).Elem(), MoveResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesPtrInput)(nil)).Elem(), MoveResourcePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponseInput)(nil)).Elem(), MoveResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponsePtrInput)(nil)).Elem(), MoveResourcePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponseErrorsInput)(nil)).Elem(), MoveResourcePropertiesResponseErrorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponseErrorsPtrInput)(nil)).Elem(), MoveResourcePropertiesResponseErrorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatusInput)(nil)).Elem(), MoveResourcePropertiesResponseMoveStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveResourcePropertiesResponseMoveStatusPtrInput)(nil)).Elem(), MoveResourcePropertiesResponseMoveStatusArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceResourceSettingsInput)(nil)).Elem(), NetworkInterfaceResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkInterfaceResourceSettingsResponseInput)(nil)).Elem(), NetworkInterfaceResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityGroupResourceSettingsInput)(nil)).Elem(), NetworkSecurityGroupResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityGroupResourceSettingsResponseInput)(nil)).Elem(), NetworkSecurityGroupResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicIpConfigurationResourceSettingsInput)(nil)).Elem(), NicIpConfigurationResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicIpConfigurationResourceSettingsArrayInput)(nil)).Elem(), NicIpConfigurationResourceSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicIpConfigurationResourceSettingsResponseInput)(nil)).Elem(), NicIpConfigurationResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicIpConfigurationResourceSettingsResponseArrayInput)(nil)).Elem(), NicIpConfigurationResourceSettingsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgReferenceInput)(nil)).Elem(), NsgReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgReferencePtrInput)(nil)).Elem(), NsgReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgReferenceResponseInput)(nil)).Elem(), NsgReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgReferenceResponsePtrInput)(nil)).Elem(), NsgReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgSecurityRuleInput)(nil)).Elem(), NsgSecurityRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgSecurityRuleArrayInput)(nil)).Elem(), NsgSecurityRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgSecurityRuleResponseInput)(nil)).Elem(), NsgSecurityRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsgSecurityRuleResponseArrayInput)(nil)).Elem(), NsgSecurityRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesPtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponseInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePtrInput)(nil)).Elem(), PrivateEndpointConnectionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPropertiesInput)(nil)).Elem(), ProjectPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPropertiesPtrInput)(nil)).Elem(), ProjectPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPropertiesResponseInput)(nil)).Elem(), ProjectPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectPropertiesResponsePtrInput)(nil)).Elem(), ProjectPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIPAddressResourceSettingsInput)(nil)).Elem(), PublicIPAddressResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIPAddressResourceSettingsResponseInput)(nil)).Elem(), PublicIPAddressResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpReferenceInput)(nil)).Elem(), PublicIpReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpReferencePtrInput)(nil)).Elem(), PublicIpReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpReferenceResponseInput)(nil)).Elem(), PublicIpReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicIpReferenceResponsePtrInput)(nil)).Elem(), PublicIpReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupResourceSettingsInput)(nil)).Elem(), ResourceGroupResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceGroupResourceSettingsResponseInput)(nil)).Elem(), ResourceGroupResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdResponseInput)(nil)).Elem(), ResourceIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdResponsePtrInput)(nil)).Elem(), ResourceIdResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServersProjectSummaryResponseInput)(nil)).Elem(), ServersProjectSummaryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServersSolutionSummaryResponseInput)(nil)).Elem(), ServersSolutionSummaryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionDetailsInput)(nil)).Elem(), SolutionDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionDetailsPtrInput)(nil)).Elem(), SolutionDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionDetailsResponseInput)(nil)).Elem(), SolutionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionDetailsResponsePtrInput)(nil)).Elem(), SolutionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionPropertiesInput)(nil)).Elem(), SolutionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionPropertiesPtrInput)(nil)).Elem(), SolutionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionPropertiesResponseInput)(nil)).Elem(), SolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SolutionPropertiesResponsePtrInput)(nil)).Elem(), SolutionPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlDatabaseResourceSettingsInput)(nil)).Elem(), SqlDatabaseResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlDatabaseResourceSettingsResponseInput)(nil)).Elem(), SqlDatabaseResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlElasticPoolResourceSettingsInput)(nil)).Elem(), SqlElasticPoolResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlElasticPoolResourceSettingsResponseInput)(nil)).Elem(), SqlElasticPoolResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerResourceSettingsInput)(nil)).Elem(), SqlServerResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SqlServerResourceSettingsResponseInput)(nil)).Elem(), SqlServerResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetReferenceInput)(nil)).Elem(), SubnetReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetReferencePtrInput)(nil)).Elem(), SubnetReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetReferenceResponseInput)(nil)).Elem(), SubnetReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetReferenceResponsePtrInput)(nil)).Elem(), SubnetReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResourceSettingsInput)(nil)).Elem(), SubnetResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResourceSettingsArrayInput)(nil)).Elem(), SubnetResourceSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResourceSettingsResponseInput)(nil)).Elem(), SubnetResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubnetResourceSettingsResponseArrayInput)(nil)).Elem(), SubnetResourceSettingsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineResourceSettingsInput)(nil)).Elem(), VirtualMachineResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualMachineResourceSettingsResponseInput)(nil)).Elem(), VirtualMachineResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkResourceSettingsInput)(nil)).Elem(), VirtualNetworkResourceSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkResourceSettingsResponseInput)(nil)).Elem(), VirtualNetworkResourceSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmUptimeInput)(nil)).Elem(), VmUptimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmUptimePtrInput)(nil)).Elem(), VmUptimeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmUptimeResponseInput)(nil)).Elem(), VmUptimeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmUptimeResponsePtrInput)(nil)).Elem(), VmUptimeResponseArgs{})
	pulumi.RegisterOutputType(AssessmentPropertiesOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AutomaticResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AvailabilitySetResourceSettingsOutput{})
	pulumi.RegisterOutputType(AvailabilitySetResourceSettingsResponseOutput{})
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
	pulumi.RegisterOutputType(DatabaseProjectSummaryResponseOutput{})
	pulumi.RegisterOutputType(DatabasesSolutionSummaryResponseOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetResourceSettingsOutput{})
	pulumi.RegisterOutputType(DiskEncryptionSetResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(GroupPropertiesOutput{})
	pulumi.RegisterOutputType(GroupPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GroupPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStatusResponseOutput{})
	pulumi.RegisterOutputType(JobStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultResourceSettingsOutput{})
	pulumi.RegisterOutputType(KeyVaultResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBBackendAddressPoolResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(LBFrontendIPConfigurationResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerBackendAddressPoolReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerNatRuleReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerResourceSettingsOutput{})
	pulumi.RegisterOutputType(LoadBalancerResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ManualResolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MigrateProjectPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectResponseTagsOutput{})
	pulumi.RegisterOutputType(MigrateProjectResponseTagsPtrOutput{})
	pulumi.RegisterOutputType(MigrateProjectTagsOutput{})
	pulumi.RegisterOutputType(MigrateProjectTagsPtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveCollectionPropertiesResponseErrorsPtrOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyOverrideResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceDependencyResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorBodyResponseArrayOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorResponseOutput{})
	pulumi.RegisterOutputType(MoveResourceErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseErrorsOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseErrorsPtrOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseMoveStatusOutput{})
	pulumi.RegisterOutputType(MoveResourcePropertiesResponseMoveStatusPtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResourceSettingsOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResourceSettingsOutput{})
	pulumi.RegisterOutputType(NetworkSecurityGroupResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(NicIpConfigurationResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(NsgReferenceOutput{})
	pulumi.RegisterOutputType(NsgReferencePtrOutput{})
	pulumi.RegisterOutputType(NsgReferenceResponseOutput{})
	pulumi.RegisterOutputType(NsgReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleArrayOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleResponseOutput{})
	pulumi.RegisterOutputType(NsgSecurityRuleResponseArrayOutput{})
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
	pulumi.RegisterOutputType(PublicIPAddressResourceSettingsOutput{})
	pulumi.RegisterOutputType(PublicIPAddressResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceOutput{})
	pulumi.RegisterOutputType(PublicIpReferencePtrOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceResponseOutput{})
	pulumi.RegisterOutputType(PublicIpReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceGroupResourceSettingsOutput{})
	pulumi.RegisterOutputType(ResourceGroupResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(ServersProjectSummaryResponseOutput{})
	pulumi.RegisterOutputType(ServersSolutionSummaryResponseOutput{})
	pulumi.RegisterOutputType(SolutionDetailsOutput{})
	pulumi.RegisterOutputType(SolutionDetailsPtrOutput{})
	pulumi.RegisterOutputType(SolutionDetailsResponseOutput{})
	pulumi.RegisterOutputType(SolutionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SolutionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlElasticPoolResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlElasticPoolResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlServerResourceSettingsOutput{})
	pulumi.RegisterOutputType(SqlServerResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubnetReferenceOutput{})
	pulumi.RegisterOutputType(SubnetReferencePtrOutput{})
	pulumi.RegisterOutputType(SubnetReferenceResponseOutput{})
	pulumi.RegisterOutputType(SubnetReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubnetResourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachineResourceSettingsOutput{})
	pulumi.RegisterOutputType(VirtualMachineResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkResourceSettingsOutput{})
	pulumi.RegisterOutputType(VirtualNetworkResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(VmUptimeOutput{})
	pulumi.RegisterOutputType(VmUptimePtrOutput{})
	pulumi.RegisterOutputType(VmUptimeResponseOutput{})
	pulumi.RegisterOutputType(VmUptimeResponsePtrOutput{})
}
