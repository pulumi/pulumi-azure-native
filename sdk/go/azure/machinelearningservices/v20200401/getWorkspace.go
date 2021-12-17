


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200401:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspaceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupWorkspaceResult struct {
	AllowPublicAccessWhenBehindVnet *bool                               `pulumi:"allowPublicAccessWhenBehindVnet"`
	ApplicationInsights             *string                             `pulumi:"applicationInsights"`
	ContainerRegistry               *string                             `pulumi:"containerRegistry"`
	CreationTime                    string                              `pulumi:"creationTime"`
	Description                     *string                             `pulumi:"description"`
	DiscoveryUrl                    *string                             `pulumi:"discoveryUrl"`
	Encryption                      *EncryptionPropertyResponse         `pulumi:"encryption"`
	FriendlyName                    *string                             `pulumi:"friendlyName"`
	HbiWorkspace                    *bool                               `pulumi:"hbiWorkspace"`
	Id                              string                              `pulumi:"id"`
	Identity                        *IdentityResponse                   `pulumi:"identity"`
	ImageBuildCompute               *string                             `pulumi:"imageBuildCompute"`
	KeyVault                        *string                             `pulumi:"keyVault"`
	Location                        *string                             `pulumi:"location"`
	Name                            string                              `pulumi:"name"`
	NotebookInfo                    NotebookResourceInfoResponse        `pulumi:"notebookInfo"`
	PrivateEndpointConnections      []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	PrivateLinkCount                int                                 `pulumi:"privateLinkCount"`
	ProvisioningState               string                              `pulumi:"provisioningState"`
	ServiceProvisionedResourceGroup string                              `pulumi:"serviceProvisionedResourceGroup"`
	SharedPrivateLinkResources      []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	Sku                             *SkuResponse                        `pulumi:"sku"`
	StorageAccount                  *string                             `pulumi:"storageAccount"`
	Tags                            map[string]string                   `pulumi:"tags"`
	Type                            string                              `pulumi:"type"`
	WorkspaceId                     string                              `pulumi:"workspaceId"`
}


func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AllowPublicAccessWhenBehindVnet) {
		allowPublicAccessWhenBehindVnet_ := false
		tmp.AllowPublicAccessWhenBehindVnet = &allowPublicAccessWhenBehindVnet_
	}
	if isZero(tmp.HbiWorkspace) {
		hbiWorkspace_ := false
		tmp.HbiWorkspace = &hbiWorkspace_
	}
	return &tmp
}
