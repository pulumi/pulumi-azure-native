


package v20160201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupResourceGroup(ctx *pulumi.Context, args *LookupResourceGroupArgs, opts ...pulumi.InvokeOption) (*LookupResourceGroupResult, error) {
	var rv LookupResourceGroupResult
	err := ctx.Invoke("azure-native:resources/v20160201:getResourceGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupResourceGroupResult struct {
	Id         string                          `pulumi:"id"`
	Location   string                          `pulumi:"location"`
	Name       *string                         `pulumi:"name"`
	Properties ResourceGroupPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string               `pulumi:"tags"`
}
