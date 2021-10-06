


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210701:getWorkspace", args, &rv, opts...)
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
	AllowPublicAccessWhenBehindVnet *bool                                    `pulumi:"allowPublicAccessWhenBehindVnet"`
	ApplicationInsights             *string                                  `pulumi:"applicationInsights"`
	ContainerRegistry               *string                                  `pulumi:"containerRegistry"`
	Description                     *string                                  `pulumi:"description"`
	DiscoveryUrl                    *string                                  `pulumi:"discoveryUrl"`
	Encryption                      *EncryptionPropertyResponse              `pulumi:"encryption"`
	FriendlyName                    *string                                  `pulumi:"friendlyName"`
	HbiWorkspace                    *bool                                    `pulumi:"hbiWorkspace"`
	Id                              string                                   `pulumi:"id"`
	Identity                        *IdentityResponse                        `pulumi:"identity"`
	ImageBuildCompute               *string                                  `pulumi:"imageBuildCompute"`
	KeyVault                        *string                                  `pulumi:"keyVault"`
	Location                        *string                                  `pulumi:"location"`
	MlFlowTrackingUri               string                                   `pulumi:"mlFlowTrackingUri"`
	Name                            string                                   `pulumi:"name"`
	NotebookInfo                    NotebookResourceInfoResponse             `pulumi:"notebookInfo"`
	PrimaryUserAssignedIdentity     *string                                  `pulumi:"primaryUserAssignedIdentity"`
	PrivateEndpointConnections      []PrivateEndpointConnectionResponse      `pulumi:"privateEndpointConnections"`
	PrivateLinkCount                int                                      `pulumi:"privateLinkCount"`
	ProvisioningState               string                                   `pulumi:"provisioningState"`
	PublicNetworkAccess             *string                                  `pulumi:"publicNetworkAccess"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettingsResponse `pulumi:"serviceManagedResourcesSettings"`
	ServiceProvisionedResourceGroup string                                   `pulumi:"serviceProvisionedResourceGroup"`
	SharedPrivateLinkResources      []SharedPrivateLinkResourceResponse      `pulumi:"sharedPrivateLinkResources"`
	Sku                             *SkuResponse                             `pulumi:"sku"`
	StorageAccount                  *string                                  `pulumi:"storageAccount"`
	StorageHnsEnabled               bool                                     `pulumi:"storageHnsEnabled"`
	SystemData                      SystemDataResponse                       `pulumi:"systemData"`
	Tags                            map[string]string                        `pulumi:"tags"`
	TenantId                        string                                   `pulumi:"tenantId"`
	Type                            string                                   `pulumi:"type"`
	WorkspaceId                     string                                   `pulumi:"workspaceId"`
}
