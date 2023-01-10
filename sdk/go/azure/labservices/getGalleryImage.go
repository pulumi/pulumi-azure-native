


package labservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:labservices:getGalleryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageArgs struct {
	Expand            *string `pulumi:"expand"`
	GalleryImageName  string  `pulumi:"galleryImageName"`
	LabAccountName    string  `pulumi:"labAccountName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupGalleryImageResult struct {
	Author                string                        `pulumi:"author"`
	CreatedDate           string                        `pulumi:"createdDate"`
	Description           string                        `pulumi:"description"`
	Icon                  string                        `pulumi:"icon"`
	Id                    string                        `pulumi:"id"`
	ImageReference        GalleryImageReferenceResponse `pulumi:"imageReference"`
	IsEnabled             *bool                         `pulumi:"isEnabled"`
	IsOverride            *bool                         `pulumi:"isOverride"`
	IsPlanAuthorized      *bool                         `pulumi:"isPlanAuthorized"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Location              *string                       `pulumi:"location"`
	Name                  string                        `pulumi:"name"`
	PlanId                string                        `pulumi:"planId"`
	ProvisioningState     *string                       `pulumi:"provisioningState"`
	Tags                  map[string]string             `pulumi:"tags"`
	Type                  string                        `pulumi:"type"`
	UniqueIdentifier      *string                       `pulumi:"uniqueIdentifier"`
}

func LookupGalleryImageOutput(ctx *pulumi.Context, args LookupGalleryImageOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryImageResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryImageResult, error) {
			args := v.(LookupGalleryImageArgs)
			r, err := LookupGalleryImage(ctx, &args, opts...)
			var s LookupGalleryImageResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryImageResultOutput)
}

type LookupGalleryImageOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	GalleryImageName  pulumi.StringInput    `pulumi:"galleryImageName"`
	LabAccountName    pulumi.StringInput    `pulumi:"labAccountName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGalleryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageArgs)(nil)).Elem()
}


type LookupGalleryImageResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageResult)(nil)).Elem()
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutput() LookupGalleryImageResultOutput {
	return o
}

func (o LookupGalleryImageResultOutput) ToLookupGalleryImageResultOutputWithContext(ctx context.Context) LookupGalleryImageResultOutput {
	return o
}

func (o LookupGalleryImageResultOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Author }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) Icon() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Icon }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) ImageReference() GalleryImageReferenceResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) GalleryImageReferenceResponse { return v.ImageReference }).(GalleryImageReferenceResponseOutput)
}

func (o LookupGalleryImageResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupGalleryImageResultOutput) IsOverride() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsOverride }).(pulumi.BoolPtrOutput)
}

func (o LookupGalleryImageResultOutput) IsPlanAuthorized() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *bool { return v.IsPlanAuthorized }).(pulumi.BoolPtrOutput)
}

func (o LookupGalleryImageResultOutput) LatestOperationResult() LatestOperationResultResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) LatestOperationResultResponse { return v.LatestOperationResult }).(LatestOperationResultResponseOutput)
}

func (o LookupGalleryImageResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) UniqueIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.UniqueIdentifier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageResultOutput{})
}
