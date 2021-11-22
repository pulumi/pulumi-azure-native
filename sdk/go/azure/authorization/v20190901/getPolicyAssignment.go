


package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyAssignment(ctx *pulumi.Context, args *LookupPolicyAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupPolicyAssignmentResult, error) {
	var rv LookupPolicyAssignmentResult
	err := ctx.Invoke("azure-native:authorization/v20190901:getPolicyAssignment", args, &rv, opts...)
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
	Description        *string                                 `pulumi:"description"`
	DisplayName        *string                                 `pulumi:"displayName"`
	EnforcementMode    *string                                 `pulumi:"enforcementMode"`
	Id                 string                                  `pulumi:"id"`
	Identity           *IdentityResponse                       `pulumi:"identity"`
	Location           *string                                 `pulumi:"location"`
	Metadata           interface{}                             `pulumi:"metadata"`
	Name               string                                  `pulumi:"name"`
	NotScopes          []string                                `pulumi:"notScopes"`
	Parameters         map[string]ParameterValuesValueResponse `pulumi:"parameters"`
	PolicyDefinitionId *string                                 `pulumi:"policyDefinitionId"`
	Scope              *string                                 `pulumi:"scope"`
	Sku                *PolicySkuResponse                      `pulumi:"sku"`
	Type               string                                  `pulumi:"type"`
}
