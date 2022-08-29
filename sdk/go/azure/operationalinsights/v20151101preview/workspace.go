


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	CustomerId        pulumi.StringOutput    `pulumi:"customerId"`
	ETag              pulumi.StringPtrOutput `pulumi:"eTag"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	PortalUrl         pulumi.StringOutput    `pulumi:"portalUrl"`
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	RetentionInDays   pulumi.IntPtrOutput    `pulumi:"retentionInDays"`
	Sku               SkuResponsePtrOutput   `pulumi:"sku"`
	Source            pulumi.StringOutput    `pulumi:"source"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
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
			Type: pulumi.String("azure-native:operationalinsights:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20201001:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20210601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:operationalinsights/v20211201preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:operationalinsights/v20151101preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:operationalinsights/v20151101preview:Workspace", name, id, state, &resource, opts...)
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
	ETag              *string           `pulumi:"eTag"`
	Location          *string           `pulumi:"location"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RetentionInDays   *int              `pulumi:"retentionInDays"`
	Sku               *Sku              `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	WorkspaceName     *string           `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	ETag              pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RetentionInDays   pulumi.IntPtrInput
	Sku               SkuPtrInput
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

func (o WorkspaceOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CustomerId }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) PortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.PortalUrl }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o WorkspaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o WorkspaceOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Source }).(pulumi.StringOutput)
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
