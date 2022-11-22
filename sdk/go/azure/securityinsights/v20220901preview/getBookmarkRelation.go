


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBookmarkRelation(ctx *pulumi.Context, args *LookupBookmarkRelationArgs, opts ...pulumi.InvokeOption) (*LookupBookmarkRelationResult, error) {
	var rv LookupBookmarkRelationResult
	err := ctx.Invoke("azure-native:securityinsights/v20220901preview:getBookmarkRelation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBookmarkRelationArgs struct {
	BookmarkId        string `pulumi:"bookmarkId"`
	RelationName      string `pulumi:"relationName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupBookmarkRelationResult struct {
	Etag                *string            `pulumi:"etag"`
	Id                  string             `pulumi:"id"`
	Name                string             `pulumi:"name"`
	RelatedResourceId   string             `pulumi:"relatedResourceId"`
	RelatedResourceKind string             `pulumi:"relatedResourceKind"`
	RelatedResourceName string             `pulumi:"relatedResourceName"`
	RelatedResourceType string             `pulumi:"relatedResourceType"`
	SystemData          SystemDataResponse `pulumi:"systemData"`
	Type                string             `pulumi:"type"`
}

func LookupBookmarkRelationOutput(ctx *pulumi.Context, args LookupBookmarkRelationOutputArgs, opts ...pulumi.InvokeOption) LookupBookmarkRelationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBookmarkRelationResult, error) {
			args := v.(LookupBookmarkRelationArgs)
			r, err := LookupBookmarkRelation(ctx, &args, opts...)
			var s LookupBookmarkRelationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBookmarkRelationResultOutput)
}

type LookupBookmarkRelationOutputArgs struct {
	BookmarkId        pulumi.StringInput `pulumi:"bookmarkId"`
	RelationName      pulumi.StringInput `pulumi:"relationName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupBookmarkRelationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkRelationArgs)(nil)).Elem()
}


type LookupBookmarkRelationResultOutput struct{ *pulumi.OutputState }

func (LookupBookmarkRelationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBookmarkRelationResult)(nil)).Elem()
}

func (o LookupBookmarkRelationResultOutput) ToLookupBookmarkRelationResultOutput() LookupBookmarkRelationResultOutput {
	return o
}

func (o LookupBookmarkRelationResultOutput) ToLookupBookmarkRelationResultOutputWithContext(ctx context.Context) LookupBookmarkRelationResultOutput {
	return o
}

func (o LookupBookmarkRelationResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupBookmarkRelationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) RelatedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceId }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) RelatedResourceKind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceKind }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) RelatedResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceName }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) RelatedResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.RelatedResourceType }).(pulumi.StringOutput)
}

func (o LookupBookmarkRelationResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupBookmarkRelationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBookmarkRelationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBookmarkRelationResultOutput{})
}
