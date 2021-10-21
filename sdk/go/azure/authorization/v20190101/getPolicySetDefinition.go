


package v20190101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicySetDefinition(ctx *pulumi.Context, args *LookupPolicySetDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionResult, error) {
	var rv LookupPolicySetDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20190101:getPolicySetDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicySetDefinitionArgs struct {
	PolicySetDefinitionName string `pulumi:"policySetDefinitionName"`
}


type LookupPolicySetDefinitionResult struct {
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
