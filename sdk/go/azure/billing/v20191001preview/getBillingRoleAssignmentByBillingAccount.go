


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBillingRoleAssignmentByBillingAccount(ctx *pulumi.Context, args *LookupBillingRoleAssignmentByBillingAccountArgs, opts ...pulumi.InvokeOption) (*LookupBillingRoleAssignmentByBillingAccountResult, error) {
	var rv LookupBillingRoleAssignmentByBillingAccountResult
	err := ctx.Invoke("azure-native:billing/v20191001preview:getBillingRoleAssignmentByBillingAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingRoleAssignmentByBillingAccountArgs struct {
	BillingAccountName        string `pulumi:"billingAccountName"`
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
}


type LookupBillingRoleAssignmentByBillingAccountResult struct {
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

func LookupBillingRoleAssignmentByBillingAccountOutput(ctx *pulumi.Context, args LookupBillingRoleAssignmentByBillingAccountOutputArgs, opts ...pulumi.InvokeOption) LookupBillingRoleAssignmentByBillingAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingRoleAssignmentByBillingAccountResult, error) {
			args := v.(LookupBillingRoleAssignmentByBillingAccountArgs)
			r, err := LookupBillingRoleAssignmentByBillingAccount(ctx, &args, opts...)
			var s LookupBillingRoleAssignmentByBillingAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingRoleAssignmentByBillingAccountResultOutput)
}

type LookupBillingRoleAssignmentByBillingAccountOutputArgs struct {
	BillingAccountName        pulumi.StringInput `pulumi:"billingAccountName"`
	BillingRoleAssignmentName pulumi.StringInput `pulumi:"billingRoleAssignmentName"`
}

func (LookupBillingRoleAssignmentByBillingAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByBillingAccountArgs)(nil)).Elem()
}


type LookupBillingRoleAssignmentByBillingAccountResultOutput struct{ *pulumi.OutputState }

func (LookupBillingRoleAssignmentByBillingAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByBillingAccountResult)(nil)).Elem()
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) ToLookupBillingRoleAssignmentByBillingAccountResultOutput() LookupBillingRoleAssignmentByBillingAccountResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) ToLookupBillingRoleAssignmentByBillingAccountResultOutputWithContext(ctx context.Context) LookupBillingRoleAssignmentByBillingAccountResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.CreatedByPrincipalTenantId }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) *string { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) *string { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByBillingAccountResultOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByBillingAccountResult) *string { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingRoleAssignmentByBillingAccountResultOutput{})
}
