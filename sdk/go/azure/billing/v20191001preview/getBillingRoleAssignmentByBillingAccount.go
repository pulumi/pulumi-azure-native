


package v20191001preview

import (
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
