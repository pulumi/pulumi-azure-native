


package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProductPolicy(ctx *pulumi.Context, args *LookupProductPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProductPolicyResult, error) {
	var rv LookupProductPolicyResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getProductPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductPolicyArgs struct {
	Format            *string `pulumi:"format"`
	PolicyId          string  `pulumi:"policyId"`
	ProductId         string  `pulumi:"productId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ServiceName       string  `pulumi:"serviceName"`
}


type LookupProductPolicyResult struct {
	Format *string `pulumi:"format"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Type   string  `pulumi:"type"`
	Value  string  `pulumi:"value"`
}
