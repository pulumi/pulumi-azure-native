


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHyperVSite(ctx *pulumi.Context, args *LookupHyperVSiteArgs, opts ...pulumi.InvokeOption) (*LookupHyperVSiteResult, error) {
	var rv LookupHyperVSiteResult
	err := ctx.Invoke("azure-native:offazure/v20200101:getHyperVSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHyperVSiteArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupHyperVSiteResult struct {
	ETag       *string                `pulumi:"eTag"`
	Id         string                 `pulumi:"id"`
	Location   *string                `pulumi:"location"`
	Name       *string                `pulumi:"name"`
	Properties SitePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string      `pulumi:"tags"`
	Type       string                 `pulumi:"type"`
}

func LookupHyperVSiteOutput(ctx *pulumi.Context, args LookupHyperVSiteOutputArgs, opts ...pulumi.InvokeOption) LookupHyperVSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHyperVSiteResult, error) {
			args := v.(LookupHyperVSiteArgs)
			r, err := LookupHyperVSite(ctx, &args, opts...)
			var s LookupHyperVSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHyperVSiteResultOutput)
}

type LookupHyperVSiteOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SiteName          pulumi.StringInput `pulumi:"siteName"`
}

func (LookupHyperVSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVSiteArgs)(nil)).Elem()
}


type LookupHyperVSiteResultOutput struct{ *pulumi.OutputState }

func (LookupHyperVSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHyperVSiteResult)(nil)).Elem()
}

func (o LookupHyperVSiteResultOutput) ToLookupHyperVSiteResultOutput() LookupHyperVSiteResultOutput {
	return o
}

func (o LookupHyperVSiteResultOutput) ToLookupHyperVSiteResultOutputWithContext(ctx context.Context) LookupHyperVSiteResultOutput {
	return o
}

func (o LookupHyperVSiteResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupHyperVSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHyperVSiteResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupHyperVSiteResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupHyperVSiteResultOutput) Properties() SitePropertiesResponseOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) SitePropertiesResponse { return v.Properties }).(SitePropertiesResponseOutput)
}

func (o LookupHyperVSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupHyperVSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHyperVSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHyperVSiteResultOutput{})
}
