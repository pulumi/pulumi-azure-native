


package quantum

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	EndpointUri       pulumi.StringOutput                       `pulumi:"endpointUri"`
	Identity          QuantumWorkspaceResponseIdentityPtrOutput `pulumi:"identity"`
	Location          pulumi.StringOutput                       `pulumi:"location"`
	Name              pulumi.StringOutput                       `pulumi:"name"`
	Providers         ProviderResponseArrayOutput               `pulumi:"providers"`
	ProvisioningState pulumi.StringOutput                       `pulumi:"provisioningState"`
	StorageAccount    pulumi.StringPtrOutput                    `pulumi:"storageAccount"`
	SystemData        SystemDataResponseOutput                  `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                    `pulumi:"tags"`
	Type              pulumi.StringOutput                       `pulumi:"type"`
	Usable            pulumi.StringOutput                       `pulumi:"usable"`
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
			Type: pulumi.String("azure-native:quantum/v20191104preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:quantum:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:quantum:Workspace", name, id, state, &resource, opts...)
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
	Identity          *QuantumWorkspaceIdentity `pulumi:"identity"`
	Location          *string                   `pulumi:"location"`
	Providers         []Provider                `pulumi:"providers"`
	ResourceGroupName string                    `pulumi:"resourceGroupName"`
	StorageAccount    *string                   `pulumi:"storageAccount"`
	Tags              map[string]string         `pulumi:"tags"`
	WorkspaceName     *string                   `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	Identity          QuantumWorkspaceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Providers         ProviderArrayInput
	ResourceGroupName pulumi.StringInput
	StorageAccount    pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	WorkspaceName     pulumi.StringPtrInput
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
