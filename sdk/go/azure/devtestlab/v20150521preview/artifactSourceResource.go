


package v20150521preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ArtifactSourceResource struct {
	pulumi.CustomResourceState

	BranchRef         pulumi.StringPtrOutput `pulumi:"branchRef"`
	DisplayName       pulumi.StringPtrOutput `pulumi:"displayName"`
	FolderPath        pulumi.StringPtrOutput `pulumi:"folderPath"`
	Location          pulumi.StringPtrOutput `pulumi:"location"`
	Name              pulumi.StringPtrOutput `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	SecurityToken     pulumi.StringPtrOutput `pulumi:"securityToken"`
	SourceType        pulumi.StringPtrOutput `pulumi:"sourceType"`
	Status            pulumi.StringPtrOutput `pulumi:"status"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringPtrOutput `pulumi:"type"`
	Uri               pulumi.StringPtrOutput `pulumi:"uri"`
}


func NewArtifactSourceResource(ctx *pulumi.Context,
	name string, args *ArtifactSourceResourceArgs, opts ...pulumi.ResourceOption) (*ArtifactSourceResource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:ArtifactSourceResource"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ArtifactSourceResource"),
		},
	})
	opts = append(opts, aliases)
	var resource ArtifactSourceResource
	err := ctx.RegisterResource("azure-native:devtestlab/v20150521preview:ArtifactSourceResource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetArtifactSourceResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ArtifactSourceResourceState, opts ...pulumi.ResourceOption) (*ArtifactSourceResource, error) {
	var resource ArtifactSourceResource
	err := ctx.ReadResource("azure-native:devtestlab/v20150521preview:ArtifactSourceResource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type artifactSourceResourceState struct {
}

type ArtifactSourceResourceState struct {
}

func (ArtifactSourceResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceResourceState)(nil)).Elem()
}

type artifactSourceResourceArgs struct {
	BranchRef         *string           `pulumi:"branchRef"`
	DisplayName       *string           `pulumi:"displayName"`
	FolderPath        *string           `pulumi:"folderPath"`
	Id                *string           `pulumi:"id"`
	LabName           string            `pulumi:"labName"`
	Location          *string           `pulumi:"location"`
	Name              *string           `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	SecurityToken     *string           `pulumi:"securityToken"`
	SourceType        *string           `pulumi:"sourceType"`
	Status            *string           `pulumi:"status"`
	Tags              map[string]string `pulumi:"tags"`
	Type              *string           `pulumi:"type"`
	Uri               *string           `pulumi:"uri"`
}


type ArtifactSourceResourceArgs struct {
	BranchRef         pulumi.StringPtrInput
	DisplayName       pulumi.StringPtrInput
	FolderPath        pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	LabName           pulumi.StringInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	SecurityToken     pulumi.StringPtrInput
	SourceType        pulumi.StringPtrInput
	Status            pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	Uri               pulumi.StringPtrInput
}

func (ArtifactSourceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*artifactSourceResourceArgs)(nil)).Elem()
}

type ArtifactSourceResourceInput interface {
	pulumi.Input

	ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput
	ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput
}

func (*ArtifactSourceResource) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSourceResource)(nil)).Elem()
}

func (i *ArtifactSourceResource) ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput {
	return i.ToArtifactSourceResourceOutputWithContext(context.Background())
}

func (i *ArtifactSourceResource) ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArtifactSourceResourceOutput)
}

type ArtifactSourceResourceOutput struct{ *pulumi.OutputState }

func (ArtifactSourceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArtifactSourceResource)(nil)).Elem()
}

func (o ArtifactSourceResourceOutput) ToArtifactSourceResourceOutput() ArtifactSourceResourceOutput {
	return o
}

func (o ArtifactSourceResourceOutput) ToArtifactSourceResourceOutputWithContext(ctx context.Context) ArtifactSourceResourceOutput {
	return o
}

func (o ArtifactSourceResourceOutput) BranchRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.BranchRef }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) FolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.FolderPath }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) SecurityToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.SecurityToken }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) SourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.SourceType }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ArtifactSourceResourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ArtifactSourceResourceOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArtifactSourceResource) pulumi.StringPtrOutput { return v.Uri }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ArtifactSourceResourceOutput{})
}
