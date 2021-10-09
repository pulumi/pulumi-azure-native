


package v20190101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicySetDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicySetDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionAtManagementGroupResult, error) {
	var rv LookupPolicySetDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20190101:getPolicySetDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionAtManagementGroupArgs struct {
	ManagementGroupId       string `pulumi:"managementGroupId"`
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}


type LookupPolicySetDefinitionAtManagementGroupResult struct {
	Description       *string                             `pulumi:"description"`
	DisplayName       *string                             `pulumi:"displayName"`
	Id                string                              `pulumi:"id"`
	Metadata          interface{}                         `pulumi:"metadata"`
	Name              string                              `pulumi:"name"`
	Parameters        interface{}                         `pulumi:"parameters"`
	PolicyDefinitions []PolicyDefinitionReferenceResponse `pulumi:"policyDefinitions"`
	PolicyType        *string                             `pulumi:"policyType"`
	Type              string                              `pulumi:"type"`
}
