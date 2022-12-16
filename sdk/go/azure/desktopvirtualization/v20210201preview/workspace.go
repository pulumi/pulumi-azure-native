


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	ApplicationGroupReferences pulumi.StringArrayOutput                                     `pulumi:"applicationGroupReferences"`
	CloudPcResource            pulumi.BoolOutput                                            `pulumi:"cloudPcResource"`
	Description                pulumi.StringPtrOutput                                       `pulumi:"description"`
	Etag                       pulumi.StringOutput                                          `pulumi:"etag"`
	FriendlyName               pulumi.StringPtrOutput                                       `pulumi:"friendlyName"`
	Identity                   ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput `pulumi:"identity"`
	Kind                       pulumi.StringPtrOutput                                       `pulumi:"kind"`
	Location                   pulumi.StringPtrOutput                                       `pulumi:"location"`
	ManagedBy                  pulumi.StringPtrOutput                                       `pulumi:"managedBy"`
	Name                       pulumi.StringOutput                                          `pulumi:"name"`
	ObjectId                   pulumi.StringOutput                                          `pulumi:"objectId"`
	Plan                       ResourceModelWithAllowedPropertySetResponsePlanPtrOutput     `pulumi:"plan"`
	Sku                        ResourceModelWithAllowedPropertySetResponseSkuPtrOutput      `pulumi:"sku"`
	Tags                       pulumi.StringMapOutput                                       `pulumi:"tags"`
	Type                       pulumi.StringOutput                                          `pulumi:"type"`
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
			Type: pulumi.String("azure-native:desktopvirtualization:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190123preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20190924preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20191210preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20200921preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201019preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201102preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20201110preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210114preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210309preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210712:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220909:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210201preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210201preview:Workspace", name, id, state, &resource, opts...)
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
	ApplicationGroupReferences []string                                     `pulumi:"applicationGroupReferences"`
	Description                *string                                      `pulumi:"description"`
	FriendlyName               *string                                      `pulumi:"friendlyName"`
	Identity                   *ResourceModelWithAllowedPropertySetIdentity `pulumi:"identity"`
	Kind                       *string                                      `pulumi:"kind"`
	Location                   *string                                      `pulumi:"location"`
	ManagedBy                  *string                                      `pulumi:"managedBy"`
	Plan                       *ResourceModelWithAllowedPropertySetPlan     `pulumi:"plan"`
	ResourceGroupName          string                                       `pulumi:"resourceGroupName"`
	Sku                        *ResourceModelWithAllowedPropertySetSku      `pulumi:"sku"`
	Tags                       map[string]string                            `pulumi:"tags"`
	WorkspaceName              *string                                      `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	ApplicationGroupReferences pulumi.StringArrayInput
	Description                pulumi.StringPtrInput
	FriendlyName               pulumi.StringPtrInput
	Identity                   ResourceModelWithAllowedPropertySetIdentityPtrInput
	Kind                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	ManagedBy                  pulumi.StringPtrInput
	Plan                       ResourceModelWithAllowedPropertySetPlanPtrInput
	ResourceGroupName          pulumi.StringInput
	Sku                        ResourceModelWithAllowedPropertySetSkuPtrInput
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

func (o WorkspaceOutput) ApplicationGroupReferences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringArrayOutput { return v.ApplicationGroupReferences }).(pulumi.StringArrayOutput)
}

func (o WorkspaceOutput) CloudPcResource() pulumi.BoolOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolOutput { return v.CloudPcResource }).(pulumi.BoolOutput)
}

func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Identity() ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput {
	return o.ApplyT(func(v *Workspace) ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput { return v.Identity }).(ResourceModelWithAllowedPropertySetResponseIdentityPtrOutput)
}

func (o WorkspaceOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Plan() ResourceModelWithAllowedPropertySetResponsePlanPtrOutput {
	return o.ApplyT(func(v *Workspace) ResourceModelWithAllowedPropertySetResponsePlanPtrOutput { return v.Plan }).(ResourceModelWithAllowedPropertySetResponsePlanPtrOutput)
}

func (o WorkspaceOutput) Sku() ResourceModelWithAllowedPropertySetResponseSkuPtrOutput {
	return o.ApplyT(func(v *Workspace) ResourceModelWithAllowedPropertySetResponseSkuPtrOutput { return v.Sku }).(ResourceModelWithAllowedPropertySetResponseSkuPtrOutput)
}

func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
