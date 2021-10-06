


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProductPolicy(ctx *pulumi.Context, args *LookupProductPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProductPolicyResult, error) {
	var rv LookupProductPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getProductPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductPolicyArgs struct {
	PolicyId          string `pulumi:"policyId"`
	ProductId         string `pulumi:"productId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupProductPolicyResult struct {
	Id            string `pulumi:"id"`
	Name          string `pulumi:"name"`
	PolicyContent string `pulumi:"policyContent"`
	Type          string `pulumi:"type"`
}
