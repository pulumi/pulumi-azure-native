


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Metadata struct {
	pulumi.CustomResourceState

	Author                   MetadataAuthorResponsePtrOutput       `pulumi:"author"`
	Categories               MetadataCategoriesResponsePtrOutput   `pulumi:"categories"`
	ContentId                pulumi.StringPtrOutput                `pulumi:"contentId"`
	ContentSchemaVersion     pulumi.StringPtrOutput                `pulumi:"contentSchemaVersion"`
	CustomVersion            pulumi.StringPtrOutput                `pulumi:"customVersion"`
	Dependencies             MetadataDependenciesResponsePtrOutput `pulumi:"dependencies"`
	Etag                     pulumi.StringPtrOutput                `pulumi:"etag"`
	FirstPublishDate         pulumi.StringPtrOutput                `pulumi:"firstPublishDate"`
	Icon                     pulumi.StringPtrOutput                `pulumi:"icon"`
	Kind                     pulumi.StringOutput                   `pulumi:"kind"`
	LastPublishDate          pulumi.StringPtrOutput                `pulumi:"lastPublishDate"`
	Name                     pulumi.StringOutput                   `pulumi:"name"`
	ParentId                 pulumi.StringOutput                   `pulumi:"parentId"`
	PreviewImages            pulumi.StringArrayOutput              `pulumi:"previewImages"`
	PreviewImagesDark        pulumi.StringArrayOutput              `pulumi:"previewImagesDark"`
	Providers                pulumi.StringArrayOutput              `pulumi:"providers"`
	Source                   MetadataSourceResponsePtrOutput       `pulumi:"source"`
	Support                  MetadataSupportResponsePtrOutput      `pulumi:"support"`
	SystemData               SystemDataResponseOutput              `pulumi:"systemData"`
	ThreatAnalysisTactics    pulumi.StringArrayOutput              `pulumi:"threatAnalysisTactics"`
	ThreatAnalysisTechniques pulumi.StringArrayOutput              `pulumi:"threatAnalysisTechniques"`
	Type                     pulumi.StringOutput                   `pulumi:"type"`
	Version                  pulumi.StringPtrOutput                `pulumi:"version"`
}


func NewMetadata(ctx *pulumi.Context,
	name string, args *MetadataArgs, opts ...pulumi.ResourceOption) (*Metadata, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ParentId == nil {
		return nil, errors.New("invalid value for required argument 'ParentId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:Metadata"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:Metadata"),
		},
	})
	opts = append(opts, aliases)
	var resource Metadata
	err := ctx.RegisterResource("azure-native:securityinsights/v20220801preview:Metadata", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMetadata(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetadataState, opts ...pulumi.ResourceOption) (*Metadata, error) {
	var resource Metadata
	err := ctx.ReadResource("azure-native:securityinsights/v20220801preview:Metadata", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type metadataState struct {
}

type MetadataState struct {
}

func (MetadataState) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataState)(nil)).Elem()
}

type metadataArgs struct {
	Author                   *MetadataAuthor       `pulumi:"author"`
	Categories               *MetadataCategories   `pulumi:"categories"`
	ContentId                *string               `pulumi:"contentId"`
	ContentSchemaVersion     *string               `pulumi:"contentSchemaVersion"`
	CustomVersion            *string               `pulumi:"customVersion"`
	Dependencies             *MetadataDependencies `pulumi:"dependencies"`
	FirstPublishDate         *string               `pulumi:"firstPublishDate"`
	Icon                     *string               `pulumi:"icon"`
	Kind                     string                `pulumi:"kind"`
	LastPublishDate          *string               `pulumi:"lastPublishDate"`
	MetadataName             *string               `pulumi:"metadataName"`
	ParentId                 string                `pulumi:"parentId"`
	PreviewImages            []string              `pulumi:"previewImages"`
	PreviewImagesDark        []string              `pulumi:"previewImagesDark"`
	Providers                []string              `pulumi:"providers"`
	ResourceGroupName        string                `pulumi:"resourceGroupName"`
	Source                   *MetadataSource       `pulumi:"source"`
	Support                  *MetadataSupport      `pulumi:"support"`
	ThreatAnalysisTactics    []string              `pulumi:"threatAnalysisTactics"`
	ThreatAnalysisTechniques []string              `pulumi:"threatAnalysisTechniques"`
	Version                  *string               `pulumi:"version"`
	WorkspaceName            string                `pulumi:"workspaceName"`
}


type MetadataArgs struct {
	Author                   MetadataAuthorPtrInput
	Categories               MetadataCategoriesPtrInput
	ContentId                pulumi.StringPtrInput
	ContentSchemaVersion     pulumi.StringPtrInput
	CustomVersion            pulumi.StringPtrInput
	Dependencies             MetadataDependenciesPtrInput
	FirstPublishDate         pulumi.StringPtrInput
	Icon                     pulumi.StringPtrInput
	Kind                     pulumi.StringInput
	LastPublishDate          pulumi.StringPtrInput
	MetadataName             pulumi.StringPtrInput
	ParentId                 pulumi.StringInput
	PreviewImages            pulumi.StringArrayInput
	PreviewImagesDark        pulumi.StringArrayInput
	Providers                pulumi.StringArrayInput
	ResourceGroupName        pulumi.StringInput
	Source                   MetadataSourcePtrInput
	Support                  MetadataSupportPtrInput
	ThreatAnalysisTactics    pulumi.StringArrayInput
	ThreatAnalysisTechniques pulumi.StringArrayInput
	Version                  pulumi.StringPtrInput
	WorkspaceName            pulumi.StringInput
}

func (MetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metadataArgs)(nil)).Elem()
}

type MetadataInput interface {
	pulumi.Input

	ToMetadataOutput() MetadataOutput
	ToMetadataOutputWithContext(ctx context.Context) MetadataOutput
}

func (*Metadata) ElementType() reflect.Type {
	return reflect.TypeOf((**Metadata)(nil)).Elem()
}

func (i *Metadata) ToMetadataOutput() MetadataOutput {
	return i.ToMetadataOutputWithContext(context.Background())
}

func (i *Metadata) ToMetadataOutputWithContext(ctx context.Context) MetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataOutput)
}

type MetadataOutput struct{ *pulumi.OutputState }

func (MetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Metadata)(nil)).Elem()
}

func (o MetadataOutput) ToMetadataOutput() MetadataOutput {
	return o
}

func (o MetadataOutput) ToMetadataOutputWithContext(ctx context.Context) MetadataOutput {
	return o
}

func (o MetadataOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataAuthorResponsePtrOutput { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

func (o MetadataOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataCategoriesResponsePtrOutput { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

func (o MetadataOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.ContentId }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) ContentSchemaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.ContentSchemaVersion }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) CustomVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.CustomVersion }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataDependenciesResponsePtrOutput { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

func (o MetadataOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) Icon() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.Icon }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MetadataOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

func (o MetadataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MetadataOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.ParentId }).(pulumi.StringOutput)
}

func (o MetadataOutput) PreviewImages() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.PreviewImages }).(pulumi.StringArrayOutput)
}

func (o MetadataOutput) PreviewImagesDark() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.PreviewImagesDark }).(pulumi.StringArrayOutput)
}

func (o MetadataOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.Providers }).(pulumi.StringArrayOutput)
}

func (o MetadataOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataSourceResponsePtrOutput { return v.Source }).(MetadataSourceResponsePtrOutput)
}

func (o MetadataOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v *Metadata) MetadataSupportResponsePtrOutput { return v.Support }).(MetadataSupportResponsePtrOutput)
}

func (o MetadataOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Metadata) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MetadataOutput) ThreatAnalysisTactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.ThreatAnalysisTactics }).(pulumi.StringArrayOutput)
}

func (o MetadataOutput) ThreatAnalysisTechniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringArrayOutput { return v.ThreatAnalysisTechniques }).(pulumi.StringArrayOutput)
}

func (o MetadataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MetadataOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Metadata) pulumi.StringPtrOutput { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MetadataOutput{})
}
