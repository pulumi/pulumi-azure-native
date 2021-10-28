


package v20171111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentStage string

const (
	AssessmentStageInProgress  = AssessmentStage("InProgress")
	AssessmentStageUnderReview = AssessmentStage("UnderReview")
	AssessmentStageApproved    = AssessmentStage("Approved")
)

func (AssessmentStage) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStage)(nil)).Elem()
}

func (e AssessmentStage) ToAssessmentStageOutput() AssessmentStageOutput {
	return pulumi.ToOutput(e).(AssessmentStageOutput)
}

func (e AssessmentStage) ToAssessmentStageOutputWithContext(ctx context.Context) AssessmentStageOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AssessmentStageOutput)
}

func (e AssessmentStage) ToAssessmentStagePtrOutput() AssessmentStagePtrOutput {
	return e.ToAssessmentStagePtrOutputWithContext(context.Background())
}

func (e AssessmentStage) ToAssessmentStagePtrOutputWithContext(ctx context.Context) AssessmentStagePtrOutput {
	return AssessmentStage(e).ToAssessmentStageOutputWithContext(ctx).ToAssessmentStagePtrOutputWithContext(ctx)
}

func (e AssessmentStage) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStage) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AssessmentStage) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AssessmentStage) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AssessmentStageOutput struct{ *pulumi.OutputState }

func (AssessmentStageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentStage)(nil)).Elem()
}

func (o AssessmentStageOutput) ToAssessmentStageOutput() AssessmentStageOutput {
	return o
}

func (o AssessmentStageOutput) ToAssessmentStageOutputWithContext(ctx context.Context) AssessmentStageOutput {
	return o
}

func (o AssessmentStageOutput) ToAssessmentStagePtrOutput() AssessmentStagePtrOutput {
	return o.ToAssessmentStagePtrOutputWithContext(context.Background())
}

func (o AssessmentStageOutput) ToAssessmentStagePtrOutputWithContext(ctx context.Context) AssessmentStagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentStage) *AssessmentStage {
		return &v
	}).(AssessmentStagePtrOutput)
}

func (o AssessmentStageOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AssessmentStageOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStage) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AssessmentStageOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStageOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AssessmentStage) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AssessmentStagePtrOutput struct{ *pulumi.OutputState }

func (AssessmentStagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentStage)(nil)).Elem()
}

func (o AssessmentStagePtrOutput) ToAssessmentStagePtrOutput() AssessmentStagePtrOutput {
	return o
}

func (o AssessmentStagePtrOutput) ToAssessmentStagePtrOutputWithContext(ctx context.Context) AssessmentStagePtrOutput {
	return o
}

func (o AssessmentStagePtrOutput) Elem() AssessmentStageOutput {
	return o.ApplyT(func(v *AssessmentStage) AssessmentStage {
		if v != nil {
			return *v
		}
		var ret AssessmentStage
		return ret
	}).(AssessmentStageOutput)
}

func (o AssessmentStagePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AssessmentStagePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AssessmentStage) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AssessmentStageInput interface {
	pulumi.Input

	ToAssessmentStageOutput() AssessmentStageOutput
	ToAssessmentStageOutputWithContext(context.Context) AssessmentStageOutput
}

var assessmentStagePtrType = reflect.TypeOf((**AssessmentStage)(nil)).Elem()

type AssessmentStagePtrInput interface {
	pulumi.Input

	ToAssessmentStagePtrOutput() AssessmentStagePtrOutput
	ToAssessmentStagePtrOutputWithContext(context.Context) AssessmentStagePtrOutput
}

type assessmentStagePtr string

func AssessmentStagePtr(v string) AssessmentStagePtrInput {
	return (*assessmentStagePtr)(&v)
}

func (*assessmentStagePtr) ElementType() reflect.Type {
	return assessmentStagePtrType
}

func (in *assessmentStagePtr) ToAssessmentStagePtrOutput() AssessmentStagePtrOutput {
	return pulumi.ToOutput(in).(AssessmentStagePtrOutput)
}

func (in *assessmentStagePtr) ToAssessmentStagePtrOutputWithContext(ctx context.Context) AssessmentStagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AssessmentStagePtrOutput)
}

type AzureHybridUseBenefit string

const (
	AzureHybridUseBenefitUnknown = AzureHybridUseBenefit("Unknown")
	AzureHybridUseBenefitYes     = AzureHybridUseBenefit("Yes")
	AzureHybridUseBenefitNo      = AzureHybridUseBenefit("No")
)

func (AzureHybridUseBenefit) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureHybridUseBenefit)(nil)).Elem()
}

func (e AzureHybridUseBenefit) ToAzureHybridUseBenefitOutput() AzureHybridUseBenefitOutput {
	return pulumi.ToOutput(e).(AzureHybridUseBenefitOutput)
}

func (e AzureHybridUseBenefit) ToAzureHybridUseBenefitOutputWithContext(ctx context.Context) AzureHybridUseBenefitOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureHybridUseBenefitOutput)
}

func (e AzureHybridUseBenefit) ToAzureHybridUseBenefitPtrOutput() AzureHybridUseBenefitPtrOutput {
	return e.ToAzureHybridUseBenefitPtrOutputWithContext(context.Background())
}

func (e AzureHybridUseBenefit) ToAzureHybridUseBenefitPtrOutputWithContext(ctx context.Context) AzureHybridUseBenefitPtrOutput {
	return AzureHybridUseBenefit(e).ToAzureHybridUseBenefitOutputWithContext(ctx).ToAzureHybridUseBenefitPtrOutputWithContext(ctx)
}

func (e AzureHybridUseBenefit) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureHybridUseBenefit) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureHybridUseBenefit) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureHybridUseBenefit) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureHybridUseBenefitOutput struct{ *pulumi.OutputState }

func (AzureHybridUseBenefitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureHybridUseBenefit)(nil)).Elem()
}

func (o AzureHybridUseBenefitOutput) ToAzureHybridUseBenefitOutput() AzureHybridUseBenefitOutput {
	return o
}

func (o AzureHybridUseBenefitOutput) ToAzureHybridUseBenefitOutputWithContext(ctx context.Context) AzureHybridUseBenefitOutput {
	return o
}

func (o AzureHybridUseBenefitOutput) ToAzureHybridUseBenefitPtrOutput() AzureHybridUseBenefitPtrOutput {
	return o.ToAzureHybridUseBenefitPtrOutputWithContext(context.Background())
}

func (o AzureHybridUseBenefitOutput) ToAzureHybridUseBenefitPtrOutputWithContext(ctx context.Context) AzureHybridUseBenefitPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureHybridUseBenefit) *AzureHybridUseBenefit {
		return &v
	}).(AzureHybridUseBenefitPtrOutput)
}

func (o AzureHybridUseBenefitOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureHybridUseBenefitOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureHybridUseBenefit) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureHybridUseBenefitOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureHybridUseBenefitOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureHybridUseBenefit) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureHybridUseBenefitPtrOutput struct{ *pulumi.OutputState }

func (AzureHybridUseBenefitPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureHybridUseBenefit)(nil)).Elem()
}

func (o AzureHybridUseBenefitPtrOutput) ToAzureHybridUseBenefitPtrOutput() AzureHybridUseBenefitPtrOutput {
	return o
}

func (o AzureHybridUseBenefitPtrOutput) ToAzureHybridUseBenefitPtrOutputWithContext(ctx context.Context) AzureHybridUseBenefitPtrOutput {
	return o
}

func (o AzureHybridUseBenefitPtrOutput) Elem() AzureHybridUseBenefitOutput {
	return o.ApplyT(func(v *AzureHybridUseBenefit) AzureHybridUseBenefit {
		if v != nil {
			return *v
		}
		var ret AzureHybridUseBenefit
		return ret
	}).(AzureHybridUseBenefitOutput)
}

func (o AzureHybridUseBenefitPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureHybridUseBenefitPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureHybridUseBenefit) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureHybridUseBenefitInput interface {
	pulumi.Input

	ToAzureHybridUseBenefitOutput() AzureHybridUseBenefitOutput
	ToAzureHybridUseBenefitOutputWithContext(context.Context) AzureHybridUseBenefitOutput
}

var azureHybridUseBenefitPtrType = reflect.TypeOf((**AzureHybridUseBenefit)(nil)).Elem()

type AzureHybridUseBenefitPtrInput interface {
	pulumi.Input

	ToAzureHybridUseBenefitPtrOutput() AzureHybridUseBenefitPtrOutput
	ToAzureHybridUseBenefitPtrOutputWithContext(context.Context) AzureHybridUseBenefitPtrOutput
}

type azureHybridUseBenefitPtr string

func AzureHybridUseBenefitPtr(v string) AzureHybridUseBenefitPtrInput {
	return (*azureHybridUseBenefitPtr)(&v)
}

func (*azureHybridUseBenefitPtr) ElementType() reflect.Type {
	return azureHybridUseBenefitPtrType
}

func (in *azureHybridUseBenefitPtr) ToAzureHybridUseBenefitPtrOutput() AzureHybridUseBenefitPtrOutput {
	return pulumi.ToOutput(in).(AzureHybridUseBenefitPtrOutput)
}

func (in *azureHybridUseBenefitPtr) ToAzureHybridUseBenefitPtrOutputWithContext(ctx context.Context) AzureHybridUseBenefitPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureHybridUseBenefitPtrOutput)
}

type AzureLocation string

const (
	AzureLocationUnknown            = AzureLocation("Unknown")
	AzureLocationEastAsia           = AzureLocation("EastAsia")
	AzureLocationSoutheastAsia      = AzureLocation("SoutheastAsia")
	AzureLocationAustraliaEast      = AzureLocation("AustraliaEast")
	AzureLocationAustraliaSoutheast = AzureLocation("AustraliaSoutheast")
	AzureLocationBrazilSouth        = AzureLocation("BrazilSouth")
	AzureLocationCanadaCentral      = AzureLocation("CanadaCentral")
	AzureLocationCanadaEast         = AzureLocation("CanadaEast")
	AzureLocationWestEurope         = AzureLocation("WestEurope")
	AzureLocationNorthEurope        = AzureLocation("NorthEurope")
	AzureLocationCentralIndia       = AzureLocation("CentralIndia")
	AzureLocationSouthIndia         = AzureLocation("SouthIndia")
	AzureLocationWestIndia          = AzureLocation("WestIndia")
	AzureLocationJapanEast          = AzureLocation("JapanEast")
	AzureLocationJapanWest          = AzureLocation("JapanWest")
	AzureLocationKoreaCentral       = AzureLocation("KoreaCentral")
	AzureLocationKoreaSouth         = AzureLocation("KoreaSouth")
	AzureLocationUkWest             = AzureLocation("UkWest")
	AzureLocationUkSouth            = AzureLocation("UkSouth")
	AzureLocationNorthCentralUs     = AzureLocation("NorthCentralUs")
	AzureLocationEastUs             = AzureLocation("EastUs")
	AzureLocationWestUs2            = AzureLocation("WestUs2")
	AzureLocationSouthCentralUs     = AzureLocation("SouthCentralUs")
	AzureLocationCentralUs          = AzureLocation("CentralUs")
	AzureLocationEastUs2            = AzureLocation("EastUs2")
	AzureLocationWestUs             = AzureLocation("WestUs")
	AzureLocationWestCentralUs      = AzureLocation("WestCentralUs")
)

func (AzureLocation) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureLocation)(nil)).Elem()
}

func (e AzureLocation) ToAzureLocationOutput() AzureLocationOutput {
	return pulumi.ToOutput(e).(AzureLocationOutput)
}

func (e AzureLocation) ToAzureLocationOutputWithContext(ctx context.Context) AzureLocationOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureLocationOutput)
}

func (e AzureLocation) ToAzureLocationPtrOutput() AzureLocationPtrOutput {
	return e.ToAzureLocationPtrOutputWithContext(context.Background())
}

func (e AzureLocation) ToAzureLocationPtrOutputWithContext(ctx context.Context) AzureLocationPtrOutput {
	return AzureLocation(e).ToAzureLocationOutputWithContext(ctx).ToAzureLocationPtrOutputWithContext(ctx)
}

func (e AzureLocation) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureLocation) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureLocation) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureLocation) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureLocationOutput struct{ *pulumi.OutputState }

func (AzureLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureLocation)(nil)).Elem()
}

func (o AzureLocationOutput) ToAzureLocationOutput() AzureLocationOutput {
	return o
}

func (o AzureLocationOutput) ToAzureLocationOutputWithContext(ctx context.Context) AzureLocationOutput {
	return o
}

func (o AzureLocationOutput) ToAzureLocationPtrOutput() AzureLocationPtrOutput {
	return o.ToAzureLocationPtrOutputWithContext(context.Background())
}

func (o AzureLocationOutput) ToAzureLocationPtrOutputWithContext(ctx context.Context) AzureLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureLocation) *AzureLocation {
		return &v
	}).(AzureLocationPtrOutput)
}

func (o AzureLocationOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureLocationOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureLocation) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureLocationOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureLocationOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureLocation) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureLocationPtrOutput struct{ *pulumi.OutputState }

func (AzureLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureLocation)(nil)).Elem()
}

func (o AzureLocationPtrOutput) ToAzureLocationPtrOutput() AzureLocationPtrOutput {
	return o
}

func (o AzureLocationPtrOutput) ToAzureLocationPtrOutputWithContext(ctx context.Context) AzureLocationPtrOutput {
	return o
}

func (o AzureLocationPtrOutput) Elem() AzureLocationOutput {
	return o.ApplyT(func(v *AzureLocation) AzureLocation {
		if v != nil {
			return *v
		}
		var ret AzureLocation
		return ret
	}).(AzureLocationOutput)
}

func (o AzureLocationPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureLocationPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureLocation) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureLocationInput interface {
	pulumi.Input

	ToAzureLocationOutput() AzureLocationOutput
	ToAzureLocationOutputWithContext(context.Context) AzureLocationOutput
}

var azureLocationPtrType = reflect.TypeOf((**AzureLocation)(nil)).Elem()

type AzureLocationPtrInput interface {
	pulumi.Input

	ToAzureLocationPtrOutput() AzureLocationPtrOutput
	ToAzureLocationPtrOutputWithContext(context.Context) AzureLocationPtrOutput
}

type azureLocationPtr string

func AzureLocationPtr(v string) AzureLocationPtrInput {
	return (*azureLocationPtr)(&v)
}

func (*azureLocationPtr) ElementType() reflect.Type {
	return azureLocationPtrType
}

func (in *azureLocationPtr) ToAzureLocationPtrOutput() AzureLocationPtrOutput {
	return pulumi.ToOutput(in).(AzureLocationPtrOutput)
}

func (in *azureLocationPtr) ToAzureLocationPtrOutputWithContext(ctx context.Context) AzureLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureLocationPtrOutput)
}

type AzureOfferCode string

const (
	AzureOfferCodeUnknown    = AzureOfferCode("Unknown")
	AzureOfferCodeMSAZR0003P = AzureOfferCode("MSAZR0003P")
	AzureOfferCodeMSAZR0044P = AzureOfferCode("MSAZR0044P")
	AzureOfferCodeMSAZR0059P = AzureOfferCode("MSAZR0059P")
	AzureOfferCodeMSAZR0060P = AzureOfferCode("MSAZR0060P")
	AzureOfferCodeMSAZR0062P = AzureOfferCode("MSAZR0062P")
	AzureOfferCodeMSAZR0063P = AzureOfferCode("MSAZR0063P")
	AzureOfferCodeMSAZR0064P = AzureOfferCode("MSAZR0064P")
	AzureOfferCodeMSAZR0029P = AzureOfferCode("MSAZR0029P")
	AzureOfferCodeMSAZR0022P = AzureOfferCode("MSAZR0022P")
	AzureOfferCodeMSAZR0023P = AzureOfferCode("MSAZR0023P")
	AzureOfferCodeMSAZR0148P = AzureOfferCode("MSAZR0148P")
	AzureOfferCodeMSAZR0025P = AzureOfferCode("MSAZR0025P")
	AzureOfferCodeMSAZR0036P = AzureOfferCode("MSAZR0036P")
	AzureOfferCodeMSAZR0120P = AzureOfferCode("MSAZR0120P")
	AzureOfferCodeMSAZR0121P = AzureOfferCode("MSAZR0121P")
	AzureOfferCodeMSAZR0122P = AzureOfferCode("MSAZR0122P")
	AzureOfferCodeMSAZR0123P = AzureOfferCode("MSAZR0123P")
	AzureOfferCodeMSAZR0124P = AzureOfferCode("MSAZR0124P")
	AzureOfferCodeMSAZR0125P = AzureOfferCode("MSAZR0125P")
	AzureOfferCodeMSAZR0126P = AzureOfferCode("MSAZR0126P")
	AzureOfferCodeMSAZR0127P = AzureOfferCode("MSAZR0127P")
	AzureOfferCodeMSAZR0128P = AzureOfferCode("MSAZR0128P")
	AzureOfferCodeMSAZR0129P = AzureOfferCode("MSAZR0129P")
	AzureOfferCodeMSAZR0130P = AzureOfferCode("MSAZR0130P")
	AzureOfferCodeMSAZR0111P = AzureOfferCode("MSAZR0111P")
	AzureOfferCodeMSAZR0144P = AzureOfferCode("MSAZR0144P")
	AzureOfferCodeMSAZR0149P = AzureOfferCode("MSAZR0149P")
)

func (AzureOfferCode) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureOfferCode)(nil)).Elem()
}

func (e AzureOfferCode) ToAzureOfferCodeOutput() AzureOfferCodeOutput {
	return pulumi.ToOutput(e).(AzureOfferCodeOutput)
}

func (e AzureOfferCode) ToAzureOfferCodeOutputWithContext(ctx context.Context) AzureOfferCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureOfferCodeOutput)
}

func (e AzureOfferCode) ToAzureOfferCodePtrOutput() AzureOfferCodePtrOutput {
	return e.ToAzureOfferCodePtrOutputWithContext(context.Background())
}

func (e AzureOfferCode) ToAzureOfferCodePtrOutputWithContext(ctx context.Context) AzureOfferCodePtrOutput {
	return AzureOfferCode(e).ToAzureOfferCodeOutputWithContext(ctx).ToAzureOfferCodePtrOutputWithContext(ctx)
}

func (e AzureOfferCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureOfferCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureOfferCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureOfferCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureOfferCodeOutput struct{ *pulumi.OutputState }

func (AzureOfferCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureOfferCode)(nil)).Elem()
}

func (o AzureOfferCodeOutput) ToAzureOfferCodeOutput() AzureOfferCodeOutput {
	return o
}

func (o AzureOfferCodeOutput) ToAzureOfferCodeOutputWithContext(ctx context.Context) AzureOfferCodeOutput {
	return o
}

func (o AzureOfferCodeOutput) ToAzureOfferCodePtrOutput() AzureOfferCodePtrOutput {
	return o.ToAzureOfferCodePtrOutputWithContext(context.Background())
}

func (o AzureOfferCodeOutput) ToAzureOfferCodePtrOutputWithContext(ctx context.Context) AzureOfferCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureOfferCode) *AzureOfferCode {
		return &v
	}).(AzureOfferCodePtrOutput)
}

func (o AzureOfferCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureOfferCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureOfferCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureOfferCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureOfferCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureOfferCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureOfferCodePtrOutput struct{ *pulumi.OutputState }

func (AzureOfferCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureOfferCode)(nil)).Elem()
}

func (o AzureOfferCodePtrOutput) ToAzureOfferCodePtrOutput() AzureOfferCodePtrOutput {
	return o
}

func (o AzureOfferCodePtrOutput) ToAzureOfferCodePtrOutputWithContext(ctx context.Context) AzureOfferCodePtrOutput {
	return o
}

func (o AzureOfferCodePtrOutput) Elem() AzureOfferCodeOutput {
	return o.ApplyT(func(v *AzureOfferCode) AzureOfferCode {
		if v != nil {
			return *v
		}
		var ret AzureOfferCode
		return ret
	}).(AzureOfferCodeOutput)
}

func (o AzureOfferCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureOfferCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureOfferCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureOfferCodeInput interface {
	pulumi.Input

	ToAzureOfferCodeOutput() AzureOfferCodeOutput
	ToAzureOfferCodeOutputWithContext(context.Context) AzureOfferCodeOutput
}

var azureOfferCodePtrType = reflect.TypeOf((**AzureOfferCode)(nil)).Elem()

type AzureOfferCodePtrInput interface {
	pulumi.Input

	ToAzureOfferCodePtrOutput() AzureOfferCodePtrOutput
	ToAzureOfferCodePtrOutputWithContext(context.Context) AzureOfferCodePtrOutput
}

type azureOfferCodePtr string

func AzureOfferCodePtr(v string) AzureOfferCodePtrInput {
	return (*azureOfferCodePtr)(&v)
}

func (*azureOfferCodePtr) ElementType() reflect.Type {
	return azureOfferCodePtrType
}

func (in *azureOfferCodePtr) ToAzureOfferCodePtrOutput() AzureOfferCodePtrOutput {
	return pulumi.ToOutput(in).(AzureOfferCodePtrOutput)
}

func (in *azureOfferCodePtr) ToAzureOfferCodePtrOutputWithContext(ctx context.Context) AzureOfferCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureOfferCodePtrOutput)
}

type AzurePricingTier string

const (
	AzurePricingTierStandard = AzurePricingTier("Standard")
	AzurePricingTierBasic    = AzurePricingTier("Basic")
)

func (AzurePricingTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePricingTier)(nil)).Elem()
}

func (e AzurePricingTier) ToAzurePricingTierOutput() AzurePricingTierOutput {
	return pulumi.ToOutput(e).(AzurePricingTierOutput)
}

func (e AzurePricingTier) ToAzurePricingTierOutputWithContext(ctx context.Context) AzurePricingTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzurePricingTierOutput)
}

func (e AzurePricingTier) ToAzurePricingTierPtrOutput() AzurePricingTierPtrOutput {
	return e.ToAzurePricingTierPtrOutputWithContext(context.Background())
}

func (e AzurePricingTier) ToAzurePricingTierPtrOutputWithContext(ctx context.Context) AzurePricingTierPtrOutput {
	return AzurePricingTier(e).ToAzurePricingTierOutputWithContext(ctx).ToAzurePricingTierPtrOutputWithContext(ctx)
}

func (e AzurePricingTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzurePricingTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzurePricingTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzurePricingTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzurePricingTierOutput struct{ *pulumi.OutputState }

func (AzurePricingTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzurePricingTier)(nil)).Elem()
}

func (o AzurePricingTierOutput) ToAzurePricingTierOutput() AzurePricingTierOutput {
	return o
}

func (o AzurePricingTierOutput) ToAzurePricingTierOutputWithContext(ctx context.Context) AzurePricingTierOutput {
	return o
}

func (o AzurePricingTierOutput) ToAzurePricingTierPtrOutput() AzurePricingTierPtrOutput {
	return o.ToAzurePricingTierPtrOutputWithContext(context.Background())
}

func (o AzurePricingTierOutput) ToAzurePricingTierPtrOutputWithContext(ctx context.Context) AzurePricingTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzurePricingTier) *AzurePricingTier {
		return &v
	}).(AzurePricingTierPtrOutput)
}

func (o AzurePricingTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzurePricingTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzurePricingTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzurePricingTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzurePricingTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzurePricingTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzurePricingTierPtrOutput struct{ *pulumi.OutputState }

func (AzurePricingTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzurePricingTier)(nil)).Elem()
}

func (o AzurePricingTierPtrOutput) ToAzurePricingTierPtrOutput() AzurePricingTierPtrOutput {
	return o
}

func (o AzurePricingTierPtrOutput) ToAzurePricingTierPtrOutputWithContext(ctx context.Context) AzurePricingTierPtrOutput {
	return o
}

func (o AzurePricingTierPtrOutput) Elem() AzurePricingTierOutput {
	return o.ApplyT(func(v *AzurePricingTier) AzurePricingTier {
		if v != nil {
			return *v
		}
		var ret AzurePricingTier
		return ret
	}).(AzurePricingTierOutput)
}

func (o AzurePricingTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzurePricingTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzurePricingTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzurePricingTierInput interface {
	pulumi.Input

	ToAzurePricingTierOutput() AzurePricingTierOutput
	ToAzurePricingTierOutputWithContext(context.Context) AzurePricingTierOutput
}

var azurePricingTierPtrType = reflect.TypeOf((**AzurePricingTier)(nil)).Elem()

type AzurePricingTierPtrInput interface {
	pulumi.Input

	ToAzurePricingTierPtrOutput() AzurePricingTierPtrOutput
	ToAzurePricingTierPtrOutputWithContext(context.Context) AzurePricingTierPtrOutput
}

type azurePricingTierPtr string

func AzurePricingTierPtr(v string) AzurePricingTierPtrInput {
	return (*azurePricingTierPtr)(&v)
}

func (*azurePricingTierPtr) ElementType() reflect.Type {
	return azurePricingTierPtrType
}

func (in *azurePricingTierPtr) ToAzurePricingTierPtrOutput() AzurePricingTierPtrOutput {
	return pulumi.ToOutput(in).(AzurePricingTierPtrOutput)
}

func (in *azurePricingTierPtr) ToAzurePricingTierPtrOutputWithContext(ctx context.Context) AzurePricingTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzurePricingTierPtrOutput)
}

type AzureStorageRedundancy string

const (
	AzureStorageRedundancyUnknown                = AzureStorageRedundancy("Unknown")
	AzureStorageRedundancyLocallyRedundant       = AzureStorageRedundancy("LocallyRedundant")
	AzureStorageRedundancyZoneRedundant          = AzureStorageRedundancy("ZoneRedundant")
	AzureStorageRedundancyGeoRedundant           = AzureStorageRedundancy("GeoRedundant")
	AzureStorageRedundancyReadAccessGeoRedundant = AzureStorageRedundancy("ReadAccessGeoRedundant")
)

func (AzureStorageRedundancy) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageRedundancy)(nil)).Elem()
}

func (e AzureStorageRedundancy) ToAzureStorageRedundancyOutput() AzureStorageRedundancyOutput {
	return pulumi.ToOutput(e).(AzureStorageRedundancyOutput)
}

func (e AzureStorageRedundancy) ToAzureStorageRedundancyOutputWithContext(ctx context.Context) AzureStorageRedundancyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureStorageRedundancyOutput)
}

func (e AzureStorageRedundancy) ToAzureStorageRedundancyPtrOutput() AzureStorageRedundancyPtrOutput {
	return e.ToAzureStorageRedundancyPtrOutputWithContext(context.Background())
}

func (e AzureStorageRedundancy) ToAzureStorageRedundancyPtrOutputWithContext(ctx context.Context) AzureStorageRedundancyPtrOutput {
	return AzureStorageRedundancy(e).ToAzureStorageRedundancyOutputWithContext(ctx).ToAzureStorageRedundancyPtrOutputWithContext(ctx)
}

func (e AzureStorageRedundancy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageRedundancy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureStorageRedundancy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureStorageRedundancy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureStorageRedundancyOutput struct{ *pulumi.OutputState }

func (AzureStorageRedundancyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageRedundancy)(nil)).Elem()
}

func (o AzureStorageRedundancyOutput) ToAzureStorageRedundancyOutput() AzureStorageRedundancyOutput {
	return o
}

func (o AzureStorageRedundancyOutput) ToAzureStorageRedundancyOutputWithContext(ctx context.Context) AzureStorageRedundancyOutput {
	return o
}

func (o AzureStorageRedundancyOutput) ToAzureStorageRedundancyPtrOutput() AzureStorageRedundancyPtrOutput {
	return o.ToAzureStorageRedundancyPtrOutputWithContext(context.Background())
}

func (o AzureStorageRedundancyOutput) ToAzureStorageRedundancyPtrOutputWithContext(ctx context.Context) AzureStorageRedundancyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureStorageRedundancy) *AzureStorageRedundancy {
		return &v
	}).(AzureStorageRedundancyPtrOutput)
}

func (o AzureStorageRedundancyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureStorageRedundancyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureStorageRedundancy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureStorageRedundancyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureStorageRedundancyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureStorageRedundancy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureStorageRedundancyPtrOutput struct{ *pulumi.OutputState }

func (AzureStorageRedundancyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureStorageRedundancy)(nil)).Elem()
}

func (o AzureStorageRedundancyPtrOutput) ToAzureStorageRedundancyPtrOutput() AzureStorageRedundancyPtrOutput {
	return o
}

func (o AzureStorageRedundancyPtrOutput) ToAzureStorageRedundancyPtrOutputWithContext(ctx context.Context) AzureStorageRedundancyPtrOutput {
	return o
}

func (o AzureStorageRedundancyPtrOutput) Elem() AzureStorageRedundancyOutput {
	return o.ApplyT(func(v *AzureStorageRedundancy) AzureStorageRedundancy {
		if v != nil {
			return *v
		}
		var ret AzureStorageRedundancy
		return ret
	}).(AzureStorageRedundancyOutput)
}

func (o AzureStorageRedundancyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureStorageRedundancyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureStorageRedundancy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureStorageRedundancyInput interface {
	pulumi.Input

	ToAzureStorageRedundancyOutput() AzureStorageRedundancyOutput
	ToAzureStorageRedundancyOutputWithContext(context.Context) AzureStorageRedundancyOutput
}

var azureStorageRedundancyPtrType = reflect.TypeOf((**AzureStorageRedundancy)(nil)).Elem()

type AzureStorageRedundancyPtrInput interface {
	pulumi.Input

	ToAzureStorageRedundancyPtrOutput() AzureStorageRedundancyPtrOutput
	ToAzureStorageRedundancyPtrOutputWithContext(context.Context) AzureStorageRedundancyPtrOutput
}

type azureStorageRedundancyPtr string

func AzureStorageRedundancyPtr(v string) AzureStorageRedundancyPtrInput {
	return (*azureStorageRedundancyPtr)(&v)
}

func (*azureStorageRedundancyPtr) ElementType() reflect.Type {
	return azureStorageRedundancyPtrType
}

func (in *azureStorageRedundancyPtr) ToAzureStorageRedundancyPtrOutput() AzureStorageRedundancyPtrOutput {
	return pulumi.ToOutput(in).(AzureStorageRedundancyPtrOutput)
}

func (in *azureStorageRedundancyPtr) ToAzureStorageRedundancyPtrOutputWithContext(ctx context.Context) AzureStorageRedundancyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureStorageRedundancyPtrOutput)
}

type Currency string

const (
	CurrencyUnknown = Currency("Unknown")
	CurrencyUSD     = Currency("USD")
	CurrencyDKK     = Currency("DKK")
	CurrencyCAD     = Currency("CAD")
	CurrencyIDR     = Currency("IDR")
	CurrencyJPY     = Currency("JPY")
	CurrencyKRW     = Currency("KRW")
	CurrencyNZD     = Currency("NZD")
	CurrencyNOK     = Currency("NOK")
	CurrencyRUB     = Currency("RUB")
	CurrencySAR     = Currency("SAR")
	CurrencyZAR     = Currency("ZAR")
	CurrencySEK     = Currency("SEK")
	CurrencyTRY     = Currency("TRY")
	CurrencyGBP     = Currency("GBP")
	CurrencyMXN     = Currency("MXN")
	CurrencyMYR     = Currency("MYR")
	CurrencyINR     = Currency("INR")
	CurrencyHKD     = Currency("HKD")
	CurrencyBRL     = Currency("BRL")
	CurrencyTWD     = Currency("TWD")
	CurrencyEUR     = Currency("EUR")
	CurrencyCHF     = Currency("CHF")
	CurrencyARS     = Currency("ARS")
	CurrencyAUD     = Currency("AUD")
)

func (Currency) ElementType() reflect.Type {
	return reflect.TypeOf((*Currency)(nil)).Elem()
}

func (e Currency) ToCurrencyOutput() CurrencyOutput {
	return pulumi.ToOutput(e).(CurrencyOutput)
}

func (e Currency) ToCurrencyOutputWithContext(ctx context.Context) CurrencyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CurrencyOutput)
}

func (e Currency) ToCurrencyPtrOutput() CurrencyPtrOutput {
	return e.ToCurrencyPtrOutputWithContext(context.Background())
}

func (e Currency) ToCurrencyPtrOutputWithContext(ctx context.Context) CurrencyPtrOutput {
	return Currency(e).ToCurrencyOutputWithContext(ctx).ToCurrencyPtrOutputWithContext(ctx)
}

func (e Currency) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Currency) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Currency) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Currency) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CurrencyOutput struct{ *pulumi.OutputState }

func (CurrencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Currency)(nil)).Elem()
}

func (o CurrencyOutput) ToCurrencyOutput() CurrencyOutput {
	return o
}

func (o CurrencyOutput) ToCurrencyOutputWithContext(ctx context.Context) CurrencyOutput {
	return o
}

func (o CurrencyOutput) ToCurrencyPtrOutput() CurrencyPtrOutput {
	return o.ToCurrencyPtrOutputWithContext(context.Background())
}

func (o CurrencyOutput) ToCurrencyPtrOutputWithContext(ctx context.Context) CurrencyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Currency) *Currency {
		return &v
	}).(CurrencyPtrOutput)
}

func (o CurrencyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CurrencyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Currency) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CurrencyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CurrencyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Currency) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CurrencyPtrOutput struct{ *pulumi.OutputState }

func (CurrencyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Currency)(nil)).Elem()
}

func (o CurrencyPtrOutput) ToCurrencyPtrOutput() CurrencyPtrOutput {
	return o
}

func (o CurrencyPtrOutput) ToCurrencyPtrOutputWithContext(ctx context.Context) CurrencyPtrOutput {
	return o
}

func (o CurrencyPtrOutput) Elem() CurrencyOutput {
	return o.ApplyT(func(v *Currency) Currency {
		if v != nil {
			return *v
		}
		var ret Currency
		return ret
	}).(CurrencyOutput)
}

func (o CurrencyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CurrencyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Currency) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CurrencyInput interface {
	pulumi.Input

	ToCurrencyOutput() CurrencyOutput
	ToCurrencyOutputWithContext(context.Context) CurrencyOutput
}

var currencyPtrType = reflect.TypeOf((**Currency)(nil)).Elem()

type CurrencyPtrInput interface {
	pulumi.Input

	ToCurrencyPtrOutput() CurrencyPtrOutput
	ToCurrencyPtrOutputWithContext(context.Context) CurrencyPtrOutput
}

type currencyPtr string

func CurrencyPtr(v string) CurrencyPtrInput {
	return (*currencyPtr)(&v)
}

func (*currencyPtr) ElementType() reflect.Type {
	return currencyPtrType
}

func (in *currencyPtr) ToCurrencyPtrOutput() CurrencyPtrOutput {
	return pulumi.ToOutput(in).(CurrencyPtrOutput)
}

func (in *currencyPtr) ToCurrencyPtrOutputWithContext(ctx context.Context) CurrencyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CurrencyPtrOutput)
}

type Percentile string

const (
	PercentilePercentile50 = Percentile("Percentile50")
	PercentilePercentile90 = Percentile("Percentile90")
	PercentilePercentile95 = Percentile("Percentile95")
	PercentilePercentile99 = Percentile("Percentile99")
)

func (Percentile) ElementType() reflect.Type {
	return reflect.TypeOf((*Percentile)(nil)).Elem()
}

func (e Percentile) ToPercentileOutput() PercentileOutput {
	return pulumi.ToOutput(e).(PercentileOutput)
}

func (e Percentile) ToPercentileOutputWithContext(ctx context.Context) PercentileOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PercentileOutput)
}

func (e Percentile) ToPercentilePtrOutput() PercentilePtrOutput {
	return e.ToPercentilePtrOutputWithContext(context.Background())
}

func (e Percentile) ToPercentilePtrOutputWithContext(ctx context.Context) PercentilePtrOutput {
	return Percentile(e).ToPercentileOutputWithContext(ctx).ToPercentilePtrOutputWithContext(ctx)
}

func (e Percentile) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Percentile) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Percentile) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Percentile) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PercentileOutput struct{ *pulumi.OutputState }

func (PercentileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Percentile)(nil)).Elem()
}

func (o PercentileOutput) ToPercentileOutput() PercentileOutput {
	return o
}

func (o PercentileOutput) ToPercentileOutputWithContext(ctx context.Context) PercentileOutput {
	return o
}

func (o PercentileOutput) ToPercentilePtrOutput() PercentilePtrOutput {
	return o.ToPercentilePtrOutputWithContext(context.Background())
}

func (o PercentileOutput) ToPercentilePtrOutputWithContext(ctx context.Context) PercentilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Percentile) *Percentile {
		return &v
	}).(PercentilePtrOutput)
}

func (o PercentileOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PercentileOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Percentile) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PercentileOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PercentileOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Percentile) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PercentilePtrOutput struct{ *pulumi.OutputState }

func (PercentilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Percentile)(nil)).Elem()
}

func (o PercentilePtrOutput) ToPercentilePtrOutput() PercentilePtrOutput {
	return o
}

func (o PercentilePtrOutput) ToPercentilePtrOutputWithContext(ctx context.Context) PercentilePtrOutput {
	return o
}

func (o PercentilePtrOutput) Elem() PercentileOutput {
	return o.ApplyT(func(v *Percentile) Percentile {
		if v != nil {
			return *v
		}
		var ret Percentile
		return ret
	}).(PercentileOutput)
}

func (o PercentilePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PercentilePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Percentile) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PercentileInput interface {
	pulumi.Input

	ToPercentileOutput() PercentileOutput
	ToPercentileOutputWithContext(context.Context) PercentileOutput
}

var percentilePtrType = reflect.TypeOf((**Percentile)(nil)).Elem()

type PercentilePtrInput interface {
	pulumi.Input

	ToPercentilePtrOutput() PercentilePtrOutput
	ToPercentilePtrOutputWithContext(context.Context) PercentilePtrOutput
}

type percentilePtr string

func PercentilePtr(v string) PercentilePtrInput {
	return (*percentilePtr)(&v)
}

func (*percentilePtr) ElementType() reflect.Type {
	return percentilePtrType
}

func (in *percentilePtr) ToPercentilePtrOutput() PercentilePtrOutput {
	return pulumi.ToOutput(in).(PercentilePtrOutput)
}

func (in *percentilePtr) ToPercentilePtrOutputWithContext(ctx context.Context) PercentilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PercentilePtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  = ProvisioningState("Accepted")
	ProvisioningStateCreating  = ProvisioningState("Creating")
	ProvisioningStateDeleting  = ProvisioningState("Deleting")
	ProvisioningStateFailed    = ProvisioningState("Failed")
	ProvisioningStateMoving    = ProvisioningState("Moving")
	ProvisioningStateSucceeded = ProvisioningState("Succeeded")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type TimeRange string

const (
	TimeRangeDay   = TimeRange("Day")
	TimeRangeWeek  = TimeRange("Week")
	TimeRangeMonth = TimeRange("Month")
)

func (TimeRange) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeRange)(nil)).Elem()
}

func (e TimeRange) ToTimeRangeOutput() TimeRangeOutput {
	return pulumi.ToOutput(e).(TimeRangeOutput)
}

func (e TimeRange) ToTimeRangeOutputWithContext(ctx context.Context) TimeRangeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeRangeOutput)
}

func (e TimeRange) ToTimeRangePtrOutput() TimeRangePtrOutput {
	return e.ToTimeRangePtrOutputWithContext(context.Background())
}

func (e TimeRange) ToTimeRangePtrOutputWithContext(ctx context.Context) TimeRangePtrOutput {
	return TimeRange(e).ToTimeRangeOutputWithContext(ctx).ToTimeRangePtrOutputWithContext(ctx)
}

func (e TimeRange) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeRange) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeRange) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeRange) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeRangeOutput struct{ *pulumi.OutputState }

func (TimeRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeRange)(nil)).Elem()
}

func (o TimeRangeOutput) ToTimeRangeOutput() TimeRangeOutput {
	return o
}

func (o TimeRangeOutput) ToTimeRangeOutputWithContext(ctx context.Context) TimeRangeOutput {
	return o
}

func (o TimeRangeOutput) ToTimeRangePtrOutput() TimeRangePtrOutput {
	return o.ToTimeRangePtrOutputWithContext(context.Background())
}

func (o TimeRangeOutput) ToTimeRangePtrOutputWithContext(ctx context.Context) TimeRangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeRange) *TimeRange {
		return &v
	}).(TimeRangePtrOutput)
}

func (o TimeRangeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeRangeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeRange) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeRangeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeRangeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeRange) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeRangePtrOutput struct{ *pulumi.OutputState }

func (TimeRangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeRange)(nil)).Elem()
}

func (o TimeRangePtrOutput) ToTimeRangePtrOutput() TimeRangePtrOutput {
	return o
}

func (o TimeRangePtrOutput) ToTimeRangePtrOutputWithContext(ctx context.Context) TimeRangePtrOutput {
	return o
}

func (o TimeRangePtrOutput) Elem() TimeRangeOutput {
	return o.ApplyT(func(v *TimeRange) TimeRange {
		if v != nil {
			return *v
		}
		var ret TimeRange
		return ret
	}).(TimeRangeOutput)
}

func (o TimeRangePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeRangePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeRange) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeRangeInput interface {
	pulumi.Input

	ToTimeRangeOutput() TimeRangeOutput
	ToTimeRangeOutputWithContext(context.Context) TimeRangeOutput
}

var timeRangePtrType = reflect.TypeOf((**TimeRange)(nil)).Elem()

type TimeRangePtrInput interface {
	pulumi.Input

	ToTimeRangePtrOutput() TimeRangePtrOutput
	ToTimeRangePtrOutputWithContext(context.Context) TimeRangePtrOutput
}

type timeRangePtr string

func TimeRangePtr(v string) TimeRangePtrInput {
	return (*timeRangePtr)(&v)
}

func (*timeRangePtr) ElementType() reflect.Type {
	return timeRangePtrType
}

func (in *timeRangePtr) ToTimeRangePtrOutput() TimeRangePtrOutput {
	return pulumi.ToOutput(in).(TimeRangePtrOutput)
}

func (in *timeRangePtr) ToTimeRangePtrOutputWithContext(ctx context.Context) TimeRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeRangePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentStageOutput{})
	pulumi.RegisterOutputType(AssessmentStagePtrOutput{})
	pulumi.RegisterOutputType(AzureHybridUseBenefitOutput{})
	pulumi.RegisterOutputType(AzureHybridUseBenefitPtrOutput{})
	pulumi.RegisterOutputType(AzureLocationOutput{})
	pulumi.RegisterOutputType(AzureLocationPtrOutput{})
	pulumi.RegisterOutputType(AzureOfferCodeOutput{})
	pulumi.RegisterOutputType(AzureOfferCodePtrOutput{})
	pulumi.RegisterOutputType(AzurePricingTierOutput{})
	pulumi.RegisterOutputType(AzurePricingTierPtrOutput{})
	pulumi.RegisterOutputType(AzureStorageRedundancyOutput{})
	pulumi.RegisterOutputType(AzureStorageRedundancyPtrOutput{})
	pulumi.RegisterOutputType(CurrencyOutput{})
	pulumi.RegisterOutputType(CurrencyPtrOutput{})
	pulumi.RegisterOutputType(PercentileOutput{})
	pulumi.RegisterOutputType(PercentilePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(TimeRangeOutput{})
	pulumi.RegisterOutputType(TimeRangePtrOutput{})
}
