


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessmentMetadataInSubscription(ctx *pulumi.Context, args *LookupAssessmentMetadataInSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentMetadataInSubscriptionResult, error) {
	var rv LookupAssessmentMetadataInSubscriptionResult
	err := ctx.Invoke("azure-native:security:getAssessmentMetadataInSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentMetadataInSubscriptionArgs struct {
	AssessmentMetadataName string `pulumi:"assessmentMetadataName"`
}


type LookupAssessmentMetadataInSubscriptionResult struct {
	AssessmentType         string                                         `pulumi:"assessmentType"`
	Categories             []string                                       `pulumi:"categories"`
	Description            *string                                        `pulumi:"description"`
	DisplayName            string                                         `pulumi:"displayName"`
	Id                     string                                         `pulumi:"id"`
	ImplementationEffort   *string                                        `pulumi:"implementationEffort"`
	Name                   string                                         `pulumi:"name"`
	PartnerData            *SecurityAssessmentMetadataPartnerDataResponse `pulumi:"partnerData"`
	PolicyDefinitionId     string                                         `pulumi:"policyDefinitionId"`
	Preview                *bool                                          `pulumi:"preview"`
	RemediationDescription *string                                        `pulumi:"remediationDescription"`
	Severity               string                                         `pulumi:"severity"`
	Threats                []string                                       `pulumi:"threats"`
	Type                   string                                         `pulumi:"type"`
	UserImpact             *string                                        `pulumi:"userImpact"`
}

func LookupAssessmentMetadataInSubscriptionOutput(ctx *pulumi.Context, args LookupAssessmentMetadataInSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAssessmentMetadataInSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssessmentMetadataInSubscriptionResult, error) {
			args := v.(LookupAssessmentMetadataInSubscriptionArgs)
			r, err := LookupAssessmentMetadataInSubscription(ctx, &args, opts...)
			var s LookupAssessmentMetadataInSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssessmentMetadataInSubscriptionResultOutput)
}

type LookupAssessmentMetadataInSubscriptionOutputArgs struct {
	AssessmentMetadataName pulumi.StringInput `pulumi:"assessmentMetadataName"`
}

func (LookupAssessmentMetadataInSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentMetadataInSubscriptionArgs)(nil)).Elem()
}


type LookupAssessmentMetadataInSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAssessmentMetadataInSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentMetadataInSubscriptionResult)(nil)).Elem()
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) ToLookupAssessmentMetadataInSubscriptionResultOutput() LookupAssessmentMetadataInSubscriptionResultOutput {
	return o
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) ToLookupAssessmentMetadataInSubscriptionResultOutputWithContext(ctx context.Context) LookupAssessmentMetadataInSubscriptionResultOutput {
	return o
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) PartnerData() SecurityAssessmentMetadataPartnerDataResponsePtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *SecurityAssessmentMetadataPartnerDataResponse {
		return v.PartnerData
	}).(SecurityAssessmentMetadataPartnerDataResponsePtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAssessmentMetadataInSubscriptionResultOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentMetadataInSubscriptionResult) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentMetadataInSubscriptionResultOutput{})
}
