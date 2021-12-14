


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyDefinitionAtManagementGroup(ctx *pulumi.Context, args *LookupPolicyDefinitionAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionAtManagementGroupResult, error) {
	var rv LookupPolicyDefinitionAtManagementGroupResult
	err := ctx.Invoke("azure-native:authorization/v20200901:getPolicyDefinitionAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyDefinitionAtManagementGroupArgs struct {
	ManagementGroupId    string `pulumi:"managementGroupId"`
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionAtManagementGroupResult struct {
	Description *string                                      `pulumi:"description"`
	DisplayName *string                                      `pulumi:"displayName"`
	Id          string                                       `pulumi:"id"`
	Metadata    interface{}                                  `pulumi:"metadata"`
	Mode        *string                                      `pulumi:"mode"`
	Name        string                                       `pulumi:"name"`
	Parameters  map[string]ParameterDefinitionsValueResponse `pulumi:"parameters"`
	PolicyRule  interface{}                                  `pulumi:"policyRule"`
	PolicyType  *string                                      `pulumi:"policyType"`
	Type        string                                       `pulumi:"type"`
}


func (val *LookupPolicyDefinitionAtManagementGroupResult) Defaults() *LookupPolicyDefinitionAtManagementGroupResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "Indexed"
		tmp.Mode = &mode_
	}
	return &tmp
}
