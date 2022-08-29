


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BookmarkRelation struct {
	pulumi.CustomResourceState

	Etag                pulumi.StringPtrOutput   `pulumi:"etag"`
	Name                pulumi.StringOutput      `pulumi:"name"`
	RelatedResourceId   pulumi.StringOutput      `pulumi:"relatedResourceId"`
	RelatedResourceKind pulumi.StringOutput      `pulumi:"relatedResourceKind"`
	RelatedResourceName pulumi.StringOutput      `pulumi:"relatedResourceName"`
	RelatedResourceType pulumi.StringOutput      `pulumi:"relatedResourceType"`
	SystemData          SystemDataResponseOutput `pulumi:"systemData"`
	Type                pulumi.StringOutput      `pulumi:"type"`
}


func NewBookmarkRelation(ctx *pulumi.Context,
	name string, args *BookmarkRelationArgs, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BookmarkId == nil {
		return nil, errors.New("invalid value for required argument 'BookmarkId'")
	}
	if args.RelatedResourceId == nil {
		return nil, errors.New("invalid value for required argument 'RelatedResourceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:BookmarkRelation"),
		},
	})
	opts = append(opts, aliases)
	var resource BookmarkRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20220401preview:BookmarkRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBookmarkRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkRelationState, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	var resource BookmarkRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20220401preview:BookmarkRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type bookmarkRelationState struct {
}

type BookmarkRelationState struct {
}

func (BookmarkRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationState)(nil)).Elem()
}

type bookmarkRelationArgs struct {
	BookmarkId        string  `pulumi:"bookmarkId"`
	RelatedResourceId string  `pulumi:"relatedResourceId"`
	RelationName      *string `pulumi:"relationName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkspaceName     string  `pulumi:"workspaceName"`
}


type BookmarkRelationArgs struct {
	BookmarkId        pulumi.StringInput
	RelatedResourceId pulumi.StringInput
	RelationName      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (BookmarkRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bookmarkRelationArgs)(nil)).Elem()
}

type BookmarkRelationInput interface {
	pulumi.Input

	ToBookmarkRelationOutput() BookmarkRelationOutput
	ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput
}

func (*BookmarkRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**BookmarkRelation)(nil)).Elem()
}

func (i *BookmarkRelation) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return i.ToBookmarkRelationOutputWithContext(context.Background())
}

func (i *BookmarkRelation) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkRelationOutput)
}

type BookmarkRelationOutput struct{ *pulumi.OutputState }

func (BookmarkRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BookmarkRelation)(nil)).Elem()
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return o
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return o
}

func (o BookmarkRelationOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o BookmarkRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BookmarkRelationOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceId }).(pulumi.StringOutput)
}

func (o BookmarkRelationOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

func (o BookmarkRelationOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceName }).(pulumi.StringOutput)
}

func (o BookmarkRelationOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.RelatedResourceType }).(pulumi.StringOutput)
}

func (o BookmarkRelationOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *BookmarkRelation) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o BookmarkRelationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *BookmarkRelation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(BookmarkRelationOutput{})
}
