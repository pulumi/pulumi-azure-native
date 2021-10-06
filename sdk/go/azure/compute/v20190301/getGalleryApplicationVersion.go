


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryApplicationVersion(ctx *pulumi.Context, args *LookupGalleryApplicationVersionArgs, opts ...pulumi.InvokeOption) (*LookupGalleryApplicationVersionResult, error) {
	var rv LookupGalleryApplicationVersionResult
	err := ctx.Invoke("azure-native:compute/v20190301:getGalleryApplicationVersion", args, &rv, opts...)
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
