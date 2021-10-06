


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlias(ctx *pulumi.Context, args *LookupAliasArgs, opts ...pulumi.InvokeOption) (*LookupAliasResult, error) {
	var rv LookupAliasResult
	err := ctx.Invoke("azure-native:subscription/v20200901:getAlias", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAliasArgs struct {
	AliasName string `pulumi:"aliasName"`
}


type LookupAliasResult struct {
	Id         string                             `pulumi:"id"`
	Name       string                             `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponse `pulumi:"properties"`
	Type       string                             `pulumi:"type"`
}
