


package v20200218preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200218preview:getWorkspace", args, &rv, opts...)
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
	ApplicationInsights             *string                     `pulumi:"applicationInsights"`
	ContainerRegistry               *string                     `pulumi:"containerRegistry"`
	CreationTime                    string                      `pulumi:"creationTime"`
	Description                     *string                     `pulumi:"description"`
	DiscoveryUrl                    *string                     `pulumi:"discoveryUrl"`
	Encryption                      *EncryptionPropertyResponse `pulumi:"encryption"`
	FriendlyName                    *string                     `pulumi:"friendlyName"`
	HbiWorkspace                    *bool                       `pulumi:"hbiWorkspace"`
	Id                              string                      `pulumi:"id"`
	Identity                        *IdentityResponse           `pulumi:"identity"`
	KeyVault                        *string                     `pulumi:"keyVault"`
	Location                        *string                     `pulumi:"location"`
	Name                            string                      `pulumi:"name"`
	ProvisioningState               string                      `pulumi:"provisioningState"`
	ServiceProvisionedResourceGroup string                      `pulumi:"serviceProvisionedResourceGroup"`
	Sku                             *SkuResponse                `pulumi:"sku"`
	StorageAccount                  *string                     `pulumi:"storageAccount"`
	Tags                            map[string]string           `pulumi:"tags"`
	Type                            string                      `pulumi:"type"`
	WorkspaceId                     string                      `pulumi:"workspaceId"`
}
