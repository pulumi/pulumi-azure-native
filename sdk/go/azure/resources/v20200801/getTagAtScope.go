


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagAtScope(ctx *pulumi.Context, args *LookupTagAtScopeArgs, opts ...pulumi.InvokeOption) (*LookupTagAtScopeResult, error) {
	var rv LookupTagAtScopeResult
	err := ctx.Invoke("azure-native:resources/v20200801:getTagAtScope", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagAtScopeArgs struct {
	Scope string `pulumi:"scope"`
}


type LookupTagAtScopeResult struct {
	Id         string       `pulumi:"id"`
	Name       string       `pulumi:"name"`
	Properties TagsResponse `pulumi:"properties"`
	Type       string       `pulumi:"type"`
}
