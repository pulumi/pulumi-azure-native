


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCustomAssessmentAutomation(ctx *pulumi.Context, args *LookupCustomAssessmentAutomationArgs, opts ...pulumi.InvokeOption) (*LookupCustomAssessmentAutomationResult, error) {
	var rv LookupCustomAssessmentAutomationResult
	err := ctx.Invoke("azure-native:security/v20210701preview:getCustomAssessmentAutomation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomAssessmentAutomationArgs struct {
	CustomAssessmentAutomationName string `pulumi:"customAssessmentAutomationName"`
	ResourceGroupName              string `pulumi:"resourceGroupName"`
}


type LookupCustomAssessmentAutomationResult struct {
	AssessmentKey          *string            `pulumi:"assessmentKey"`
	CompressedQuery        *string            `pulumi:"compressedQuery"`
	Description            *string            `pulumi:"description"`
	DisplayName            *string            `pulumi:"displayName"`
	Id                     string             `pulumi:"id"`
	Name                   string             `pulumi:"name"`
	RemediationDescription *string            `pulumi:"remediationDescription"`
	Severity               *string            `pulumi:"severity"`
	SupportedCloud         *string            `pulumi:"supportedCloud"`
	SystemData             SystemDataResponse `pulumi:"systemData"`
	Type                   string             `pulumi:"type"`
}

func LookupCustomAssessmentAutomationOutput(ctx *pulumi.Context, args LookupCustomAssessmentAutomationOutputArgs, opts ...pulumi.InvokeOption) LookupCustomAssessmentAutomationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomAssessmentAutomationResult, error) {
			args := v.(LookupCustomAssessmentAutomationArgs)
			r, err := LookupCustomAssessmentAutomation(ctx, &args, opts...)
			var s LookupCustomAssessmentAutomationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomAssessmentAutomationResultOutput)
}

type LookupCustomAssessmentAutomationOutputArgs struct {
	CustomAssessmentAutomationName pulumi.StringInput `pulumi:"customAssessmentAutomationName"`
	ResourceGroupName              pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCustomAssessmentAutomationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomAssessmentAutomationArgs)(nil)).Elem()
}


type LookupCustomAssessmentAutomationResultOutput struct{ *pulumi.OutputState }

func (LookupCustomAssessmentAutomationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomAssessmentAutomationResult)(nil)).Elem()
}

func (o LookupCustomAssessmentAutomationResultOutput) ToLookupCustomAssessmentAutomationResultOutput() LookupCustomAssessmentAutomationResultOutput {
	return o
}

func (o LookupCustomAssessmentAutomationResultOutput) ToLookupCustomAssessmentAutomationResultOutputWithContext(ctx context.Context) LookupCustomAssessmentAutomationResultOutput {
	return o
}

func (o LookupCustomAssessmentAutomationResultOutput) AssessmentKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.AssessmentKey }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) CompressedQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.CompressedQuery }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) RemediationDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.RemediationDescription }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) SupportedCloud() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) *string { return v.SupportedCloud }).(pulumi.StringPtrOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupCustomAssessmentAutomationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomAssessmentAutomationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomAssessmentAutomationResultOutput{})
}
