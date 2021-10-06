


package billing

import (
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
