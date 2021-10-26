


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGallery(ctx *pulumi.Context, args *LookupGalleryArgs, opts ...pulumi.InvokeOption) (*LookupGalleryResult, error) {
	var rv LookupGalleryResult
	err := ctx.Invoke("azure-native:compute/v20210701:getGallery", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryArgs struct {
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
	SoftDeletePolicy  *SoftDeletePolicyResponse  `pulumi:"softDeletePolicy"`
	Tags              map[string]string          `pulumi:"tags"`
	Type              string                     `pulumi:"type"`
}
