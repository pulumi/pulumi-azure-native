


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGalleryApplication(ctx *pulumi.Context, args *LookupGalleryApplicationArgs, opts ...pulumi.InvokeOption) (*LookupGalleryApplicationResult, error) {
	var rv LookupGalleryApplicationResult
	err := ctx.Invoke("azure-native:compute/v20190301:getGalleryApplication", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGalleryApplicationArgs struct {
	GalleryApplicationName string `pulumi:"galleryApplicationName"`
	GalleryName            string `pulumi:"galleryName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type LookupGalleryApplicationResult struct {
	Description         *string           `pulumi:"description"`
	EndOfLifeDate       *string           `pulumi:"endOfLifeDate"`
	Eula                *string           `pulumi:"eula"`
	Id                  string            `pulumi:"id"`
	Location            string            `pulumi:"location"`
	Name                string            `pulumi:"name"`
	PrivacyStatementUri *string           `pulumi:"privacyStatementUri"`
	ReleaseNoteUri      *string           `pulumi:"releaseNoteUri"`
	SupportedOSType     string            `pulumi:"supportedOSType"`
	Tags                map[string]string `pulumi:"tags"`
	Type                string            `pulumi:"type"`
}
