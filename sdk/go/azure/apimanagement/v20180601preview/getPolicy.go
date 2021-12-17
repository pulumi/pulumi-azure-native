


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPolicy(ctx *pulumi.Context, args *LookupPolicyArgs, opts ...pulumi.InvokeOption) (*LookupPolicyResult, error) {
	var rv LookupPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPolicyArgs struct {
	PolicyId          string `pulumi:"policyId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupPolicyResult struct {
	ContentFormat *string `pulumi:"contentFormat"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyContent string  `pulumi:"policyContent"`
	Type          string  `pulumi:"type"`
}


func (val *LookupPolicyResult) Defaults() *LookupPolicyResult {
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
