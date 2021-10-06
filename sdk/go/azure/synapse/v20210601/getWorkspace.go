


package v20210601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:synapse/v20210601:getWorkspace", args, &rv, opts...)
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
	AdlaResourceId                   string                                    `pulumi:"adlaResourceId"`
	AzureADOnlyAuthentication        *bool                                     `pulumi:"azureADOnlyAuthentication"`
	ConnectivityEndpoints            map[string]string                         `pulumi:"connectivityEndpoints"`
	CspWorkspaceAdminProperties      *CspWorkspaceAdminPropertiesResponse      `pulumi:"cspWorkspaceAdminProperties"`
	DefaultDataLakeStorage           *DataLakeStorageAccountDetailsResponse    `pulumi:"defaultDataLakeStorage"`
	Encryption                       *EncryptionDetailsResponse                `pulumi:"encryption"`
	ExtraProperties                  map[string]interface{}                    `pulumi:"extraProperties"`
	Id                               string                                    `pulumi:"id"`
	Identity                         *ManagedIdentityResponse                  `pulumi:"identity"`
	Location                         string                                    `pulumi:"location"`
	ManagedResourceGroupName         *string                                   `pulumi:"managedResourceGroupName"`
	ManagedVirtualNetwork            *string                                   `pulumi:"managedVirtualNetwork"`
	ManagedVirtualNetworkSettings    *ManagedVirtualNetworkSettingsResponse    `pulumi:"managedVirtualNetworkSettings"`
	Name                             string                                    `pulumi:"name"`
	PrivateEndpointConnections       []PrivateEndpointConnectionResponse       `pulumi:"privateEndpointConnections"`
	ProvisioningState                string                                    `pulumi:"provisioningState"`
	PublicNetworkAccess              *string                                   `pulumi:"publicNetworkAccess"`
	PurviewConfiguration             *PurviewConfigurationResponse             `pulumi:"purviewConfiguration"`
	Settings                         map[string]interface{}                    `pulumi:"settings"`
	SqlAdministratorLogin            *string                                   `pulumi:"sqlAdministratorLogin"`
	SqlAdministratorLoginPassword    *string                                   `pulumi:"sqlAdministratorLoginPassword"`
	Tags                             map[string]string                         `pulumi:"tags"`
	Type                             string                                    `pulumi:"type"`
	VirtualNetworkProfile            *VirtualNetworkProfileResponse            `pulumi:"virtualNetworkProfile"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfigurationResponse `pulumi:"workspaceRepositoryConfiguration"`
	WorkspaceUID                     string                                    `pulumi:"workspaceUID"`
}
