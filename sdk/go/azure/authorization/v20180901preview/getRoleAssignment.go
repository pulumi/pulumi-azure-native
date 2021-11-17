


package v20180901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20180901preview:getRoleAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRoleAssignmentArgs struct {
	RoleAssignmentName string `pulumi:"roleAssignmentName"`
	Scope              string `pulumi:"scope"`
}


type LookupRoleAssignmentResult struct {
	CanDelegate      *bool   `pulumi:"canDelegate"`
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	PrincipalId      *string `pulumi:"principalId"`
	PrincipalType    *string `pulumi:"principalType"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
	Type             string  `pulumi:"type"`
}
