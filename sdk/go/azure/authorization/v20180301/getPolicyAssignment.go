


package v20180301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20180301:getPolicyAssignment", args, &rv, opts...)
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
	Description        *string            `pulumi:"description"`
	DisplayName        *string            `pulumi:"displayName"`
	Id                 string             `pulumi:"id"`
	Metadata           interface{}        `pulumi:"metadata"`
	Name               string             `pulumi:"name"`
	NotScopes          []string           `pulumi:"notScopes"`
	Parameters         interface{}        `pulumi:"parameters"`
	PolicyDefinitionId *string            `pulumi:"policyDefinitionId"`
	Scope              *string            `pulumi:"scope"`
	Sku                *PolicySkuResponse `pulumi:"sku"`
	Type               string             `pulumi:"type"`
}
