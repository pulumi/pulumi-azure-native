


package v20210101

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
	Identity                        IdentityResponsePtrOutput                        `pulumi:"identity"`
	ImageBuildCompute               pulumi.StringPtrOutput                           `pulumi:"imageBuildCompute"`
	KeyVault                        pulumi.StringPtrOutput                           `pulumi:"keyVault"`
	Location                        pulumi.StringPtrOutput                           `pulumi:"location"`
	Name                            pulumi.StringOutput                              `pulumi:"name"`
	NotebookInfo                    NotebookResourceInfoResponseOutput               `pulumi:"notebookInfo"`
	PrimaryUserAssignedIdentity     pulumi.StringPtrOutput                           `pulumi:"primaryUserAssignedIdentity"`
	PrivateEndpointConnections      PrivateEndpointConnectionResponseArrayOutput     `pulumi:"privateEndpointConnections"`
	PrivateLinkCount                pulumi.IntOutput                                 `pulumi:"privateLinkCount"`
	ProvisioningState               pulumi.StringOutput                              `pulumi:"provisioningState"`
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsResponsePtrOutput `pulumi:"serviceManagedResourcesSettings"`
	ServiceProvisionedResourceGroup pulumi.StringOutput                              `pulumi:"serviceProvisionedResourceGroup"`
	SharedPrivateLinkResources      SharedPrivateLinkResourceResponseArrayOutput     `pulumi:"sharedPrivateLinkResources"`
	Sku                             SkuResponsePtrOutput                             `pulumi:"sku"`
	StorageAccount                  pulumi.StringPtrOutput                           `pulumi:"storageAccount"`
	SystemData                      SystemDataResponseOutput                         `pulumi:"systemData"`
	Tags                            pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                            pulumi.StringOutput                              `pulumi:"type"`
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
	if args.AllowPublicAccessWhenBehindVnet == nil {
		args.AllowPublicAccessWhenBehindVnet = pulumi.BoolPtr(false)
	}
	if args.HbiWorkspace == nil {
		args.HbiWorkspace = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20180301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20181119:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20190501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20190601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20191101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200218preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200301:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200501preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200515preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210701:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210101:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210101:Workspace", name, id, state, &resource, opts...)
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
	Identity                        *Identity                        `pulumi:"identity"`
	ImageBuildCompute               *string                          `pulumi:"imageBuildCompute"`
	KeyVault                        *string                          `pulumi:"keyVault"`
	Location                        *string                          `pulumi:"location"`
	PrimaryUserAssignedIdentity     *string                          `pulumi:"primaryUserAssignedIdentity"`
	ResourceGroupName               string                           `pulumi:"resourceGroupName"`
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings `pulumi:"serviceManagedResourcesSettings"`
	SharedPrivateLinkResources      []SharedPrivateLinkResource      `pulumi:"sharedPrivateLinkResources"`
	Sku                             *Sku                             `pulumi:"sku"`
	StorageAccount                  *string                          `pulumi:"storageAccount"`
	Tags                            map[string]string                `pulumi:"tags"`
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
	Identity                        IdentityPtrInput
	ImageBuildCompute               pulumi.StringPtrInput
	KeyVault                        pulumi.StringPtrInput
	Location                        pulumi.StringPtrInput
	PrimaryUserAssignedIdentity     pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsPtrInput
	SharedPrivateLinkResources      SharedPrivateLinkResourceArrayInput
	Sku                             SkuPtrInput
	StorageAccount                  pulumi.StringPtrInput
	Tags                            pulumi.StringMapInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*WorkspaceInput)(nil)).Elem(), &Workspace{})
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
