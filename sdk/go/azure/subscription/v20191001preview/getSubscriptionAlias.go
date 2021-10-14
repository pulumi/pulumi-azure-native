


package v20191001preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscriptionAlias(ctx *pulumi.Context, args *LookupSubscriptionAliasArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionAliasResult, error) {
	var rv LookupSubscriptionAliasResult
	err := ctx.Invoke("azure-native:subscription/v20191001preview:getSubscriptionAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionAliasArgs struct {
	AliasName string `pulumi:"aliasName"`
}


type LookupSubscriptionAliasResult struct {
	Id         string                             `pulumi:"id"`
	Name       string                             `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponse `pulumi:"properties"`
	Type       string                             `pulumi:"type"`
}
