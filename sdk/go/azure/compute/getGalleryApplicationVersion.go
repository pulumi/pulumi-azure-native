


package compute

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryApplicationVersion(ctx *pulumi.Context, args *LookupGalleryApplicationVersionArgs, opts ...pulumi.InvokeOption) (*LookupGalleryApplicationVersionResult, error) {
	var rv LookupGalleryApplicationVersionResult
	err := ctx.Invoke("azure-native:compute:getGalleryApplicationVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryApplicationVersionArgs struct {
	Expand                        *string `pulumi:"expand"`
	GalleryApplicationName        string  `pulumi:"galleryApplicationName"`
	GalleryApplicationVersionName string  `pulumi:"galleryApplicationVersionName"`
	GalleryName                   string  `pulumi:"galleryName"`
	ResourceGroupName             string  `pulumi:"resourceGroupName"`
}


type LookupGalleryApplicationVersionResult struct {
	Id                string                                             `pulumi:"id"`
	Location          string                                             `pulumi:"location"`
	Name              string                                             `pulumi:"name"`
	ProvisioningState string                                             `pulumi:"provisioningState"`
	PublishingProfile GalleryApplicationVersionPublishingProfileResponse `pulumi:"publishingProfile"`
	ReplicationStatus ReplicationStatusResponse                          `pulumi:"replicationStatus"`
	Tags              map[string]string                                  `pulumi:"tags"`
	Type              string                                             `pulumi:"type"`
}

func LookupGalleryApplicationVersionOutput(ctx *pulumi.Context, args LookupGalleryApplicationVersionOutputArgs, opts ...pulumi.InvokeOption) LookupGalleryApplicationVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGalleryApplicationVersionResult, error) {
			args := v.(LookupGalleryApplicationVersionArgs)
			r, err := LookupGalleryApplicationVersion(ctx, &args, opts...)
			var s LookupGalleryApplicationVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGalleryApplicationVersionResultOutput)
}

type LookupGalleryApplicationVersionOutputArgs struct {
	Expand                        pulumi.StringPtrInput `pulumi:"expand"`
	GalleryApplicationName        pulumi.StringInput    `pulumi:"galleryApplicationName"`
	GalleryApplicationVersionName pulumi.StringInput    `pulumi:"galleryApplicationVersionName"`
	GalleryName                   pulumi.StringInput    `pulumi:"galleryName"`
	ResourceGroupName             pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (LookupGalleryApplicationVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryApplicationVersionArgs)(nil)).Elem()
}


type LookupGalleryApplicationVersionResultOutput struct{ *pulumi.OutputState }

func (LookupGalleryApplicationVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGalleryApplicationVersionResult)(nil)).Elem()
}

func (o LookupGalleryApplicationVersionResultOutput) ToLookupGalleryApplicationVersionResultOutput() LookupGalleryApplicationVersionResultOutput {
	return o
}

func (o LookupGalleryApplicationVersionResultOutput) ToLookupGalleryApplicationVersionResultOutputWithContext(ctx context.Context) LookupGalleryApplicationVersionResultOutput {
	return o
}

func (o LookupGalleryApplicationVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) PublishingProfile() GalleryApplicationVersionPublishingProfileResponseOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) GalleryApplicationVersionPublishingProfileResponse {
		return v.PublishingProfile
	}).(GalleryApplicationVersionPublishingProfileResponseOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) ReplicationStatus() ReplicationStatusResponseOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) ReplicationStatusResponse { return v.ReplicationStatus }).(ReplicationStatusResponseOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupGalleryApplicationVersionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGalleryApplicationVersionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGalleryApplicationVersionResultOutput{})
}
