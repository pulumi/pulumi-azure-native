


package education

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStudent(ctx *pulumi.Context, args *LookupStudentArgs, opts ...pulumi.InvokeOption) (*LookupStudentResult, error) {
	var rv LookupStudentResult
	err := ctx.Invoke("azure-native:education:getStudent", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStudentArgs struct {
	BillingAccountName string `pulumi:"billingAccountName"`
	BillingProfileName string `pulumi:"billingProfileName"`
	InvoiceSectionName string `pulumi:"invoiceSectionName"`
	StudentAlias       string `pulumi:"studentAlias"`
}


type LookupStudentResult struct {
	Budget                         AmountResponse     `pulumi:"budget"`
	EffectiveDate                  string             `pulumi:"effectiveDate"`
	Email                          string             `pulumi:"email"`
	ExpirationDate                 string             `pulumi:"expirationDate"`
	FirstName                      string             `pulumi:"firstName"`
	Id                             string             `pulumi:"id"`
	LastName                       string             `pulumi:"lastName"`
	Name                           string             `pulumi:"name"`
	Role                           string             `pulumi:"role"`
	Status                         string             `pulumi:"status"`
	SubscriptionAlias              *string            `pulumi:"subscriptionAlias"`
	SubscriptionId                 string             `pulumi:"subscriptionId"`
	SubscriptionInviteLastSentDate *string            `pulumi:"subscriptionInviteLastSentDate"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Type                           string             `pulumi:"type"`
}

func LookupStudentOutput(ctx *pulumi.Context, args LookupStudentOutputArgs, opts ...pulumi.InvokeOption) LookupStudentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStudentResult, error) {
			args := v.(LookupStudentArgs)
			r, err := LookupStudent(ctx, &args, opts...)
			var s LookupStudentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStudentResultOutput)
}

type LookupStudentOutputArgs struct {
	BillingAccountName pulumi.StringInput `pulumi:"billingAccountName"`
	BillingProfileName pulumi.StringInput `pulumi:"billingProfileName"`
	InvoiceSectionName pulumi.StringInput `pulumi:"invoiceSectionName"`
	StudentAlias       pulumi.StringInput `pulumi:"studentAlias"`
}

func (LookupStudentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudentArgs)(nil)).Elem()
}


type LookupStudentResultOutput struct{ *pulumi.OutputState }

func (LookupStudentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStudentResult)(nil)).Elem()
}

func (o LookupStudentResultOutput) ToLookupStudentResultOutput() LookupStudentResultOutput {
	return o
}

func (o LookupStudentResultOutput) ToLookupStudentResultOutputWithContext(ctx context.Context) LookupStudentResultOutput {
	return o
}

func (o LookupStudentResultOutput) Budget() AmountResponseOutput {
	return o.ApplyT(func(v LookupStudentResult) AmountResponse { return v.Budget }).(AmountResponseOutput)
}

func (o LookupStudentResultOutput) EffectiveDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.EffectiveDate }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) ExpirationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.ExpirationDate }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.FirstName }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.LastName }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Role }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) SubscriptionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudentResult) *string { return v.SubscriptionAlias }).(pulumi.StringPtrOutput)
}

func (o LookupStudentResultOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o LookupStudentResultOutput) SubscriptionInviteLastSentDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStudentResult) *string { return v.SubscriptionInviteLastSentDate }).(pulumi.StringPtrOutput)
}

func (o LookupStudentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStudentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStudentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStudentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStudentResultOutput{})
}
