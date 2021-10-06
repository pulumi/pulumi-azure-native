


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20201001preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	RoleAssignmentName string  `pulumi:"roleAssignmentName"`
	Scope              string  `pulumi:"scope"`
	TenantId           *string `pulumi:"tenantId"`
}


type LookupRoleAssignmentResult struct {
	Condition                          *string `pulumi:"condition"`
	ConditionVersion                   *string `pulumi:"conditionVersion"`
	CreatedBy                          string  `pulumi:"createdBy"`
	CreatedOn                          string  `pulumi:"createdOn"`
	DelegatedManagedIdentityResourceId *string `pulumi:"delegatedManagedIdentityResourceId"`
	Description                        *string `pulumi:"description"`
	Id                                 string  `pulumi:"id"`
	Name                               string  `pulumi:"name"`
	PrincipalId                        string  `pulumi:"principalId"`
	PrincipalType                      *string `pulumi:"principalType"`
	RoleDefinitionId                   string  `pulumi:"roleDefinitionId"`
	Scope                              string  `pulumi:"scope"`
	Type                               string  `pulumi:"type"`
	UpdatedBy                          string  `pulumi:"updatedBy"`
	UpdatedOn                          string  `pulumi:"updatedOn"`
}
