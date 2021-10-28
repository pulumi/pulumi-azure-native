


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BookmarkRelation struct {
	pulumi.CustomResourceState

	Etag                pulumi.StringPtrOutput `pulumi:"etag"`
	Name                pulumi.StringOutput    `pulumi:"name"`
	RelatedResourceId   pulumi.StringOutput    `pulumi:"relatedResourceId"`
	RelatedResourceKind pulumi.StringOutput    `pulumi:"relatedResourceKind"`
	RelatedResourceName pulumi.StringOutput    `pulumi:"relatedResourceName"`
	RelatedResourceType pulumi.StringOutput    `pulumi:"relatedResourceType"`
	Type                pulumi.StringOutput    `pulumi:"type"`
}


func NewBookmarkRelation(ctx *pulumi.Context,
	name string, args *BookmarkRelationArgs, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BookmarkId == nil {
		return nil, errors.New("invalid value for required argument 'BookmarkId'")
	}
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
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
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:BookmarkRelation"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:BookmarkRelation"),
		},
	})
	opts = append(opts, aliases)
	var resource BookmarkRelation
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:BookmarkRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetBookmarkRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BookmarkRelationState, opts ...pulumi.ResourceOption) (*BookmarkRelation, error) {
	var resource BookmarkRelation
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:BookmarkRelation", name, id, state, &resource, opts...)
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
	BookmarkId                          string  `pulumi:"bookmarkId"`
	Etag                                *string `pulumi:"etag"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	RelatedResourceId                   string  `pulumi:"relatedResourceId"`
	RelationName                        *string `pulumi:"relationName"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type BookmarkRelationArgs struct {
	BookmarkId                          pulumi.StringInput
	Etag                                pulumi.StringPtrInput
	OperationalInsightsResourceProvider pulumi.StringInput
	RelatedResourceId                   pulumi.StringInput
	RelationName                        pulumi.StringPtrInput
	ResourceGroupName                   pulumi.StringInput
	WorkspaceName                       pulumi.StringInput
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
	return reflect.TypeOf((*BookmarkRelation)(nil))
}

func (i *BookmarkRelation) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return i.ToBookmarkRelationOutputWithContext(context.Background())
}

func (i *BookmarkRelation) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkRelationOutput)
}

type BookmarkRelationOutput struct{ *pulumi.OutputState }

func (BookmarkRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkRelation)(nil))
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutput() BookmarkRelationOutput {
	return o
}

func (o BookmarkRelationOutput) ToBookmarkRelationOutputWithContext(ctx context.Context) BookmarkRelationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BookmarkRelationOutput{})
}
