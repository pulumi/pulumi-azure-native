


package v20220601

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

type CompatibilityResponseOutput struct{ *pulumi.OutputState }

func (CompatibilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompatibilityResponse)(nil)).Elem()
}

func (o CompatibilityResponseOutput) ToCompatibilityResponseOutput() CompatibilityResponseOutput {
	return o
}

func (o CompatibilityResponseOutput) ToCompatibilityResponseOutputWithContext(ctx context.Context) CompatibilityResponseOutput {
	return o
}

func (o CompatibilityResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompatibilityResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CompatibilityResponseOutput) IsCompatible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CompatibilityResponse) *bool { return v.IsCompatible }).(pulumi.BoolPtrOutput)
}

func (o CompatibilityResponseOutput) Issues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CompatibilityResponse) []string { return v.Issues }).(pulumi.StringArrayOutput)
}

func (o CompatibilityResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompatibilityResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type CompatibilityResponsePtrOutput struct{ *pulumi.OutputState }

func (CompatibilityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompatibilityResponse)(nil)).Elem()
}

func (o CompatibilityResponsePtrOutput) ToCompatibilityResponsePtrOutput() CompatibilityResponsePtrOutput {
	return o
}

func (o CompatibilityResponsePtrOutput) ToCompatibilityResponsePtrOutputWithContext(ctx context.Context) CompatibilityResponsePtrOutput {
	return o
}

func (o CompatibilityResponsePtrOutput) Elem() CompatibilityResponseOutput {
	return o.ApplyT(func(v *CompatibilityResponse) CompatibilityResponse {
		if v != nil {
			return *v
		}
		var ret CompatibilityResponse
		return ret
	}).(CompatibilityResponseOutput)
}

func (o CompatibilityResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompatibilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CompatibilityResponsePtrOutput) IsCompatible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CompatibilityResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCompatible
	}).(pulumi.BoolPtrOutput)
}

func (o CompatibilityResponsePtrOutput) Issues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CompatibilityResponse) []string {
		if v == nil {
			return nil
		}
		return v.Issues
	}).(pulumi.StringArrayOutput)
}

func (o CompatibilityResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CompatibilityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type DataDiskImageResponse struct {
	Lun              int    `pulumi:"lun"`
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

type DataDiskImageResponseOutput struct{ *pulumi.OutputState }

func (DataDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageResponse)(nil)).Elem()
}

func (o DataDiskImageResponseOutput) ToDataDiskImageResponseOutput() DataDiskImageResponseOutput {
	return o
}

func (o DataDiskImageResponseOutput) ToDataDiskImageResponseOutputWithContext(ctx context.Context) DataDiskImageResponseOutput {
	return o
}

func (o DataDiskImageResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskImageResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o DataDiskImageResponseOutput) SourceBlobSasUri() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskImageResponse) string { return v.SourceBlobSasUri }).(pulumi.StringOutput)
}

type DataDiskImageResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskImageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageResponse)(nil)).Elem()
}

func (o DataDiskImageResponseArrayOutput) ToDataDiskImageResponseArrayOutput() DataDiskImageResponseArrayOutput {
	return o
}

func (o DataDiskImageResponseArrayOutput) ToDataDiskImageResponseArrayOutputWithContext(ctx context.Context) DataDiskImageResponseArrayOutput {
	return o
}

func (o DataDiskImageResponseArrayOutput) Index(i pulumi.IntInput) DataDiskImageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskImageResponse {
		return vs[0].([]DataDiskImageResponse)[vs[1].(int)]
	}).(DataDiskImageResponseOutput)
}

type IconUrisResponse struct {
	Hero   *string `pulumi:"hero"`
	Large  *string `pulumi:"large"`
	Medium *string `pulumi:"medium"`
	Small  *string `pulumi:"small"`
	Wide   *string `pulumi:"wide"`
}

type IconUrisResponseOutput struct{ *pulumi.OutputState }

func (IconUrisResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IconUrisResponse)(nil)).Elem()
}

func (o IconUrisResponseOutput) ToIconUrisResponseOutput() IconUrisResponseOutput {
	return o
}

func (o IconUrisResponseOutput) ToIconUrisResponseOutputWithContext(ctx context.Context) IconUrisResponseOutput {
	return o
}

func (o IconUrisResponseOutput) Hero() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IconUrisResponse) *string { return v.Hero }).(pulumi.StringPtrOutput)
}

func (o IconUrisResponseOutput) Large() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IconUrisResponse) *string { return v.Large }).(pulumi.StringPtrOutput)
}

func (o IconUrisResponseOutput) Medium() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IconUrisResponse) *string { return v.Medium }).(pulumi.StringPtrOutput)
}

func (o IconUrisResponseOutput) Small() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IconUrisResponse) *string { return v.Small }).(pulumi.StringPtrOutput)
}

func (o IconUrisResponseOutput) Wide() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IconUrisResponse) *string { return v.Wide }).(pulumi.StringPtrOutput)
}

type IconUrisResponsePtrOutput struct{ *pulumi.OutputState }

func (IconUrisResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IconUrisResponse)(nil)).Elem()
}

func (o IconUrisResponsePtrOutput) ToIconUrisResponsePtrOutput() IconUrisResponsePtrOutput {
	return o
}

func (o IconUrisResponsePtrOutput) ToIconUrisResponsePtrOutputWithContext(ctx context.Context) IconUrisResponsePtrOutput {
	return o
}

func (o IconUrisResponsePtrOutput) Elem() IconUrisResponseOutput {
	return o.ApplyT(func(v *IconUrisResponse) IconUrisResponse {
		if v != nil {
			return *v
		}
		var ret IconUrisResponse
		return ret
	}).(IconUrisResponseOutput)
}

func (o IconUrisResponsePtrOutput) Hero() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IconUrisResponse) *string {
		if v == nil {
			return nil
		}
		return v.Hero
	}).(pulumi.StringPtrOutput)
}

func (o IconUrisResponsePtrOutput) Large() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IconUrisResponse) *string {
		if v == nil {
			return nil
		}
		return v.Large
	}).(pulumi.StringPtrOutput)
}

func (o IconUrisResponsePtrOutput) Medium() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IconUrisResponse) *string {
		if v == nil {
			return nil
		}
		return v.Medium
	}).(pulumi.StringPtrOutput)
}

func (o IconUrisResponsePtrOutput) Small() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IconUrisResponse) *string {
		if v == nil {
			return nil
		}
		return v.Small
	}).(pulumi.StringPtrOutput)
}

func (o IconUrisResponsePtrOutput) Wide() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IconUrisResponse) *string {
		if v == nil {
			return nil
		}
		return v.Wide
	}).(pulumi.StringPtrOutput)
}

type OsDiskImageResponse struct {
	OperatingSystem  string `pulumi:"operatingSystem"`
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

type OsDiskImageResponseOutput struct{ *pulumi.OutputState }

func (OsDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDiskImageResponse)(nil)).Elem()
}

func (o OsDiskImageResponseOutput) ToOsDiskImageResponseOutput() OsDiskImageResponseOutput {
	return o
}

func (o OsDiskImageResponseOutput) ToOsDiskImageResponseOutputWithContext(ctx context.Context) OsDiskImageResponseOutput {
	return o
}

func (o OsDiskImageResponseOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v OsDiskImageResponse) string { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o OsDiskImageResponseOutput) SourceBlobSasUri() pulumi.StringOutput {
	return o.ApplyT(func(v OsDiskImageResponse) string { return v.SourceBlobSasUri }).(pulumi.StringOutput)
}

type ProductLinkResponse struct {
	DisplayName *string `pulumi:"displayName"`
	Uri         *string `pulumi:"uri"`
}

type ProductLinkResponseOutput struct{ *pulumi.OutputState }

func (ProductLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductLinkResponse)(nil)).Elem()
}

func (o ProductLinkResponseOutput) ToProductLinkResponseOutput() ProductLinkResponseOutput {
	return o
}

func (o ProductLinkResponseOutput) ToProductLinkResponseOutputWithContext(ctx context.Context) ProductLinkResponseOutput {
	return o
}

func (o ProductLinkResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductLinkResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ProductLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ProductLinkResponseArrayOutput struct{ *pulumi.OutputState }

func (ProductLinkResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductLinkResponse)(nil)).Elem()
}

func (o ProductLinkResponseArrayOutput) ToProductLinkResponseArrayOutput() ProductLinkResponseArrayOutput {
	return o
}

func (o ProductLinkResponseArrayOutput) ToProductLinkResponseArrayOutputWithContext(ctx context.Context) ProductLinkResponseArrayOutput {
	return o
}

func (o ProductLinkResponseArrayOutput) Index(i pulumi.IntInput) ProductLinkResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductLinkResponse {
		return vs[0].([]ProductLinkResponse)[vs[1].(int)]
	}).(ProductLinkResponseOutput)
}

type ProductPropertiesResponse struct {
	Version *string `pulumi:"version"`
}

type ProductPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProductPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductPropertiesResponse)(nil)).Elem()
}

func (o ProductPropertiesResponseOutput) ToProductPropertiesResponseOutput() ProductPropertiesResponseOutput {
	return o
}

func (o ProductPropertiesResponseOutput) ToProductPropertiesResponseOutputWithContext(ctx context.Context) ProductPropertiesResponseOutput {
	return o
}

func (o ProductPropertiesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductPropertiesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ProductPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProductPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProductPropertiesResponse)(nil)).Elem()
}

func (o ProductPropertiesResponsePtrOutput) ToProductPropertiesResponsePtrOutput() ProductPropertiesResponsePtrOutput {
	return o
}

func (o ProductPropertiesResponsePtrOutput) ToProductPropertiesResponsePtrOutputWithContext(ctx context.Context) ProductPropertiesResponsePtrOutput {
	return o
}

func (o ProductPropertiesResponsePtrOutput) Elem() ProductPropertiesResponseOutput {
	return o.ApplyT(func(v *ProductPropertiesResponse) ProductPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProductPropertiesResponse
		return ret
	}).(ProductPropertiesResponseOutput)
}

func (o ProductPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProductPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
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

type ProductResponseOutput struct{ *pulumi.OutputState }

func (ProductResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProductResponse)(nil)).Elem()
}

func (o ProductResponseOutput) ToProductResponseOutput() ProductResponseOutput {
	return o
}

func (o ProductResponseOutput) ToProductResponseOutputWithContext(ctx context.Context) ProductResponseOutput {
	return o
}

func (o ProductResponseOutput) BillingPartNumber() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.BillingPartNumber }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) Compatibility() CompatibilityResponsePtrOutput {
	return o.ApplyT(func(v ProductResponse) *CompatibilityResponse { return v.Compatibility }).(CompatibilityResponsePtrOutput)
}

func (o ProductResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) GalleryItemIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.GalleryItemIdentity }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) IconUris() IconUrisResponsePtrOutput {
	return o.ApplyT(func(v ProductResponse) *IconUrisResponse { return v.IconUris }).(IconUrisResponsePtrOutput)
}

func (o ProductResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ProductResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ProductResponseOutput) LegalTerms() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.LegalTerms }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) Links() ProductLinkResponseArrayOutput {
	return o.ApplyT(func(v ProductResponse) []ProductLinkResponse { return v.Links }).(ProductLinkResponseArrayOutput)
}

func (o ProductResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ProductResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ProductResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) OfferVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.OfferVersion }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) PayloadLength() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ProductResponse) *float64 { return v.PayloadLength }).(pulumi.Float64PtrOutput)
}

func (o ProductResponseOutput) PrivacyPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.PrivacyPolicy }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) ProductKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.ProductKind }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) ProductProperties() ProductPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ProductResponse) *ProductPropertiesResponse { return v.ProductProperties }).(ProductPropertiesResponsePtrOutput)
}

func (o ProductResponseOutput) PublisherDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.PublisherDisplayName }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) PublisherIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.PublisherIdentifier }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ProductResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ProductResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ProductResponseOutput) VmExtensionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProductResponse) *string { return v.VmExtensionType }).(pulumi.StringPtrOutput)
}

type ProductResponseArrayOutput struct{ *pulumi.OutputState }

func (ProductResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProductResponse)(nil)).Elem()
}

func (o ProductResponseArrayOutput) ToProductResponseArrayOutput() ProductResponseArrayOutput {
	return o
}

func (o ProductResponseArrayOutput) ToProductResponseArrayOutputWithContext(ctx context.Context) ProductResponseArrayOutput {
	return o
}

func (o ProductResponseArrayOutput) Index(i pulumi.IntInput) ProductResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProductResponse {
		return vs[0].([]ProductResponse)[vs[1].(int)]
	}).(ProductResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CompatibilityResponseOutput{})
	pulumi.RegisterOutputType(CompatibilityResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskImageResponseOutput{})
	pulumi.RegisterOutputType(DataDiskImageResponseArrayOutput{})
	pulumi.RegisterOutputType(IconUrisResponseOutput{})
	pulumi.RegisterOutputType(IconUrisResponsePtrOutput{})
	pulumi.RegisterOutputType(OsDiskImageResponseOutput{})
	pulumi.RegisterOutputType(ProductLinkResponseOutput{})
	pulumi.RegisterOutputType(ProductLinkResponseArrayOutput{})
	pulumi.RegisterOutputType(ProductPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProductPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ProductResponseOutput{})
	pulumi.RegisterOutputType(ProductResponseArrayOutput{})
}
