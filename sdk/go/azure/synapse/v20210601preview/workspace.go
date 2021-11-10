


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	AdlaResourceId                   pulumi.StringOutput                               `pulumi:"adlaResourceId"`
	AzureADOnlyAuthentication        pulumi.BoolPtrOutput                              `pulumi:"azureADOnlyAuthentication"`
	ConnectivityEndpoints            pulumi.StringMapOutput                            `pulumi:"connectivityEndpoints"`
	CspWorkspaceAdminProperties      CspWorkspaceAdminPropertiesResponsePtrOutput      `pulumi:"cspWorkspaceAdminProperties"`
	DefaultDataLakeStorage           DataLakeStorageAccountDetailsResponsePtrOutput    `pulumi:"defaultDataLakeStorage"`
	Encryption                       EncryptionDetailsResponsePtrOutput                `pulumi:"encryption"`
	ExtraProperties                  pulumi.MapOutput                                  `pulumi:"extraProperties"`
	Identity                         ManagedIdentityResponsePtrOutput                  `pulumi:"identity"`
	Location                         pulumi.StringOutput                               `pulumi:"location"`
	ManagedResourceGroupName         pulumi.StringPtrOutput                            `pulumi:"managedResourceGroupName"`
	ManagedVirtualNetwork            pulumi.StringPtrOutput                            `pulumi:"managedVirtualNetwork"`
	ManagedVirtualNetworkSettings    ManagedVirtualNetworkSettingsResponsePtrOutput    `pulumi:"managedVirtualNetworkSettings"`
	Name                             pulumi.StringOutput                               `pulumi:"name"`
	PrivateEndpointConnections       PrivateEndpointConnectionResponseArrayOutput      `pulumi:"privateEndpointConnections"`
	ProvisioningState                pulumi.StringOutput                               `pulumi:"provisioningState"`
	PublicNetworkAccess              pulumi.StringPtrOutput                            `pulumi:"publicNetworkAccess"`
	PurviewConfiguration             PurviewConfigurationResponsePtrOutput             `pulumi:"purviewConfiguration"`
	Settings                         pulumi.MapOutput                                  `pulumi:"settings"`
	SqlAdministratorLogin            pulumi.StringPtrOutput                            `pulumi:"sqlAdministratorLogin"`
	SqlAdministratorLoginPassword    pulumi.StringPtrOutput                            `pulumi:"sqlAdministratorLoginPassword"`
	Tags                             pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                             pulumi.StringOutput                               `pulumi:"type"`
	VirtualNetworkProfile            VirtualNetworkProfileResponsePtrOutput            `pulumi:"virtualNetworkProfile"`
	WorkspaceRepositoryConfiguration WorkspaceRepositoryConfigurationResponsePtrOutput `pulumi:"workspaceRepositoryConfiguration"`
	WorkspaceUID                     pulumi.StringOutput                               `pulumi:"workspaceUID"`
}


func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:synapse/v20210601preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:synapse/v20210601preview:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	AzureADOnlyAuthentication        *bool                             `pulumi:"azureADOnlyAuthentication"`
	ConnectivityEndpoints            map[string]string                 `pulumi:"connectivityEndpoints"`
	CspWorkspaceAdminProperties      *CspWorkspaceAdminProperties      `pulumi:"cspWorkspaceAdminProperties"`
	DefaultDataLakeStorage           *DataLakeStorageAccountDetails    `pulumi:"defaultDataLakeStorage"`
	Encryption                       *EncryptionDetails                `pulumi:"encryption"`
	Identity                         *ManagedIdentity                  `pulumi:"identity"`
	Location                         *string                           `pulumi:"location"`
	ManagedResourceGroupName         *string                           `pulumi:"managedResourceGroupName"`
	ManagedVirtualNetwork            *string                           `pulumi:"managedVirtualNetwork"`
	ManagedVirtualNetworkSettings    *ManagedVirtualNetworkSettings    `pulumi:"managedVirtualNetworkSettings"`
	PrivateEndpointConnections       []PrivateEndpointConnectionType   `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess              *string                           `pulumi:"publicNetworkAccess"`
	PurviewConfiguration             *PurviewConfiguration             `pulumi:"purviewConfiguration"`
	ResourceGroupName                string                            `pulumi:"resourceGroupName"`
	SqlAdministratorLogin            *string                           `pulumi:"sqlAdministratorLogin"`
	SqlAdministratorLoginPassword    *string                           `pulumi:"sqlAdministratorLoginPassword"`
	Tags                             map[string]string                 `pulumi:"tags"`
	VirtualNetworkProfile            *VirtualNetworkProfile            `pulumi:"virtualNetworkProfile"`
	WorkspaceName                    *string                           `pulumi:"workspaceName"`
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration `pulumi:"workspaceRepositoryConfiguration"`
}


type WorkspaceArgs struct {
	AzureADOnlyAuthentication        pulumi.BoolPtrInput
	ConnectivityEndpoints            pulumi.StringMapInput
	CspWorkspaceAdminProperties      CspWorkspaceAdminPropertiesPtrInput
	DefaultDataLakeStorage           DataLakeStorageAccountDetailsPtrInput
	Encryption                       EncryptionDetailsPtrInput
	Identity                         ManagedIdentityPtrInput
	Location                         pulumi.StringPtrInput
	ManagedResourceGroupName         pulumi.StringPtrInput
	ManagedVirtualNetwork            pulumi.StringPtrInput
	ManagedVirtualNetworkSettings    ManagedVirtualNetworkSettingsPtrInput
	PrivateEndpointConnections       PrivateEndpointConnectionTypeArrayInput
	PublicNetworkAccess              pulumi.StringPtrInput
	PurviewConfiguration             PurviewConfigurationPtrInput
	ResourceGroupName                pulumi.StringInput
	SqlAdministratorLogin            pulumi.StringPtrInput
	SqlAdministratorLoginPassword    pulumi.StringPtrInput
	Tags                             pulumi.StringMapInput
	VirtualNetworkProfile            VirtualNetworkProfilePtrInput
	WorkspaceName                    pulumi.StringPtrInput
	WorkspaceRepositoryConfiguration WorkspaceRepositoryConfigurationPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
