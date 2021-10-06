


package v20151001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20151001preview:getPolicyAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyAssignmentArgs struct {
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	Scope                string `pulumi:"scope"`
}


type LookupPolicyAssignmentResult struct {
	DisplayName        *string `pulumi:"displayName"`
	Id                 *string `pulumi:"id"`
	Name               *string `pulumi:"name"`
	PolicyDefinitionId *string `pulumi:"policyDefinitionId"`
	Scope              *string `pulumi:"scope"`
	Type               *string `pulumi:"type"`
}
