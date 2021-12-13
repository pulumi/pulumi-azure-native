


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

func (o GroupPropertiesOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupProperties) *string { return v.GroupType }).(pulumi.StringPtrOutput)
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

func (o ImportCollectorPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o ImportCollectorPropertiesResponseOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

func (o ImportCollectorPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
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

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   ResourceIdResponse                         `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
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

type PrivateEndpointConnectionResponse struct {
	ETag       *string                                     `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
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

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
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

func (o ResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
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

func (o VmUptimeOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptime) *float64 { return v.DaysPerMonth }).(pulumi.Float64PtrOutput)
}

func (o VmUptimeOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptime) *float64 { return v.HoursPerDay }).(pulumi.Float64PtrOutput)
}

type VmUptimeResponse struct {
	DaysPerMonth *float64 `pulumi:"daysPerMonth"`
	HoursPerDay  *float64 `pulumi:"hoursPerDay"`
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

func (o VmUptimeResponseOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptimeResponse) *float64 { return v.DaysPerMonth }).(pulumi.Float64PtrOutput)
}

func (o VmUptimeResponseOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptimeResponse) *float64 { return v.HoursPerDay }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentPropertiesOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(GroupPropertiesOutput{})
	pulumi.RegisterOutputType(GroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(VmUptimeOutput{})
	pulumi.RegisterOutputType(VmUptimeResponseOutput{})
}
