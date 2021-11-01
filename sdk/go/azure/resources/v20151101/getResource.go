


package v20151101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	var rv LookupResourceResult
	err := ctx.Invoke("azure-native:resources/v20151101:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	ParentResourcePath        string `pulumi:"parentResourcePath"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	ResourceName              string `pulumi:"resourceName"`
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	ResourceType              string `pulumi:"resourceType"`
}


type LookupResourceResult struct {
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Plan       *PlanResponse     `pulumi:"plan"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
