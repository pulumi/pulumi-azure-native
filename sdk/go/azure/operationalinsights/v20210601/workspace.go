


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workspace struct {
	pulumi.CustomResourceState

	CreatedDate                     pulumi.StringOutput                          `pulumi:"createdDate"`
	CustomerId                      pulumi.StringOutput                          `pulumi:"customerId"`
	Etag                            pulumi.StringPtrOutput                       `pulumi:"etag"`
	Features                        WorkspaceFeaturesResponsePtrOutput           `pulumi:"features"`
	ForceCmkForQuery                pulumi.BoolPtrOutput                         `pulumi:"forceCmkForQuery"`
	Location                        pulumi.StringOutput                          `pulumi:"location"`
	ModifiedDate                    pulumi.StringOutput                          `pulumi:"modifiedDate"`
	Name                            pulumi.StringOutput                          `pulumi:"name"`
	PrivateLinkScopedResources      PrivateLinkScopedResourceResponseArrayOutput `pulumi:"privateLinkScopedResources"`
	ProvisioningState               pulumi.StringPtrOutput                       `pulumi:"provisioningState"`
	PublicNetworkAccessForIngestion pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery     pulumi.StringPtrOutput                       `pulumi:"publicNetworkAccessForQuery"`
	RetentionInDays                 pulumi.IntPtrOutput                          `pulumi:"retentionInDays"`
	Sku                             WorkspaceSkuResponsePtrOutput                `pulumi:"sku"`
	Tags                            pulumi.StringMapOutput                       `pulumi:"tags"`
	Type                            pulumi.StringOutput                          `pulumi:"type"`
	WorkspaceCapping                WorkspaceCappingResponsePtrOutput            `pulumi:"workspaceCapping"`
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
			Type: pulumi.String("azure-native:operationalinsights/v20151101preview:Workspace"),
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
			Type: pulumi.String("azure-native:operationalinsights/v20211201preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:operationalinsights/v20210601:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:operationalinsights/v20210601:Workspace", name, id, state, &resource, opts...)
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
	Features                        *WorkspaceFeatures `pulumi:"features"`
	ForceCmkForQuery                *bool              `pulumi:"forceCmkForQuery"`
	Location                        *string            `pulumi:"location"`
	ProvisioningState               *string            `pulumi:"provisioningState"`
	PublicNetworkAccessForIngestion *string            `pulumi:"publicNetworkAccessForIngestion"`
	PublicNetworkAccessForQuery     *string            `pulumi:"publicNetworkAccessForQuery"`
	ResourceGroupName               string             `pulumi:"resourceGroupName"`
	RetentionInDays                 *int               `pulumi:"retentionInDays"`
	Sku                             *WorkspaceSku      `pulumi:"sku"`
	Tags                            map[string]string  `pulumi:"tags"`
	WorkspaceCapping                *WorkspaceCapping  `pulumi:"workspaceCapping"`
	WorkspaceName                   *string            `pulumi:"workspaceName"`
}


type WorkspaceArgs struct {
	Features                        WorkspaceFeaturesPtrInput
	ForceCmkForQuery                pulumi.BoolPtrInput
	Location                        pulumi.StringPtrInput
	ProvisioningState               pulumi.StringPtrInput
	PublicNetworkAccessForIngestion pulumi.StringPtrInput
	PublicNetworkAccessForQuery     pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	RetentionInDays                 pulumi.IntPtrInput
	Sku                             WorkspaceSkuPtrInput
	Tags                            pulumi.StringMapInput
	WorkspaceCapping                WorkspaceCappingPtrInput
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

func (o WorkspaceOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) CustomerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CustomerId }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) Features() WorkspaceFeaturesResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceFeaturesResponsePtrOutput { return v.Features }).(WorkspaceFeaturesResponsePtrOutput)
}

func (o WorkspaceOutput) ForceCmkForQuery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.ForceCmkForQuery }).(pulumi.BoolPtrOutput)
}

func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) ModifiedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ModifiedDate }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) PrivateLinkScopedResources() PrivateLinkScopedResourceResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateLinkScopedResourceResponseArrayOutput { return v.PrivateLinkScopedResources }).(PrivateLinkScopedResourceResponseArrayOutput)
}

func (o WorkspaceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) PublicNetworkAccessForIngestion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccessForIngestion }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) PublicNetworkAccessForQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccessForQuery }).(pulumi.StringPtrOutput)
}

func (o WorkspaceOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.IntPtrOutput { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o WorkspaceOutput) Sku() WorkspaceSkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceSkuResponsePtrOutput { return v.Sku }).(WorkspaceSkuResponsePtrOutput)
}

func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkspaceOutput) WorkspaceCapping() WorkspaceCappingResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceCappingResponsePtrOutput { return v.WorkspaceCapping }).(WorkspaceCappingResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
