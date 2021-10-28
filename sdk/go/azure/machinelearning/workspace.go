


package machinelearning

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	CreationTime         pulumi.StringOutput    `pulumi:"creationTime"`
	KeyVaultIdentifierId pulumi.StringPtrOutput `pulumi:"keyVaultIdentifierId"`
	Location             pulumi.StringOutput    `pulumi:"location"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	OwnerEmail           pulumi.StringOutput    `pulumi:"ownerEmail"`
	StudioEndpoint       pulumi.StringOutput    `pulumi:"studioEndpoint"`
	Tags                 pulumi.StringMapOutput `pulumi:"tags"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	UserStorageAccountId pulumi.StringOutput    `pulumi:"userStorageAccountId"`
	WorkspaceId          pulumi.StringOutput    `pulumi:"workspaceId"`
	WorkspaceState       pulumi.StringOutput    `pulumi:"workspaceState"`
	WorkspaceType        pulumi.StringOutput    `pulumi:"workspaceType"`
}


func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OwnerEmail == nil {
		return nil, errors.New("invalid value for required argument 'OwnerEmail'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.UserStorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'UserStorageAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearning:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20160401:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20160401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearning/v20191001:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearning/v20191001:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearning:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearning:Workspace", name, id, state, &resource, opts...)
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
	KeyVaultIdentifierId *string           `pulumi:"keyVaultIdentifierId"`
	Location             *string           `pulumi:"location"`
	OwnerEmail           string            `pulumi:"ownerEmail"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Tags                 map[string]string `pulumi:"tags"`
	UserStorageAccountId string            `pulumi:"userStorageAccountId"`
	WorkspaceName        *string           `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	KeyVaultIdentifierId pulumi.StringPtrInput
	Location             pulumi.StringPtrInput
	OwnerEmail           pulumi.StringInput
	ResourceGroupName    pulumi.StringInput
	Tags                 pulumi.StringMapInput
	UserStorageAccountId pulumi.StringInput
	WorkspaceName        pulumi.StringPtrInput
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
