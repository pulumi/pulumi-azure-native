


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiOperationPolicy(ctx *pulumi.Context, args *LookupApiOperationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiOperationPolicyResult, error) {
	var rv LookupApiOperationPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiOperationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiOperationPolicyArgs struct {
	ApiId             string `pulumi:"apiId"`
	OperationId       string `pulumi:"operationId"`
	PolicyId          string `pulumi:"policyId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiOperationPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupApiOperationPolicyResult) Defaults() *LookupApiOperationPolicyResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ContentFormat) {
		contentFormat_ := "xml"
		tmp.ContentFormat = &contentFormat_
	}
	return &tmp
}
