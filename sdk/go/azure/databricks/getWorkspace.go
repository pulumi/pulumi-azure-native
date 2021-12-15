


package databricks

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:databricks:getWorkspace", args, &rv, opts...)
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
	Authorizations         []WorkspaceProviderAuthorizationResponse `pulumi:"authorizations"`
	CreatedBy              *CreatedByResponse                       `pulumi:"createdBy"`
	CreatedDateTime        string                                   `pulumi:"createdDateTime"`
	Id                     string                                   `pulumi:"id"`
	Location               string                                   `pulumi:"location"`
	ManagedResourceGroupId string                                   `pulumi:"managedResourceGroupId"`
	Name                   string                                   `pulumi:"name"`
	Parameters             *WorkspaceCustomParametersResponse       `pulumi:"parameters"`
	ProvisioningState      string                                   `pulumi:"provisioningState"`
	Sku                    *SkuResponse                             `pulumi:"sku"`
	StorageAccountIdentity *ManagedIdentityConfigurationResponse    `pulumi:"storageAccountIdentity"`
	Tags                   map[string]string                        `pulumi:"tags"`
	Type                   string                                   `pulumi:"type"`
	UiDefinitionUri        *string                                  `pulumi:"uiDefinitionUri"`
	UpdatedBy              *CreatedByResponse                       `pulumi:"updatedBy"`
	WorkspaceId            string                                   `pulumi:"workspaceId"`
	WorkspaceUrl           string                                   `pulumi:"workspaceUrl"`
}


func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Parameters = tmp.Parameters.Defaults()

	return &tmp
}
