


package v20210501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupEdgeModule(ctx *pulumi.Context, args *LookupEdgeModuleArgs, opts ...pulumi.InvokeOption) (*LookupEdgeModuleResult, error) {
	var rv LookupEdgeModuleResult
	err := ctx.Invoke("azure-native:videoanalyzer/v20210501preview:getEdgeModule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEdgeModuleArgs struct {
	AccountName       string `pulumi:"accountName"`
	EdgeModuleName    string `pulumi:"edgeModuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupEdgeModuleResult struct {
	EdgeModuleId string             `pulumi:"edgeModuleId"`
	Id           string             `pulumi:"id"`
	Name         string             `pulumi:"name"`
	SystemData   SystemDataResponse `pulumi:"systemData"`
	Type         string             `pulumi:"type"`
}
