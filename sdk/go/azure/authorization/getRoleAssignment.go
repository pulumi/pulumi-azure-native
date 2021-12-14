


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupRoleAssignmentResult) Defaults() *LookupRoleAssignmentResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PrincipalType) {
		principalType_ := "User"
		tmp.PrincipalType = &principalType_
	}
	return &tmp
}
