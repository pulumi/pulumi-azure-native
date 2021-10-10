


package v20201001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleManagementPolicyAssignment(ctx *pulumi.Context, args *LookupRoleManagementPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleManagementPolicyAssignmentResult, error) {
	var rv LookupRoleManagementPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20201001preview:getRoleManagementPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleManagementPolicyAssignmentArgs struct {
	RoleManagementPolicyAssignmentName string `pulumi:"roleManagementPolicyAssignmentName"`
	Scope                              string `pulumi:"scope"`
}


type LookupRoleManagementPolicyAssignmentResult struct {
	Id                         string                             `pulumi:"id"`
	Name                       string                             `pulumi:"name"`
	PolicyAssignmentProperties PolicyAssignmentPropertiesResponse `pulumi:"policyAssignmentProperties"`
	PolicyId                   *string                            `pulumi:"policyId"`
	RoleDefinitionId           *string                            `pulumi:"roleDefinitionId"`
	Scope                      *string                            `pulumi:"scope"`
	Type                       string                             `pulumi:"type"`
}
