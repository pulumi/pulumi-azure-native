


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBillingRoleAssignmentByEnrollmentAccount(ctx *pulumi.Context, args *LookupBillingRoleAssignmentByEnrollmentAccountArgs, opts ...pulumi.InvokeOption) (*LookupBillingRoleAssignmentByEnrollmentAccountResult, error) {
	var rv LookupBillingRoleAssignmentByEnrollmentAccountResult
	err := ctx.Invoke("azure-native:billing/v20191001preview:getBillingRoleAssignmentByEnrollmentAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingRoleAssignmentByEnrollmentAccountArgs struct {
	BillingAccountName        string `pulumi:"billingAccountName"`
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
	EnrollmentAccountName     string `pulumi:"enrollmentAccountName"`
}


type LookupBillingRoleAssignmentByEnrollmentAccountResult struct {
	CreatedByPrincipalId       string  `pulumi:"createdByPrincipalId"`
	CreatedByPrincipalTenantId string  `pulumi:"createdByPrincipalTenantId"`
	CreatedByUserEmailAddress  string  `pulumi:"createdByUserEmailAddress"`
	CreatedOn                  string  `pulumi:"createdOn"`
	Id                         string  `pulumi:"id"`
	Name                       string  `pulumi:"name"`
	PrincipalId                *string `pulumi:"principalId"`
	PrincipalTenantId          *string `pulumi:"principalTenantId"`
	RoleDefinitionId           *string `pulumi:"roleDefinitionId"`
	Scope                      string  `pulumi:"scope"`
	Type                       string  `pulumi:"type"`
	UserAuthenticationType     *string `pulumi:"userAuthenticationType"`
	UserEmailAddress           *string `pulumi:"userEmailAddress"`
}

func LookupBillingRoleAssignmentByEnrollmentAccountOutput(ctx *pulumi.Context, args LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs, opts ...pulumi.InvokeOption) LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingRoleAssignmentByEnrollmentAccountResult, error) {
			args := v.(LookupBillingRoleAssignmentByEnrollmentAccountArgs)
			r, err := LookupBillingRoleAssignmentByEnrollmentAccount(ctx, &args, opts...)
			var s LookupBillingRoleAssignmentByEnrollmentAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingRoleAssignmentByEnrollmentAccountResultOutput)
}

type LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs struct {
	BillingAccountName        pulumi.StringInput `pulumi:"billingAccountName"`
	BillingRoleAssignmentName pulumi.StringInput `pulumi:"billingRoleAssignmentName"`
	EnrollmentAccountName     pulumi.StringInput `pulumi:"enrollmentAccountName"`
}

func (LookupBillingRoleAssignmentByEnrollmentAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByEnrollmentAccountArgs)(nil)).Elem()
}


type LookupBillingRoleAssignmentByEnrollmentAccountResultOutput struct{ *pulumi.OutputState }

func (LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByEnrollmentAccountResult)(nil)).Elem()
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ToLookupBillingRoleAssignmentByEnrollmentAccountResultOutput() LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) ToLookupBillingRoleAssignmentByEnrollmentAccountResultOutputWithContext(ctx context.Context) LookupBillingRoleAssignmentByEnrollmentAccountResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string {
		return v.CreatedByPrincipalTenantId
	}).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string {
		return v.CreatedByUserEmailAddress
	}).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByEnrollmentAccountResultOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByEnrollmentAccountResult) *string { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingRoleAssignmentByEnrollmentAccountResultOutput{})
}
