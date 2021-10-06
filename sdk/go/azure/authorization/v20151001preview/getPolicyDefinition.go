


package v20151001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicyDefinition(ctx *pulumi.Context, args *LookupPolicyDefinitionArgs, opts ...pulumi.InvokeOption) (*LookupPolicyDefinitionResult, error) {
	var rv LookupPolicyDefinitionResult
	err := ctx.Invoke("azure-native:authorization/v20151001preview:getPolicyDefinition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPolicyDefinitionArgs struct {
	PolicyDefinitionName string `pulumi:"policyDefinitionName"`
}


type LookupPolicyDefinitionResult struct {
	Description *string     `pulumi:"description"`
	DisplayName *string     `pulumi:"displayName"`
	Id          string      `pulumi:"id"`
	Name        *string     `pulumi:"name"`
	PolicyRule  interface{} `pulumi:"policyRule"`
	PolicyType  *string     `pulumi:"policyType"`
}
