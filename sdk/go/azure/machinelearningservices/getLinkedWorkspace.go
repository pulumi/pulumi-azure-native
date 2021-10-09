


package machinelearningservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLinkedWorkspace(ctx *pulumi.Context, args *LookupLinkedWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupLinkedWorkspaceResult, error) {
	var rv LookupLinkedWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices:getLinkedWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLinkedWorkspaceArgs struct {
	LinkName          string `pulumi:"linkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupLinkedWorkspaceResult struct {
	Id         string                       `pulumi:"id"`
	Name       string                       `pulumi:"name"`
	Properties LinkedWorkspacePropsResponse `pulumi:"properties"`
	Type       string                       `pulumi:"type"`
}
