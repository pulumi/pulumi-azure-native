


package education

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLab(ctx *pulumi.Context, args *LookupLabArgs, opts ...pulumi.InvokeOption) (*LookupLabResult, error) {
	var rv LookupLabResult
	err := ctx.Invoke("azure-native:education:getLab", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLabArgs struct {
	BillingAccountName string `pulumi:"billingAccountName"`
	BillingProfileName string `pulumi:"billingProfileName"`
	IncludeBudget      *bool  `pulumi:"includeBudget"`
	InvoiceSectionName string `pulumi:"invoiceSectionName"`
}


type LookupLabResult struct {
	BudgetPerStudent AmountResponse     `pulumi:"budgetPerStudent"`
	Currency         *string            `pulumi:"currency"`
	Description      string             `pulumi:"description"`
	DisplayName      string             `pulumi:"displayName"`
	EffectiveDate    string             `pulumi:"effectiveDate"`
	ExpirationDate   string             `pulumi:"expirationDate"`
	Id               string             `pulumi:"id"`
	InvitationCode   string             `pulumi:"invitationCode"`
	MaxStudentCount  float64            `pulumi:"maxStudentCount"`
	Name             string             `pulumi:"name"`
	Status           string             `pulumi:"status"`
	SystemData       SystemDataResponse `pulumi:"systemData"`
	Type             string             `pulumi:"type"`
	Value            *float64           `pulumi:"value"`
}

func LookupLabOutput(ctx *pulumi.Context, args LookupLabOutputArgs, opts ...pulumi.InvokeOption) LookupLabResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLabResult, error) {
			args := v.(LookupLabArgs)
			r, err := LookupLab(ctx, &args, opts...)
			var s LookupLabResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLabResultOutput)
}

type LookupLabOutputArgs struct {
	BillingAccountName pulumi.StringInput  `pulumi:"billingAccountName"`
	BillingProfileName pulumi.StringInput  `pulumi:"billingProfileName"`
	IncludeBudget      pulumi.BoolPtrInput `pulumi:"includeBudget"`
	InvoiceSectionName pulumi.StringInput  `pulumi:"invoiceSectionName"`
}

func (LookupLabOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabArgs)(nil)).Elem()
}


type LookupLabResultOutput struct{ *pulumi.OutputState }

func (LookupLabResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLabResult)(nil)).Elem()
}

func (o LookupLabResultOutput) ToLookupLabResultOutput() LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) ToLookupLabResultOutputWithContext(ctx context.Context) LookupLabResultOutput {
	return o
}

func (o LookupLabResultOutput) BudgetPerStudent() AmountResponseOutput {
	return o.ApplyT(func(v LookupLabResult) AmountResponse { return v.BudgetPerStudent }).(AmountResponseOutput)
}

func (o LookupLabResultOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLabResult) *string { return v.Currency }).(pulumi.StringPtrOutput)
}

func (o LookupLabResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.EffectiveDate }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) InvitationCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.InvitationCode }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) MaxStudentCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupLabResult) float64 { return v.MaxStudentCount }).(pulumi.Float64Output)
}

func (o LookupLabResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLabResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLabResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLabResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLabResultOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupLabResult) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLabResultOutput{})
}
