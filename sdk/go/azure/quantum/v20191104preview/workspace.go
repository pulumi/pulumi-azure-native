


package v20191104preview

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
			Type: pulumi.String("azure-native:quantum:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:quantum/v20220110preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:quantum/v20191104preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:quantum/v20191104preview:Workspace", name, id, state, &resource, opts...)
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

func (o WorkspaceOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Identity() QuantumWorkspaceResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Workspace) QuantumWorkspaceResponseIdentityPtrOutput { return v.Identity }).(QuantumWorkspaceResponseIdentityPtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) ProviderResponseArrayOutput { return v.Providers }).(ProviderResponseArrayOutput)
}

func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Usable() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Usable }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
