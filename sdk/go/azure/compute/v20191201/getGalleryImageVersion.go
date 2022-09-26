


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupGalleryImageVersion(ctx *pulumi.Context, args *LookupGalleryImageVersionArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageVersionResult, error) {
	var rv LookupGalleryImageVersionResult
	err := ctx.Invoke("azure-native:compute/v20191201:getGalleryImageVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryImageVersionArgs struct {
	Expand                  *string `pulumi:"expand"`
	GalleryImageName        string  `pulumi:"galleryImageName"`
	GalleryImageVersionName string  `pulumi:"galleryImageVersionName"`
	GalleryName             string  `pulumi:"galleryName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
}


type LookupGalleryImageVersionResult struct {
	Id                string                                        `pulumi:"id"`
	Location          string                                        `pulumi:"location"`
	Name              string                                        `pulumi:"name"`
	ProvisioningState string                                        `pulumi:"provisioningState"`
	PublishingProfile *GalleryImageVersionPublishingProfileResponse `pulumi:"publishingProfile"`
	ReplicationStatus ReplicationStatusResponse                     `pulumi:"replicationStatus"`
	StorageProfile    GalleryImageVersionStorageProfileResponse     `pulumi:"storageProfile"`
	Tags              map[string]string                             `pulumi:"tags"`
	Type              string                                        `pulumi:"type"`
}

func LookupGalleryImageVersionOutput(ctx *pulumi.Context, args LookupGalleryImageVersionOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryImageVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryImageVersionResult, error) {
			args := v.(LookupGalleryImageVersionArgs)
			r, err := LookupGalleryImageVersion(ctx, &args, opts...)
			var s LookupGalleryImageVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryImageVersionResultOutput)
}

type LookupGalleryImageVersionOutputArgs struct {
	Expand                  pulumi.StringPtrInput `pulumi:"expand"`
	GalleryImageName        pulumi.StringInput    `pulumi:"galleryImageName"`
	GalleryImageVersionName pulumi.StringInput    `pulumi:"galleryImageVersionName"`
	GalleryName             pulumi.StringInput    `pulumi:"galleryName"`
	ResourceGroupName       pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGalleryImageVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageVersionArgs)(nil)).Elem()
}


type LookupGalleryImageVersionResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryImageVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryImageVersionResult)(nil)).Elem()
}

func (o LookupGalleryImageVersionResultOutput) ToLookupGalleryImageVersionResultOutput() LookupGalleryImageVersionResultOutput {
	return o
}

func (o LookupGalleryImageVersionResultOutput) ToLookupGalleryImageVersionResultOutputWithContext(ctx context.Context) LookupGalleryImageVersionResultOutput {
	return o
}

func (o LookupGalleryImageVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryImageVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGalleryImageVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryImageVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGalleryImageVersionResultOutput) PublishingProfile() GalleryImageVersionPublishingProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) *GalleryImageVersionPublishingProfileResponse {
		return v.PublishingProfile
	}).(GalleryImageVersionPublishingProfileResponsePtrOutput)
}

func (o LookupGalleryImageVersionResultOutput) ReplicationStatus() ReplicationStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) ReplicationStatusResponse { return v.ReplicationStatus }).(ReplicationStatusResponseOutput)
}

func (o LookupGalleryImageVersionResultOutput) StorageProfile() GalleryImageVersionStorageProfileResponseOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) GalleryImageVersionStorageProfileResponse {
		return v.StorageProfile
	}).(GalleryImageVersionStorageProfileResponseOutput)
}

func (o LookupGalleryImageVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryImageVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryImageVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryImageVersionResultOutput{})
}
