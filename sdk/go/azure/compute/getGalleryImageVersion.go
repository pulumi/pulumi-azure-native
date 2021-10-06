


package compute

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryImageVersion(ctx *pulumi.Context, args *LookupGalleryImageVersionArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageVersionResult, error) {
	var rv LookupGalleryImageVersionResult
	err := ctx.Invoke("azure-native:compute:getGalleryImageVersion", args, &rv, opts...)
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
