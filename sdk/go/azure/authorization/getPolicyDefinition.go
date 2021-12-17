


package authorization

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyDefinition(ctx *pulumi.Context, args *LookupPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionResult, error) {
	var rv LookupPolicyDefinitionResult
	err := ctx.Invoke("azure-native:authorization:getPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyDefinitionArgs struct {
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionResult struct {
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


func (val *LookupPolicyDefinitionResult) Defaults() *LookupPolicyDefinitionResult {
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
