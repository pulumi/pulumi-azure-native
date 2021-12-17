


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProductPolicy(ctx *pulumi.Context, args *LookupProductPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProductPolicyResult, error) {
	var rv LookupProductPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getProductPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupProductPolicyArgs struct {
	PolicyId          string `pulumi:"policyId"`
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupProductPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupProductPolicyResult) Defaults() *LookupProductPolicyResult {
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
