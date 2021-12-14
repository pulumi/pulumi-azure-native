


package databricks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	Authorizations         WorkspaceProviderAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	CreatedBy              CreatedByResponsePtrOutput                        `pulumi:"createdBy"`
	CreatedDateTime        pulumi.StringOutput                               `pulumi:"createdDateTime"`
	Location               pulumi.StringOutput                               `pulumi:"location"`
	ManagedResourceGroupId pulumi.StringOutput                               `pulumi:"managedResourceGroupId"`
	Name                   pulumi.StringOutput                               `pulumi:"name"`
	Parameters             WorkspaceCustomParametersResponsePtrOutput        `pulumi:"parameters"`
	ProvisioningState      pulumi.StringOutput                               `pulumi:"provisioningState"`
	Sku                    SkuResponsePtrOutput                              `pulumi:"sku"`
	StorageAccountIdentity ManagedIdentityConfigurationResponsePtrOutput     `pulumi:"storageAccountIdentity"`
	Tags                   pulumi.StringMapOutput                            `pulumi:"tags"`
	Type                   pulumi.StringOutput                               `pulumi:"type"`
	UiDefinitionUri        pulumi.StringPtrOutput                            `pulumi:"uiDefinitionUri"`
	UpdatedBy              CreatedByResponsePtrOutput                        `pulumi:"updatedBy"`
	WorkspaceId            pulumi.StringOutput                               `pulumi:"workspaceId"`
	WorkspaceUrl           pulumi.StringOutput                               `pulumi:"workspaceUrl"`
}


func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedResourceGroupId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	parametersApplier := func(v WorkspaceCustomParameters) *WorkspaceCustomParameters { return v.Defaults() }
	if args.Parameters != nil {
		args.Parameters = args.Parameters.ToWorkspaceCustomParametersPtrOutput().Elem().ApplyT(parametersApplier).(WorkspaceCustomParametersPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databricks/v20180401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20210401preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:databricks:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:databricks:Workspace", name, id, state, &resource, opts...)
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
	Authorizations         []WorkspaceProviderAuthorization `pulumi:"authorizations"`
	Location               *string                          `pulumi:"location"`
	ManagedResourceGroupId string                           `pulumi:"managedResourceGroupId"`
	Parameters             *WorkspaceCustomParameters       `pulumi:"parameters"`
	ResourceGroupName      string                           `pulumi:"resourceGroupName"`
	Sku                    *Sku                             `pulumi:"sku"`
	Tags                   map[string]string                `pulumi:"tags"`
	UiDefinitionUri        *string                          `pulumi:"uiDefinitionUri"`
	WorkspaceName          *string                          `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	Authorizations         WorkspaceProviderAuthorizationArrayInput
	Location               pulumi.StringPtrInput
	ManagedResourceGroupId pulumi.StringInput
	Parameters             WorkspaceCustomParametersPtrInput
	ResourceGroupName      pulumi.StringInput
	Sku                    SkuPtrInput
	Tags                   pulumi.StringMapInput
	UiDefinitionUri        pulumi.StringPtrInput
	WorkspaceName          pulumi.StringPtrInput
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

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
