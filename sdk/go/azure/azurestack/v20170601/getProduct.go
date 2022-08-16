


package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	var rv GetProductResult
	err := ctx.Invoke("azure-native:azurestack/v20170601:getProduct", args, &rv, opts...)
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

func GetProductOutput(ctx *pulumi.Context, args GetProductOutputArgs, opts ...pulumi.InvokeOption) GetProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductResult, error) {
			args := v.(GetProductArgs)
			r, err := GetProduct(ctx, &args, opts...)
			var s GetProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductResultOutput)
}

type GetProductOutputArgs struct {
	ProductName      pulumi.StringInput `pulumi:"productName"`
	RegistrationName pulumi.StringInput `pulumi:"registrationName"`
	ResourceGroup    pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductArgs)(nil)).Elem()
}


type GetProductResultOutput struct{ *pulumi.OutputState }

func (GetProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductResult)(nil)).Elem()
}

func (o GetProductResultOutput) ToGetProductResultOutput() GetProductResultOutput {
	return o
}

func (o GetProductResultOutput) ToGetProductResultOutputWithContext(ctx context.Context) GetProductResultOutput {
	return o
}

func (o GetProductResultOutput) BillingPartNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.BillingPartNumber }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) Compatibility() CompatibilityResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *CompatibilityResponse { return v.Compatibility }).(CompatibilityResponsePtrOutput)
}

func (o GetProductResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) GalleryItemIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.GalleryItemIdentity }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) IconUris() IconUrisResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *IconUrisResponse { return v.IconUris }).(IconUrisResponsePtrOutput)
}

func (o GetProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) LegalTerms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.LegalTerms }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) Links() ProductLinkResponseArrayOutput {
	return o.ApplyT(func(v GetProductResult) []ProductLinkResponse { return v.Links }).(ProductLinkResponseArrayOutput)
}

func (o GetProductResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) OfferVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.OfferVersion }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) PayloadLength() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v GetProductResult) *float64 { return v.PayloadLength }).(pulumi.Float64PtrOutput)
}

func (o GetProductResultOutput) PrivacyPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PrivacyPolicy }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) ProductKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.ProductKind }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) ProductProperties() ProductPropertiesResponsePtrOutput {
	return o.ApplyT(func(v GetProductResult) *ProductPropertiesResponse { return v.ProductProperties }).(ProductPropertiesResponsePtrOutput)
}

func (o GetProductResultOutput) PublisherDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PublisherDisplayName }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) PublisherIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.PublisherIdentifier }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GetProductResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) VmExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetProductResult) *string { return v.VmExtensionType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductResultOutput{})
}
