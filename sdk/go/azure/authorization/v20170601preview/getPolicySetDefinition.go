


package v20170601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicySetDefinition(ctx *pulumi.Context, args *LookupPolicySetDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetDefinitionResult, error) {
	var rv LookupPolicySetDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20170601preview:getPolicySetDefinition", args, &rv, opts...)
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
