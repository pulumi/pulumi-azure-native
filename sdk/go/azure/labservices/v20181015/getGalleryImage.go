


package v20181015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGalleryImage", args, &rv, opts...)
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
