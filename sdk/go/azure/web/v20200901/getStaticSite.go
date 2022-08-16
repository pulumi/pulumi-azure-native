


package v20200901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStaticSite(ctx *pulumi.Context, args *LookupStaticSiteArgs, opts ...pulumi.InvokeOption) (*LookupStaticSiteResult, error) {
	var rv LookupStaticSiteResult
	err := ctx.Invoke("azure-native:web/v20200901:getStaticSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStaticSiteArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupStaticSiteResult struct {
	Branch          *string                            `pulumi:"branch"`
	BuildProperties *StaticSiteBuildPropertiesResponse `pulumi:"buildProperties"`
	CustomDomains   []string                           `pulumi:"customDomains"`
	DefaultHostname string                             `pulumi:"defaultHostname"`
	Id              string                             `pulumi:"id"`
	Kind            *string                            `pulumi:"kind"`
	Location        string                             `pulumi:"location"`
	Name            string                             `pulumi:"name"`
	RepositoryToken *string                            `pulumi:"repositoryToken"`
	RepositoryUrl   *string                            `pulumi:"repositoryUrl"`
	Sku             *SkuDescriptionResponse            `pulumi:"sku"`
	SystemData      SystemDataResponse                 `pulumi:"systemData"`
	Tags            map[string]string                  `pulumi:"tags"`
	Type            string                             `pulumi:"type"`
}

func LookupStaticSiteOutput(ctx *pulumi.Context, args LookupStaticSiteOutputArgs, opts ...pulumi.InvokeOption) LookupStaticSiteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStaticSiteResult, error) {
			args := v.(LookupStaticSiteArgs)
			r, err := LookupStaticSite(ctx, &args, opts...)
			var s LookupStaticSiteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupStaticSiteResultOutput)
}

type LookupStaticSiteOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupStaticSiteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteArgs)(nil)).Elem()
}


type LookupStaticSiteResultOutput struct{ *pulumi.OutputState }

func (LookupStaticSiteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStaticSiteResult)(nil)).Elem()
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutput() LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) ToLookupStaticSiteResultOutputWithContext(ctx context.Context) LookupStaticSiteResultOutput {
	return o
}

func (o LookupStaticSiteResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) BuildProperties() StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *StaticSiteBuildPropertiesResponse { return v.BuildProperties }).(StaticSiteBuildPropertiesResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) CustomDomains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) []string { return v.CustomDomains }).(pulumi.StringArrayOutput)
}

func (o LookupStaticSiteResultOutput) DefaultHostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.DefaultHostname }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupStaticSiteResultOutput) RepositoryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryToken }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) RepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *string { return v.RepositoryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupStaticSiteResultOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) *SkuDescriptionResponse { return v.Sku }).(SkuDescriptionResponsePtrOutput)
}

func (o LookupStaticSiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupStaticSiteResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupStaticSiteResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStaticSiteResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStaticSiteResultOutput{})
}
