


package v20200921preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	ApplicationGroupReferences pulumi.StringArrayOutput `pulumi:"applicationGroupReferences"`
	Description                pulumi.StringPtrOutput   `pulumi:"description"`
	FriendlyName               pulumi.StringPtrOutput   `pulumi:"friendlyName"`
	Location                   pulumi.StringOutput      `pulumi:"location"`
	Name                       pulumi.StringOutput      `pulumi:"name"`
	Tags                       pulumi.StringMapOutput   `pulumi:"tags"`
	Type                       pulumi.StringOutput      `pulumi:"type"`
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
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20200921preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20190123preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20190924preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20191210preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201019preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201102preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20201110preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210114preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210309preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210712:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:desktopvirtualization/v20210903preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20200921preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20200921preview:Workspace", name, id, state, &resource, opts...)
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
	ApplicationGroupReferences []string          `pulumi:"applicationGroupReferences"`
	Description                *string           `pulumi:"description"`
	FriendlyName               *string           `pulumi:"friendlyName"`
	Location                   *string           `pulumi:"location"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	Tags                       map[string]string `pulumi:"tags"`
	WorkspaceName              *string           `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	ApplicationGroupReferences pulumi.StringArrayInput
	Description                pulumi.StringPtrInput
	FriendlyName               pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
	WorkspaceName              pulumi.StringPtrInput
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
