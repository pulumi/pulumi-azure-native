


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessmentsMetadataSubscription(ctx *pulumi.Context, args *LookupAssessmentsMetadataSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentsMetadataSubscriptionResult, error) {
	var rv LookupAssessmentsMetadataSubscriptionResult
	err := ctx.Invoke("azure-native:security/v20190101preview:getAssessmentsMetadataSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentsMetadataSubscriptionArgs struct {
	AssessmentMetadataName string `pulumi:"assessmentMetadataName"`
}


type LookupAssessmentsMetadataSubscriptionResult struct {
	AssessmentType         string   `pulumi:"assessmentType"`
	Categories             []string `pulumi:"categories"`
	Description            *string  `pulumi:"description"`
	DisplayName            string   `pulumi:"displayName"`
	Id                     string   `pulumi:"id"`
	ImplementationEffort   *string  `pulumi:"implementationEffort"`
	Name                   string   `pulumi:"name"`
	PolicyDefinitionId     string   `pulumi:"policyDefinitionId"`
	Preview                *bool    `pulumi:"preview"`
	RemediationDescription *string  `pulumi:"remediationDescription"`
	Severity               string   `pulumi:"severity"`
	Threats                []string `pulumi:"threats"`
	Type                   string   `pulumi:"type"`
	UserImpact             *string  `pulumi:"userImpact"`
}

func LookupAssessmentsMetadataSubscriptionOutput(ctx *pulumi.Context, args LookupAssessmentsMetadataSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupAssessmentsMetadataSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssessmentsMetadataSubscriptionResult, error) {
			args := v.(LookupAssessmentsMetadataSubscriptionArgs)
			r, err := LookupAssessmentsMetadataSubscription(ctx, &args, opts...)
			var s LookupAssessmentsMetadataSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssessmentsMetadataSubscriptionResultOutput)
}

type LookupAssessmentsMetadataSubscriptionOutputArgs struct {
	AssessmentMetadataName pulumi.StringInput `pulumi:"assessmentMetadataName"`
}

func (LookupAssessmentsMetadataSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentsMetadataSubscriptionArgs)(nil)).Elem()
}


type LookupAssessmentsMetadataSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupAssessmentsMetadataSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentsMetadataSubscriptionResult)(nil)).Elem()
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) ToLookupAssessmentsMetadataSubscriptionResultOutput() LookupAssessmentsMetadataSubscriptionResultOutput {
	return o
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) ToLookupAssessmentsMetadataSubscriptionResultOutputWithContext(ctx context.Context) LookupAssessmentsMetadataSubscriptionResultOutput {
	return o
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) AssessmentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.AssessmentType }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) ImplementationEffort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) *string { return v.ImplementationEffort }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) PolicyDefinitionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.PolicyDefinitionId }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Preview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) *bool { return v.Preview }).(pulumi.BoolPtrOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Threats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) []string { return v.Threats }).(pulumi.StringArrayOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAssessmentsMetadataSubscriptionResultOutput) UserImpact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentsMetadataSubscriptionResult) *string { return v.UserImpact }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentsMetadataSubscriptionResultOutput{})
}
