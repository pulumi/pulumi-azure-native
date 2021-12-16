


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperationPolicy(ctx *pulumi.Context, args *LookupApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationPolicyResult, error) {
	var rv LookupApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20210401preview:getApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupApiOperationPolicyResult) Defaults() *LookupApiOperationPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Format) {
		format_ := "xml"
		tmp.Format = &format_
	}
	return &tmp
}
