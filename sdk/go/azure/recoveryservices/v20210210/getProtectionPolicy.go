


package v20210210

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProtectionPolicy(ctx *pulumi.Context, args *LookupProtectionPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProtectionPolicyResult, error) {
	var rv LookupProtectionPolicyResult
	err := ctx.Invoke("azure-native:recoveryservices/v20210210:getProtectionPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProtectionPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VaultName         string `pulumi:"vaultName"`
}


type LookupProtectionPolicyResult struct {
	ETag       *string           `pulumi:"eTag"`
	Id         string            `pulumi:"id"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
