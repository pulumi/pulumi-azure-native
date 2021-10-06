


package v20160101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	var rv GetProductResult
	err := ctx.Invoke("azure-native:azurestack/v20160101:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProductArgs struct {
	ProductName      string `pulumi:"productName"`
	RegistrationName string `pulumi:"registrationName"`
	ResourceGroup    string `pulumi:"resourceGroup"`
}


type GetProductResult struct {
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
