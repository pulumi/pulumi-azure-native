


package v20221111preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGallery(ctx *pulumi.Context, args *LookupGalleryArgs, opts ...pulumi.InvokeOption) (*LookupGalleryResult, error) {
	var rv LookupGalleryResult
	err := ctx.Invoke("azure-native:devcenter/v20221111preview:getGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryArgs struct {
	DevCenterName     string `pulumi:"devCenterName"`
	GalleryName       string `pulumi:"galleryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGalleryResult struct {
	GalleryResourceId string             `pulumi:"galleryResourceId"`
	Id                string             `pulumi:"id"`
	Name              string             `pulumi:"name"`
	ProvisioningState string             `pulumi:"provisioningState"`
	SystemData        SystemDataResponse `pulumi:"systemData"`
	Type              string             `pulumi:"type"`
}

func LookupGalleryOutput(ctx *pulumi.Context, args LookupGalleryOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryResult, error) {
			args := v.(LookupGalleryArgs)
			r, err := LookupGallery(ctx, &args, opts...)
			var s LookupGalleryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryResultOutput)
}

type LookupGalleryOutputArgs struct {
	DevCenterName     pulumi.StringInput `pulumi:"devCenterName"`
	GalleryName       pulumi.StringInput `pulumi:"galleryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGalleryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryArgs)(nil)).Elem()
}


type LookupGalleryResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryResult)(nil)).Elem()
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutput() LookupGalleryResultOutput {
	return o
}

func (o LookupGalleryResultOutput) ToLookupGalleryResultOutputWithContext(ctx context.Context) LookupGalleryResultOutput {
	return o
}

func (o LookupGalleryResultOutput) GalleryResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.GalleryResourceId }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupGalleryResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupGalleryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryResultOutput{})
}
