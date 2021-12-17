


package v20160101

type CompatibilityResponse struct {
	Description  *string  `pulumi:"description"`
	IsCompatible *bool    `pulumi:"isCompatible"`
	Issues       []string `pulumi:"issues"`
	Message      *string  `pulumi:"message"`
}

type DataDiskImageResponse struct {
	Lun              int    `pulumi:"lun"`
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

type IconUrisResponse struct {
	Hero   *string `pulumi:"hero"`
	Large  *string `pulumi:"large"`
	Medium *string `pulumi:"medium"`
	Small  *string `pulumi:"small"`
	Wide   *string `pulumi:"wide"`
}

type OsDiskImageResponse struct {
	OperatingSystem  string `pulumi:"operatingSystem"`
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

type ProductLinkResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Uri         *string `pulumi:"uri"`
}

type ProductPropertiesResponse struct {
	Version *string `pulumi:"version"`
}

type ProductResponse struct {
	BillingPartNumber    *string                    `pulumi:"billingPartNumber"`
	Compatibility        *CompatibilityResponse     `pulumi:"compatibility"`
	Description          *string                    `pulumi:"description"`
	DisplayName          *string                    `pulumi:"displayName"`
	GalleryItemIdentity  *string                    `pulumi:"galleryItemIdentity"`
	IconUris             *IconUrisResponse          `pulumi:"iconUris"`
	Id                   string                     `pulumi:"id"`
	LegalTerms           *string                    `pulumi:"legalTerms"`
	Links                []ProductLinkResponse      `pulumi:"links"`
	Name                 string                     `pulumi:"name"`
	Offer                *string                    `pulumi:"offer"`
	OfferVersion         *string                    `pulumi:"offerVersion"`
	PayloadLength        *float64                   `pulumi:"payloadLength"`
	PrivacyPolicy        *string                    `pulumi:"privacyPolicy"`
	ProductKind          *string                    `pulumi:"productKind"`
	ProductProperties    *ProductPropertiesResponse `pulumi:"productProperties"`
	PublisherDisplayName *string                    `pulumi:"publisherDisplayName"`
	PublisherIdentifier  *string                    `pulumi:"publisherIdentifier"`
	Sku                  *string                    `pulumi:"sku"`
	Type                 string                     `pulumi:"type"`
	VmExtensionType      *string                    `pulumi:"vmExtensionType"`
}

func init() {
}
