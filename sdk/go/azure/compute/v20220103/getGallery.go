


package v20220103

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGallery(ctx *pulumi.Context, args *LookupGalleryArgs, opts ...pulumi.InvokeOption) (*LookupGalleryResult, error) {
	var rv LookupGalleryResult
	err := ctx.Invoke("azure-native:compute/v20220103:getGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryArgs struct {
	Expand            *string `pulumi:"expand"`
	GalleryName       string  `pulumi:"galleryName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	Select            *string `pulumi:"select"`
}


type LookupGalleryResult struct {
	Description       *string                    `pulumi:"description"`
	Id                string                     `pulumi:"id"`
	Identifier        *GalleryIdentifierResponse `pulumi:"identifier"`
	Location          string                     `pulumi:"location"`
	Name              string                     `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	SharingProfile    *SharingProfileResponse    `pulumi:"sharingProfile"`
	SharingStatus     SharingStatusResponse      `pulumi:"sharingStatus"`
	SoftDeletePolicy  *SoftDeletePolicyResponse  `pulumi:"softDeletePolicy"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
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
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	GalleryName       pulumi.StringInput    `pulumi:"galleryName"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	Select            pulumi.StringPtrInput `pulumi:"select"`
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

func (o LookupGalleryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGalleryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) Identifier() GalleryIdentifierResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *GalleryIdentifierResponse { return v.Identifier }).(GalleryIdentifierResponsePtrOutput)
}

func (o LookupGalleryResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGalleryResultOutput) SharingProfile() SharingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *SharingProfileResponse { return v.SharingProfile }).(SharingProfileResponsePtrOutput)
}

func (o LookupGalleryResultOutput) SharingStatus() SharingStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryResult) SharingStatusResponse { return v.SharingStatus }).(SharingStatusResponseOutput)
}

func (o LookupGalleryResultOutput) SoftDeletePolicy() SoftDeletePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryResult) *SoftDeletePolicyResponse { return v.SoftDeletePolicy }).(SoftDeletePolicyResponsePtrOutput)
}

func (o LookupGalleryResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryResultOutput{})
}
