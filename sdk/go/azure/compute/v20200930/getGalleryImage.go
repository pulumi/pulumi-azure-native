


package v20200930

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryImage(ctx *pulumi.Context, args *LookupGalleryImageArgs, opts ...pulumi.InvokeOption) (*LookupGalleryImageResult, error) {
	var rv LookupGalleryImageResult
	err := ctx.Invoke("azure-native:compute/v20200930:getGalleryImage", args, &rv, opts...)
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
