


package v20151101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkedServiceName string `pulumi:"linkedServiceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedServiceResult struct {
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	ResourceId string `pulumi:"resourceId"`
	Type       string `pulumi:"type"`
}
