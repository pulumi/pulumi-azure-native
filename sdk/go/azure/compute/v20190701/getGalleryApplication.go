


package v20190701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupGalleryApplication(ctx *pulumi.Context, args *LookupGalleryApplicationArgs, opts ...pulumi.InvokeOption) (*LookupGalleryApplicationResult, error) {
	var rv LookupGalleryApplicationResult
	err := ctx.Invoke("azure-native:compute/v20190701:getGalleryApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryApplicationArgs struct {
	GalleryApplicationName string `pulumi:"galleryApplicationName"`
	GalleryName            string `pulumi:"galleryName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupGalleryApplicationResult struct {
	Description         *string           `pulumi:"description"`
	EndOfLifeDate       *string           `pulumi:"endOfLifeDate"`
	Eula                *string           `pulumi:"eula"`
	Id                  string            `pulumi:"id"`
	Location            string            `pulumi:"location"`
	Name                string            `pulumi:"name"`
	PrivacyStatementUri *string           `pulumi:"privacyStatementUri"`
	ReleaseNoteUri      *string           `pulumi:"releaseNoteUri"`
	SupportedOSType     string            `pulumi:"supportedOSType"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
}

func LookupGalleryApplicationOutput(ctx *pulumi.Context, args LookupGalleryApplicationOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryApplicationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryApplicationResult, error) {
			args := v.(LookupGalleryApplicationArgs)
			r, err := LookupGalleryApplication(ctx, &args, opts...)
			var s LookupGalleryApplicationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryApplicationResultOutput)
}

type LookupGalleryApplicationOutputArgs struct {
	GalleryApplicationName pulumi.StringInput `pulumi:"galleryApplicationName"`
	GalleryName            pulumi.StringInput `pulumi:"galleryName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryApplicationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryApplicationArgs)(nil)).Elem()
}


type LookupGalleryApplicationResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryApplicationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryApplicationResult)(nil)).Elem()
}

func (o LookupGalleryApplicationResultOutput) ToLookupGalleryApplicationResultOutput() LookupGalleryApplicationResultOutput {
	return o
}

func (o LookupGalleryApplicationResultOutput) ToLookupGalleryApplicationResultOutputWithContext(ctx context.Context) LookupGalleryApplicationResultOutput {
	return o
}

func (o LookupGalleryApplicationResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryApplicationResultOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryApplicationResultOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) *string { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryApplicationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationResultOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) *string { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryApplicationResultOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) *string { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryApplicationResultOutput) SupportedOSType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) string { return v.SupportedOSType }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryApplicationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryApplicationResultOutput{})
}
