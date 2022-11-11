


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	AdlaResourceId                   pulumi.StringOutput                               `pulumi:"adlaResourceId"`
	ConnectivityEndpoints            pulumi.StringMapOutput                            `pulumi:"connectivityEndpoints"`
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
			Type: pulumi.String("azure-native:synapse/v20210501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:synapse/v20210401preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:synapse/v20210401preview:Workspace", name, id, state, &resource, opts...)
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
	ConnectivityEndpoints            map[string]string                 `pulumi:"connectivityEndpoints"`
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
	ConnectivityEndpoints            pulumi.StringMapInput
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
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) AdlaResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.AdlaResourceId }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ConnectivityEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.ConnectivityEndpoints }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) DefaultDataLakeStorage() DataLakeStorageAccountDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) DataLakeStorageAccountDetailsResponsePtrOutput { return v.DefaultDataLakeStorage }).(DataLakeStorageAccountDetailsResponsePtrOutput)
}

func (o WorkspaceOutput) Encryption() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) EncryptionDetailsResponsePtrOutput { return v.Encryption }).(EncryptionDetailsResponsePtrOutput)
}

func (o WorkspaceOutput) ExtraProperties() pulumi.MapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.MapOutput { return v.ExtraProperties }).(pulumi.MapOutput)
}

func (o WorkspaceOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) ManagedVirtualNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ManagedVirtualNetwork }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) ManagedVirtualNetworkSettings() ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedVirtualNetworkSettingsResponsePtrOutput {
		return v.ManagedVirtualNetworkSettings
	}).(ManagedVirtualNetworkSettingsResponsePtrOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) PurviewConfigurationResponsePtrOutput { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

func (o WorkspaceOutput) SqlAdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.SqlAdministratorLogin }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) SqlAdministratorLoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.SqlAdministratorLoginPassword }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) VirtualNetworkProfile() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) VirtualNetworkProfileResponsePtrOutput { return v.VirtualNetworkProfile }).(VirtualNetworkProfileResponsePtrOutput)
}

func (o WorkspaceOutput) WorkspaceRepositoryConfiguration() WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceRepositoryConfigurationResponsePtrOutput {
		return v.WorkspaceRepositoryConfiguration
	}).(WorkspaceRepositoryConfigurationResponsePtrOutput)
}

func (o WorkspaceOutput) WorkspaceUID() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceUID }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
