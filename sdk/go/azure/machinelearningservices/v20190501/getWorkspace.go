


package v20190501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20190501:getWorkspace", args, &rv, opts...)
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
	ApplicationInsights *string           `pulumi:"applicationInsights"`
	ContainerRegistry   *string           `pulumi:"containerRegistry"`
	CreationTime        string            `pulumi:"creationTime"`
	Description         *string           `pulumi:"description"`
	DiscoveryUrl        *string           `pulumi:"discoveryUrl"`
	FriendlyName        *string           `pulumi:"friendlyName"`
	Id                  string            `pulumi:"id"`
	Identity            *IdentityResponse `pulumi:"identity"`
	KeyVault            *string           `pulumi:"keyVault"`
	Location            *string           `pulumi:"location"`
	Name                string            `pulumi:"name"`
	ProvisioningState   string            `pulumi:"provisioningState"`
	StorageAccount      *string           `pulumi:"storageAccount"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
	WorkspaceId         string            `pulumi:"workspaceId"`
}
