


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:compute/v20210701:getGalleryImage", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageArgs struct {
	GalleryImageName  string `pulumi:"galleryImageName"`
	GalleryName       string `pulumi:"galleryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupGalleryImageResult struct {
	Description         *string                                  `pulumi:"description"`
	Disallowed          *DisallowedResponse                      `pulumi:"disallowed"`
	EndOfLifeDate       *string                                  `pulumi:"endOfLifeDate"`
	Eula                *string                                  `pulumi:"eula"`
	Features            []GalleryImageFeatureResponse            `pulumi:"features"`
	HyperVGeneration    *string                                  `pulumi:"hyperVGeneration"`
	Id                  string                                   `pulumi:"id"`
	Identifier          GalleryImageIdentifierResponse           `pulumi:"identifier"`
	Location            string                                   `pulumi:"location"`
	Name                string                                   `pulumi:"name"`
	OsState             string                                   `pulumi:"osState"`
	OsType              string                                   `pulumi:"osType"`
	PrivacyStatementUri *string                                  `pulumi:"privacyStatementUri"`
	ProvisioningState   string                                   `pulumi:"provisioningState"`
	PurchasePlan        *ImagePurchasePlanResponse               `pulumi:"purchasePlan"`
	Recommended         *RecommendedMachineConfigurationResponse `pulumi:"recommended"`
	ReleaseNoteUri      *string                                  `pulumi:"releaseNoteUri"`
	Tags                map[string]string                        `pulumi:"tags"`
	Type                string                                   `pulumi:"type"`
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
	GalleryImageName  pulumi.StringInput `pulumi:"galleryImageName"`
	GalleryName       pulumi.StringInput `pulumi:"galleryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
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

func (o LookupGalleryImageResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Disallowed() DisallowedResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *DisallowedResponse { return v.Disallowed }).(DisallowedResponsePtrOutput)
}

func (o LookupGalleryImageResultOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Features() GalleryImageFeatureResponseArrayOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) []GalleryImageFeatureResponse { return v.Features }).(GalleryImageFeatureResponseArrayOutput)
}

func (o LookupGalleryImageResultOutput) HyperVGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.HyperVGeneration }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) Identifier() GalleryImageIdentifierResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) GalleryImageIdentifierResponse { return v.Identifier }).(GalleryImageIdentifierResponseOutput)
}

func (o LookupGalleryImageResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) OsState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.OsState }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) PrivacyStatementUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.PrivacyStatementUri }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGalleryImageResultOutput) PurchasePlan() ImagePurchasePlanResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *ImagePurchasePlanResponse { return v.PurchasePlan }).(ImagePurchasePlanResponsePtrOutput)
}

func (o LookupGalleryImageResultOutput) Recommended() RecommendedMachineConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *RecommendedMachineConfigurationResponse { return v.Recommended }).(RecommendedMachineConfigurationResponsePtrOutput)
}

func (o LookupGalleryImageResultOutput) ReleaseNoteUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) *string { return v.ReleaseNoteUri }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryImageResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryImageResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageResultOutput{})
}
