


package v20201201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiPolicy(ctx *pulumi.Context, args *LookupApiPolicyArgs, opts ...pulumi.InvokeOption) (*LookupApiPolicyResult, error) {
	var rv LookupApiPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getApiPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApiPolicyArgs struct {
	ApiId             string  `pulumi:"apiId"`
	Format            *string `pulumi:"format"`
	PolicyId          string  `pulumi:"policyId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupApiPolicyResult struct {
	Format *string `pulumi:"format"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Type   string  `pulumi:"type"`
	Value  string  `pulumi:"value"`
}


func (val *LookupApiPolicyResult) Defaults() *LookupApiPolicyResult {
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
