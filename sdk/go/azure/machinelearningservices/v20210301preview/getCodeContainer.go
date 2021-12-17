


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCodeContainer(ctx *pulumi.Context, args *LookupCodeContainerArgs, opts ...pulumi.InvokeOption) (*LookupCodeContainerResult, error) {
	var rv LookupCodeContainerResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:getCodeContainer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCodeContainerArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupCodeContainerResult struct {
	Id         string                `pulumi:"id"`
	Name       string                `pulumi:"name"`
	Properties CodeContainerResponse `pulumi:"properties"`
	SystemData SystemDataResponse    `pulumi:"systemData"`
	Type       string                `pulumi:"type"`
}
