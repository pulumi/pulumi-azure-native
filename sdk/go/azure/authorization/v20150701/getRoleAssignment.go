


package v20150701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRoleAssignment(ctx *pulumi.Context, args *LookupRoleAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupRoleAssignmentResult, error) {
	var rv LookupRoleAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20150701:getRoleAssignment", args, &rv, opts...)
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
	Id         string                                    `pulumi:"id"`
	Name       string                                    `pulumi:"name"`
	Properties RoleAssignmentPropertiesWithScopeResponse `pulumi:"properties"`
	Type       string                                    `pulumi:"type"`
}
