


package v20210401preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:databricks/v20210401preview:getWorkspace", args, &rv, opts...)
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
	Authorizations             []WorkspaceProviderAuthorizationResponse `pulumi:"authorizations"`
	CreatedBy                  *CreatedByResponse                       `pulumi:"createdBy"`
	CreatedDateTime            string                                   `pulumi:"createdDateTime"`
	Encryption                 *WorkspacePropertiesResponseEncryption   `pulumi:"encryption"`
	Id                         string                                   `pulumi:"id"`
	Location                   string                                   `pulumi:"location"`
	ManagedResourceGroupId     string                                   `pulumi:"managedResourceGroupId"`
	Name                       string                                   `pulumi:"name"`
	Parameters                 *WorkspaceCustomParametersResponse       `pulumi:"parameters"`
	PrivateEndpointConnections []PrivateEndpointConnectionResponse      `pulumi:"privateEndpointConnections"`
	ProvisioningState          string                                   `pulumi:"provisioningState"`
	PublicNetworkAccess        *string                                  `pulumi:"publicNetworkAccess"`
	RequiredNsgRules           *string                                  `pulumi:"requiredNsgRules"`
	Sku                        *SkuResponse                             `pulumi:"sku"`
	StorageAccountIdentity     *ManagedIdentityConfigurationResponse    `pulumi:"storageAccountIdentity"`
	SystemData                 SystemDataResponse                       `pulumi:"systemData"`
	Tags                       map[string]string                        `pulumi:"tags"`
	Type                       string                                   `pulumi:"type"`
	UiDefinitionUri            *string                                  `pulumi:"uiDefinitionUri"`
	UpdatedBy                  *CreatedByResponse                       `pulumi:"updatedBy"`
	WorkspaceId                string                                   `pulumi:"workspaceId"`
	WorkspaceUrl               string                                   `pulumi:"workspaceUrl"`
}
