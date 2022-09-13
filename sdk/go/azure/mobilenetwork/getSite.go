


package mobilenetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSite(ctx *pulumi.Context, args *LookupSiteArgs, opts ...pulumi.InvokeOption) (*LookupSiteResult, error) {
	var rv LookupSiteResult
	err := ctx.Invoke("azure-native:mobilenetwork:getSite", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SiteName          string `pulumi:"siteName"`
}


type LookupSiteResult struct {
	CreatedAt          *string               `pulumi:"createdAt"`
	CreatedBy          *string               `pulumi:"createdBy"`
	CreatedByType      *string               `pulumi:"createdByType"`
	Id                 string                `pulumi:"id"`
	LastModifiedAt     *string               `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string               `pulumi:"lastModifiedBy"`
	LastModifiedByType *string               `pulumi:"lastModifiedByType"`
	Location           string                `pulumi:"location"`
	Name               string                `pulumi:"name"`
	NetworkFunctions   []SubResourceResponse `pulumi:"networkFunctions"`
	ProvisioningState  string                `pulumi:"provisioningState"`
	SystemData         SystemDataResponse    `pulumi:"systemData"`
	Tags               map[string]string     `pulumi:"tags"`
	Type               string                `pulumi:"type"`
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
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
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

func (o LookupSiteResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteResult) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func (o LookupSiteResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) NetworkFunctions() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupSiteResult) []SubResourceResponse { return v.NetworkFunctions }).(SubResourceResponseArrayOutput)
}

func (o LookupSiteResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupSiteResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSiteResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
