


package v20161201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicyDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionAtManagementGroupResult, error) {
	var rv LookupPolicyDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20161201:getPolicyDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyDefinitionAtManagementGroupArgs struct {
	ManagementGroupId    string `pulumi:"managementGroupId"`
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionAtManagementGroupResult struct {
	Description *string     `pulumi:"description"`
	DisplayName *string     `pulumi:"displayName"`
	Id          string      `pulumi:"id"`
	Metadata    interface{} `pulumi:"metadata"`
	Mode        *string     `pulumi:"mode"`
	Name        string      `pulumi:"name"`
	Parameters  interface{} `pulumi:"parameters"`
	PolicyRule  interface{} `pulumi:"policyRule"`
	PolicyType  *string     `pulumi:"policyType"`
}
