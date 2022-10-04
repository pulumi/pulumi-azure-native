


package v20220501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	AllowPublicAccessWhenBehindVnet pulumi.BoolPtrOutput                             `pulumi:"allowPublicAccessWhenBehindVnet"`
	ApplicationInsights             pulumi.StringPtrOutput                           `pulumi:"applicationInsights"`
	ContainerRegistry               pulumi.StringPtrOutput                           `pulumi:"containerRegistry"`
	Description                     pulumi.StringPtrOutput                           `pulumi:"description"`
	DiscoveryUrl                    pulumi.StringPtrOutput                           `pulumi:"discoveryUrl"`
	Encryption                      EncryptionPropertyResponsePtrOutput              `pulumi:"encryption"`
	FriendlyName                    pulumi.StringPtrOutput                           `pulumi:"friendlyName"`
	HbiWorkspace                    pulumi.BoolPtrOutput                             `pulumi:"hbiWorkspace"`
	Identity                        ManagedServiceIdentityResponsePtrOutput          `pulumi:"identity"`
	ImageBuildCompute               pulumi.StringPtrOutput                           `pulumi:"imageBuildCompute"`
	KeyVault                        pulumi.StringPtrOutput                           `pulumi:"keyVault"`
	Location                        pulumi.StringPtrOutput                           `pulumi:"location"`
	MlFlowTrackingUri               pulumi.StringOutput                              `pulumi:"mlFlowTrackingUri"`
	Name                            pulumi.StringOutput                              `pulumi:"name"`
	NotebookInfo                    NotebookResourceInfoResponseOutput               `pulumi:"notebookInfo"`
	PrimaryUserAssignedIdentity     pulumi.StringPtrOutput                           `pulumi:"primaryUserAssignedIdentity"`
	PrivateEndpointConnections      PrivateEndpointConnectionResponseArrayOutput     `pulumi:"privateEndpointConnections"`
	PrivateLinkCount                pulumi.IntOutput                                 `pulumi:"privateLinkCount"`
	ProvisioningState               pulumi.StringOutput                              `pulumi:"provisioningState"`
	PublicNetworkAccess             pulumi.StringPtrOutput                           `pulumi:"publicNetworkAccess"`
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsResponsePtrOutput `pulumi:"serviceManagedResourcesSettings"`
	ServiceProvisionedResourceGroup pulumi.StringOutput                              `pulumi:"serviceProvisionedResourceGroup"`
	SharedPrivateLinkResources      SharedPrivateLinkResourceResponseArrayOutput     `pulumi:"sharedPrivateLinkResources"`
	Sku                             SkuResponsePtrOutput                             `pulumi:"sku"`
	StorageAccount                  pulumi.StringPtrOutput                           `pulumi:"storageAccount"`
	StorageHnsEnabled               pulumi.BoolOutput                                `pulumi:"storageHnsEnabled"`
	SystemData                      SystemDataResponseOutput                         `pulumi:"systemData"`
	Tags                            pulumi.StringMapOutput                           `pulumi:"tags"`
	TenantId                        pulumi.StringOutput                              `pulumi:"tenantId"`
	Type                            pulumi.StringOutput                              `pulumi:"type"`
	V1LegacyMode                    pulumi.BoolPtrOutput                             `pulumi:"v1LegacyMode"`
	WorkspaceId                     pulumi.StringOutput                              `pulumi:"workspaceId"`
}


func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if isZero(args.AllowPublicAccessWhenBehindVnet) {
		args.AllowPublicAccessWhenBehindVnet = pulumi.BoolPtr(false)
	}
	if isZero(args.HbiWorkspace) {
		args.HbiWorkspace = pulumi.BoolPtr(false)
	}
	if isZero(args.V1LegacyMode) {
		args.V1LegacyMode = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220501:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220501:Workspace", name, id, state, &resource, opts...)
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
	AllowPublicAccessWhenBehindVnet *bool                            `pulumi:"allowPublicAccessWhenBehindVnet"`
	ApplicationInsights             *string                          `pulumi:"applicationInsights"`
	ContainerRegistry               *string                          `pulumi:"containerRegistry"`
	Description                     *string                          `pulumi:"description"`
	DiscoveryUrl                    *string                          `pulumi:"discoveryUrl"`
	Encryption                      *EncryptionProperty              `pulumi:"encryption"`
	FriendlyName                    *string                          `pulumi:"friendlyName"`
	HbiWorkspace                    *bool                            `pulumi:"hbiWorkspace"`
	Identity                        *ManagedServiceIdentity          `pulumi:"identity"`
	ImageBuildCompute               *string                          `pulumi:"imageBuildCompute"`
	KeyVault                        *string                          `pulumi:"keyVault"`
	Location                        *string                          `pulumi:"location"`
	PrimaryUserAssignedIdentity     *string                          `pulumi:"primaryUserAssignedIdentity"`
	PublicNetworkAccess             *string                          `pulumi:"publicNetworkAccess"`
	ResourceGroupName               string                           `pulumi:"resourceGroupName"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings `pulumi:"serviceManagedResourcesSettings"`
	SharedPrivateLinkResources      []SharedPrivateLinkResource      `pulumi:"sharedPrivateLinkResources"`
	Sku                             *Sku                             `pulumi:"sku"`
	StorageAccount                  *string                          `pulumi:"storageAccount"`
	Tags                            map[string]string                `pulumi:"tags"`
	V1LegacyMode                    *bool                            `pulumi:"v1LegacyMode"`
	WorkspaceName                   *string                          `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	AllowPublicAccessWhenBehindVnet pulumi.BoolPtrInput
	ApplicationInsights             pulumi.StringPtrInput
	ContainerRegistry               pulumi.StringPtrInput
	Description                     pulumi.StringPtrInput
	DiscoveryUrl                    pulumi.StringPtrInput
	Encryption                      EncryptionPropertyPtrInput
	FriendlyName                    pulumi.StringPtrInput
	HbiWorkspace                    pulumi.BoolPtrInput
	Identity                        ManagedServiceIdentityPtrInput
	ImageBuildCompute               pulumi.StringPtrInput
	KeyVault                        pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	PrimaryUserAssignedIdentity     pulumi.StringPtrInput
	PublicNetworkAccess             pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsPtrInput
	SharedPrivateLinkResources      SharedPrivateLinkResourceArrayInput
	Sku                             SkuPtrInput
	StorageAccount                  pulumi.StringPtrInput
	Tags                            pulumi.StringMapInput
	V1LegacyMode                    pulumi.BoolPtrInput
	WorkspaceName                   pulumi.StringPtrInput
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

func (o WorkspaceOutput) AllowPublicAccessWhenBehindVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.AllowPublicAccessWhenBehindVnet }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) EncryptionPropertyResponsePtrOutput { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

func (o WorkspaceOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) HbiWorkspace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.HbiWorkspace }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o WorkspaceOutput) ImageBuildCompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ImageBuildCompute }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.KeyVault }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) MlFlowTrackingUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.MlFlowTrackingUri }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) NotebookInfo() NotebookResourceInfoResponseOutput {
	return o.ApplyT(func(v *Workspace) NotebookResourceInfoResponseOutput { return v.NotebookInfo }).(NotebookResourceInfoResponseOutput)
}

func (o WorkspaceOutput) PrimaryUserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PrimaryUserAssignedIdentity }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o WorkspaceOutput) PrivateLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Workspace) pulumi.IntOutput { return v.PrivateLinkCount }).(pulumi.IntOutput)
}

func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) ServiceManagedResourcesSettings() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ServiceManagedResourcesSettingsResponsePtrOutput {
		return v.ServiceManagedResourcesSettings
	}).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

func (o WorkspaceOutput) ServiceProvisionedResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ServiceProvisionedResourceGroup }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

func (o WorkspaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o WorkspaceOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) StorageHnsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolOutput { return v.StorageHnsEnabled }).(pulumi.BoolOutput)
}

func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) V1LegacyMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.V1LegacyMode }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
