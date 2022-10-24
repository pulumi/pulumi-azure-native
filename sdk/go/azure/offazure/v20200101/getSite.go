


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:offazure/v20200101:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupSiteResult struct {
	ETag       *string                `pulumi:"eTag"`
	Id         string                 `pulumi:"id"`
	Location   *string                `pulumi:"location"`
	Name       *string                `pulumi:"name"`
	Properties SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}

func LookupSiteOutput(ctx *pulumi.Context, args LookupSiteOutputArgs, opts ...pulumi.InvokeOption) LookupSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteResult, error) {
			args := v.(LookupSiteArgs)
			r, err := LookupSite(ctx, &args, opts...)
			var s LookupSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteResultOutput)
}

type LookupSiteOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteName          pulumi.StringInput `pulumi:"siteName"`
}

func (LookupSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteArgs)(nil)).Elem()
}


type LookupSiteResultOutput struct{ *pulumi.OutputState }

func (LookupSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteResult)(nil)).Elem()
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutput() LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ToLookupSiteResultOutputWithContext(ctx context.Context) LookupSiteResultOutput {
	return o
}

func (o LookupSiteResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Properties() SitePropertiesResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SitePropertiesResponse { return v.Properties }).(SitePropertiesResponseOutput)
}

func (o LookupSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteResultOutput{})
}
