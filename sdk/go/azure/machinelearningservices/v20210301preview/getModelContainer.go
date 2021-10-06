


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupModelContainer(ctx *pulumi.Context, args *LookupModelContainerArgs, opts ...pulumi.InvokeOption) (*LookupModelContainerResult, error) {
	var rv LookupModelContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getModelContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupModelContainerResult struct {
	Id         string                 `pulumi:"id"`
	Name       string                 `pulumi:"name"`
	Properties ModelContainerResponse `pulumi:"properties"`
	SystemData SystemDataResponse     `pulumi:"systemData"`
	Type       string                 `pulumi:"type"`
}
