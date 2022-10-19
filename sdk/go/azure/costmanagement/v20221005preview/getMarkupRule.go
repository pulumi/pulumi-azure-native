


package v20221005preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMarkupRule(ctx *pulumi.Context, args *LookupMarkupRuleArgs, opts ...pulumi.InvokeOption) (*LookupMarkupRuleResult, error) {
	var rv LookupMarkupRuleResult
	err := ctx.Invoke("azure-native:costmanagement/v20221005preview:getMarkupRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMarkupRuleArgs struct {
	BillingAccountId string `pulumi:"billingAccountId"`
	BillingProfileId string `pulumi:"billingProfileId"`
	Name             string `pulumi:"name"`
}


type LookupMarkupRuleResult struct {
	CustomerDetails CustomerMetadataResponse `pulumi:"customerDetails"`
	Description     *string                  `pulumi:"description"`
	ETag            *string                  `pulumi:"eTag"`
	EndDate         *string                  `pulumi:"endDate"`
	Id              string                   `pulumi:"id"`
	Name            string                   `pulumi:"name"`
	Percentage      float64                  `pulumi:"percentage"`
	StartDate       string                   `pulumi:"startDate"`
	Type            string                   `pulumi:"type"`
}

func LookupMarkupRuleOutput(ctx *pulumi.Context, args LookupMarkupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMarkupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMarkupRuleResult, error) {
			args := v.(LookupMarkupRuleArgs)
			r, err := LookupMarkupRule(ctx, &args, opts...)
			var s LookupMarkupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMarkupRuleResultOutput)
}

type LookupMarkupRuleOutputArgs struct {
	BillingAccountId pulumi.StringInput `pulumi:"billingAccountId"`
	BillingProfileId pulumi.StringInput `pulumi:"billingProfileId"`
	Name             pulumi.StringInput `pulumi:"name"`
}

func (LookupMarkupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarkupRuleArgs)(nil)).Elem()
}


type LookupMarkupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMarkupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMarkupRuleResult)(nil)).Elem()
}

func (o LookupMarkupRuleResultOutput) ToLookupMarkupRuleResultOutput() LookupMarkupRuleResultOutput {
	return o
}

func (o LookupMarkupRuleResultOutput) ToLookupMarkupRuleResultOutputWithContext(ctx context.Context) LookupMarkupRuleResultOutput {
	return o
}

func (o LookupMarkupRuleResultOutput) CustomerDetails() CustomerMetadataResponseOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) CustomerMetadataResponse { return v.CustomerDetails }).(CustomerMetadataResponseOutput)
}

func (o LookupMarkupRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupMarkupRuleResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupMarkupRuleResultOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o LookupMarkupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMarkupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMarkupRuleResultOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v LookupMarkupRuleResult) float64 { return v.Percentage }).(pulumi.Float64Output)
}

func (o LookupMarkupRuleResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o LookupMarkupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMarkupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMarkupRuleResultOutput{})
}
