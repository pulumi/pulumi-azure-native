


package v20190123preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:desktopvirtualization/v20190123preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	ApplicationGroupReferences []string          `pulumi:"applicationGroupReferences"`
	Description                *string           `pulumi:"description"`
	FriendlyName               *string           `pulumi:"friendlyName"`
	Id                         string            `pulumi:"id"`
	Location                   string            `pulumi:"location"`
	Name                       string            `pulumi:"name"`
	Tags                       map[string]string `pulumi:"tags"`
	Type                       string            `pulumi:"type"`
}
