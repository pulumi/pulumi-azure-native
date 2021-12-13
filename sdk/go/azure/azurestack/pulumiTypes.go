


package azurestack

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Etag                 *string                    `pulumi:"etag"`
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
