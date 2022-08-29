


package billing

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBillingRoleAssignmentByDepartment(ctx *pulumi.Context, args *LookupBillingRoleAssignmentByDepartmentArgs, opts ...pulumi.InvokeOption) (*LookupBillingRoleAssignmentByDepartmentResult, error) {
	var rv LookupBillingRoleAssignmentByDepartmentResult
	err := ctx.Invoke("azure-native:billing:getBillingRoleAssignmentByDepartment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBillingRoleAssignmentByDepartmentArgs struct {
	BillingAccountName        string `pulumi:"billingAccountName"`
	BillingRoleAssignmentName string `pulumi:"billingRoleAssignmentName"`
	DepartmentName            string `pulumi:"departmentName"`
}


type LookupBillingRoleAssignmentByDepartmentResult struct {
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

func LookupBillingRoleAssignmentByDepartmentOutput(ctx *pulumi.Context, args LookupBillingRoleAssignmentByDepartmentOutputArgs, opts ...pulumi.InvokeOption) LookupBillingRoleAssignmentByDepartmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBillingRoleAssignmentByDepartmentResult, error) {
			args := v.(LookupBillingRoleAssignmentByDepartmentArgs)
			r, err := LookupBillingRoleAssignmentByDepartment(ctx, &args, opts...)
			var s LookupBillingRoleAssignmentByDepartmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBillingRoleAssignmentByDepartmentResultOutput)
}

type LookupBillingRoleAssignmentByDepartmentOutputArgs struct {
	BillingAccountName        pulumi.StringInput `pulumi:"billingAccountName"`
	BillingRoleAssignmentName pulumi.StringInput `pulumi:"billingRoleAssignmentName"`
	DepartmentName            pulumi.StringInput `pulumi:"departmentName"`
}

func (LookupBillingRoleAssignmentByDepartmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByDepartmentArgs)(nil)).Elem()
}


type LookupBillingRoleAssignmentByDepartmentResultOutput struct{ *pulumi.OutputState }

func (LookupBillingRoleAssignmentByDepartmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBillingRoleAssignmentByDepartmentResult)(nil)).Elem()
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) ToLookupBillingRoleAssignmentByDepartmentResultOutput() LookupBillingRoleAssignmentByDepartmentResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) ToLookupBillingRoleAssignmentByDepartmentResultOutputWithContext(ctx context.Context) LookupBillingRoleAssignmentByDepartmentResultOutput {
	return o
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByPrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByPrincipalId }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByPrincipalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByPrincipalTenantId }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedByUserEmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedByUserEmailAddress }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) PrincipalTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.PrincipalTenantId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Scope }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) UserAuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.UserAuthenticationType }).(pulumi.StringPtrOutput)
}

func (o LookupBillingRoleAssignmentByDepartmentResultOutput) UserEmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBillingRoleAssignmentByDepartmentResult) *string { return v.UserEmailAddress }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBillingRoleAssignmentByDepartmentResultOutput{})
}
