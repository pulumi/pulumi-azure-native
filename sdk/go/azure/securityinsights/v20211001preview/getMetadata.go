


package v20211001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetadata(ctx *pulumi.Context, args *LookupMetadataArgs, opts ...pulumi.InvokeOption) (*LookupMetadataResult, error) {
	var rv LookupMetadataResult
	err := ctx.Invoke("azure-native:securityinsights/v20211001preview:getMetadata", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetadataArgs struct {
	MetadataName      string `pulumi:"metadataName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMetadataResult struct {
	Author           *MetadataAuthorResponse       `pulumi:"author"`
	Categories       *MetadataCategoriesResponse   `pulumi:"categories"`
	ContentId        *string                       `pulumi:"contentId"`
	Dependencies     *MetadataDependenciesResponse `pulumi:"dependencies"`
	Etag             *string                       `pulumi:"etag"`
	FirstPublishDate *string                       `pulumi:"firstPublishDate"`
	Id               string                        `pulumi:"id"`
	Kind             string                        `pulumi:"kind"`
	LastPublishDate  *string                       `pulumi:"lastPublishDate"`
	Name             string                        `pulumi:"name"`
	ParentId         string                        `pulumi:"parentId"`
	Providers        []string                      `pulumi:"providers"`
	Source           *MetadataSourceResponse       `pulumi:"source"`
	Support          *MetadataSupportResponse      `pulumi:"support"`
	SystemData       SystemDataResponse            `pulumi:"systemData"`
	Type             string                        `pulumi:"type"`
	Version          *string                       `pulumi:"version"`
}

func LookupMetadataOutput(ctx *pulumi.Context, args LookupMetadataOutputArgs, opts ...pulumi.InvokeOption) LookupMetadataResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetadataResult, error) {
			args := v.(LookupMetadataArgs)
			r, err := LookupMetadata(ctx, &args, opts...)
			var s LookupMetadataResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetadataResultOutput)
}

type LookupMetadataOutputArgs struct {
	MetadataName      pulumi.StringInput `pulumi:"metadataName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMetadataOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataArgs)(nil)).Elem()
}


type LookupMetadataResultOutput struct{ *pulumi.OutputState }

func (LookupMetadataResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetadataResult)(nil)).Elem()
}

func (o LookupMetadataResultOutput) ToLookupMetadataResultOutput() LookupMetadataResultOutput {
	return o
}

func (o LookupMetadataResultOutput) ToLookupMetadataResultOutputWithContext(ctx context.Context) LookupMetadataResultOutput {
	return o
}

func (o LookupMetadataResultOutput) Author() MetadataAuthorResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataAuthorResponse { return v.Author }).(MetadataAuthorResponsePtrOutput)
}

func (o LookupMetadataResultOutput) Categories() MetadataCategoriesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataCategoriesResponse { return v.Categories }).(MetadataCategoriesResponsePtrOutput)
}

func (o LookupMetadataResultOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.ContentId }).(pulumi.StringPtrOutput)
}

func (o LookupMetadataResultOutput) Dependencies() MetadataDependenciesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataDependenciesResponse { return v.Dependencies }).(MetadataDependenciesResponsePtrOutput)
}

func (o LookupMetadataResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMetadataResultOutput) FirstPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.FirstPublishDate }).(pulumi.StringPtrOutput)
}

func (o LookupMetadataResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetadataResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMetadataResultOutput) LastPublishDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.LastPublishDate }).(pulumi.StringPtrOutput)
}

func (o LookupMetadataResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMetadataResultOutput) ParentId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.ParentId }).(pulumi.StringOutput)
}

func (o LookupMetadataResultOutput) Providers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMetadataResult) []string { return v.Providers }).(pulumi.StringArrayOutput)
}

func (o LookupMetadataResultOutput) Source() MetadataSourceResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataSourceResponse { return v.Source }).(MetadataSourceResponsePtrOutput)
}

func (o LookupMetadataResultOutput) Support() MetadataSupportResponsePtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *MetadataSupportResponse { return v.Support }).(MetadataSupportResponsePtrOutput)
}

func (o LookupMetadataResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetadataResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMetadataResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetadataResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupMetadataResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMetadataResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetadataResultOutput{})
}
