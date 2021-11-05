


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperationPolicy(ctx *pulumi.Context, args *LookupApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationPolicyResult, error) {
	var rv LookupApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20210801:getApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiOperationPolicyArgs struct {
	ApiId             string  `pulumi:"apiId"`
	Format            *string `pulumi:"format"`
	OperationId       string  `pulumi:"operationId"`
	PolicyId          string  `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupApiOperationPolicyResult struct {
	Format *string `pulumi:"format"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Type   string  `pulumi:"type"`
	Value  string  `pulumi:"value"`
}
