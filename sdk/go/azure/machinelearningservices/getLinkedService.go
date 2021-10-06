


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedService(ctx *pulumi.Context, args *LookupLinkedServiceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedServiceResult, error) {
	var rv LookupLinkedServiceResult
	err := ctx.Invoke("azure-native:machinelearningservices:getLinkedService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedServiceArgs struct {
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedServiceResult struct {
	Id         string                     `pulumi:"id"`
	Identity   *IdentityResponse          `pulumi:"identity"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Properties LinkedServicePropsResponse `pulumi:"properties"`
	Type       string                     `pulumi:"type"`
}
