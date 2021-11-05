


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataContainer(ctx *pulumi.Context, args *LookupDataContainerArgs, opts ...pulumi.InvokeOption) (*LookupDataContainerResult, error) {
	var rv LookupDataContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices:getDataContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDataContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupDataContainerResult struct {
	Id         string                `pulumi:"id"`
	Name       string                `pulumi:"name"`
	Properties DataContainerResponse `pulumi:"properties"`
	SystemData SystemDataResponse    `pulumi:"systemData"`
	Type       string                `pulumi:"type"`
}
