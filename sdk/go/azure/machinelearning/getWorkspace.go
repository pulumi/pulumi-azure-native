


package machinelearning

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearning:getWorkspace", args, &rv, opts...)
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
	CreationTime         string            `pulumi:"creationTime"`
	Id                   string            `pulumi:"id"`
	KeyVaultIdentifierId *string           `pulumi:"keyVaultIdentifierId"`
	Location             string            `pulumi:"location"`
	Name                 string            `pulumi:"name"`
	OwnerEmail           string            `pulumi:"ownerEmail"`
	StudioEndpoint       string            `pulumi:"studioEndpoint"`
	Tags                 map[string]string `pulumi:"tags"`
	Type                 string            `pulumi:"type"`
	UserStorageAccountId string            `pulumi:"userStorageAccountId"`
	WorkspaceId          string            `pulumi:"workspaceId"`
	WorkspaceState       string            `pulumi:"workspaceState"`
	WorkspaceType        string            `pulumi:"workspaceType"`
}
