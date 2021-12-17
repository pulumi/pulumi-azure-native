


package v20180101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiPolicy(ctx *pulumi.Context, args *LookupApiPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiPolicyResult, error) {
	var rv LookupApiPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPolicyArgs struct {
	ApiId             string `pulumi:"apiId"`
	PolicyId          string `pulumi:"policyId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupApiPolicyResult) Defaults() *LookupApiPolicyResult {
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
