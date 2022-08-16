


package v20220103

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommunityGalleryInfo struct {
	Eula             *string `pulumi:"eula"`
	PublicNamePrefix *string `pulumi:"publicNamePrefix"`
	PublisherContact *string `pulumi:"publisherContact"`
	PublisherUri     *string `pulumi:"publisherUri"`
}





type CommunityGalleryInfoInput interface {
	pulumi.Input

	ToCommunityGalleryInfoOutput() CommunityGalleryInfoOutput
	ToCommunityGalleryInfoOutputWithContext(context.Context) CommunityGalleryInfoOutput
}

type CommunityGalleryInfoArgs struct {
	Eula             pulumi.StringPtrInput `pulumi:"eula"`
	PublicNamePrefix pulumi.StringPtrInput `pulumi:"publicNamePrefix"`
	PublisherContact pulumi.StringPtrInput `pulumi:"publisherContact"`
	PublisherUri     pulumi.StringPtrInput `pulumi:"publisherUri"`
}

func (CommunityGalleryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunityGalleryInfo)(nil)).Elem()
}

func (i CommunityGalleryInfoArgs) ToCommunityGalleryInfoOutput() CommunityGalleryInfoOutput {
	return i.ToCommunityGalleryInfoOutputWithContext(context.Background())
}

func (i CommunityGalleryInfoArgs) ToCommunityGalleryInfoOutputWithContext(ctx context.Context) CommunityGalleryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityGalleryInfoOutput)
}

func (i CommunityGalleryInfoArgs) ToCommunityGalleryInfoPtrOutput() CommunityGalleryInfoPtrOutput {
	return i.ToCommunityGalleryInfoPtrOutputWithContext(context.Background())
}

func (i CommunityGalleryInfoArgs) ToCommunityGalleryInfoPtrOutputWithContext(ctx context.Context) CommunityGalleryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityGalleryInfoOutput).ToCommunityGalleryInfoPtrOutputWithContext(ctx)
}









type CommunityGalleryInfoPtrInput interface {
	pulumi.Input

	ToCommunityGalleryInfoPtrOutput() CommunityGalleryInfoPtrOutput
	ToCommunityGalleryInfoPtrOutputWithContext(context.Context) CommunityGalleryInfoPtrOutput
}

type communityGalleryInfoPtrType CommunityGalleryInfoArgs

func CommunityGalleryInfoPtr(v *CommunityGalleryInfoArgs) CommunityGalleryInfoPtrInput {
	return (*communityGalleryInfoPtrType)(v)
}

func (*communityGalleryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunityGalleryInfo)(nil)).Elem()
}

func (i *communityGalleryInfoPtrType) ToCommunityGalleryInfoPtrOutput() CommunityGalleryInfoPtrOutput {
	return i.ToCommunityGalleryInfoPtrOutputWithContext(context.Background())
}

func (i *communityGalleryInfoPtrType) ToCommunityGalleryInfoPtrOutputWithContext(ctx context.Context) CommunityGalleryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommunityGalleryInfoPtrOutput)
}

type CommunityGalleryInfoOutput struct{ *pulumi.OutputState }

func (CommunityGalleryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunityGalleryInfo)(nil)).Elem()
}

func (o CommunityGalleryInfoOutput) ToCommunityGalleryInfoOutput() CommunityGalleryInfoOutput {
	return o
}

func (o CommunityGalleryInfoOutput) ToCommunityGalleryInfoOutputWithContext(ctx context.Context) CommunityGalleryInfoOutput {
	return o
}

func (o CommunityGalleryInfoOutput) ToCommunityGalleryInfoPtrOutput() CommunityGalleryInfoPtrOutput {
	return o.ToCommunityGalleryInfoPtrOutputWithContext(context.Background())
}

func (o CommunityGalleryInfoOutput) ToCommunityGalleryInfoPtrOutputWithContext(ctx context.Context) CommunityGalleryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CommunityGalleryInfo) *CommunityGalleryInfo {
		return &v
	}).(CommunityGalleryInfoPtrOutput)
}

func (o CommunityGalleryInfoOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfo) *string { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoOutput) PublicNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfo) *string { return v.PublicNamePrefix }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoOutput) PublisherContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfo) *string { return v.PublisherContact }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoOutput) PublisherUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfo) *string { return v.PublisherUri }).(pulumi.StringPtrOutput)
}

type CommunityGalleryInfoPtrOutput struct{ *pulumi.OutputState }

func (CommunityGalleryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunityGalleryInfo)(nil)).Elem()
}

func (o CommunityGalleryInfoPtrOutput) ToCommunityGalleryInfoPtrOutput() CommunityGalleryInfoPtrOutput {
	return o
}

func (o CommunityGalleryInfoPtrOutput) ToCommunityGalleryInfoPtrOutputWithContext(ctx context.Context) CommunityGalleryInfoPtrOutput {
	return o
}

func (o CommunityGalleryInfoPtrOutput) Elem() CommunityGalleryInfoOutput {
	return o.ApplyT(func(v *CommunityGalleryInfo) CommunityGalleryInfo {
		if v != nil {
			return *v
		}
		var ret CommunityGalleryInfo
		return ret
	}).(CommunityGalleryInfoOutput)
}

func (o CommunityGalleryInfoPtrOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfo) *string {
		if v == nil {
			return nil
		}
		return v.Eula
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoPtrOutput) PublicNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfo) *string {
		if v == nil {
			return nil
		}
		return v.PublicNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoPtrOutput) PublisherContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfo) *string {
		if v == nil {
			return nil
		}
		return v.PublisherContact
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoPtrOutput) PublisherUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfo) *string {
		if v == nil {
			return nil
		}
		return v.PublisherUri
	}).(pulumi.StringPtrOutput)
}

type CommunityGalleryInfoResponse struct {
	CommunityGalleryEnabled bool     `pulumi:"communityGalleryEnabled"`
	Eula                    *string  `pulumi:"eula"`
	PublicNamePrefix        *string  `pulumi:"publicNamePrefix"`
	PublicNames             []string `pulumi:"publicNames"`
	PublisherContact        *string  `pulumi:"publisherContact"`
	PublisherUri            *string  `pulumi:"publisherUri"`
}

type CommunityGalleryInfoResponseOutput struct{ *pulumi.OutputState }

func (CommunityGalleryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommunityGalleryInfoResponse)(nil)).Elem()
}

func (o CommunityGalleryInfoResponseOutput) ToCommunityGalleryInfoResponseOutput() CommunityGalleryInfoResponseOutput {
	return o
}

func (o CommunityGalleryInfoResponseOutput) ToCommunityGalleryInfoResponseOutputWithContext(ctx context.Context) CommunityGalleryInfoResponseOutput {
	return o
}

func (o CommunityGalleryInfoResponseOutput) CommunityGalleryEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) bool { return v.CommunityGalleryEnabled }).(pulumi.BoolOutput)
}

func (o CommunityGalleryInfoResponseOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) *string { return v.Eula }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponseOutput) PublicNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) *string { return v.PublicNamePrefix }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponseOutput) PublicNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) []string { return v.PublicNames }).(pulumi.StringArrayOutput)
}

func (o CommunityGalleryInfoResponseOutput) PublisherContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) *string { return v.PublisherContact }).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponseOutput) PublisherUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommunityGalleryInfoResponse) *string { return v.PublisherUri }).(pulumi.StringPtrOutput)
}

type CommunityGalleryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (CommunityGalleryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommunityGalleryInfoResponse)(nil)).Elem()
}

func (o CommunityGalleryInfoResponsePtrOutput) ToCommunityGalleryInfoResponsePtrOutput() CommunityGalleryInfoResponsePtrOutput {
	return o
}

func (o CommunityGalleryInfoResponsePtrOutput) ToCommunityGalleryInfoResponsePtrOutputWithContext(ctx context.Context) CommunityGalleryInfoResponsePtrOutput {
	return o
}

func (o CommunityGalleryInfoResponsePtrOutput) Elem() CommunityGalleryInfoResponseOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) CommunityGalleryInfoResponse {
		if v != nil {
			return *v
		}
		var ret CommunityGalleryInfoResponse
		return ret
	}).(CommunityGalleryInfoResponseOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) CommunityGalleryEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CommunityGalleryEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) Eula() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Eula
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) PublicNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNamePrefix
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) PublicNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.PublicNames
	}).(pulumi.StringArrayOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) PublisherContact() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublisherContact
	}).(pulumi.StringPtrOutput)
}

func (o CommunityGalleryInfoResponsePtrOutput) PublisherUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommunityGalleryInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublisherUri
	}).(pulumi.StringPtrOutput)
}

type DataDiskImageEncryption struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	Lun                 int     `pulumi:"lun"`
}





type DataDiskImageEncryptionInput interface {
	pulumi.Input

	ToDataDiskImageEncryptionOutput() DataDiskImageEncryptionOutput
	ToDataDiskImageEncryptionOutputWithContext(context.Context) DataDiskImageEncryptionOutput
}

type DataDiskImageEncryptionArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
	Lun                 pulumi.IntInput       `pulumi:"lun"`
}

func (DataDiskImageEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageEncryption)(nil)).Elem()
}

func (i DataDiskImageEncryptionArgs) ToDataDiskImageEncryptionOutput() DataDiskImageEncryptionOutput {
	return i.ToDataDiskImageEncryptionOutputWithContext(context.Background())
}

func (i DataDiskImageEncryptionArgs) ToDataDiskImageEncryptionOutputWithContext(ctx context.Context) DataDiskImageEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskImageEncryptionOutput)
}





type DataDiskImageEncryptionArrayInput interface {
	pulumi.Input

	ToDataDiskImageEncryptionArrayOutput() DataDiskImageEncryptionArrayOutput
	ToDataDiskImageEncryptionArrayOutputWithContext(context.Context) DataDiskImageEncryptionArrayOutput
}

type DataDiskImageEncryptionArray []DataDiskImageEncryptionInput

func (DataDiskImageEncryptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageEncryption)(nil)).Elem()
}

func (i DataDiskImageEncryptionArray) ToDataDiskImageEncryptionArrayOutput() DataDiskImageEncryptionArrayOutput {
	return i.ToDataDiskImageEncryptionArrayOutputWithContext(context.Background())
}

func (i DataDiskImageEncryptionArray) ToDataDiskImageEncryptionArrayOutputWithContext(ctx context.Context) DataDiskImageEncryptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskImageEncryptionArrayOutput)
}

type DataDiskImageEncryptionOutput struct{ *pulumi.OutputState }

func (DataDiskImageEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageEncryption)(nil)).Elem()
}

func (o DataDiskImageEncryptionOutput) ToDataDiskImageEncryptionOutput() DataDiskImageEncryptionOutput {
	return o
}

func (o DataDiskImageEncryptionOutput) ToDataDiskImageEncryptionOutputWithContext(ctx context.Context) DataDiskImageEncryptionOutput {
	return o
}

func (o DataDiskImageEncryptionOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskImageEncryption) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o DataDiskImageEncryptionOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskImageEncryption) int { return v.Lun }).(pulumi.IntOutput)
}

type DataDiskImageEncryptionArrayOutput struct{ *pulumi.OutputState }

func (DataDiskImageEncryptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageEncryption)(nil)).Elem()
}

func (o DataDiskImageEncryptionArrayOutput) ToDataDiskImageEncryptionArrayOutput() DataDiskImageEncryptionArrayOutput {
	return o
}

func (o DataDiskImageEncryptionArrayOutput) ToDataDiskImageEncryptionArrayOutputWithContext(ctx context.Context) DataDiskImageEncryptionArrayOutput {
	return o
}

func (o DataDiskImageEncryptionArrayOutput) Index(i pulumi.IntInput) DataDiskImageEncryptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskImageEncryption {
		return vs[0].([]DataDiskImageEncryption)[vs[1].(int)]
	}).(DataDiskImageEncryptionOutput)
}

type DataDiskImageEncryptionResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	Lun                 int     `pulumi:"lun"`
}

type DataDiskImageEncryptionResponseOutput struct{ *pulumi.OutputState }

func (DataDiskImageEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageEncryptionResponse)(nil)).Elem()
}

func (o DataDiskImageEncryptionResponseOutput) ToDataDiskImageEncryptionResponseOutput() DataDiskImageEncryptionResponseOutput {
	return o
}

func (o DataDiskImageEncryptionResponseOutput) ToDataDiskImageEncryptionResponseOutputWithContext(ctx context.Context) DataDiskImageEncryptionResponseOutput {
	return o
}

func (o DataDiskImageEncryptionResponseOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskImageEncryptionResponse) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o DataDiskImageEncryptionResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskImageEncryptionResponse) int { return v.Lun }).(pulumi.IntOutput)
}

type DataDiskImageEncryptionResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskImageEncryptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageEncryptionResponse)(nil)).Elem()
}

func (o DataDiskImageEncryptionResponseArrayOutput) ToDataDiskImageEncryptionResponseArrayOutput() DataDiskImageEncryptionResponseArrayOutput {
	return o
}

func (o DataDiskImageEncryptionResponseArrayOutput) ToDataDiskImageEncryptionResponseArrayOutputWithContext(ctx context.Context) DataDiskImageEncryptionResponseArrayOutput {
	return o
}

func (o DataDiskImageEncryptionResponseArrayOutput) Index(i pulumi.IntInput) DataDiskImageEncryptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskImageEncryptionResponse {
		return vs[0].([]DataDiskImageEncryptionResponse)[vs[1].(int)]
	}).(DataDiskImageEncryptionResponseOutput)
}

type Disallowed struct {
	DiskTypes []string `pulumi:"diskTypes"`
}





type DisallowedInput interface {
	pulumi.Input

	ToDisallowedOutput() DisallowedOutput
	ToDisallowedOutputWithContext(context.Context) DisallowedOutput
}

type DisallowedArgs struct {
	DiskTypes pulumi.StringArrayInput `pulumi:"diskTypes"`
}

func (DisallowedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Disallowed)(nil)).Elem()
}

func (i DisallowedArgs) ToDisallowedOutput() DisallowedOutput {
	return i.ToDisallowedOutputWithContext(context.Background())
}

func (i DisallowedArgs) ToDisallowedOutputWithContext(ctx context.Context) DisallowedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisallowedOutput)
}

func (i DisallowedArgs) ToDisallowedPtrOutput() DisallowedPtrOutput {
	return i.ToDisallowedPtrOutputWithContext(context.Background())
}

func (i DisallowedArgs) ToDisallowedPtrOutputWithContext(ctx context.Context) DisallowedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisallowedOutput).ToDisallowedPtrOutputWithContext(ctx)
}









type DisallowedPtrInput interface {
	pulumi.Input

	ToDisallowedPtrOutput() DisallowedPtrOutput
	ToDisallowedPtrOutputWithContext(context.Context) DisallowedPtrOutput
}

type disallowedPtrType DisallowedArgs

func DisallowedPtr(v *DisallowedArgs) DisallowedPtrInput {
	return (*disallowedPtrType)(v)
}

func (*disallowedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Disallowed)(nil)).Elem()
}

func (i *disallowedPtrType) ToDisallowedPtrOutput() DisallowedPtrOutput {
	return i.ToDisallowedPtrOutputWithContext(context.Background())
}

func (i *disallowedPtrType) ToDisallowedPtrOutputWithContext(ctx context.Context) DisallowedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DisallowedPtrOutput)
}

type DisallowedOutput struct{ *pulumi.OutputState }

func (DisallowedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Disallowed)(nil)).Elem()
}

func (o DisallowedOutput) ToDisallowedOutput() DisallowedOutput {
	return o
}

func (o DisallowedOutput) ToDisallowedOutputWithContext(ctx context.Context) DisallowedOutput {
	return o
}

func (o DisallowedOutput) ToDisallowedPtrOutput() DisallowedPtrOutput {
	return o.ToDisallowedPtrOutputWithContext(context.Background())
}

func (o DisallowedOutput) ToDisallowedPtrOutputWithContext(ctx context.Context) DisallowedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Disallowed) *Disallowed {
		return &v
	}).(DisallowedPtrOutput)
}

func (o DisallowedOutput) DiskTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Disallowed) []string { return v.DiskTypes }).(pulumi.StringArrayOutput)
}

type DisallowedPtrOutput struct{ *pulumi.OutputState }

func (DisallowedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Disallowed)(nil)).Elem()
}

func (o DisallowedPtrOutput) ToDisallowedPtrOutput() DisallowedPtrOutput {
	return o
}

func (o DisallowedPtrOutput) ToDisallowedPtrOutputWithContext(ctx context.Context) DisallowedPtrOutput {
	return o
}

func (o DisallowedPtrOutput) Elem() DisallowedOutput {
	return o.ApplyT(func(v *Disallowed) Disallowed {
		if v != nil {
			return *v
		}
		var ret Disallowed
		return ret
	}).(DisallowedOutput)
}

func (o DisallowedPtrOutput) DiskTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Disallowed) []string {
		if v == nil {
			return nil
		}
		return v.DiskTypes
	}).(pulumi.StringArrayOutput)
}

type DisallowedResponse struct {
	DiskTypes []string `pulumi:"diskTypes"`
}

type DisallowedResponseOutput struct{ *pulumi.OutputState }

func (DisallowedResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DisallowedResponse)(nil)).Elem()
}

func (o DisallowedResponseOutput) ToDisallowedResponseOutput() DisallowedResponseOutput {
	return o
}

func (o DisallowedResponseOutput) ToDisallowedResponseOutputWithContext(ctx context.Context) DisallowedResponseOutput {
	return o
}

func (o DisallowedResponseOutput) DiskTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DisallowedResponse) []string { return v.DiskTypes }).(pulumi.StringArrayOutput)
}

type DisallowedResponsePtrOutput struct{ *pulumi.OutputState }

func (DisallowedResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DisallowedResponse)(nil)).Elem()
}

func (o DisallowedResponsePtrOutput) ToDisallowedResponsePtrOutput() DisallowedResponsePtrOutput {
	return o
}

func (o DisallowedResponsePtrOutput) ToDisallowedResponsePtrOutputWithContext(ctx context.Context) DisallowedResponsePtrOutput {
	return o
}

func (o DisallowedResponsePtrOutput) Elem() DisallowedResponseOutput {
	return o.ApplyT(func(v *DisallowedResponse) DisallowedResponse {
		if v != nil {
			return *v
		}
		var ret DisallowedResponse
		return ret
	}).(DisallowedResponseOutput)
}

func (o DisallowedResponsePtrOutput) DiskTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DisallowedResponse) []string {
		if v == nil {
			return nil
		}
		return v.DiskTypes
	}).(pulumi.StringArrayOutput)
}

type EncryptionImages struct {
	DataDiskImages []DataDiskImageEncryption `pulumi:"dataDiskImages"`
	OsDiskImage    *OSDiskImageEncryption    `pulumi:"osDiskImage"`
}





type EncryptionImagesInput interface {
	pulumi.Input

	ToEncryptionImagesOutput() EncryptionImagesOutput
	ToEncryptionImagesOutputWithContext(context.Context) EncryptionImagesOutput
}

type EncryptionImagesArgs struct {
	DataDiskImages DataDiskImageEncryptionArrayInput `pulumi:"dataDiskImages"`
	OsDiskImage    OSDiskImageEncryptionPtrInput     `pulumi:"osDiskImage"`
}

func (EncryptionImagesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionImages)(nil)).Elem()
}

func (i EncryptionImagesArgs) ToEncryptionImagesOutput() EncryptionImagesOutput {
	return i.ToEncryptionImagesOutputWithContext(context.Background())
}

func (i EncryptionImagesArgs) ToEncryptionImagesOutputWithContext(ctx context.Context) EncryptionImagesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionImagesOutput)
}

func (i EncryptionImagesArgs) ToEncryptionImagesPtrOutput() EncryptionImagesPtrOutput {
	return i.ToEncryptionImagesPtrOutputWithContext(context.Background())
}

func (i EncryptionImagesArgs) ToEncryptionImagesPtrOutputWithContext(ctx context.Context) EncryptionImagesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionImagesOutput).ToEncryptionImagesPtrOutputWithContext(ctx)
}









type EncryptionImagesPtrInput interface {
	pulumi.Input

	ToEncryptionImagesPtrOutput() EncryptionImagesPtrOutput
	ToEncryptionImagesPtrOutputWithContext(context.Context) EncryptionImagesPtrOutput
}

type encryptionImagesPtrType EncryptionImagesArgs

func EncryptionImagesPtr(v *EncryptionImagesArgs) EncryptionImagesPtrInput {
	return (*encryptionImagesPtrType)(v)
}

func (*encryptionImagesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionImages)(nil)).Elem()
}

func (i *encryptionImagesPtrType) ToEncryptionImagesPtrOutput() EncryptionImagesPtrOutput {
	return i.ToEncryptionImagesPtrOutputWithContext(context.Background())
}

func (i *encryptionImagesPtrType) ToEncryptionImagesPtrOutputWithContext(ctx context.Context) EncryptionImagesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionImagesPtrOutput)
}

type EncryptionImagesOutput struct{ *pulumi.OutputState }

func (EncryptionImagesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionImages)(nil)).Elem()
}

func (o EncryptionImagesOutput) ToEncryptionImagesOutput() EncryptionImagesOutput {
	return o
}

func (o EncryptionImagesOutput) ToEncryptionImagesOutputWithContext(ctx context.Context) EncryptionImagesOutput {
	return o
}

func (o EncryptionImagesOutput) ToEncryptionImagesPtrOutput() EncryptionImagesPtrOutput {
	return o.ToEncryptionImagesPtrOutputWithContext(context.Background())
}

func (o EncryptionImagesOutput) ToEncryptionImagesPtrOutputWithContext(ctx context.Context) EncryptionImagesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionImages) *EncryptionImages {
		return &v
	}).(EncryptionImagesPtrOutput)
}

func (o EncryptionImagesOutput) DataDiskImages() DataDiskImageEncryptionArrayOutput {
	return o.ApplyT(func(v EncryptionImages) []DataDiskImageEncryption { return v.DataDiskImages }).(DataDiskImageEncryptionArrayOutput)
}

func (o EncryptionImagesOutput) OsDiskImage() OSDiskImageEncryptionPtrOutput {
	return o.ApplyT(func(v EncryptionImages) *OSDiskImageEncryption { return v.OsDiskImage }).(OSDiskImageEncryptionPtrOutput)
}

type EncryptionImagesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionImagesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionImages)(nil)).Elem()
}

func (o EncryptionImagesPtrOutput) ToEncryptionImagesPtrOutput() EncryptionImagesPtrOutput {
	return o
}

func (o EncryptionImagesPtrOutput) ToEncryptionImagesPtrOutputWithContext(ctx context.Context) EncryptionImagesPtrOutput {
	return o
}

func (o EncryptionImagesPtrOutput) Elem() EncryptionImagesOutput {
	return o.ApplyT(func(v *EncryptionImages) EncryptionImages {
		if v != nil {
			return *v
		}
		var ret EncryptionImages
		return ret
	}).(EncryptionImagesOutput)
}

func (o EncryptionImagesPtrOutput) DataDiskImages() DataDiskImageEncryptionArrayOutput {
	return o.ApplyT(func(v *EncryptionImages) []DataDiskImageEncryption {
		if v == nil {
			return nil
		}
		return v.DataDiskImages
	}).(DataDiskImageEncryptionArrayOutput)
}

func (o EncryptionImagesPtrOutput) OsDiskImage() OSDiskImageEncryptionPtrOutput {
	return o.ApplyT(func(v *EncryptionImages) *OSDiskImageEncryption {
		if v == nil {
			return nil
		}
		return v.OsDiskImage
	}).(OSDiskImageEncryptionPtrOutput)
}

type EncryptionImagesResponse struct {
	DataDiskImages []DataDiskImageEncryptionResponse `pulumi:"dataDiskImages"`
	OsDiskImage    *OSDiskImageEncryptionResponse    `pulumi:"osDiskImage"`
}

type EncryptionImagesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionImagesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionImagesResponse)(nil)).Elem()
}

func (o EncryptionImagesResponseOutput) ToEncryptionImagesResponseOutput() EncryptionImagesResponseOutput {
	return o
}

func (o EncryptionImagesResponseOutput) ToEncryptionImagesResponseOutputWithContext(ctx context.Context) EncryptionImagesResponseOutput {
	return o
}

func (o EncryptionImagesResponseOutput) DataDiskImages() DataDiskImageEncryptionResponseArrayOutput {
	return o.ApplyT(func(v EncryptionImagesResponse) []DataDiskImageEncryptionResponse { return v.DataDiskImages }).(DataDiskImageEncryptionResponseArrayOutput)
}

func (o EncryptionImagesResponseOutput) OsDiskImage() OSDiskImageEncryptionResponsePtrOutput {
	return o.ApplyT(func(v EncryptionImagesResponse) *OSDiskImageEncryptionResponse { return v.OsDiskImage }).(OSDiskImageEncryptionResponsePtrOutput)
}

type EncryptionImagesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionImagesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionImagesResponse)(nil)).Elem()
}

func (o EncryptionImagesResponsePtrOutput) ToEncryptionImagesResponsePtrOutput() EncryptionImagesResponsePtrOutput {
	return o
}

func (o EncryptionImagesResponsePtrOutput) ToEncryptionImagesResponsePtrOutputWithContext(ctx context.Context) EncryptionImagesResponsePtrOutput {
	return o
}

func (o EncryptionImagesResponsePtrOutput) Elem() EncryptionImagesResponseOutput {
	return o.ApplyT(func(v *EncryptionImagesResponse) EncryptionImagesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionImagesResponse
		return ret
	}).(EncryptionImagesResponseOutput)
}

func (o EncryptionImagesResponsePtrOutput) DataDiskImages() DataDiskImageEncryptionResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionImagesResponse) []DataDiskImageEncryptionResponse {
		if v == nil {
			return nil
		}
		return v.DataDiskImages
	}).(DataDiskImageEncryptionResponseArrayOutput)
}

func (o EncryptionImagesResponsePtrOutput) OsDiskImage() OSDiskImageEncryptionResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionImagesResponse) *OSDiskImageEncryptionResponse {
		if v == nil {
			return nil
		}
		return v.OsDiskImage
	}).(OSDiskImageEncryptionResponsePtrOutput)
}

type GalleryApplicationVersionPublishingProfile struct {
	AdvancedSettings        map[string]string               `pulumi:"advancedSettings"`
	EnableHealthCheck       *bool                           `pulumi:"enableHealthCheck"`
	EndOfLifeDate           *string                         `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       *bool                           `pulumi:"excludeFromLatest"`
	ManageActions           *UserArtifactManage             `pulumi:"manageActions"`
	ReplicaCount            *int                            `pulumi:"replicaCount"`
	ReplicationMode         *string                         `pulumi:"replicationMode"`
	Settings                *UserArtifactSettings           `pulumi:"settings"`
	Source                  UserArtifactSource              `pulumi:"source"`
	StorageAccountType      *string                         `pulumi:"storageAccountType"`
	TargetExtendedLocations []GalleryTargetExtendedLocation `pulumi:"targetExtendedLocations"`
	TargetRegions           []TargetRegion                  `pulumi:"targetRegions"`
}





type GalleryApplicationVersionPublishingProfileInput interface {
	pulumi.Input

	ToGalleryApplicationVersionPublishingProfileOutput() GalleryApplicationVersionPublishingProfileOutput
	ToGalleryApplicationVersionPublishingProfileOutputWithContext(context.Context) GalleryApplicationVersionPublishingProfileOutput
}

type GalleryApplicationVersionPublishingProfileArgs struct {
	AdvancedSettings        pulumi.StringMapInput                   `pulumi:"advancedSettings"`
	EnableHealthCheck       pulumi.BoolPtrInput                     `pulumi:"enableHealthCheck"`
	EndOfLifeDate           pulumi.StringPtrInput                   `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       pulumi.BoolPtrInput                     `pulumi:"excludeFromLatest"`
	ManageActions           UserArtifactManagePtrInput              `pulumi:"manageActions"`
	ReplicaCount            pulumi.IntPtrInput                      `pulumi:"replicaCount"`
	ReplicationMode         pulumi.StringPtrInput                   `pulumi:"replicationMode"`
	Settings                UserArtifactSettingsPtrInput            `pulumi:"settings"`
	Source                  UserArtifactSourceInput                 `pulumi:"source"`
	StorageAccountType      pulumi.StringPtrInput                   `pulumi:"storageAccountType"`
	TargetExtendedLocations GalleryTargetExtendedLocationArrayInput `pulumi:"targetExtendedLocations"`
	TargetRegions           TargetRegionArrayInput                  `pulumi:"targetRegions"`
}

func (GalleryApplicationVersionPublishingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryApplicationVersionPublishingProfile)(nil)).Elem()
}

func (i GalleryApplicationVersionPublishingProfileArgs) ToGalleryApplicationVersionPublishingProfileOutput() GalleryApplicationVersionPublishingProfileOutput {
	return i.ToGalleryApplicationVersionPublishingProfileOutputWithContext(context.Background())
}

func (i GalleryApplicationVersionPublishingProfileArgs) ToGalleryApplicationVersionPublishingProfileOutputWithContext(ctx context.Context) GalleryApplicationVersionPublishingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryApplicationVersionPublishingProfileOutput)
}

type GalleryApplicationVersionPublishingProfileOutput struct{ *pulumi.OutputState }

func (GalleryApplicationVersionPublishingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryApplicationVersionPublishingProfile)(nil)).Elem()
}

func (o GalleryApplicationVersionPublishingProfileOutput) ToGalleryApplicationVersionPublishingProfileOutput() GalleryApplicationVersionPublishingProfileOutput {
	return o
}

func (o GalleryApplicationVersionPublishingProfileOutput) ToGalleryApplicationVersionPublishingProfileOutputWithContext(ctx context.Context) GalleryApplicationVersionPublishingProfileOutput {
	return o
}

func (o GalleryApplicationVersionPublishingProfileOutput) AdvancedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) map[string]string { return v.AdvancedSettings }).(pulumi.StringMapOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) EnableHealthCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *bool { return v.EnableHealthCheck }).(pulumi.BoolPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) ManageActions() UserArtifactManagePtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *UserArtifactManage { return v.ManageActions }).(UserArtifactManagePtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *string { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) Settings() UserArtifactSettingsPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *UserArtifactSettings { return v.Settings }).(UserArtifactSettingsPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) Source() UserArtifactSourceOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) UserArtifactSource { return v.Source }).(UserArtifactSourceOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) TargetExtendedLocations() GalleryTargetExtendedLocationArrayOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) []GalleryTargetExtendedLocation {
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationArrayOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) TargetRegions() TargetRegionArrayOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) []TargetRegion { return v.TargetRegions }).(TargetRegionArrayOutput)
}

type GalleryApplicationVersionPublishingProfileResponse struct {
	AdvancedSettings        map[string]string                       `pulumi:"advancedSettings"`
	EnableHealthCheck       *bool                                   `pulumi:"enableHealthCheck"`
	EndOfLifeDate           *string                                 `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       *bool                                   `pulumi:"excludeFromLatest"`
	ManageActions           *UserArtifactManageResponse             `pulumi:"manageActions"`
	PublishedDate           string                                  `pulumi:"publishedDate"`
	ReplicaCount            *int                                    `pulumi:"replicaCount"`
	ReplicationMode         *string                                 `pulumi:"replicationMode"`
	Settings                *UserArtifactSettingsResponse           `pulumi:"settings"`
	Source                  UserArtifactSourceResponse              `pulumi:"source"`
	StorageAccountType      *string                                 `pulumi:"storageAccountType"`
	TargetExtendedLocations []GalleryTargetExtendedLocationResponse `pulumi:"targetExtendedLocations"`
	TargetRegions           []TargetRegionResponse                  `pulumi:"targetRegions"`
}

type GalleryApplicationVersionPublishingProfileResponseOutput struct{ *pulumi.OutputState }

func (GalleryApplicationVersionPublishingProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryApplicationVersionPublishingProfileResponse)(nil)).Elem()
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ToGalleryApplicationVersionPublishingProfileResponseOutput() GalleryApplicationVersionPublishingProfileResponseOutput {
	return o
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ToGalleryApplicationVersionPublishingProfileResponseOutputWithContext(ctx context.Context) GalleryApplicationVersionPublishingProfileResponseOutput {
	return o
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) AdvancedSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) map[string]string {
		return v.AdvancedSettings
	}).(pulumi.StringMapOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) EnableHealthCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *bool { return v.EnableHealthCheck }).(pulumi.BoolPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ManageActions() UserArtifactManageResponsePtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *UserArtifactManageResponse {
		return v.ManageActions
	}).(UserArtifactManageResponsePtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) PublishedDate() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) string { return v.PublishedDate }).(pulumi.StringOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *string { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) Settings() UserArtifactSettingsResponsePtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *UserArtifactSettingsResponse {
		return v.Settings
	}).(UserArtifactSettingsResponsePtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) Source() UserArtifactSourceResponseOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) UserArtifactSourceResponse { return v.Source }).(UserArtifactSourceResponseOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) TargetExtendedLocations() GalleryTargetExtendedLocationResponseArrayOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) []GalleryTargetExtendedLocationResponse {
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationResponseArrayOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) TargetRegions() TargetRegionResponseArrayOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) []TargetRegionResponse {
		return v.TargetRegions
	}).(TargetRegionResponseArrayOutput)
}

type GalleryArtifactVersionSource struct {
	Id  *string `pulumi:"id"`
	Uri *string `pulumi:"uri"`
}





type GalleryArtifactVersionSourceInput interface {
	pulumi.Input

	ToGalleryArtifactVersionSourceOutput() GalleryArtifactVersionSourceOutput
	ToGalleryArtifactVersionSourceOutputWithContext(context.Context) GalleryArtifactVersionSourceOutput
}

type GalleryArtifactVersionSourceArgs struct {
	Id  pulumi.StringPtrInput `pulumi:"id"`
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (GalleryArtifactVersionSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryArtifactVersionSource)(nil)).Elem()
}

func (i GalleryArtifactVersionSourceArgs) ToGalleryArtifactVersionSourceOutput() GalleryArtifactVersionSourceOutput {
	return i.ToGalleryArtifactVersionSourceOutputWithContext(context.Background())
}

func (i GalleryArtifactVersionSourceArgs) ToGalleryArtifactVersionSourceOutputWithContext(ctx context.Context) GalleryArtifactVersionSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryArtifactVersionSourceOutput)
}

func (i GalleryArtifactVersionSourceArgs) ToGalleryArtifactVersionSourcePtrOutput() GalleryArtifactVersionSourcePtrOutput {
	return i.ToGalleryArtifactVersionSourcePtrOutputWithContext(context.Background())
}

func (i GalleryArtifactVersionSourceArgs) ToGalleryArtifactVersionSourcePtrOutputWithContext(ctx context.Context) GalleryArtifactVersionSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryArtifactVersionSourceOutput).ToGalleryArtifactVersionSourcePtrOutputWithContext(ctx)
}









type GalleryArtifactVersionSourcePtrInput interface {
	pulumi.Input

	ToGalleryArtifactVersionSourcePtrOutput() GalleryArtifactVersionSourcePtrOutput
	ToGalleryArtifactVersionSourcePtrOutputWithContext(context.Context) GalleryArtifactVersionSourcePtrOutput
}

type galleryArtifactVersionSourcePtrType GalleryArtifactVersionSourceArgs

func GalleryArtifactVersionSourcePtr(v *GalleryArtifactVersionSourceArgs) GalleryArtifactVersionSourcePtrInput {
	return (*galleryArtifactVersionSourcePtrType)(v)
}

func (*galleryArtifactVersionSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryArtifactVersionSource)(nil)).Elem()
}

func (i *galleryArtifactVersionSourcePtrType) ToGalleryArtifactVersionSourcePtrOutput() GalleryArtifactVersionSourcePtrOutput {
	return i.ToGalleryArtifactVersionSourcePtrOutputWithContext(context.Background())
}

func (i *galleryArtifactVersionSourcePtrType) ToGalleryArtifactVersionSourcePtrOutputWithContext(ctx context.Context) GalleryArtifactVersionSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryArtifactVersionSourcePtrOutput)
}

type GalleryArtifactVersionSourceOutput struct{ *pulumi.OutputState }

func (GalleryArtifactVersionSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryArtifactVersionSource)(nil)).Elem()
}

func (o GalleryArtifactVersionSourceOutput) ToGalleryArtifactVersionSourceOutput() GalleryArtifactVersionSourceOutput {
	return o
}

func (o GalleryArtifactVersionSourceOutput) ToGalleryArtifactVersionSourceOutputWithContext(ctx context.Context) GalleryArtifactVersionSourceOutput {
	return o
}

func (o GalleryArtifactVersionSourceOutput) ToGalleryArtifactVersionSourcePtrOutput() GalleryArtifactVersionSourcePtrOutput {
	return o.ToGalleryArtifactVersionSourcePtrOutputWithContext(context.Background())
}

func (o GalleryArtifactVersionSourceOutput) ToGalleryArtifactVersionSourcePtrOutputWithContext(ctx context.Context) GalleryArtifactVersionSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryArtifactVersionSource) *GalleryArtifactVersionSource {
		return &v
	}).(GalleryArtifactVersionSourcePtrOutput)
}

func (o GalleryArtifactVersionSourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryArtifactVersionSource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GalleryArtifactVersionSourceOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryArtifactVersionSource) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GalleryArtifactVersionSourcePtrOutput struct{ *pulumi.OutputState }

func (GalleryArtifactVersionSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryArtifactVersionSource)(nil)).Elem()
}

func (o GalleryArtifactVersionSourcePtrOutput) ToGalleryArtifactVersionSourcePtrOutput() GalleryArtifactVersionSourcePtrOutput {
	return o
}

func (o GalleryArtifactVersionSourcePtrOutput) ToGalleryArtifactVersionSourcePtrOutputWithContext(ctx context.Context) GalleryArtifactVersionSourcePtrOutput {
	return o
}

func (o GalleryArtifactVersionSourcePtrOutput) Elem() GalleryArtifactVersionSourceOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSource) GalleryArtifactVersionSource {
		if v != nil {
			return *v
		}
		var ret GalleryArtifactVersionSource
		return ret
	}).(GalleryArtifactVersionSourceOutput)
}

func (o GalleryArtifactVersionSourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GalleryArtifactVersionSourcePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSource) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type GalleryArtifactVersionSourceResponse struct {
	Id  *string `pulumi:"id"`
	Uri *string `pulumi:"uri"`
}

type GalleryArtifactVersionSourceResponseOutput struct{ *pulumi.OutputState }

func (GalleryArtifactVersionSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryArtifactVersionSourceResponse)(nil)).Elem()
}

func (o GalleryArtifactVersionSourceResponseOutput) ToGalleryArtifactVersionSourceResponseOutput() GalleryArtifactVersionSourceResponseOutput {
	return o
}

func (o GalleryArtifactVersionSourceResponseOutput) ToGalleryArtifactVersionSourceResponseOutputWithContext(ctx context.Context) GalleryArtifactVersionSourceResponseOutput {
	return o
}

func (o GalleryArtifactVersionSourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryArtifactVersionSourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o GalleryArtifactVersionSourceResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryArtifactVersionSourceResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type GalleryArtifactVersionSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryArtifactVersionSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryArtifactVersionSourceResponse)(nil)).Elem()
}

func (o GalleryArtifactVersionSourceResponsePtrOutput) ToGalleryArtifactVersionSourceResponsePtrOutput() GalleryArtifactVersionSourceResponsePtrOutput {
	return o
}

func (o GalleryArtifactVersionSourceResponsePtrOutput) ToGalleryArtifactVersionSourceResponsePtrOutputWithContext(ctx context.Context) GalleryArtifactVersionSourceResponsePtrOutput {
	return o
}

func (o GalleryArtifactVersionSourceResponsePtrOutput) Elem() GalleryArtifactVersionSourceResponseOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSourceResponse) GalleryArtifactVersionSourceResponse {
		if v != nil {
			return *v
		}
		var ret GalleryArtifactVersionSourceResponse
		return ret
	}).(GalleryArtifactVersionSourceResponseOutput)
}

func (o GalleryArtifactVersionSourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GalleryArtifactVersionSourceResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryArtifactVersionSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type GalleryDataDiskImage struct {
	HostCaching *HostCaching                  `pulumi:"hostCaching"`
	Lun         int                           `pulumi:"lun"`
	Source      *GalleryArtifactVersionSource `pulumi:"source"`
}





type GalleryDataDiskImageInput interface {
	pulumi.Input

	ToGalleryDataDiskImageOutput() GalleryDataDiskImageOutput
	ToGalleryDataDiskImageOutputWithContext(context.Context) GalleryDataDiskImageOutput
}

type GalleryDataDiskImageArgs struct {
	HostCaching HostCachingPtrInput                  `pulumi:"hostCaching"`
	Lun         pulumi.IntInput                      `pulumi:"lun"`
	Source      GalleryArtifactVersionSourcePtrInput `pulumi:"source"`
}

func (GalleryDataDiskImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryDataDiskImage)(nil)).Elem()
}

func (i GalleryDataDiskImageArgs) ToGalleryDataDiskImageOutput() GalleryDataDiskImageOutput {
	return i.ToGalleryDataDiskImageOutputWithContext(context.Background())
}

func (i GalleryDataDiskImageArgs) ToGalleryDataDiskImageOutputWithContext(ctx context.Context) GalleryDataDiskImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryDataDiskImageOutput)
}





type GalleryDataDiskImageArrayInput interface {
	pulumi.Input

	ToGalleryDataDiskImageArrayOutput() GalleryDataDiskImageArrayOutput
	ToGalleryDataDiskImageArrayOutputWithContext(context.Context) GalleryDataDiskImageArrayOutput
}

type GalleryDataDiskImageArray []GalleryDataDiskImageInput

func (GalleryDataDiskImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryDataDiskImage)(nil)).Elem()
}

func (i GalleryDataDiskImageArray) ToGalleryDataDiskImageArrayOutput() GalleryDataDiskImageArrayOutput {
	return i.ToGalleryDataDiskImageArrayOutputWithContext(context.Background())
}

func (i GalleryDataDiskImageArray) ToGalleryDataDiskImageArrayOutputWithContext(ctx context.Context) GalleryDataDiskImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryDataDiskImageArrayOutput)
}

type GalleryDataDiskImageOutput struct{ *pulumi.OutputState }

func (GalleryDataDiskImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryDataDiskImage)(nil)).Elem()
}

func (o GalleryDataDiskImageOutput) ToGalleryDataDiskImageOutput() GalleryDataDiskImageOutput {
	return o
}

func (o GalleryDataDiskImageOutput) ToGalleryDataDiskImageOutputWithContext(ctx context.Context) GalleryDataDiskImageOutput {
	return o
}

func (o GalleryDataDiskImageOutput) HostCaching() HostCachingPtrOutput {
	return o.ApplyT(func(v GalleryDataDiskImage) *HostCaching { return v.HostCaching }).(HostCachingPtrOutput)
}

func (o GalleryDataDiskImageOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v GalleryDataDiskImage) int { return v.Lun }).(pulumi.IntOutput)
}

func (o GalleryDataDiskImageOutput) Source() GalleryArtifactVersionSourcePtrOutput {
	return o.ApplyT(func(v GalleryDataDiskImage) *GalleryArtifactVersionSource { return v.Source }).(GalleryArtifactVersionSourcePtrOutput)
}

type GalleryDataDiskImageArrayOutput struct{ *pulumi.OutputState }

func (GalleryDataDiskImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryDataDiskImage)(nil)).Elem()
}

func (o GalleryDataDiskImageArrayOutput) ToGalleryDataDiskImageArrayOutput() GalleryDataDiskImageArrayOutput {
	return o
}

func (o GalleryDataDiskImageArrayOutput) ToGalleryDataDiskImageArrayOutputWithContext(ctx context.Context) GalleryDataDiskImageArrayOutput {
	return o
}

func (o GalleryDataDiskImageArrayOutput) Index(i pulumi.IntInput) GalleryDataDiskImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryDataDiskImage {
		return vs[0].([]GalleryDataDiskImage)[vs[1].(int)]
	}).(GalleryDataDiskImageOutput)
}

type GalleryDataDiskImageResponse struct {
	HostCaching *string                               `pulumi:"hostCaching"`
	Lun         int                                   `pulumi:"lun"`
	SizeInGB    int                                   `pulumi:"sizeInGB"`
	Source      *GalleryArtifactVersionSourceResponse `pulumi:"source"`
}

type GalleryDataDiskImageResponseOutput struct{ *pulumi.OutputState }

func (GalleryDataDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryDataDiskImageResponse)(nil)).Elem()
}

func (o GalleryDataDiskImageResponseOutput) ToGalleryDataDiskImageResponseOutput() GalleryDataDiskImageResponseOutput {
	return o
}

func (o GalleryDataDiskImageResponseOutput) ToGalleryDataDiskImageResponseOutputWithContext(ctx context.Context) GalleryDataDiskImageResponseOutput {
	return o
}

func (o GalleryDataDiskImageResponseOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryDataDiskImageResponse) *string { return v.HostCaching }).(pulumi.StringPtrOutput)
}

func (o GalleryDataDiskImageResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v GalleryDataDiskImageResponse) int { return v.Lun }).(pulumi.IntOutput)
}

func (o GalleryDataDiskImageResponseOutput) SizeInGB() pulumi.IntOutput {
	return o.ApplyT(func(v GalleryDataDiskImageResponse) int { return v.SizeInGB }).(pulumi.IntOutput)
}

func (o GalleryDataDiskImageResponseOutput) Source() GalleryArtifactVersionSourceResponsePtrOutput {
	return o.ApplyT(func(v GalleryDataDiskImageResponse) *GalleryArtifactVersionSourceResponse { return v.Source }).(GalleryArtifactVersionSourceResponsePtrOutput)
}

type GalleryDataDiskImageResponseArrayOutput struct{ *pulumi.OutputState }

func (GalleryDataDiskImageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryDataDiskImageResponse)(nil)).Elem()
}

func (o GalleryDataDiskImageResponseArrayOutput) ToGalleryDataDiskImageResponseArrayOutput() GalleryDataDiskImageResponseArrayOutput {
	return o
}

func (o GalleryDataDiskImageResponseArrayOutput) ToGalleryDataDiskImageResponseArrayOutputWithContext(ctx context.Context) GalleryDataDiskImageResponseArrayOutput {
	return o
}

func (o GalleryDataDiskImageResponseArrayOutput) Index(i pulumi.IntInput) GalleryDataDiskImageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryDataDiskImageResponse {
		return vs[0].([]GalleryDataDiskImageResponse)[vs[1].(int)]
	}).(GalleryDataDiskImageResponseOutput)
}

type GalleryExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type GalleryExtendedLocationInput interface {
	pulumi.Input

	ToGalleryExtendedLocationOutput() GalleryExtendedLocationOutput
	ToGalleryExtendedLocationOutputWithContext(context.Context) GalleryExtendedLocationOutput
}

type GalleryExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (GalleryExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryExtendedLocation)(nil)).Elem()
}

func (i GalleryExtendedLocationArgs) ToGalleryExtendedLocationOutput() GalleryExtendedLocationOutput {
	return i.ToGalleryExtendedLocationOutputWithContext(context.Background())
}

func (i GalleryExtendedLocationArgs) ToGalleryExtendedLocationOutputWithContext(ctx context.Context) GalleryExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryExtendedLocationOutput)
}

func (i GalleryExtendedLocationArgs) ToGalleryExtendedLocationPtrOutput() GalleryExtendedLocationPtrOutput {
	return i.ToGalleryExtendedLocationPtrOutputWithContext(context.Background())
}

func (i GalleryExtendedLocationArgs) ToGalleryExtendedLocationPtrOutputWithContext(ctx context.Context) GalleryExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryExtendedLocationOutput).ToGalleryExtendedLocationPtrOutputWithContext(ctx)
}









type GalleryExtendedLocationPtrInput interface {
	pulumi.Input

	ToGalleryExtendedLocationPtrOutput() GalleryExtendedLocationPtrOutput
	ToGalleryExtendedLocationPtrOutputWithContext(context.Context) GalleryExtendedLocationPtrOutput
}

type galleryExtendedLocationPtrType GalleryExtendedLocationArgs

func GalleryExtendedLocationPtr(v *GalleryExtendedLocationArgs) GalleryExtendedLocationPtrInput {
	return (*galleryExtendedLocationPtrType)(v)
}

func (*galleryExtendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryExtendedLocation)(nil)).Elem()
}

func (i *galleryExtendedLocationPtrType) ToGalleryExtendedLocationPtrOutput() GalleryExtendedLocationPtrOutput {
	return i.ToGalleryExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *galleryExtendedLocationPtrType) ToGalleryExtendedLocationPtrOutputWithContext(ctx context.Context) GalleryExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryExtendedLocationPtrOutput)
}

type GalleryExtendedLocationOutput struct{ *pulumi.OutputState }

func (GalleryExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryExtendedLocation)(nil)).Elem()
}

func (o GalleryExtendedLocationOutput) ToGalleryExtendedLocationOutput() GalleryExtendedLocationOutput {
	return o
}

func (o GalleryExtendedLocationOutput) ToGalleryExtendedLocationOutputWithContext(ctx context.Context) GalleryExtendedLocationOutput {
	return o
}

func (o GalleryExtendedLocationOutput) ToGalleryExtendedLocationPtrOutput() GalleryExtendedLocationPtrOutput {
	return o.ToGalleryExtendedLocationPtrOutputWithContext(context.Background())
}

func (o GalleryExtendedLocationOutput) ToGalleryExtendedLocationPtrOutputWithContext(ctx context.Context) GalleryExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryExtendedLocation) *GalleryExtendedLocation {
		return &v
	}).(GalleryExtendedLocationPtrOutput)
}

func (o GalleryExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GalleryExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (GalleryExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryExtendedLocation)(nil)).Elem()
}

func (o GalleryExtendedLocationPtrOutput) ToGalleryExtendedLocationPtrOutput() GalleryExtendedLocationPtrOutput {
	return o
}

func (o GalleryExtendedLocationPtrOutput) ToGalleryExtendedLocationPtrOutputWithContext(ctx context.Context) GalleryExtendedLocationPtrOutput {
	return o
}

func (o GalleryExtendedLocationPtrOutput) Elem() GalleryExtendedLocationOutput {
	return o.ApplyT(func(v *GalleryExtendedLocation) GalleryExtendedLocation {
		if v != nil {
			return *v
		}
		var ret GalleryExtendedLocation
		return ret
	}).(GalleryExtendedLocationOutput)
}

func (o GalleryExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GalleryExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GalleryExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type GalleryExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (GalleryExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryExtendedLocationResponse)(nil)).Elem()
}

func (o GalleryExtendedLocationResponseOutput) ToGalleryExtendedLocationResponseOutput() GalleryExtendedLocationResponseOutput {
	return o
}

func (o GalleryExtendedLocationResponseOutput) ToGalleryExtendedLocationResponseOutputWithContext(ctx context.Context) GalleryExtendedLocationResponseOutput {
	return o
}

func (o GalleryExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type GalleryExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryExtendedLocationResponse)(nil)).Elem()
}

func (o GalleryExtendedLocationResponsePtrOutput) ToGalleryExtendedLocationResponsePtrOutput() GalleryExtendedLocationResponsePtrOutput {
	return o
}

func (o GalleryExtendedLocationResponsePtrOutput) ToGalleryExtendedLocationResponsePtrOutputWithContext(ctx context.Context) GalleryExtendedLocationResponsePtrOutput {
	return o
}

func (o GalleryExtendedLocationResponsePtrOutput) Elem() GalleryExtendedLocationResponseOutput {
	return o.ApplyT(func(v *GalleryExtendedLocationResponse) GalleryExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret GalleryExtendedLocationResponse
		return ret
	}).(GalleryExtendedLocationResponseOutput)
}

func (o GalleryExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GalleryExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GalleryIdentifierResponse struct {
	UniqueName string `pulumi:"uniqueName"`
}

type GalleryIdentifierResponseOutput struct{ *pulumi.OutputState }

func (GalleryIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryIdentifierResponse)(nil)).Elem()
}

func (o GalleryIdentifierResponseOutput) ToGalleryIdentifierResponseOutput() GalleryIdentifierResponseOutput {
	return o
}

func (o GalleryIdentifierResponseOutput) ToGalleryIdentifierResponseOutputWithContext(ctx context.Context) GalleryIdentifierResponseOutput {
	return o
}

func (o GalleryIdentifierResponseOutput) UniqueName() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryIdentifierResponse) string { return v.UniqueName }).(pulumi.StringOutput)
}

type GalleryIdentifierResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryIdentifierResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryIdentifierResponse)(nil)).Elem()
}

func (o GalleryIdentifierResponsePtrOutput) ToGalleryIdentifierResponsePtrOutput() GalleryIdentifierResponsePtrOutput {
	return o
}

func (o GalleryIdentifierResponsePtrOutput) ToGalleryIdentifierResponsePtrOutputWithContext(ctx context.Context) GalleryIdentifierResponsePtrOutput {
	return o
}

func (o GalleryIdentifierResponsePtrOutput) Elem() GalleryIdentifierResponseOutput {
	return o.ApplyT(func(v *GalleryIdentifierResponse) GalleryIdentifierResponse {
		if v != nil {
			return *v
		}
		var ret GalleryIdentifierResponse
		return ret
	}).(GalleryIdentifierResponseOutput)
}

func (o GalleryIdentifierResponsePtrOutput) UniqueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryIdentifierResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UniqueName
	}).(pulumi.StringPtrOutput)
}

type GalleryImageFeature struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type GalleryImageFeatureInput interface {
	pulumi.Input

	ToGalleryImageFeatureOutput() GalleryImageFeatureOutput
	ToGalleryImageFeatureOutputWithContext(context.Context) GalleryImageFeatureOutput
}

type GalleryImageFeatureArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (GalleryImageFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageFeature)(nil)).Elem()
}

func (i GalleryImageFeatureArgs) ToGalleryImageFeatureOutput() GalleryImageFeatureOutput {
	return i.ToGalleryImageFeatureOutputWithContext(context.Background())
}

func (i GalleryImageFeatureArgs) ToGalleryImageFeatureOutputWithContext(ctx context.Context) GalleryImageFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageFeatureOutput)
}





type GalleryImageFeatureArrayInput interface {
	pulumi.Input

	ToGalleryImageFeatureArrayOutput() GalleryImageFeatureArrayOutput
	ToGalleryImageFeatureArrayOutputWithContext(context.Context) GalleryImageFeatureArrayOutput
}

type GalleryImageFeatureArray []GalleryImageFeatureInput

func (GalleryImageFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryImageFeature)(nil)).Elem()
}

func (i GalleryImageFeatureArray) ToGalleryImageFeatureArrayOutput() GalleryImageFeatureArrayOutput {
	return i.ToGalleryImageFeatureArrayOutputWithContext(context.Background())
}

func (i GalleryImageFeatureArray) ToGalleryImageFeatureArrayOutputWithContext(ctx context.Context) GalleryImageFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageFeatureArrayOutput)
}

type GalleryImageFeatureOutput struct{ *pulumi.OutputState }

func (GalleryImageFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageFeature)(nil)).Elem()
}

func (o GalleryImageFeatureOutput) ToGalleryImageFeatureOutput() GalleryImageFeatureOutput {
	return o
}

func (o GalleryImageFeatureOutput) ToGalleryImageFeatureOutputWithContext(ctx context.Context) GalleryImageFeatureOutput {
	return o
}

func (o GalleryImageFeatureOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageFeature) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryImageFeatureOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageFeature) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GalleryImageFeatureArrayOutput struct{ *pulumi.OutputState }

func (GalleryImageFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryImageFeature)(nil)).Elem()
}

func (o GalleryImageFeatureArrayOutput) ToGalleryImageFeatureArrayOutput() GalleryImageFeatureArrayOutput {
	return o
}

func (o GalleryImageFeatureArrayOutput) ToGalleryImageFeatureArrayOutputWithContext(ctx context.Context) GalleryImageFeatureArrayOutput {
	return o
}

func (o GalleryImageFeatureArrayOutput) Index(i pulumi.IntInput) GalleryImageFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryImageFeature {
		return vs[0].([]GalleryImageFeature)[vs[1].(int)]
	}).(GalleryImageFeatureOutput)
}

type GalleryImageFeatureResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type GalleryImageFeatureResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageFeatureResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageFeatureResponse)(nil)).Elem()
}

func (o GalleryImageFeatureResponseOutput) ToGalleryImageFeatureResponseOutput() GalleryImageFeatureResponseOutput {
	return o
}

func (o GalleryImageFeatureResponseOutput) ToGalleryImageFeatureResponseOutputWithContext(ctx context.Context) GalleryImageFeatureResponseOutput {
	return o
}

func (o GalleryImageFeatureResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageFeatureResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryImageFeatureResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageFeatureResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type GalleryImageFeatureResponseArrayOutput struct{ *pulumi.OutputState }

func (GalleryImageFeatureResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryImageFeatureResponse)(nil)).Elem()
}

func (o GalleryImageFeatureResponseArrayOutput) ToGalleryImageFeatureResponseArrayOutput() GalleryImageFeatureResponseArrayOutput {
	return o
}

func (o GalleryImageFeatureResponseArrayOutput) ToGalleryImageFeatureResponseArrayOutputWithContext(ctx context.Context) GalleryImageFeatureResponseArrayOutput {
	return o
}

func (o GalleryImageFeatureResponseArrayOutput) Index(i pulumi.IntInput) GalleryImageFeatureResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryImageFeatureResponse {
		return vs[0].([]GalleryImageFeatureResponse)[vs[1].(int)]
	}).(GalleryImageFeatureResponseOutput)
}

type GalleryImageIdentifier struct {
	Offer     string `pulumi:"offer"`
	Publisher string `pulumi:"publisher"`
	Sku       string `pulumi:"sku"`
}





type GalleryImageIdentifierInput interface {
	pulumi.Input

	ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput
	ToGalleryImageIdentifierOutputWithContext(context.Context) GalleryImageIdentifierOutput
}

type GalleryImageIdentifierArgs struct {
	Offer     pulumi.StringInput `pulumi:"offer"`
	Publisher pulumi.StringInput `pulumi:"publisher"`
	Sku       pulumi.StringInput `pulumi:"sku"`
}

func (GalleryImageIdentifierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifier)(nil)).Elem()
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput {
	return i.ToGalleryImageIdentifierOutputWithContext(context.Background())
}

func (i GalleryImageIdentifierArgs) ToGalleryImageIdentifierOutputWithContext(ctx context.Context) GalleryImageIdentifierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageIdentifierOutput)
}

type GalleryImageIdentifierOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifier)(nil)).Elem()
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierOutput() GalleryImageIdentifierOutput {
	return o
}

func (o GalleryImageIdentifierOutput) ToGalleryImageIdentifierOutputWithContext(ctx context.Context) GalleryImageIdentifierOutput {
	return o
}

func (o GalleryImageIdentifierOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Offer }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifier) string { return v.Sku }).(pulumi.StringOutput)
}

type GalleryImageIdentifierResponse struct {
	Offer     string `pulumi:"offer"`
	Publisher string `pulumi:"publisher"`
	Sku       string `pulumi:"sku"`
}

type GalleryImageIdentifierResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageIdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageIdentifierResponse)(nil)).Elem()
}

func (o GalleryImageIdentifierResponseOutput) ToGalleryImageIdentifierResponseOutput() GalleryImageIdentifierResponseOutput {
	return o
}

func (o GalleryImageIdentifierResponseOutput) ToGalleryImageIdentifierResponseOutputWithContext(ctx context.Context) GalleryImageIdentifierResponseOutput {
	return o
}

func (o GalleryImageIdentifierResponseOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Offer }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o GalleryImageIdentifierResponseOutput) Sku() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageIdentifierResponse) string { return v.Sku }).(pulumi.StringOutput)
}

type GalleryImageVersionPublishingProfile struct {
	EndOfLifeDate           *string                         `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       *bool                           `pulumi:"excludeFromLatest"`
	ReplicaCount            *int                            `pulumi:"replicaCount"`
	ReplicationMode         *string                         `pulumi:"replicationMode"`
	StorageAccountType      *string                         `pulumi:"storageAccountType"`
	TargetExtendedLocations []GalleryTargetExtendedLocation `pulumi:"targetExtendedLocations"`
	TargetRegions           []TargetRegion                  `pulumi:"targetRegions"`
}





type GalleryImageVersionPublishingProfileInput interface {
	pulumi.Input

	ToGalleryImageVersionPublishingProfileOutput() GalleryImageVersionPublishingProfileOutput
	ToGalleryImageVersionPublishingProfileOutputWithContext(context.Context) GalleryImageVersionPublishingProfileOutput
}

type GalleryImageVersionPublishingProfileArgs struct {
	EndOfLifeDate           pulumi.StringPtrInput                   `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       pulumi.BoolPtrInput                     `pulumi:"excludeFromLatest"`
	ReplicaCount            pulumi.IntPtrInput                      `pulumi:"replicaCount"`
	ReplicationMode         pulumi.StringPtrInput                   `pulumi:"replicationMode"`
	StorageAccountType      pulumi.StringPtrInput                   `pulumi:"storageAccountType"`
	TargetExtendedLocations GalleryTargetExtendedLocationArrayInput `pulumi:"targetExtendedLocations"`
	TargetRegions           TargetRegionArrayInput                  `pulumi:"targetRegions"`
}

func (GalleryImageVersionPublishingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionPublishingProfile)(nil)).Elem()
}

func (i GalleryImageVersionPublishingProfileArgs) ToGalleryImageVersionPublishingProfileOutput() GalleryImageVersionPublishingProfileOutput {
	return i.ToGalleryImageVersionPublishingProfileOutputWithContext(context.Background())
}

func (i GalleryImageVersionPublishingProfileArgs) ToGalleryImageVersionPublishingProfileOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionPublishingProfileOutput)
}

func (i GalleryImageVersionPublishingProfileArgs) ToGalleryImageVersionPublishingProfilePtrOutput() GalleryImageVersionPublishingProfilePtrOutput {
	return i.ToGalleryImageVersionPublishingProfilePtrOutputWithContext(context.Background())
}

func (i GalleryImageVersionPublishingProfileArgs) ToGalleryImageVersionPublishingProfilePtrOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionPublishingProfileOutput).ToGalleryImageVersionPublishingProfilePtrOutputWithContext(ctx)
}









type GalleryImageVersionPublishingProfilePtrInput interface {
	pulumi.Input

	ToGalleryImageVersionPublishingProfilePtrOutput() GalleryImageVersionPublishingProfilePtrOutput
	ToGalleryImageVersionPublishingProfilePtrOutputWithContext(context.Context) GalleryImageVersionPublishingProfilePtrOutput
}

type galleryImageVersionPublishingProfilePtrType GalleryImageVersionPublishingProfileArgs

func GalleryImageVersionPublishingProfilePtr(v *GalleryImageVersionPublishingProfileArgs) GalleryImageVersionPublishingProfilePtrInput {
	return (*galleryImageVersionPublishingProfilePtrType)(v)
}

func (*galleryImageVersionPublishingProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersionPublishingProfile)(nil)).Elem()
}

func (i *galleryImageVersionPublishingProfilePtrType) ToGalleryImageVersionPublishingProfilePtrOutput() GalleryImageVersionPublishingProfilePtrOutput {
	return i.ToGalleryImageVersionPublishingProfilePtrOutputWithContext(context.Background())
}

func (i *galleryImageVersionPublishingProfilePtrType) ToGalleryImageVersionPublishingProfilePtrOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionPublishingProfilePtrOutput)
}

type GalleryImageVersionPublishingProfileOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionPublishingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionPublishingProfile)(nil)).Elem()
}

func (o GalleryImageVersionPublishingProfileOutput) ToGalleryImageVersionPublishingProfileOutput() GalleryImageVersionPublishingProfileOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileOutput) ToGalleryImageVersionPublishingProfileOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfileOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileOutput) ToGalleryImageVersionPublishingProfilePtrOutput() GalleryImageVersionPublishingProfilePtrOutput {
	return o.ToGalleryImageVersionPublishingProfilePtrOutputWithContext(context.Background())
}

func (o GalleryImageVersionPublishingProfileOutput) ToGalleryImageVersionPublishingProfilePtrOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryImageVersionPublishingProfile) *GalleryImageVersionPublishingProfile {
		return &v
	}).(GalleryImageVersionPublishingProfilePtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *string { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) TargetExtendedLocations() GalleryTargetExtendedLocationArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) []GalleryTargetExtendedLocation {
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationArrayOutput)
}

func (o GalleryImageVersionPublishingProfileOutput) TargetRegions() TargetRegionArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) []TargetRegion { return v.TargetRegions }).(TargetRegionArrayOutput)
}

type GalleryImageVersionPublishingProfilePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionPublishingProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersionPublishingProfile)(nil)).Elem()
}

func (o GalleryImageVersionPublishingProfilePtrOutput) ToGalleryImageVersionPublishingProfilePtrOutput() GalleryImageVersionPublishingProfilePtrOutput {
	return o
}

func (o GalleryImageVersionPublishingProfilePtrOutput) ToGalleryImageVersionPublishingProfilePtrOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfilePtrOutput {
	return o
}

func (o GalleryImageVersionPublishingProfilePtrOutput) Elem() GalleryImageVersionPublishingProfileOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) GalleryImageVersionPublishingProfile {
		if v != nil {
			return *v
		}
		var ret GalleryImageVersionPublishingProfile
		return ret
	}).(GalleryImageVersionPublishingProfileOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *string {
		if v == nil {
			return nil
		}
		return v.EndOfLifeDate
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *bool {
		if v == nil {
			return nil
		}
		return v.ExcludeFromLatest
	}).(pulumi.BoolPtrOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *int {
		if v == nil {
			return nil
		}
		return v.ReplicaCount
	}).(pulumi.IntPtrOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationMode
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) TargetExtendedLocations() GalleryTargetExtendedLocationArrayOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) []GalleryTargetExtendedLocation {
		if v == nil {
			return nil
		}
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationArrayOutput)
}

func (o GalleryImageVersionPublishingProfilePtrOutput) TargetRegions() TargetRegionArrayOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) []TargetRegion {
		if v == nil {
			return nil
		}
		return v.TargetRegions
	}).(TargetRegionArrayOutput)
}

type GalleryImageVersionPublishingProfileResponse struct {
	EndOfLifeDate           *string                                 `pulumi:"endOfLifeDate"`
	ExcludeFromLatest       *bool                                   `pulumi:"excludeFromLatest"`
	PublishedDate           string                                  `pulumi:"publishedDate"`
	ReplicaCount            *int                                    `pulumi:"replicaCount"`
	ReplicationMode         *string                                 `pulumi:"replicationMode"`
	StorageAccountType      *string                                 `pulumi:"storageAccountType"`
	TargetExtendedLocations []GalleryTargetExtendedLocationResponse `pulumi:"targetExtendedLocations"`
	TargetRegions           []TargetRegionResponse                  `pulumi:"targetRegions"`
}

type GalleryImageVersionPublishingProfileResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionPublishingProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionPublishingProfileResponse)(nil)).Elem()
}

func (o GalleryImageVersionPublishingProfileResponseOutput) ToGalleryImageVersionPublishingProfileResponseOutput() GalleryImageVersionPublishingProfileResponseOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileResponseOutput) ToGalleryImageVersionPublishingProfileResponseOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfileResponseOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileResponseOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *string { return v.EndOfLifeDate }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *bool { return v.ExcludeFromLatest }).(pulumi.BoolPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) PublishedDate() pulumi.StringOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) string { return v.PublishedDate }).(pulumi.StringOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *string { return v.ReplicationMode }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) TargetExtendedLocations() GalleryTargetExtendedLocationResponseArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) []GalleryTargetExtendedLocationResponse {
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationResponseArrayOutput)
}

func (o GalleryImageVersionPublishingProfileResponseOutput) TargetRegions() TargetRegionResponseArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) []TargetRegionResponse { return v.TargetRegions }).(TargetRegionResponseArrayOutput)
}

type GalleryImageVersionPublishingProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionPublishingProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryImageVersionPublishingProfileResponse)(nil)).Elem()
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) ToGalleryImageVersionPublishingProfileResponsePtrOutput() GalleryImageVersionPublishingProfileResponsePtrOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) ToGalleryImageVersionPublishingProfileResponsePtrOutputWithContext(ctx context.Context) GalleryImageVersionPublishingProfileResponsePtrOutput {
	return o
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) Elem() GalleryImageVersionPublishingProfileResponseOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) GalleryImageVersionPublishingProfileResponse {
		if v != nil {
			return *v
		}
		var ret GalleryImageVersionPublishingProfileResponse
		return ret
	}).(GalleryImageVersionPublishingProfileResponseOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) EndOfLifeDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndOfLifeDate
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) ExcludeFromLatest() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ExcludeFromLatest
	}).(pulumi.BoolPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) PublishedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublishedDate
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.ReplicaCount
	}).(pulumi.IntPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) ReplicationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationMode
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) TargetExtendedLocations() GalleryTargetExtendedLocationResponseArrayOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) []GalleryTargetExtendedLocationResponse {
		if v == nil {
			return nil
		}
		return v.TargetExtendedLocations
	}).(GalleryTargetExtendedLocationResponseArrayOutput)
}

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) TargetRegions() TargetRegionResponseArrayOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) []TargetRegionResponse {
		if v == nil {
			return nil
		}
		return v.TargetRegions
	}).(TargetRegionResponseArrayOutput)
}

type GalleryImageVersionStorageProfile struct {
	DataDiskImages []GalleryDataDiskImage        `pulumi:"dataDiskImages"`
	OsDiskImage    *GalleryOSDiskImage           `pulumi:"osDiskImage"`
	Source         *GalleryArtifactVersionSource `pulumi:"source"`
}





type GalleryImageVersionStorageProfileInput interface {
	pulumi.Input

	ToGalleryImageVersionStorageProfileOutput() GalleryImageVersionStorageProfileOutput
	ToGalleryImageVersionStorageProfileOutputWithContext(context.Context) GalleryImageVersionStorageProfileOutput
}

type GalleryImageVersionStorageProfileArgs struct {
	DataDiskImages GalleryDataDiskImageArrayInput       `pulumi:"dataDiskImages"`
	OsDiskImage    GalleryOSDiskImagePtrInput           `pulumi:"osDiskImage"`
	Source         GalleryArtifactVersionSourcePtrInput `pulumi:"source"`
}

func (GalleryImageVersionStorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionStorageProfile)(nil)).Elem()
}

func (i GalleryImageVersionStorageProfileArgs) ToGalleryImageVersionStorageProfileOutput() GalleryImageVersionStorageProfileOutput {
	return i.ToGalleryImageVersionStorageProfileOutputWithContext(context.Background())
}

func (i GalleryImageVersionStorageProfileArgs) ToGalleryImageVersionStorageProfileOutputWithContext(ctx context.Context) GalleryImageVersionStorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryImageVersionStorageProfileOutput)
}

type GalleryImageVersionStorageProfileOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionStorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionStorageProfile)(nil)).Elem()
}

func (o GalleryImageVersionStorageProfileOutput) ToGalleryImageVersionStorageProfileOutput() GalleryImageVersionStorageProfileOutput {
	return o
}

func (o GalleryImageVersionStorageProfileOutput) ToGalleryImageVersionStorageProfileOutputWithContext(ctx context.Context) GalleryImageVersionStorageProfileOutput {
	return o
}

func (o GalleryImageVersionStorageProfileOutput) DataDiskImages() GalleryDataDiskImageArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfile) []GalleryDataDiskImage { return v.DataDiskImages }).(GalleryDataDiskImageArrayOutput)
}

func (o GalleryImageVersionStorageProfileOutput) OsDiskImage() GalleryOSDiskImagePtrOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfile) *GalleryOSDiskImage { return v.OsDiskImage }).(GalleryOSDiskImagePtrOutput)
}

func (o GalleryImageVersionStorageProfileOutput) Source() GalleryArtifactVersionSourcePtrOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfile) *GalleryArtifactVersionSource { return v.Source }).(GalleryArtifactVersionSourcePtrOutput)
}

type GalleryImageVersionStorageProfileResponse struct {
	DataDiskImages []GalleryDataDiskImageResponse        `pulumi:"dataDiskImages"`
	OsDiskImage    *GalleryOSDiskImageResponse           `pulumi:"osDiskImage"`
	Source         *GalleryArtifactVersionSourceResponse `pulumi:"source"`
}

type GalleryImageVersionStorageProfileResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageVersionStorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageVersionStorageProfileResponse)(nil)).Elem()
}

func (o GalleryImageVersionStorageProfileResponseOutput) ToGalleryImageVersionStorageProfileResponseOutput() GalleryImageVersionStorageProfileResponseOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponseOutput) ToGalleryImageVersionStorageProfileResponseOutputWithContext(ctx context.Context) GalleryImageVersionStorageProfileResponseOutput {
	return o
}

func (o GalleryImageVersionStorageProfileResponseOutput) DataDiskImages() GalleryDataDiskImageResponseArrayOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfileResponse) []GalleryDataDiskImageResponse {
		return v.DataDiskImages
	}).(GalleryDataDiskImageResponseArrayOutput)
}

func (o GalleryImageVersionStorageProfileResponseOutput) OsDiskImage() GalleryOSDiskImageResponsePtrOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfileResponse) *GalleryOSDiskImageResponse { return v.OsDiskImage }).(GalleryOSDiskImageResponsePtrOutput)
}

func (o GalleryImageVersionStorageProfileResponseOutput) Source() GalleryArtifactVersionSourceResponsePtrOutput {
	return o.ApplyT(func(v GalleryImageVersionStorageProfileResponse) *GalleryArtifactVersionSourceResponse {
		return v.Source
	}).(GalleryArtifactVersionSourceResponsePtrOutput)
}

type GalleryOSDiskImage struct {
	HostCaching *HostCaching                  `pulumi:"hostCaching"`
	Source      *GalleryArtifactVersionSource `pulumi:"source"`
}





type GalleryOSDiskImageInput interface {
	pulumi.Input

	ToGalleryOSDiskImageOutput() GalleryOSDiskImageOutput
	ToGalleryOSDiskImageOutputWithContext(context.Context) GalleryOSDiskImageOutput
}

type GalleryOSDiskImageArgs struct {
	HostCaching HostCachingPtrInput                  `pulumi:"hostCaching"`
	Source      GalleryArtifactVersionSourcePtrInput `pulumi:"source"`
}

func (GalleryOSDiskImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryOSDiskImage)(nil)).Elem()
}

func (i GalleryOSDiskImageArgs) ToGalleryOSDiskImageOutput() GalleryOSDiskImageOutput {
	return i.ToGalleryOSDiskImageOutputWithContext(context.Background())
}

func (i GalleryOSDiskImageArgs) ToGalleryOSDiskImageOutputWithContext(ctx context.Context) GalleryOSDiskImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOSDiskImageOutput)
}

func (i GalleryOSDiskImageArgs) ToGalleryOSDiskImagePtrOutput() GalleryOSDiskImagePtrOutput {
	return i.ToGalleryOSDiskImagePtrOutputWithContext(context.Background())
}

func (i GalleryOSDiskImageArgs) ToGalleryOSDiskImagePtrOutputWithContext(ctx context.Context) GalleryOSDiskImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOSDiskImageOutput).ToGalleryOSDiskImagePtrOutputWithContext(ctx)
}









type GalleryOSDiskImagePtrInput interface {
	pulumi.Input

	ToGalleryOSDiskImagePtrOutput() GalleryOSDiskImagePtrOutput
	ToGalleryOSDiskImagePtrOutputWithContext(context.Context) GalleryOSDiskImagePtrOutput
}

type galleryOSDiskImagePtrType GalleryOSDiskImageArgs

func GalleryOSDiskImagePtr(v *GalleryOSDiskImageArgs) GalleryOSDiskImagePtrInput {
	return (*galleryOSDiskImagePtrType)(v)
}

func (*galleryOSDiskImagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryOSDiskImage)(nil)).Elem()
}

func (i *galleryOSDiskImagePtrType) ToGalleryOSDiskImagePtrOutput() GalleryOSDiskImagePtrOutput {
	return i.ToGalleryOSDiskImagePtrOutputWithContext(context.Background())
}

func (i *galleryOSDiskImagePtrType) ToGalleryOSDiskImagePtrOutputWithContext(ctx context.Context) GalleryOSDiskImagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryOSDiskImagePtrOutput)
}

type GalleryOSDiskImageOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryOSDiskImage)(nil)).Elem()
}

func (o GalleryOSDiskImageOutput) ToGalleryOSDiskImageOutput() GalleryOSDiskImageOutput {
	return o
}

func (o GalleryOSDiskImageOutput) ToGalleryOSDiskImageOutputWithContext(ctx context.Context) GalleryOSDiskImageOutput {
	return o
}

func (o GalleryOSDiskImageOutput) ToGalleryOSDiskImagePtrOutput() GalleryOSDiskImagePtrOutput {
	return o.ToGalleryOSDiskImagePtrOutputWithContext(context.Background())
}

func (o GalleryOSDiskImageOutput) ToGalleryOSDiskImagePtrOutputWithContext(ctx context.Context) GalleryOSDiskImagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GalleryOSDiskImage) *GalleryOSDiskImage {
		return &v
	}).(GalleryOSDiskImagePtrOutput)
}

func (o GalleryOSDiskImageOutput) HostCaching() HostCachingPtrOutput {
	return o.ApplyT(func(v GalleryOSDiskImage) *HostCaching { return v.HostCaching }).(HostCachingPtrOutput)
}

func (o GalleryOSDiskImageOutput) Source() GalleryArtifactVersionSourcePtrOutput {
	return o.ApplyT(func(v GalleryOSDiskImage) *GalleryArtifactVersionSource { return v.Source }).(GalleryArtifactVersionSourcePtrOutput)
}

type GalleryOSDiskImagePtrOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryOSDiskImage)(nil)).Elem()
}

func (o GalleryOSDiskImagePtrOutput) ToGalleryOSDiskImagePtrOutput() GalleryOSDiskImagePtrOutput {
	return o
}

func (o GalleryOSDiskImagePtrOutput) ToGalleryOSDiskImagePtrOutputWithContext(ctx context.Context) GalleryOSDiskImagePtrOutput {
	return o
}

func (o GalleryOSDiskImagePtrOutput) Elem() GalleryOSDiskImageOutput {
	return o.ApplyT(func(v *GalleryOSDiskImage) GalleryOSDiskImage {
		if v != nil {
			return *v
		}
		var ret GalleryOSDiskImage
		return ret
	}).(GalleryOSDiskImageOutput)
}

func (o GalleryOSDiskImagePtrOutput) HostCaching() HostCachingPtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImage) *HostCaching {
		if v == nil {
			return nil
		}
		return v.HostCaching
	}).(HostCachingPtrOutput)
}

func (o GalleryOSDiskImagePtrOutput) Source() GalleryArtifactVersionSourcePtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImage) *GalleryArtifactVersionSource {
		if v == nil {
			return nil
		}
		return v.Source
	}).(GalleryArtifactVersionSourcePtrOutput)
}

type GalleryOSDiskImageResponse struct {
	HostCaching *string                               `pulumi:"hostCaching"`
	SizeInGB    int                                   `pulumi:"sizeInGB"`
	Source      *GalleryArtifactVersionSourceResponse `pulumi:"source"`
}

type GalleryOSDiskImageResponseOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryOSDiskImageResponse)(nil)).Elem()
}

func (o GalleryOSDiskImageResponseOutput) ToGalleryOSDiskImageResponseOutput() GalleryOSDiskImageResponseOutput {
	return o
}

func (o GalleryOSDiskImageResponseOutput) ToGalleryOSDiskImageResponseOutputWithContext(ctx context.Context) GalleryOSDiskImageResponseOutput {
	return o
}

func (o GalleryOSDiskImageResponseOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryOSDiskImageResponse) *string { return v.HostCaching }).(pulumi.StringPtrOutput)
}

func (o GalleryOSDiskImageResponseOutput) SizeInGB() pulumi.IntOutput {
	return o.ApplyT(func(v GalleryOSDiskImageResponse) int { return v.SizeInGB }).(pulumi.IntOutput)
}

func (o GalleryOSDiskImageResponseOutput) Source() GalleryArtifactVersionSourceResponsePtrOutput {
	return o.ApplyT(func(v GalleryOSDiskImageResponse) *GalleryArtifactVersionSourceResponse { return v.Source }).(GalleryArtifactVersionSourceResponsePtrOutput)
}

type GalleryOSDiskImageResponsePtrOutput struct{ *pulumi.OutputState }

func (GalleryOSDiskImageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GalleryOSDiskImageResponse)(nil)).Elem()
}

func (o GalleryOSDiskImageResponsePtrOutput) ToGalleryOSDiskImageResponsePtrOutput() GalleryOSDiskImageResponsePtrOutput {
	return o
}

func (o GalleryOSDiskImageResponsePtrOutput) ToGalleryOSDiskImageResponsePtrOutputWithContext(ctx context.Context) GalleryOSDiskImageResponsePtrOutput {
	return o
}

func (o GalleryOSDiskImageResponsePtrOutput) Elem() GalleryOSDiskImageResponseOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) GalleryOSDiskImageResponse {
		if v != nil {
			return *v
		}
		var ret GalleryOSDiskImageResponse
		return ret
	}).(GalleryOSDiskImageResponseOutput)
}

func (o GalleryOSDiskImageResponsePtrOutput) HostCaching() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) *string {
		if v == nil {
			return nil
		}
		return v.HostCaching
	}).(pulumi.StringPtrOutput)
}

func (o GalleryOSDiskImageResponsePtrOutput) SizeInGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) *int {
		if v == nil {
			return nil
		}
		return &v.SizeInGB
	}).(pulumi.IntPtrOutput)
}

func (o GalleryOSDiskImageResponsePtrOutput) Source() GalleryArtifactVersionSourceResponsePtrOutput {
	return o.ApplyT(func(v *GalleryOSDiskImageResponse) *GalleryArtifactVersionSourceResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(GalleryArtifactVersionSourceResponsePtrOutput)
}

type GalleryTargetExtendedLocation struct {
	Encryption                   *EncryptionImages        `pulumi:"encryption"`
	ExtendedLocation             *GalleryExtendedLocation `pulumi:"extendedLocation"`
	ExtendedLocationReplicaCount *int                     `pulumi:"extendedLocationReplicaCount"`
	Name                         *string                  `pulumi:"name"`
	StorageAccountType           *string                  `pulumi:"storageAccountType"`
}





type GalleryTargetExtendedLocationInput interface {
	pulumi.Input

	ToGalleryTargetExtendedLocationOutput() GalleryTargetExtendedLocationOutput
	ToGalleryTargetExtendedLocationOutputWithContext(context.Context) GalleryTargetExtendedLocationOutput
}

type GalleryTargetExtendedLocationArgs struct {
	Encryption                   EncryptionImagesPtrInput        `pulumi:"encryption"`
	ExtendedLocation             GalleryExtendedLocationPtrInput `pulumi:"extendedLocation"`
	ExtendedLocationReplicaCount pulumi.IntPtrInput              `pulumi:"extendedLocationReplicaCount"`
	Name                         pulumi.StringPtrInput           `pulumi:"name"`
	StorageAccountType           pulumi.StringPtrInput           `pulumi:"storageAccountType"`
}

func (GalleryTargetExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryTargetExtendedLocation)(nil)).Elem()
}

func (i GalleryTargetExtendedLocationArgs) ToGalleryTargetExtendedLocationOutput() GalleryTargetExtendedLocationOutput {
	return i.ToGalleryTargetExtendedLocationOutputWithContext(context.Background())
}

func (i GalleryTargetExtendedLocationArgs) ToGalleryTargetExtendedLocationOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryTargetExtendedLocationOutput)
}





type GalleryTargetExtendedLocationArrayInput interface {
	pulumi.Input

	ToGalleryTargetExtendedLocationArrayOutput() GalleryTargetExtendedLocationArrayOutput
	ToGalleryTargetExtendedLocationArrayOutputWithContext(context.Context) GalleryTargetExtendedLocationArrayOutput
}

type GalleryTargetExtendedLocationArray []GalleryTargetExtendedLocationInput

func (GalleryTargetExtendedLocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryTargetExtendedLocation)(nil)).Elem()
}

func (i GalleryTargetExtendedLocationArray) ToGalleryTargetExtendedLocationArrayOutput() GalleryTargetExtendedLocationArrayOutput {
	return i.ToGalleryTargetExtendedLocationArrayOutputWithContext(context.Background())
}

func (i GalleryTargetExtendedLocationArray) ToGalleryTargetExtendedLocationArrayOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryTargetExtendedLocationArrayOutput)
}

type GalleryTargetExtendedLocationOutput struct{ *pulumi.OutputState }

func (GalleryTargetExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryTargetExtendedLocation)(nil)).Elem()
}

func (o GalleryTargetExtendedLocationOutput) ToGalleryTargetExtendedLocationOutput() GalleryTargetExtendedLocationOutput {
	return o
}

func (o GalleryTargetExtendedLocationOutput) ToGalleryTargetExtendedLocationOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationOutput {
	return o
}

func (o GalleryTargetExtendedLocationOutput) Encryption() EncryptionImagesPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocation) *EncryptionImages { return v.Encryption }).(EncryptionImagesPtrOutput)
}

func (o GalleryTargetExtendedLocationOutput) ExtendedLocation() GalleryExtendedLocationPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocation) *GalleryExtendedLocation { return v.ExtendedLocation }).(GalleryExtendedLocationPtrOutput)
}

func (o GalleryTargetExtendedLocationOutput) ExtendedLocationReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocation) *int { return v.ExtendedLocationReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryTargetExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryTargetExtendedLocationOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocation) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type GalleryTargetExtendedLocationArrayOutput struct{ *pulumi.OutputState }

func (GalleryTargetExtendedLocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryTargetExtendedLocation)(nil)).Elem()
}

func (o GalleryTargetExtendedLocationArrayOutput) ToGalleryTargetExtendedLocationArrayOutput() GalleryTargetExtendedLocationArrayOutput {
	return o
}

func (o GalleryTargetExtendedLocationArrayOutput) ToGalleryTargetExtendedLocationArrayOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationArrayOutput {
	return o
}

func (o GalleryTargetExtendedLocationArrayOutput) Index(i pulumi.IntInput) GalleryTargetExtendedLocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryTargetExtendedLocation {
		return vs[0].([]GalleryTargetExtendedLocation)[vs[1].(int)]
	}).(GalleryTargetExtendedLocationOutput)
}

type GalleryTargetExtendedLocationResponse struct {
	Encryption                   *EncryptionImagesResponse        `pulumi:"encryption"`
	ExtendedLocation             *GalleryExtendedLocationResponse `pulumi:"extendedLocation"`
	ExtendedLocationReplicaCount *int                             `pulumi:"extendedLocationReplicaCount"`
	Name                         *string                          `pulumi:"name"`
	StorageAccountType           *string                          `pulumi:"storageAccountType"`
}

type GalleryTargetExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (GalleryTargetExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryTargetExtendedLocationResponse)(nil)).Elem()
}

func (o GalleryTargetExtendedLocationResponseOutput) ToGalleryTargetExtendedLocationResponseOutput() GalleryTargetExtendedLocationResponseOutput {
	return o
}

func (o GalleryTargetExtendedLocationResponseOutput) ToGalleryTargetExtendedLocationResponseOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationResponseOutput {
	return o
}

func (o GalleryTargetExtendedLocationResponseOutput) Encryption() EncryptionImagesResponsePtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocationResponse) *EncryptionImagesResponse { return v.Encryption }).(EncryptionImagesResponsePtrOutput)
}

func (o GalleryTargetExtendedLocationResponseOutput) ExtendedLocation() GalleryExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocationResponse) *GalleryExtendedLocationResponse {
		return v.ExtendedLocation
	}).(GalleryExtendedLocationResponsePtrOutput)
}

func (o GalleryTargetExtendedLocationResponseOutput) ExtendedLocationReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocationResponse) *int { return v.ExtendedLocationReplicaCount }).(pulumi.IntPtrOutput)
}

func (o GalleryTargetExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GalleryTargetExtendedLocationResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryTargetExtendedLocationResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type GalleryTargetExtendedLocationResponseArrayOutput struct{ *pulumi.OutputState }

func (GalleryTargetExtendedLocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GalleryTargetExtendedLocationResponse)(nil)).Elem()
}

func (o GalleryTargetExtendedLocationResponseArrayOutput) ToGalleryTargetExtendedLocationResponseArrayOutput() GalleryTargetExtendedLocationResponseArrayOutput {
	return o
}

func (o GalleryTargetExtendedLocationResponseArrayOutput) ToGalleryTargetExtendedLocationResponseArrayOutputWithContext(ctx context.Context) GalleryTargetExtendedLocationResponseArrayOutput {
	return o
}

func (o GalleryTargetExtendedLocationResponseArrayOutput) Index(i pulumi.IntInput) GalleryTargetExtendedLocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GalleryTargetExtendedLocationResponse {
		return vs[0].([]GalleryTargetExtendedLocationResponse)[vs[1].(int)]
	}).(GalleryTargetExtendedLocationResponseOutput)
}

type ImagePurchasePlan struct {
	Name      *string `pulumi:"name"`
	Product   *string `pulumi:"product"`
	Publisher *string `pulumi:"publisher"`
}





type ImagePurchasePlanInput interface {
	pulumi.Input

	ToImagePurchasePlanOutput() ImagePurchasePlanOutput
	ToImagePurchasePlanOutputWithContext(context.Context) ImagePurchasePlanOutput
}

type ImagePurchasePlanArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	Product   pulumi.StringPtrInput `pulumi:"product"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
}

func (ImagePurchasePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImagePurchasePlan)(nil)).Elem()
}

func (i ImagePurchasePlanArgs) ToImagePurchasePlanOutput() ImagePurchasePlanOutput {
	return i.ToImagePurchasePlanOutputWithContext(context.Background())
}

func (i ImagePurchasePlanArgs) ToImagePurchasePlanOutputWithContext(ctx context.Context) ImagePurchasePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePurchasePlanOutput)
}

func (i ImagePurchasePlanArgs) ToImagePurchasePlanPtrOutput() ImagePurchasePlanPtrOutput {
	return i.ToImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (i ImagePurchasePlanArgs) ToImagePurchasePlanPtrOutputWithContext(ctx context.Context) ImagePurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePurchasePlanOutput).ToImagePurchasePlanPtrOutputWithContext(ctx)
}









type ImagePurchasePlanPtrInput interface {
	pulumi.Input

	ToImagePurchasePlanPtrOutput() ImagePurchasePlanPtrOutput
	ToImagePurchasePlanPtrOutputWithContext(context.Context) ImagePurchasePlanPtrOutput
}

type imagePurchasePlanPtrType ImagePurchasePlanArgs

func ImagePurchasePlanPtr(v *ImagePurchasePlanArgs) ImagePurchasePlanPtrInput {
	return (*imagePurchasePlanPtrType)(v)
}

func (*imagePurchasePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagePurchasePlan)(nil)).Elem()
}

func (i *imagePurchasePlanPtrType) ToImagePurchasePlanPtrOutput() ImagePurchasePlanPtrOutput {
	return i.ToImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (i *imagePurchasePlanPtrType) ToImagePurchasePlanPtrOutputWithContext(ctx context.Context) ImagePurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImagePurchasePlanPtrOutput)
}

type ImagePurchasePlanOutput struct{ *pulumi.OutputState }

func (ImagePurchasePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImagePurchasePlan)(nil)).Elem()
}

func (o ImagePurchasePlanOutput) ToImagePurchasePlanOutput() ImagePurchasePlanOutput {
	return o
}

func (o ImagePurchasePlanOutput) ToImagePurchasePlanOutputWithContext(ctx context.Context) ImagePurchasePlanOutput {
	return o
}

func (o ImagePurchasePlanOutput) ToImagePurchasePlanPtrOutput() ImagePurchasePlanPtrOutput {
	return o.ToImagePurchasePlanPtrOutputWithContext(context.Background())
}

func (o ImagePurchasePlanOutput) ToImagePurchasePlanPtrOutputWithContext(ctx context.Context) ImagePurchasePlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImagePurchasePlan) *ImagePurchasePlan {
		return &v
	}).(ImagePurchasePlanPtrOutput)
}

func (o ImagePurchasePlanOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlan) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlan) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlan) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type ImagePurchasePlanPtrOutput struct{ *pulumi.OutputState }

func (ImagePurchasePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagePurchasePlan)(nil)).Elem()
}

func (o ImagePurchasePlanPtrOutput) ToImagePurchasePlanPtrOutput() ImagePurchasePlanPtrOutput {
	return o
}

func (o ImagePurchasePlanPtrOutput) ToImagePurchasePlanPtrOutputWithContext(ctx context.Context) ImagePurchasePlanPtrOutput {
	return o
}

func (o ImagePurchasePlanPtrOutput) Elem() ImagePurchasePlanOutput {
	return o.ApplyT(func(v *ImagePurchasePlan) ImagePurchasePlan {
		if v != nil {
			return *v
		}
		var ret ImagePurchasePlan
		return ret
	}).(ImagePurchasePlanOutput)
}

func (o ImagePurchasePlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type ImagePurchasePlanResponse struct {
	Name      *string `pulumi:"name"`
	Product   *string `pulumi:"product"`
	Publisher *string `pulumi:"publisher"`
}

type ImagePurchasePlanResponseOutput struct{ *pulumi.OutputState }

func (ImagePurchasePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImagePurchasePlanResponse)(nil)).Elem()
}

func (o ImagePurchasePlanResponseOutput) ToImagePurchasePlanResponseOutput() ImagePurchasePlanResponseOutput {
	return o
}

func (o ImagePurchasePlanResponseOutput) ToImagePurchasePlanResponseOutputWithContext(ctx context.Context) ImagePurchasePlanResponseOutput {
	return o
}

func (o ImagePurchasePlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImagePurchasePlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

type ImagePurchasePlanResponsePtrOutput struct{ *pulumi.OutputState }

func (ImagePurchasePlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImagePurchasePlanResponse)(nil)).Elem()
}

func (o ImagePurchasePlanResponsePtrOutput) ToImagePurchasePlanResponsePtrOutput() ImagePurchasePlanResponsePtrOutput {
	return o
}

func (o ImagePurchasePlanResponsePtrOutput) ToImagePurchasePlanResponsePtrOutputWithContext(ctx context.Context) ImagePurchasePlanResponsePtrOutput {
	return o
}

func (o ImagePurchasePlanResponsePtrOutput) Elem() ImagePurchasePlanResponseOutput {
	return o.ApplyT(func(v *ImagePurchasePlanResponse) ImagePurchasePlanResponse {
		if v != nil {
			return *v
		}
		var ret ImagePurchasePlanResponse
		return ret
	}).(ImagePurchasePlanResponseOutput)
}

func (o ImagePurchasePlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ImagePurchasePlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImagePurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

type OSDiskImageEncryption struct {
	DiskEncryptionSetId *string                     `pulumi:"diskEncryptionSetId"`
	SecurityProfile     *OSDiskImageSecurityProfile `pulumi:"securityProfile"`
}





type OSDiskImageEncryptionInput interface {
	pulumi.Input

	ToOSDiskImageEncryptionOutput() OSDiskImageEncryptionOutput
	ToOSDiskImageEncryptionOutputWithContext(context.Context) OSDiskImageEncryptionOutput
}

type OSDiskImageEncryptionArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput              `pulumi:"diskEncryptionSetId"`
	SecurityProfile     OSDiskImageSecurityProfilePtrInput `pulumi:"securityProfile"`
}

func (OSDiskImageEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageEncryption)(nil)).Elem()
}

func (i OSDiskImageEncryptionArgs) ToOSDiskImageEncryptionOutput() OSDiskImageEncryptionOutput {
	return i.ToOSDiskImageEncryptionOutputWithContext(context.Background())
}

func (i OSDiskImageEncryptionArgs) ToOSDiskImageEncryptionOutputWithContext(ctx context.Context) OSDiskImageEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageEncryptionOutput)
}

func (i OSDiskImageEncryptionArgs) ToOSDiskImageEncryptionPtrOutput() OSDiskImageEncryptionPtrOutput {
	return i.ToOSDiskImageEncryptionPtrOutputWithContext(context.Background())
}

func (i OSDiskImageEncryptionArgs) ToOSDiskImageEncryptionPtrOutputWithContext(ctx context.Context) OSDiskImageEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageEncryptionOutput).ToOSDiskImageEncryptionPtrOutputWithContext(ctx)
}









type OSDiskImageEncryptionPtrInput interface {
	pulumi.Input

	ToOSDiskImageEncryptionPtrOutput() OSDiskImageEncryptionPtrOutput
	ToOSDiskImageEncryptionPtrOutputWithContext(context.Context) OSDiskImageEncryptionPtrOutput
}

type osdiskImageEncryptionPtrType OSDiskImageEncryptionArgs

func OSDiskImageEncryptionPtr(v *OSDiskImageEncryptionArgs) OSDiskImageEncryptionPtrInput {
	return (*osdiskImageEncryptionPtrType)(v)
}

func (*osdiskImageEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageEncryption)(nil)).Elem()
}

func (i *osdiskImageEncryptionPtrType) ToOSDiskImageEncryptionPtrOutput() OSDiskImageEncryptionPtrOutput {
	return i.ToOSDiskImageEncryptionPtrOutputWithContext(context.Background())
}

func (i *osdiskImageEncryptionPtrType) ToOSDiskImageEncryptionPtrOutputWithContext(ctx context.Context) OSDiskImageEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageEncryptionPtrOutput)
}

type OSDiskImageEncryptionOutput struct{ *pulumi.OutputState }

func (OSDiskImageEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageEncryption)(nil)).Elem()
}

func (o OSDiskImageEncryptionOutput) ToOSDiskImageEncryptionOutput() OSDiskImageEncryptionOutput {
	return o
}

func (o OSDiskImageEncryptionOutput) ToOSDiskImageEncryptionOutputWithContext(ctx context.Context) OSDiskImageEncryptionOutput {
	return o
}

func (o OSDiskImageEncryptionOutput) ToOSDiskImageEncryptionPtrOutput() OSDiskImageEncryptionPtrOutput {
	return o.ToOSDiskImageEncryptionPtrOutputWithContext(context.Background())
}

func (o OSDiskImageEncryptionOutput) ToOSDiskImageEncryptionPtrOutputWithContext(ctx context.Context) OSDiskImageEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDiskImageEncryption) *OSDiskImageEncryption {
		return &v
	}).(OSDiskImageEncryptionPtrOutput)
}

func (o OSDiskImageEncryptionOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageEncryption) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o OSDiskImageEncryptionOutput) SecurityProfile() OSDiskImageSecurityProfilePtrOutput {
	return o.ApplyT(func(v OSDiskImageEncryption) *OSDiskImageSecurityProfile { return v.SecurityProfile }).(OSDiskImageSecurityProfilePtrOutput)
}

type OSDiskImageEncryptionPtrOutput struct{ *pulumi.OutputState }

func (OSDiskImageEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageEncryption)(nil)).Elem()
}

func (o OSDiskImageEncryptionPtrOutput) ToOSDiskImageEncryptionPtrOutput() OSDiskImageEncryptionPtrOutput {
	return o
}

func (o OSDiskImageEncryptionPtrOutput) ToOSDiskImageEncryptionPtrOutputWithContext(ctx context.Context) OSDiskImageEncryptionPtrOutput {
	return o
}

func (o OSDiskImageEncryptionPtrOutput) Elem() OSDiskImageEncryptionOutput {
	return o.ApplyT(func(v *OSDiskImageEncryption) OSDiskImageEncryption {
		if v != nil {
			return *v
		}
		var ret OSDiskImageEncryption
		return ret
	}).(OSDiskImageEncryptionOutput)
}

func (o OSDiskImageEncryptionPtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageEncryption) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskImageEncryptionPtrOutput) SecurityProfile() OSDiskImageSecurityProfilePtrOutput {
	return o.ApplyT(func(v *OSDiskImageEncryption) *OSDiskImageSecurityProfile {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(OSDiskImageSecurityProfilePtrOutput)
}

type OSDiskImageEncryptionResponse struct {
	DiskEncryptionSetId *string                             `pulumi:"diskEncryptionSetId"`
	SecurityProfile     *OSDiskImageSecurityProfileResponse `pulumi:"securityProfile"`
}

type OSDiskImageEncryptionResponseOutput struct{ *pulumi.OutputState }

func (OSDiskImageEncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageEncryptionResponse)(nil)).Elem()
}

func (o OSDiskImageEncryptionResponseOutput) ToOSDiskImageEncryptionResponseOutput() OSDiskImageEncryptionResponseOutput {
	return o
}

func (o OSDiskImageEncryptionResponseOutput) ToOSDiskImageEncryptionResponseOutputWithContext(ctx context.Context) OSDiskImageEncryptionResponseOutput {
	return o
}

func (o OSDiskImageEncryptionResponseOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageEncryptionResponse) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o OSDiskImageEncryptionResponseOutput) SecurityProfile() OSDiskImageSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v OSDiskImageEncryptionResponse) *OSDiskImageSecurityProfileResponse { return v.SecurityProfile }).(OSDiskImageSecurityProfileResponsePtrOutput)
}

type OSDiskImageEncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskImageEncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageEncryptionResponse)(nil)).Elem()
}

func (o OSDiskImageEncryptionResponsePtrOutput) ToOSDiskImageEncryptionResponsePtrOutput() OSDiskImageEncryptionResponsePtrOutput {
	return o
}

func (o OSDiskImageEncryptionResponsePtrOutput) ToOSDiskImageEncryptionResponsePtrOutputWithContext(ctx context.Context) OSDiskImageEncryptionResponsePtrOutput {
	return o
}

func (o OSDiskImageEncryptionResponsePtrOutput) Elem() OSDiskImageEncryptionResponseOutput {
	return o.ApplyT(func(v *OSDiskImageEncryptionResponse) OSDiskImageEncryptionResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskImageEncryptionResponse
		return ret
	}).(OSDiskImageEncryptionResponseOutput)
}

func (o OSDiskImageEncryptionResponsePtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageEncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskImageEncryptionResponsePtrOutput) SecurityProfile() OSDiskImageSecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *OSDiskImageEncryptionResponse) *OSDiskImageSecurityProfileResponse {
		if v == nil {
			return nil
		}
		return v.SecurityProfile
	}).(OSDiskImageSecurityProfileResponsePtrOutput)
}

type OSDiskImageSecurityProfile struct {
	ConfidentialVMEncryptionType *string `pulumi:"confidentialVMEncryptionType"`
	SecureVMDiskEncryptionSetId  *string `pulumi:"secureVMDiskEncryptionSetId"`
}





type OSDiskImageSecurityProfileInput interface {
	pulumi.Input

	ToOSDiskImageSecurityProfileOutput() OSDiskImageSecurityProfileOutput
	ToOSDiskImageSecurityProfileOutputWithContext(context.Context) OSDiskImageSecurityProfileOutput
}

type OSDiskImageSecurityProfileArgs struct {
	ConfidentialVMEncryptionType pulumi.StringPtrInput `pulumi:"confidentialVMEncryptionType"`
	SecureVMDiskEncryptionSetId  pulumi.StringPtrInput `pulumi:"secureVMDiskEncryptionSetId"`
}

func (OSDiskImageSecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageSecurityProfile)(nil)).Elem()
}

func (i OSDiskImageSecurityProfileArgs) ToOSDiskImageSecurityProfileOutput() OSDiskImageSecurityProfileOutput {
	return i.ToOSDiskImageSecurityProfileOutputWithContext(context.Background())
}

func (i OSDiskImageSecurityProfileArgs) ToOSDiskImageSecurityProfileOutputWithContext(ctx context.Context) OSDiskImageSecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageSecurityProfileOutput)
}

func (i OSDiskImageSecurityProfileArgs) ToOSDiskImageSecurityProfilePtrOutput() OSDiskImageSecurityProfilePtrOutput {
	return i.ToOSDiskImageSecurityProfilePtrOutputWithContext(context.Background())
}

func (i OSDiskImageSecurityProfileArgs) ToOSDiskImageSecurityProfilePtrOutputWithContext(ctx context.Context) OSDiskImageSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageSecurityProfileOutput).ToOSDiskImageSecurityProfilePtrOutputWithContext(ctx)
}









type OSDiskImageSecurityProfilePtrInput interface {
	pulumi.Input

	ToOSDiskImageSecurityProfilePtrOutput() OSDiskImageSecurityProfilePtrOutput
	ToOSDiskImageSecurityProfilePtrOutputWithContext(context.Context) OSDiskImageSecurityProfilePtrOutput
}

type osdiskImageSecurityProfilePtrType OSDiskImageSecurityProfileArgs

func OSDiskImageSecurityProfilePtr(v *OSDiskImageSecurityProfileArgs) OSDiskImageSecurityProfilePtrInput {
	return (*osdiskImageSecurityProfilePtrType)(v)
}

func (*osdiskImageSecurityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageSecurityProfile)(nil)).Elem()
}

func (i *osdiskImageSecurityProfilePtrType) ToOSDiskImageSecurityProfilePtrOutput() OSDiskImageSecurityProfilePtrOutput {
	return i.ToOSDiskImageSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *osdiskImageSecurityProfilePtrType) ToOSDiskImageSecurityProfilePtrOutputWithContext(ctx context.Context) OSDiskImageSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskImageSecurityProfilePtrOutput)
}

type OSDiskImageSecurityProfileOutput struct{ *pulumi.OutputState }

func (OSDiskImageSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageSecurityProfile)(nil)).Elem()
}

func (o OSDiskImageSecurityProfileOutput) ToOSDiskImageSecurityProfileOutput() OSDiskImageSecurityProfileOutput {
	return o
}

func (o OSDiskImageSecurityProfileOutput) ToOSDiskImageSecurityProfileOutputWithContext(ctx context.Context) OSDiskImageSecurityProfileOutput {
	return o
}

func (o OSDiskImageSecurityProfileOutput) ToOSDiskImageSecurityProfilePtrOutput() OSDiskImageSecurityProfilePtrOutput {
	return o.ToOSDiskImageSecurityProfilePtrOutputWithContext(context.Background())
}

func (o OSDiskImageSecurityProfileOutput) ToOSDiskImageSecurityProfilePtrOutputWithContext(ctx context.Context) OSDiskImageSecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDiskImageSecurityProfile) *OSDiskImageSecurityProfile {
		return &v
	}).(OSDiskImageSecurityProfilePtrOutput)
}

func (o OSDiskImageSecurityProfileOutput) ConfidentialVMEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageSecurityProfile) *string { return v.ConfidentialVMEncryptionType }).(pulumi.StringPtrOutput)
}

func (o OSDiskImageSecurityProfileOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageSecurityProfile) *string { return v.SecureVMDiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

type OSDiskImageSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (OSDiskImageSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageSecurityProfile)(nil)).Elem()
}

func (o OSDiskImageSecurityProfilePtrOutput) ToOSDiskImageSecurityProfilePtrOutput() OSDiskImageSecurityProfilePtrOutput {
	return o
}

func (o OSDiskImageSecurityProfilePtrOutput) ToOSDiskImageSecurityProfilePtrOutputWithContext(ctx context.Context) OSDiskImageSecurityProfilePtrOutput {
	return o
}

func (o OSDiskImageSecurityProfilePtrOutput) Elem() OSDiskImageSecurityProfileOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfile) OSDiskImageSecurityProfile {
		if v != nil {
			return *v
		}
		var ret OSDiskImageSecurityProfile
		return ret
	}).(OSDiskImageSecurityProfileOutput)
}

func (o OSDiskImageSecurityProfilePtrOutput) ConfidentialVMEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.ConfidentialVMEncryptionType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskImageSecurityProfilePtrOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.SecureVMDiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

type OSDiskImageSecurityProfileResponse struct {
	ConfidentialVMEncryptionType *string `pulumi:"confidentialVMEncryptionType"`
	SecureVMDiskEncryptionSetId  *string `pulumi:"secureVMDiskEncryptionSetId"`
}

type OSDiskImageSecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (OSDiskImageSecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskImageSecurityProfileResponse)(nil)).Elem()
}

func (o OSDiskImageSecurityProfileResponseOutput) ToOSDiskImageSecurityProfileResponseOutput() OSDiskImageSecurityProfileResponseOutput {
	return o
}

func (o OSDiskImageSecurityProfileResponseOutput) ToOSDiskImageSecurityProfileResponseOutputWithContext(ctx context.Context) OSDiskImageSecurityProfileResponseOutput {
	return o
}

func (o OSDiskImageSecurityProfileResponseOutput) ConfidentialVMEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageSecurityProfileResponse) *string { return v.ConfidentialVMEncryptionType }).(pulumi.StringPtrOutput)
}

func (o OSDiskImageSecurityProfileResponseOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskImageSecurityProfileResponse) *string { return v.SecureVMDiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

type OSDiskImageSecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskImageSecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskImageSecurityProfileResponse)(nil)).Elem()
}

func (o OSDiskImageSecurityProfileResponsePtrOutput) ToOSDiskImageSecurityProfileResponsePtrOutput() OSDiskImageSecurityProfileResponsePtrOutput {
	return o
}

func (o OSDiskImageSecurityProfileResponsePtrOutput) ToOSDiskImageSecurityProfileResponsePtrOutputWithContext(ctx context.Context) OSDiskImageSecurityProfileResponsePtrOutput {
	return o
}

func (o OSDiskImageSecurityProfileResponsePtrOutput) Elem() OSDiskImageSecurityProfileResponseOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfileResponse) OSDiskImageSecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskImageSecurityProfileResponse
		return ret
	}).(OSDiskImageSecurityProfileResponseOutput)
}

func (o OSDiskImageSecurityProfileResponsePtrOutput) ConfidentialVMEncryptionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfidentialVMEncryptionType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskImageSecurityProfileResponsePtrOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskImageSecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecureVMDiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

type RecommendedMachineConfiguration struct {
	Memory *ResourceRange `pulumi:"memory"`
	VCPUs  *ResourceRange `pulumi:"vCPUs"`
}





type RecommendedMachineConfigurationInput interface {
	pulumi.Input

	ToRecommendedMachineConfigurationOutput() RecommendedMachineConfigurationOutput
	ToRecommendedMachineConfigurationOutputWithContext(context.Context) RecommendedMachineConfigurationOutput
}

type RecommendedMachineConfigurationArgs struct {
	Memory ResourceRangePtrInput `pulumi:"memory"`
	VCPUs  ResourceRangePtrInput `pulumi:"vCPUs"`
}

func (RecommendedMachineConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedMachineConfiguration)(nil)).Elem()
}

func (i RecommendedMachineConfigurationArgs) ToRecommendedMachineConfigurationOutput() RecommendedMachineConfigurationOutput {
	return i.ToRecommendedMachineConfigurationOutputWithContext(context.Background())
}

func (i RecommendedMachineConfigurationArgs) ToRecommendedMachineConfigurationOutputWithContext(ctx context.Context) RecommendedMachineConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedMachineConfigurationOutput)
}

func (i RecommendedMachineConfigurationArgs) ToRecommendedMachineConfigurationPtrOutput() RecommendedMachineConfigurationPtrOutput {
	return i.ToRecommendedMachineConfigurationPtrOutputWithContext(context.Background())
}

func (i RecommendedMachineConfigurationArgs) ToRecommendedMachineConfigurationPtrOutputWithContext(ctx context.Context) RecommendedMachineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedMachineConfigurationOutput).ToRecommendedMachineConfigurationPtrOutputWithContext(ctx)
}









type RecommendedMachineConfigurationPtrInput interface {
	pulumi.Input

	ToRecommendedMachineConfigurationPtrOutput() RecommendedMachineConfigurationPtrOutput
	ToRecommendedMachineConfigurationPtrOutputWithContext(context.Context) RecommendedMachineConfigurationPtrOutput
}

type recommendedMachineConfigurationPtrType RecommendedMachineConfigurationArgs

func RecommendedMachineConfigurationPtr(v *RecommendedMachineConfigurationArgs) RecommendedMachineConfigurationPtrInput {
	return (*recommendedMachineConfigurationPtrType)(v)
}

func (*recommendedMachineConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecommendedMachineConfiguration)(nil)).Elem()
}

func (i *recommendedMachineConfigurationPtrType) ToRecommendedMachineConfigurationPtrOutput() RecommendedMachineConfigurationPtrOutput {
	return i.ToRecommendedMachineConfigurationPtrOutputWithContext(context.Background())
}

func (i *recommendedMachineConfigurationPtrType) ToRecommendedMachineConfigurationPtrOutputWithContext(ctx context.Context) RecommendedMachineConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecommendedMachineConfigurationPtrOutput)
}

type RecommendedMachineConfigurationOutput struct{ *pulumi.OutputState }

func (RecommendedMachineConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedMachineConfiguration)(nil)).Elem()
}

func (o RecommendedMachineConfigurationOutput) ToRecommendedMachineConfigurationOutput() RecommendedMachineConfigurationOutput {
	return o
}

func (o RecommendedMachineConfigurationOutput) ToRecommendedMachineConfigurationOutputWithContext(ctx context.Context) RecommendedMachineConfigurationOutput {
	return o
}

func (o RecommendedMachineConfigurationOutput) ToRecommendedMachineConfigurationPtrOutput() RecommendedMachineConfigurationPtrOutput {
	return o.ToRecommendedMachineConfigurationPtrOutputWithContext(context.Background())
}

func (o RecommendedMachineConfigurationOutput) ToRecommendedMachineConfigurationPtrOutputWithContext(ctx context.Context) RecommendedMachineConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecommendedMachineConfiguration) *RecommendedMachineConfiguration {
		return &v
	}).(RecommendedMachineConfigurationPtrOutput)
}

func (o RecommendedMachineConfigurationOutput) Memory() ResourceRangePtrOutput {
	return o.ApplyT(func(v RecommendedMachineConfiguration) *ResourceRange { return v.Memory }).(ResourceRangePtrOutput)
}

func (o RecommendedMachineConfigurationOutput) VCPUs() ResourceRangePtrOutput {
	return o.ApplyT(func(v RecommendedMachineConfiguration) *ResourceRange { return v.VCPUs }).(ResourceRangePtrOutput)
}

type RecommendedMachineConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RecommendedMachineConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecommendedMachineConfiguration)(nil)).Elem()
}

func (o RecommendedMachineConfigurationPtrOutput) ToRecommendedMachineConfigurationPtrOutput() RecommendedMachineConfigurationPtrOutput {
	return o
}

func (o RecommendedMachineConfigurationPtrOutput) ToRecommendedMachineConfigurationPtrOutputWithContext(ctx context.Context) RecommendedMachineConfigurationPtrOutput {
	return o
}

func (o RecommendedMachineConfigurationPtrOutput) Elem() RecommendedMachineConfigurationOutput {
	return o.ApplyT(func(v *RecommendedMachineConfiguration) RecommendedMachineConfiguration {
		if v != nil {
			return *v
		}
		var ret RecommendedMachineConfiguration
		return ret
	}).(RecommendedMachineConfigurationOutput)
}

func (o RecommendedMachineConfigurationPtrOutput) Memory() ResourceRangePtrOutput {
	return o.ApplyT(func(v *RecommendedMachineConfiguration) *ResourceRange {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(ResourceRangePtrOutput)
}

func (o RecommendedMachineConfigurationPtrOutput) VCPUs() ResourceRangePtrOutput {
	return o.ApplyT(func(v *RecommendedMachineConfiguration) *ResourceRange {
		if v == nil {
			return nil
		}
		return v.VCPUs
	}).(ResourceRangePtrOutput)
}

type RecommendedMachineConfigurationResponse struct {
	Memory *ResourceRangeResponse `pulumi:"memory"`
	VCPUs  *ResourceRangeResponse `pulumi:"vCPUs"`
}

type RecommendedMachineConfigurationResponseOutput struct{ *pulumi.OutputState }

func (RecommendedMachineConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendedMachineConfigurationResponse)(nil)).Elem()
}

func (o RecommendedMachineConfigurationResponseOutput) ToRecommendedMachineConfigurationResponseOutput() RecommendedMachineConfigurationResponseOutput {
	return o
}

func (o RecommendedMachineConfigurationResponseOutput) ToRecommendedMachineConfigurationResponseOutputWithContext(ctx context.Context) RecommendedMachineConfigurationResponseOutput {
	return o
}

func (o RecommendedMachineConfigurationResponseOutput) Memory() ResourceRangeResponsePtrOutput {
	return o.ApplyT(func(v RecommendedMachineConfigurationResponse) *ResourceRangeResponse { return v.Memory }).(ResourceRangeResponsePtrOutput)
}

func (o RecommendedMachineConfigurationResponseOutput) VCPUs() ResourceRangeResponsePtrOutput {
	return o.ApplyT(func(v RecommendedMachineConfigurationResponse) *ResourceRangeResponse { return v.VCPUs }).(ResourceRangeResponsePtrOutput)
}

type RecommendedMachineConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (RecommendedMachineConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecommendedMachineConfigurationResponse)(nil)).Elem()
}

func (o RecommendedMachineConfigurationResponsePtrOutput) ToRecommendedMachineConfigurationResponsePtrOutput() RecommendedMachineConfigurationResponsePtrOutput {
	return o
}

func (o RecommendedMachineConfigurationResponsePtrOutput) ToRecommendedMachineConfigurationResponsePtrOutputWithContext(ctx context.Context) RecommendedMachineConfigurationResponsePtrOutput {
	return o
}

func (o RecommendedMachineConfigurationResponsePtrOutput) Elem() RecommendedMachineConfigurationResponseOutput {
	return o.ApplyT(func(v *RecommendedMachineConfigurationResponse) RecommendedMachineConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret RecommendedMachineConfigurationResponse
		return ret
	}).(RecommendedMachineConfigurationResponseOutput)
}

func (o RecommendedMachineConfigurationResponsePtrOutput) Memory() ResourceRangeResponsePtrOutput {
	return o.ApplyT(func(v *RecommendedMachineConfigurationResponse) *ResourceRangeResponse {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(ResourceRangeResponsePtrOutput)
}

func (o RecommendedMachineConfigurationResponsePtrOutput) VCPUs() ResourceRangeResponsePtrOutput {
	return o.ApplyT(func(v *RecommendedMachineConfigurationResponse) *ResourceRangeResponse {
		if v == nil {
			return nil
		}
		return v.VCPUs
	}).(ResourceRangeResponsePtrOutput)
}

type RegionalReplicationStatusResponse struct {
	Details  string `pulumi:"details"`
	Progress int    `pulumi:"progress"`
	Region   string `pulumi:"region"`
	State    string `pulumi:"state"`
}

type RegionalReplicationStatusResponseOutput struct{ *pulumi.OutputState }

func (RegionalReplicationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionalReplicationStatusResponse)(nil)).Elem()
}

func (o RegionalReplicationStatusResponseOutput) ToRegionalReplicationStatusResponseOutput() RegionalReplicationStatusResponseOutput {
	return o
}

func (o RegionalReplicationStatusResponseOutput) ToRegionalReplicationStatusResponseOutputWithContext(ctx context.Context) RegionalReplicationStatusResponseOutput {
	return o
}

func (o RegionalReplicationStatusResponseOutput) Details() pulumi.StringOutput {
	return o.ApplyT(func(v RegionalReplicationStatusResponse) string { return v.Details }).(pulumi.StringOutput)
}

func (o RegionalReplicationStatusResponseOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v RegionalReplicationStatusResponse) int { return v.Progress }).(pulumi.IntOutput)
}

func (o RegionalReplicationStatusResponseOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v RegionalReplicationStatusResponse) string { return v.Region }).(pulumi.StringOutput)
}

func (o RegionalReplicationStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v RegionalReplicationStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type RegionalReplicationStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (RegionalReplicationStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionalReplicationStatusResponse)(nil)).Elem()
}

func (o RegionalReplicationStatusResponseArrayOutput) ToRegionalReplicationStatusResponseArrayOutput() RegionalReplicationStatusResponseArrayOutput {
	return o
}

func (o RegionalReplicationStatusResponseArrayOutput) ToRegionalReplicationStatusResponseArrayOutputWithContext(ctx context.Context) RegionalReplicationStatusResponseArrayOutput {
	return o
}

func (o RegionalReplicationStatusResponseArrayOutput) Index(i pulumi.IntInput) RegionalReplicationStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegionalReplicationStatusResponse {
		return vs[0].([]RegionalReplicationStatusResponse)[vs[1].(int)]
	}).(RegionalReplicationStatusResponseOutput)
}

type RegionalSharingStatusResponse struct {
	Details *string `pulumi:"details"`
	Region  *string `pulumi:"region"`
	State   string  `pulumi:"state"`
}

type RegionalSharingStatusResponseOutput struct{ *pulumi.OutputState }

func (RegionalSharingStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegionalSharingStatusResponse)(nil)).Elem()
}

func (o RegionalSharingStatusResponseOutput) ToRegionalSharingStatusResponseOutput() RegionalSharingStatusResponseOutput {
	return o
}

func (o RegionalSharingStatusResponseOutput) ToRegionalSharingStatusResponseOutputWithContext(ctx context.Context) RegionalSharingStatusResponseOutput {
	return o
}

func (o RegionalSharingStatusResponseOutput) Details() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegionalSharingStatusResponse) *string { return v.Details }).(pulumi.StringPtrOutput)
}

func (o RegionalSharingStatusResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegionalSharingStatusResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o RegionalSharingStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v RegionalSharingStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type RegionalSharingStatusResponseArrayOutput struct{ *pulumi.OutputState }

func (RegionalSharingStatusResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegionalSharingStatusResponse)(nil)).Elem()
}

func (o RegionalSharingStatusResponseArrayOutput) ToRegionalSharingStatusResponseArrayOutput() RegionalSharingStatusResponseArrayOutput {
	return o
}

func (o RegionalSharingStatusResponseArrayOutput) ToRegionalSharingStatusResponseArrayOutputWithContext(ctx context.Context) RegionalSharingStatusResponseArrayOutput {
	return o
}

func (o RegionalSharingStatusResponseArrayOutput) Index(i pulumi.IntInput) RegionalSharingStatusResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegionalSharingStatusResponse {
		return vs[0].([]RegionalSharingStatusResponse)[vs[1].(int)]
	}).(RegionalSharingStatusResponseOutput)
}

type ReplicationStatusResponse struct {
	AggregatedState string                              `pulumi:"aggregatedState"`
	Summary         []RegionalReplicationStatusResponse `pulumi:"summary"`
}

type ReplicationStatusResponseOutput struct{ *pulumi.OutputState }

func (ReplicationStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationStatusResponse)(nil)).Elem()
}

func (o ReplicationStatusResponseOutput) ToReplicationStatusResponseOutput() ReplicationStatusResponseOutput {
	return o
}

func (o ReplicationStatusResponseOutput) ToReplicationStatusResponseOutputWithContext(ctx context.Context) ReplicationStatusResponseOutput {
	return o
}

func (o ReplicationStatusResponseOutput) AggregatedState() pulumi.StringOutput {
	return o.ApplyT(func(v ReplicationStatusResponse) string { return v.AggregatedState }).(pulumi.StringOutput)
}

func (o ReplicationStatusResponseOutput) Summary() RegionalReplicationStatusResponseArrayOutput {
	return o.ApplyT(func(v ReplicationStatusResponse) []RegionalReplicationStatusResponse { return v.Summary }).(RegionalReplicationStatusResponseArrayOutput)
}

type ResourceRange struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}





type ResourceRangeInput interface {
	pulumi.Input

	ToResourceRangeOutput() ResourceRangeOutput
	ToResourceRangeOutputWithContext(context.Context) ResourceRangeOutput
}

type ResourceRangeArgs struct {
	Max pulumi.IntPtrInput `pulumi:"max"`
	Min pulumi.IntPtrInput `pulumi:"min"`
}

func (ResourceRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRange)(nil)).Elem()
}

func (i ResourceRangeArgs) ToResourceRangeOutput() ResourceRangeOutput {
	return i.ToResourceRangeOutputWithContext(context.Background())
}

func (i ResourceRangeArgs) ToResourceRangeOutputWithContext(ctx context.Context) ResourceRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRangeOutput)
}

func (i ResourceRangeArgs) ToResourceRangePtrOutput() ResourceRangePtrOutput {
	return i.ToResourceRangePtrOutputWithContext(context.Background())
}

func (i ResourceRangeArgs) ToResourceRangePtrOutputWithContext(ctx context.Context) ResourceRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRangeOutput).ToResourceRangePtrOutputWithContext(ctx)
}









type ResourceRangePtrInput interface {
	pulumi.Input

	ToResourceRangePtrOutput() ResourceRangePtrOutput
	ToResourceRangePtrOutputWithContext(context.Context) ResourceRangePtrOutput
}

type resourceRangePtrType ResourceRangeArgs

func ResourceRangePtr(v *ResourceRangeArgs) ResourceRangePtrInput {
	return (*resourceRangePtrType)(v)
}

func (*resourceRangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRange)(nil)).Elem()
}

func (i *resourceRangePtrType) ToResourceRangePtrOutput() ResourceRangePtrOutput {
	return i.ToResourceRangePtrOutputWithContext(context.Background())
}

func (i *resourceRangePtrType) ToResourceRangePtrOutputWithContext(ctx context.Context) ResourceRangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRangePtrOutput)
}

type ResourceRangeOutput struct{ *pulumi.OutputState }

func (ResourceRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRange)(nil)).Elem()
}

func (o ResourceRangeOutput) ToResourceRangeOutput() ResourceRangeOutput {
	return o
}

func (o ResourceRangeOutput) ToResourceRangeOutputWithContext(ctx context.Context) ResourceRangeOutput {
	return o
}

func (o ResourceRangeOutput) ToResourceRangePtrOutput() ResourceRangePtrOutput {
	return o.ToResourceRangePtrOutputWithContext(context.Background())
}

func (o ResourceRangeOutput) ToResourceRangePtrOutputWithContext(ctx context.Context) ResourceRangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceRange) *ResourceRange {
		return &v
	}).(ResourceRangePtrOutput)
}

func (o ResourceRangeOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceRange) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o ResourceRangeOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceRange) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type ResourceRangePtrOutput struct{ *pulumi.OutputState }

func (ResourceRangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRange)(nil)).Elem()
}

func (o ResourceRangePtrOutput) ToResourceRangePtrOutput() ResourceRangePtrOutput {
	return o
}

func (o ResourceRangePtrOutput) ToResourceRangePtrOutputWithContext(ctx context.Context) ResourceRangePtrOutput {
	return o
}

func (o ResourceRangePtrOutput) Elem() ResourceRangeOutput {
	return o.ApplyT(func(v *ResourceRange) ResourceRange {
		if v != nil {
			return *v
		}
		var ret ResourceRange
		return ret
	}).(ResourceRangeOutput)
}

func (o ResourceRangePtrOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceRange) *int {
		if v == nil {
			return nil
		}
		return v.Max
	}).(pulumi.IntPtrOutput)
}

func (o ResourceRangePtrOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceRange) *int {
		if v == nil {
			return nil
		}
		return v.Min
	}).(pulumi.IntPtrOutput)
}

type ResourceRangeResponse struct {
	Max *int `pulumi:"max"`
	Min *int `pulumi:"min"`
}

type ResourceRangeResponseOutput struct{ *pulumi.OutputState }

func (ResourceRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRangeResponse)(nil)).Elem()
}

func (o ResourceRangeResponseOutput) ToResourceRangeResponseOutput() ResourceRangeResponseOutput {
	return o
}

func (o ResourceRangeResponseOutput) ToResourceRangeResponseOutputWithContext(ctx context.Context) ResourceRangeResponseOutput {
	return o
}

func (o ResourceRangeResponseOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceRangeResponse) *int { return v.Max }).(pulumi.IntPtrOutput)
}

func (o ResourceRangeResponseOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceRangeResponse) *int { return v.Min }).(pulumi.IntPtrOutput)
}

type ResourceRangeResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceRangeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceRangeResponse)(nil)).Elem()
}

func (o ResourceRangeResponsePtrOutput) ToResourceRangeResponsePtrOutput() ResourceRangeResponsePtrOutput {
	return o
}

func (o ResourceRangeResponsePtrOutput) ToResourceRangeResponsePtrOutputWithContext(ctx context.Context) ResourceRangeResponsePtrOutput {
	return o
}

func (o ResourceRangeResponsePtrOutput) Elem() ResourceRangeResponseOutput {
	return o.ApplyT(func(v *ResourceRangeResponse) ResourceRangeResponse {
		if v != nil {
			return *v
		}
		var ret ResourceRangeResponse
		return ret
	}).(ResourceRangeResponseOutput)
}

func (o ResourceRangeResponsePtrOutput) Max() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceRangeResponse) *int {
		if v == nil {
			return nil
		}
		return v.Max
	}).(pulumi.IntPtrOutput)
}

func (o ResourceRangeResponsePtrOutput) Min() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceRangeResponse) *int {
		if v == nil {
			return nil
		}
		return v.Min
	}).(pulumi.IntPtrOutput)
}

type SharingProfile struct {
	CommunityGalleryInfo *CommunityGalleryInfo `pulumi:"communityGalleryInfo"`
	Permissions          *string               `pulumi:"permissions"`
}





type SharingProfileInput interface {
	pulumi.Input

	ToSharingProfileOutput() SharingProfileOutput
	ToSharingProfileOutputWithContext(context.Context) SharingProfileOutput
}

type SharingProfileArgs struct {
	CommunityGalleryInfo CommunityGalleryInfoPtrInput `pulumi:"communityGalleryInfo"`
	Permissions          pulumi.StringPtrInput        `pulumi:"permissions"`
}

func (SharingProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharingProfile)(nil)).Elem()
}

func (i SharingProfileArgs) ToSharingProfileOutput() SharingProfileOutput {
	return i.ToSharingProfileOutputWithContext(context.Background())
}

func (i SharingProfileArgs) ToSharingProfileOutputWithContext(ctx context.Context) SharingProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharingProfileOutput)
}

func (i SharingProfileArgs) ToSharingProfilePtrOutput() SharingProfilePtrOutput {
	return i.ToSharingProfilePtrOutputWithContext(context.Background())
}

func (i SharingProfileArgs) ToSharingProfilePtrOutputWithContext(ctx context.Context) SharingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharingProfileOutput).ToSharingProfilePtrOutputWithContext(ctx)
}









type SharingProfilePtrInput interface {
	pulumi.Input

	ToSharingProfilePtrOutput() SharingProfilePtrOutput
	ToSharingProfilePtrOutputWithContext(context.Context) SharingProfilePtrOutput
}

type sharingProfilePtrType SharingProfileArgs

func SharingProfilePtr(v *SharingProfileArgs) SharingProfilePtrInput {
	return (*sharingProfilePtrType)(v)
}

func (*sharingProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SharingProfile)(nil)).Elem()
}

func (i *sharingProfilePtrType) ToSharingProfilePtrOutput() SharingProfilePtrOutput {
	return i.ToSharingProfilePtrOutputWithContext(context.Background())
}

func (i *sharingProfilePtrType) ToSharingProfilePtrOutputWithContext(ctx context.Context) SharingProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharingProfilePtrOutput)
}

type SharingProfileOutput struct{ *pulumi.OutputState }

func (SharingProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharingProfile)(nil)).Elem()
}

func (o SharingProfileOutput) ToSharingProfileOutput() SharingProfileOutput {
	return o
}

func (o SharingProfileOutput) ToSharingProfileOutputWithContext(ctx context.Context) SharingProfileOutput {
	return o
}

func (o SharingProfileOutput) ToSharingProfilePtrOutput() SharingProfilePtrOutput {
	return o.ToSharingProfilePtrOutputWithContext(context.Background())
}

func (o SharingProfileOutput) ToSharingProfilePtrOutputWithContext(ctx context.Context) SharingProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SharingProfile) *SharingProfile {
		return &v
	}).(SharingProfilePtrOutput)
}

func (o SharingProfileOutput) CommunityGalleryInfo() CommunityGalleryInfoPtrOutput {
	return o.ApplyT(func(v SharingProfile) *CommunityGalleryInfo { return v.CommunityGalleryInfo }).(CommunityGalleryInfoPtrOutput)
}

func (o SharingProfileOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharingProfile) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

type SharingProfilePtrOutput struct{ *pulumi.OutputState }

func (SharingProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharingProfile)(nil)).Elem()
}

func (o SharingProfilePtrOutput) ToSharingProfilePtrOutput() SharingProfilePtrOutput {
	return o
}

func (o SharingProfilePtrOutput) ToSharingProfilePtrOutputWithContext(ctx context.Context) SharingProfilePtrOutput {
	return o
}

func (o SharingProfilePtrOutput) Elem() SharingProfileOutput {
	return o.ApplyT(func(v *SharingProfile) SharingProfile {
		if v != nil {
			return *v
		}
		var ret SharingProfile
		return ret
	}).(SharingProfileOutput)
}

func (o SharingProfilePtrOutput) CommunityGalleryInfo() CommunityGalleryInfoPtrOutput {
	return o.ApplyT(func(v *SharingProfile) *CommunityGalleryInfo {
		if v == nil {
			return nil
		}
		return v.CommunityGalleryInfo
	}).(CommunityGalleryInfoPtrOutput)
}

func (o SharingProfilePtrOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharingProfile) *string {
		if v == nil {
			return nil
		}
		return v.Permissions
	}).(pulumi.StringPtrOutput)
}

type SharingProfileGroupResponse struct {
	Ids  []string `pulumi:"ids"`
	Type *string  `pulumi:"type"`
}

type SharingProfileGroupResponseOutput struct{ *pulumi.OutputState }

func (SharingProfileGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharingProfileGroupResponse)(nil)).Elem()
}

func (o SharingProfileGroupResponseOutput) ToSharingProfileGroupResponseOutput() SharingProfileGroupResponseOutput {
	return o
}

func (o SharingProfileGroupResponseOutput) ToSharingProfileGroupResponseOutputWithContext(ctx context.Context) SharingProfileGroupResponseOutput {
	return o
}

func (o SharingProfileGroupResponseOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SharingProfileGroupResponse) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o SharingProfileGroupResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharingProfileGroupResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type SharingProfileGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (SharingProfileGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharingProfileGroupResponse)(nil)).Elem()
}

func (o SharingProfileGroupResponseArrayOutput) ToSharingProfileGroupResponseArrayOutput() SharingProfileGroupResponseArrayOutput {
	return o
}

func (o SharingProfileGroupResponseArrayOutput) ToSharingProfileGroupResponseArrayOutputWithContext(ctx context.Context) SharingProfileGroupResponseArrayOutput {
	return o
}

func (o SharingProfileGroupResponseArrayOutput) Index(i pulumi.IntInput) SharingProfileGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharingProfileGroupResponse {
		return vs[0].([]SharingProfileGroupResponse)[vs[1].(int)]
	}).(SharingProfileGroupResponseOutput)
}

type SharingProfileResponse struct {
	CommunityGalleryInfo *CommunityGalleryInfoResponse `pulumi:"communityGalleryInfo"`
	Groups               []SharingProfileGroupResponse `pulumi:"groups"`
	Permissions          *string                       `pulumi:"permissions"`
}

type SharingProfileResponseOutput struct{ *pulumi.OutputState }

func (SharingProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharingProfileResponse)(nil)).Elem()
}

func (o SharingProfileResponseOutput) ToSharingProfileResponseOutput() SharingProfileResponseOutput {
	return o
}

func (o SharingProfileResponseOutput) ToSharingProfileResponseOutputWithContext(ctx context.Context) SharingProfileResponseOutput {
	return o
}

func (o SharingProfileResponseOutput) CommunityGalleryInfo() CommunityGalleryInfoResponsePtrOutput {
	return o.ApplyT(func(v SharingProfileResponse) *CommunityGalleryInfoResponse { return v.CommunityGalleryInfo }).(CommunityGalleryInfoResponsePtrOutput)
}

func (o SharingProfileResponseOutput) Groups() SharingProfileGroupResponseArrayOutput {
	return o.ApplyT(func(v SharingProfileResponse) []SharingProfileGroupResponse { return v.Groups }).(SharingProfileGroupResponseArrayOutput)
}

func (o SharingProfileResponseOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharingProfileResponse) *string { return v.Permissions }).(pulumi.StringPtrOutput)
}

type SharingProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (SharingProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SharingProfileResponse)(nil)).Elem()
}

func (o SharingProfileResponsePtrOutput) ToSharingProfileResponsePtrOutput() SharingProfileResponsePtrOutput {
	return o
}

func (o SharingProfileResponsePtrOutput) ToSharingProfileResponsePtrOutputWithContext(ctx context.Context) SharingProfileResponsePtrOutput {
	return o
}

func (o SharingProfileResponsePtrOutput) Elem() SharingProfileResponseOutput {
	return o.ApplyT(func(v *SharingProfileResponse) SharingProfileResponse {
		if v != nil {
			return *v
		}
		var ret SharingProfileResponse
		return ret
	}).(SharingProfileResponseOutput)
}

func (o SharingProfileResponsePtrOutput) CommunityGalleryInfo() CommunityGalleryInfoResponsePtrOutput {
	return o.ApplyT(func(v *SharingProfileResponse) *CommunityGalleryInfoResponse {
		if v == nil {
			return nil
		}
		return v.CommunityGalleryInfo
	}).(CommunityGalleryInfoResponsePtrOutput)
}

func (o SharingProfileResponsePtrOutput) Groups() SharingProfileGroupResponseArrayOutput {
	return o.ApplyT(func(v *SharingProfileResponse) []SharingProfileGroupResponse {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(SharingProfileGroupResponseArrayOutput)
}

func (o SharingProfileResponsePtrOutput) Permissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SharingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Permissions
	}).(pulumi.StringPtrOutput)
}

type SharingStatusResponse struct {
	AggregatedState string                          `pulumi:"aggregatedState"`
	Summary         []RegionalSharingStatusResponse `pulumi:"summary"`
}

type SharingStatusResponseOutput struct{ *pulumi.OutputState }

func (SharingStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharingStatusResponse)(nil)).Elem()
}

func (o SharingStatusResponseOutput) ToSharingStatusResponseOutput() SharingStatusResponseOutput {
	return o
}

func (o SharingStatusResponseOutput) ToSharingStatusResponseOutputWithContext(ctx context.Context) SharingStatusResponseOutput {
	return o
}

func (o SharingStatusResponseOutput) AggregatedState() pulumi.StringOutput {
	return o.ApplyT(func(v SharingStatusResponse) string { return v.AggregatedState }).(pulumi.StringOutput)
}

func (o SharingStatusResponseOutput) Summary() RegionalSharingStatusResponseArrayOutput {
	return o.ApplyT(func(v SharingStatusResponse) []RegionalSharingStatusResponse { return v.Summary }).(RegionalSharingStatusResponseArrayOutput)
}

type SoftDeletePolicy struct {
	IsSoftDeleteEnabled *bool `pulumi:"isSoftDeleteEnabled"`
}





type SoftDeletePolicyInput interface {
	pulumi.Input

	ToSoftDeletePolicyOutput() SoftDeletePolicyOutput
	ToSoftDeletePolicyOutputWithContext(context.Context) SoftDeletePolicyOutput
}

type SoftDeletePolicyArgs struct {
	IsSoftDeleteEnabled pulumi.BoolPtrInput `pulumi:"isSoftDeleteEnabled"`
}

func (SoftDeletePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftDeletePolicy)(nil)).Elem()
}

func (i SoftDeletePolicyArgs) ToSoftDeletePolicyOutput() SoftDeletePolicyOutput {
	return i.ToSoftDeletePolicyOutputWithContext(context.Background())
}

func (i SoftDeletePolicyArgs) ToSoftDeletePolicyOutputWithContext(ctx context.Context) SoftDeletePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftDeletePolicyOutput)
}

func (i SoftDeletePolicyArgs) ToSoftDeletePolicyPtrOutput() SoftDeletePolicyPtrOutput {
	return i.ToSoftDeletePolicyPtrOutputWithContext(context.Background())
}

func (i SoftDeletePolicyArgs) ToSoftDeletePolicyPtrOutputWithContext(ctx context.Context) SoftDeletePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftDeletePolicyOutput).ToSoftDeletePolicyPtrOutputWithContext(ctx)
}









type SoftDeletePolicyPtrInput interface {
	pulumi.Input

	ToSoftDeletePolicyPtrOutput() SoftDeletePolicyPtrOutput
	ToSoftDeletePolicyPtrOutputWithContext(context.Context) SoftDeletePolicyPtrOutput
}

type softDeletePolicyPtrType SoftDeletePolicyArgs

func SoftDeletePolicyPtr(v *SoftDeletePolicyArgs) SoftDeletePolicyPtrInput {
	return (*softDeletePolicyPtrType)(v)
}

func (*softDeletePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftDeletePolicy)(nil)).Elem()
}

func (i *softDeletePolicyPtrType) ToSoftDeletePolicyPtrOutput() SoftDeletePolicyPtrOutput {
	return i.ToSoftDeletePolicyPtrOutputWithContext(context.Background())
}

func (i *softDeletePolicyPtrType) ToSoftDeletePolicyPtrOutputWithContext(ctx context.Context) SoftDeletePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftDeletePolicyPtrOutput)
}

type SoftDeletePolicyOutput struct{ *pulumi.OutputState }

func (SoftDeletePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftDeletePolicy)(nil)).Elem()
}

func (o SoftDeletePolicyOutput) ToSoftDeletePolicyOutput() SoftDeletePolicyOutput {
	return o
}

func (o SoftDeletePolicyOutput) ToSoftDeletePolicyOutputWithContext(ctx context.Context) SoftDeletePolicyOutput {
	return o
}

func (o SoftDeletePolicyOutput) ToSoftDeletePolicyPtrOutput() SoftDeletePolicyPtrOutput {
	return o.ToSoftDeletePolicyPtrOutputWithContext(context.Background())
}

func (o SoftDeletePolicyOutput) ToSoftDeletePolicyPtrOutputWithContext(ctx context.Context) SoftDeletePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoftDeletePolicy) *SoftDeletePolicy {
		return &v
	}).(SoftDeletePolicyPtrOutput)
}

func (o SoftDeletePolicyOutput) IsSoftDeleteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SoftDeletePolicy) *bool { return v.IsSoftDeleteEnabled }).(pulumi.BoolPtrOutput)
}

type SoftDeletePolicyPtrOutput struct{ *pulumi.OutputState }

func (SoftDeletePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftDeletePolicy)(nil)).Elem()
}

func (o SoftDeletePolicyPtrOutput) ToSoftDeletePolicyPtrOutput() SoftDeletePolicyPtrOutput {
	return o
}

func (o SoftDeletePolicyPtrOutput) ToSoftDeletePolicyPtrOutputWithContext(ctx context.Context) SoftDeletePolicyPtrOutput {
	return o
}

func (o SoftDeletePolicyPtrOutput) Elem() SoftDeletePolicyOutput {
	return o.ApplyT(func(v *SoftDeletePolicy) SoftDeletePolicy {
		if v != nil {
			return *v
		}
		var ret SoftDeletePolicy
		return ret
	}).(SoftDeletePolicyOutput)
}

func (o SoftDeletePolicyPtrOutput) IsSoftDeleteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SoftDeletePolicy) *bool {
		if v == nil {
			return nil
		}
		return v.IsSoftDeleteEnabled
	}).(pulumi.BoolPtrOutput)
}

type SoftDeletePolicyResponse struct {
	IsSoftDeleteEnabled *bool `pulumi:"isSoftDeleteEnabled"`
}

type SoftDeletePolicyResponseOutput struct{ *pulumi.OutputState }

func (SoftDeletePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftDeletePolicyResponse)(nil)).Elem()
}

func (o SoftDeletePolicyResponseOutput) ToSoftDeletePolicyResponseOutput() SoftDeletePolicyResponseOutput {
	return o
}

func (o SoftDeletePolicyResponseOutput) ToSoftDeletePolicyResponseOutputWithContext(ctx context.Context) SoftDeletePolicyResponseOutput {
	return o
}

func (o SoftDeletePolicyResponseOutput) IsSoftDeleteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SoftDeletePolicyResponse) *bool { return v.IsSoftDeleteEnabled }).(pulumi.BoolPtrOutput)
}

type SoftDeletePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (SoftDeletePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftDeletePolicyResponse)(nil)).Elem()
}

func (o SoftDeletePolicyResponsePtrOutput) ToSoftDeletePolicyResponsePtrOutput() SoftDeletePolicyResponsePtrOutput {
	return o
}

func (o SoftDeletePolicyResponsePtrOutput) ToSoftDeletePolicyResponsePtrOutputWithContext(ctx context.Context) SoftDeletePolicyResponsePtrOutput {
	return o
}

func (o SoftDeletePolicyResponsePtrOutput) Elem() SoftDeletePolicyResponseOutput {
	return o.ApplyT(func(v *SoftDeletePolicyResponse) SoftDeletePolicyResponse {
		if v != nil {
			return *v
		}
		var ret SoftDeletePolicyResponse
		return ret
	}).(SoftDeletePolicyResponseOutput)
}

func (o SoftDeletePolicyResponsePtrOutput) IsSoftDeleteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SoftDeletePolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSoftDeleteEnabled
	}).(pulumi.BoolPtrOutput)
}

type TargetRegion struct {
	Encryption           *EncryptionImages `pulumi:"encryption"`
	Name                 string            `pulumi:"name"`
	RegionalReplicaCount *int              `pulumi:"regionalReplicaCount"`
	StorageAccountType   *string           `pulumi:"storageAccountType"`
}





type TargetRegionInput interface {
	pulumi.Input

	ToTargetRegionOutput() TargetRegionOutput
	ToTargetRegionOutputWithContext(context.Context) TargetRegionOutput
}

type TargetRegionArgs struct {
	Encryption           EncryptionImagesPtrInput `pulumi:"encryption"`
	Name                 pulumi.StringInput       `pulumi:"name"`
	RegionalReplicaCount pulumi.IntPtrInput       `pulumi:"regionalReplicaCount"`
	StorageAccountType   pulumi.StringPtrInput    `pulumi:"storageAccountType"`
}

func (TargetRegionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetRegion)(nil)).Elem()
}

func (i TargetRegionArgs) ToTargetRegionOutput() TargetRegionOutput {
	return i.ToTargetRegionOutputWithContext(context.Background())
}

func (i TargetRegionArgs) ToTargetRegionOutputWithContext(ctx context.Context) TargetRegionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetRegionOutput)
}





type TargetRegionArrayInput interface {
	pulumi.Input

	ToTargetRegionArrayOutput() TargetRegionArrayOutput
	ToTargetRegionArrayOutputWithContext(context.Context) TargetRegionArrayOutput
}

type TargetRegionArray []TargetRegionInput

func (TargetRegionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetRegion)(nil)).Elem()
}

func (i TargetRegionArray) ToTargetRegionArrayOutput() TargetRegionArrayOutput {
	return i.ToTargetRegionArrayOutputWithContext(context.Background())
}

func (i TargetRegionArray) ToTargetRegionArrayOutputWithContext(ctx context.Context) TargetRegionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetRegionArrayOutput)
}

type TargetRegionOutput struct{ *pulumi.OutputState }

func (TargetRegionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetRegion)(nil)).Elem()
}

func (o TargetRegionOutput) ToTargetRegionOutput() TargetRegionOutput {
	return o
}

func (o TargetRegionOutput) ToTargetRegionOutputWithContext(ctx context.Context) TargetRegionOutput {
	return o
}

func (o TargetRegionOutput) Encryption() EncryptionImagesPtrOutput {
	return o.ApplyT(func(v TargetRegion) *EncryptionImages { return v.Encryption }).(EncryptionImagesPtrOutput)
}

func (o TargetRegionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetRegion) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetRegionOutput) RegionalReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetRegion) *int { return v.RegionalReplicaCount }).(pulumi.IntPtrOutput)
}

func (o TargetRegionOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetRegion) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type TargetRegionArrayOutput struct{ *pulumi.OutputState }

func (TargetRegionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetRegion)(nil)).Elem()
}

func (o TargetRegionArrayOutput) ToTargetRegionArrayOutput() TargetRegionArrayOutput {
	return o
}

func (o TargetRegionArrayOutput) ToTargetRegionArrayOutputWithContext(ctx context.Context) TargetRegionArrayOutput {
	return o
}

func (o TargetRegionArrayOutput) Index(i pulumi.IntInput) TargetRegionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetRegion {
		return vs[0].([]TargetRegion)[vs[1].(int)]
	}).(TargetRegionOutput)
}

type TargetRegionResponse struct {
	Encryption           *EncryptionImagesResponse `pulumi:"encryption"`
	Name                 string                    `pulumi:"name"`
	RegionalReplicaCount *int                      `pulumi:"regionalReplicaCount"`
	StorageAccountType   *string                   `pulumi:"storageAccountType"`
}

type TargetRegionResponseOutput struct{ *pulumi.OutputState }

func (TargetRegionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetRegionResponse)(nil)).Elem()
}

func (o TargetRegionResponseOutput) ToTargetRegionResponseOutput() TargetRegionResponseOutput {
	return o
}

func (o TargetRegionResponseOutput) ToTargetRegionResponseOutputWithContext(ctx context.Context) TargetRegionResponseOutput {
	return o
}

func (o TargetRegionResponseOutput) Encryption() EncryptionImagesResponsePtrOutput {
	return o.ApplyT(func(v TargetRegionResponse) *EncryptionImagesResponse { return v.Encryption }).(EncryptionImagesResponsePtrOutput)
}

func (o TargetRegionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetRegionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetRegionResponseOutput) RegionalReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TargetRegionResponse) *int { return v.RegionalReplicaCount }).(pulumi.IntPtrOutput)
}

func (o TargetRegionResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetRegionResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

type TargetRegionResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetRegionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetRegionResponse)(nil)).Elem()
}

func (o TargetRegionResponseArrayOutput) ToTargetRegionResponseArrayOutput() TargetRegionResponseArrayOutput {
	return o
}

func (o TargetRegionResponseArrayOutput) ToTargetRegionResponseArrayOutputWithContext(ctx context.Context) TargetRegionResponseArrayOutput {
	return o
}

func (o TargetRegionResponseArrayOutput) Index(i pulumi.IntInput) TargetRegionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetRegionResponse {
		return vs[0].([]TargetRegionResponse)[vs[1].(int)]
	}).(TargetRegionResponseOutput)
}

type UserArtifactManage struct {
	Install string  `pulumi:"install"`
	Remove  string  `pulumi:"remove"`
	Update  *string `pulumi:"update"`
}





type UserArtifactManageInput interface {
	pulumi.Input

	ToUserArtifactManageOutput() UserArtifactManageOutput
	ToUserArtifactManageOutputWithContext(context.Context) UserArtifactManageOutput
}

type UserArtifactManageArgs struct {
	Install pulumi.StringInput    `pulumi:"install"`
	Remove  pulumi.StringInput    `pulumi:"remove"`
	Update  pulumi.StringPtrInput `pulumi:"update"`
}

func (UserArtifactManageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactManage)(nil)).Elem()
}

func (i UserArtifactManageArgs) ToUserArtifactManageOutput() UserArtifactManageOutput {
	return i.ToUserArtifactManageOutputWithContext(context.Background())
}

func (i UserArtifactManageArgs) ToUserArtifactManageOutputWithContext(ctx context.Context) UserArtifactManageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactManageOutput)
}

func (i UserArtifactManageArgs) ToUserArtifactManagePtrOutput() UserArtifactManagePtrOutput {
	return i.ToUserArtifactManagePtrOutputWithContext(context.Background())
}

func (i UserArtifactManageArgs) ToUserArtifactManagePtrOutputWithContext(ctx context.Context) UserArtifactManagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactManageOutput).ToUserArtifactManagePtrOutputWithContext(ctx)
}









type UserArtifactManagePtrInput interface {
	pulumi.Input

	ToUserArtifactManagePtrOutput() UserArtifactManagePtrOutput
	ToUserArtifactManagePtrOutputWithContext(context.Context) UserArtifactManagePtrOutput
}

type userArtifactManagePtrType UserArtifactManageArgs

func UserArtifactManagePtr(v *UserArtifactManageArgs) UserArtifactManagePtrInput {
	return (*userArtifactManagePtrType)(v)
}

func (*userArtifactManagePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactManage)(nil)).Elem()
}

func (i *userArtifactManagePtrType) ToUserArtifactManagePtrOutput() UserArtifactManagePtrOutput {
	return i.ToUserArtifactManagePtrOutputWithContext(context.Background())
}

func (i *userArtifactManagePtrType) ToUserArtifactManagePtrOutputWithContext(ctx context.Context) UserArtifactManagePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactManagePtrOutput)
}

type UserArtifactManageOutput struct{ *pulumi.OutputState }

func (UserArtifactManageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactManage)(nil)).Elem()
}

func (o UserArtifactManageOutput) ToUserArtifactManageOutput() UserArtifactManageOutput {
	return o
}

func (o UserArtifactManageOutput) ToUserArtifactManageOutputWithContext(ctx context.Context) UserArtifactManageOutput {
	return o
}

func (o UserArtifactManageOutput) ToUserArtifactManagePtrOutput() UserArtifactManagePtrOutput {
	return o.ToUserArtifactManagePtrOutputWithContext(context.Background())
}

func (o UserArtifactManageOutput) ToUserArtifactManagePtrOutputWithContext(ctx context.Context) UserArtifactManagePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserArtifactManage) *UserArtifactManage {
		return &v
	}).(UserArtifactManagePtrOutput)
}

func (o UserArtifactManageOutput) Install() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactManage) string { return v.Install }).(pulumi.StringOutput)
}

func (o UserArtifactManageOutput) Remove() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactManage) string { return v.Remove }).(pulumi.StringOutput)
}

func (o UserArtifactManageOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactManage) *string { return v.Update }).(pulumi.StringPtrOutput)
}

type UserArtifactManagePtrOutput struct{ *pulumi.OutputState }

func (UserArtifactManagePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactManage)(nil)).Elem()
}

func (o UserArtifactManagePtrOutput) ToUserArtifactManagePtrOutput() UserArtifactManagePtrOutput {
	return o
}

func (o UserArtifactManagePtrOutput) ToUserArtifactManagePtrOutputWithContext(ctx context.Context) UserArtifactManagePtrOutput {
	return o
}

func (o UserArtifactManagePtrOutput) Elem() UserArtifactManageOutput {
	return o.ApplyT(func(v *UserArtifactManage) UserArtifactManage {
		if v != nil {
			return *v
		}
		var ret UserArtifactManage
		return ret
	}).(UserArtifactManageOutput)
}

func (o UserArtifactManagePtrOutput) Install() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManage) *string {
		if v == nil {
			return nil
		}
		return &v.Install
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactManagePtrOutput) Remove() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManage) *string {
		if v == nil {
			return nil
		}
		return &v.Remove
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactManagePtrOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManage) *string {
		if v == nil {
			return nil
		}
		return v.Update
	}).(pulumi.StringPtrOutput)
}

type UserArtifactManageResponse struct {
	Install string  `pulumi:"install"`
	Remove  string  `pulumi:"remove"`
	Update  *string `pulumi:"update"`
}

type UserArtifactManageResponseOutput struct{ *pulumi.OutputState }

func (UserArtifactManageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactManageResponse)(nil)).Elem()
}

func (o UserArtifactManageResponseOutput) ToUserArtifactManageResponseOutput() UserArtifactManageResponseOutput {
	return o
}

func (o UserArtifactManageResponseOutput) ToUserArtifactManageResponseOutputWithContext(ctx context.Context) UserArtifactManageResponseOutput {
	return o
}

func (o UserArtifactManageResponseOutput) Install() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactManageResponse) string { return v.Install }).(pulumi.StringOutput)
}

func (o UserArtifactManageResponseOutput) Remove() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactManageResponse) string { return v.Remove }).(pulumi.StringOutput)
}

func (o UserArtifactManageResponseOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactManageResponse) *string { return v.Update }).(pulumi.StringPtrOutput)
}

type UserArtifactManageResponsePtrOutput struct{ *pulumi.OutputState }

func (UserArtifactManageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactManageResponse)(nil)).Elem()
}

func (o UserArtifactManageResponsePtrOutput) ToUserArtifactManageResponsePtrOutput() UserArtifactManageResponsePtrOutput {
	return o
}

func (o UserArtifactManageResponsePtrOutput) ToUserArtifactManageResponsePtrOutputWithContext(ctx context.Context) UserArtifactManageResponsePtrOutput {
	return o
}

func (o UserArtifactManageResponsePtrOutput) Elem() UserArtifactManageResponseOutput {
	return o.ApplyT(func(v *UserArtifactManageResponse) UserArtifactManageResponse {
		if v != nil {
			return *v
		}
		var ret UserArtifactManageResponse
		return ret
	}).(UserArtifactManageResponseOutput)
}

func (o UserArtifactManageResponsePtrOutput) Install() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManageResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Install
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactManageResponsePtrOutput) Remove() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManageResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Remove
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactManageResponsePtrOutput) Update() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactManageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Update
	}).(pulumi.StringPtrOutput)
}

type UserArtifactSettings struct {
	ConfigFileName  *string `pulumi:"configFileName"`
	PackageFileName *string `pulumi:"packageFileName"`
}





type UserArtifactSettingsInput interface {
	pulumi.Input

	ToUserArtifactSettingsOutput() UserArtifactSettingsOutput
	ToUserArtifactSettingsOutputWithContext(context.Context) UserArtifactSettingsOutput
}

type UserArtifactSettingsArgs struct {
	ConfigFileName  pulumi.StringPtrInput `pulumi:"configFileName"`
	PackageFileName pulumi.StringPtrInput `pulumi:"packageFileName"`
}

func (UserArtifactSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSettings)(nil)).Elem()
}

func (i UserArtifactSettingsArgs) ToUserArtifactSettingsOutput() UserArtifactSettingsOutput {
	return i.ToUserArtifactSettingsOutputWithContext(context.Background())
}

func (i UserArtifactSettingsArgs) ToUserArtifactSettingsOutputWithContext(ctx context.Context) UserArtifactSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactSettingsOutput)
}

func (i UserArtifactSettingsArgs) ToUserArtifactSettingsPtrOutput() UserArtifactSettingsPtrOutput {
	return i.ToUserArtifactSettingsPtrOutputWithContext(context.Background())
}

func (i UserArtifactSettingsArgs) ToUserArtifactSettingsPtrOutputWithContext(ctx context.Context) UserArtifactSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactSettingsOutput).ToUserArtifactSettingsPtrOutputWithContext(ctx)
}









type UserArtifactSettingsPtrInput interface {
	pulumi.Input

	ToUserArtifactSettingsPtrOutput() UserArtifactSettingsPtrOutput
	ToUserArtifactSettingsPtrOutputWithContext(context.Context) UserArtifactSettingsPtrOutput
}

type userArtifactSettingsPtrType UserArtifactSettingsArgs

func UserArtifactSettingsPtr(v *UserArtifactSettingsArgs) UserArtifactSettingsPtrInput {
	return (*userArtifactSettingsPtrType)(v)
}

func (*userArtifactSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactSettings)(nil)).Elem()
}

func (i *userArtifactSettingsPtrType) ToUserArtifactSettingsPtrOutput() UserArtifactSettingsPtrOutput {
	return i.ToUserArtifactSettingsPtrOutputWithContext(context.Background())
}

func (i *userArtifactSettingsPtrType) ToUserArtifactSettingsPtrOutputWithContext(ctx context.Context) UserArtifactSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactSettingsPtrOutput)
}

type UserArtifactSettingsOutput struct{ *pulumi.OutputState }

func (UserArtifactSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSettings)(nil)).Elem()
}

func (o UserArtifactSettingsOutput) ToUserArtifactSettingsOutput() UserArtifactSettingsOutput {
	return o
}

func (o UserArtifactSettingsOutput) ToUserArtifactSettingsOutputWithContext(ctx context.Context) UserArtifactSettingsOutput {
	return o
}

func (o UserArtifactSettingsOutput) ToUserArtifactSettingsPtrOutput() UserArtifactSettingsPtrOutput {
	return o.ToUserArtifactSettingsPtrOutputWithContext(context.Background())
}

func (o UserArtifactSettingsOutput) ToUserArtifactSettingsPtrOutputWithContext(ctx context.Context) UserArtifactSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserArtifactSettings) *UserArtifactSettings {
		return &v
	}).(UserArtifactSettingsPtrOutput)
}

func (o UserArtifactSettingsOutput) ConfigFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSettings) *string { return v.ConfigFileName }).(pulumi.StringPtrOutput)
}

func (o UserArtifactSettingsOutput) PackageFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSettings) *string { return v.PackageFileName }).(pulumi.StringPtrOutput)
}

type UserArtifactSettingsPtrOutput struct{ *pulumi.OutputState }

func (UserArtifactSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactSettings)(nil)).Elem()
}

func (o UserArtifactSettingsPtrOutput) ToUserArtifactSettingsPtrOutput() UserArtifactSettingsPtrOutput {
	return o
}

func (o UserArtifactSettingsPtrOutput) ToUserArtifactSettingsPtrOutputWithContext(ctx context.Context) UserArtifactSettingsPtrOutput {
	return o
}

func (o UserArtifactSettingsPtrOutput) Elem() UserArtifactSettingsOutput {
	return o.ApplyT(func(v *UserArtifactSettings) UserArtifactSettings {
		if v != nil {
			return *v
		}
		var ret UserArtifactSettings
		return ret
	}).(UserArtifactSettingsOutput)
}

func (o UserArtifactSettingsPtrOutput) ConfigFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactSettings) *string {
		if v == nil {
			return nil
		}
		return v.ConfigFileName
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactSettingsPtrOutput) PackageFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactSettings) *string {
		if v == nil {
			return nil
		}
		return v.PackageFileName
	}).(pulumi.StringPtrOutput)
}

type UserArtifactSettingsResponse struct {
	ConfigFileName  *string `pulumi:"configFileName"`
	PackageFileName *string `pulumi:"packageFileName"`
}

type UserArtifactSettingsResponseOutput struct{ *pulumi.OutputState }

func (UserArtifactSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSettingsResponse)(nil)).Elem()
}

func (o UserArtifactSettingsResponseOutput) ToUserArtifactSettingsResponseOutput() UserArtifactSettingsResponseOutput {
	return o
}

func (o UserArtifactSettingsResponseOutput) ToUserArtifactSettingsResponseOutputWithContext(ctx context.Context) UserArtifactSettingsResponseOutput {
	return o
}

func (o UserArtifactSettingsResponseOutput) ConfigFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSettingsResponse) *string { return v.ConfigFileName }).(pulumi.StringPtrOutput)
}

func (o UserArtifactSettingsResponseOutput) PackageFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSettingsResponse) *string { return v.PackageFileName }).(pulumi.StringPtrOutput)
}

type UserArtifactSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (UserArtifactSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserArtifactSettingsResponse)(nil)).Elem()
}

func (o UserArtifactSettingsResponsePtrOutput) ToUserArtifactSettingsResponsePtrOutput() UserArtifactSettingsResponsePtrOutput {
	return o
}

func (o UserArtifactSettingsResponsePtrOutput) ToUserArtifactSettingsResponsePtrOutputWithContext(ctx context.Context) UserArtifactSettingsResponsePtrOutput {
	return o
}

func (o UserArtifactSettingsResponsePtrOutput) Elem() UserArtifactSettingsResponseOutput {
	return o.ApplyT(func(v *UserArtifactSettingsResponse) UserArtifactSettingsResponse {
		if v != nil {
			return *v
		}
		var ret UserArtifactSettingsResponse
		return ret
	}).(UserArtifactSettingsResponseOutput)
}

func (o UserArtifactSettingsResponsePtrOutput) ConfigFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigFileName
	}).(pulumi.StringPtrOutput)
}

func (o UserArtifactSettingsResponsePtrOutput) PackageFileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserArtifactSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PackageFileName
	}).(pulumi.StringPtrOutput)
}

type UserArtifactSource struct {
	DefaultConfigurationLink *string `pulumi:"defaultConfigurationLink"`
	MediaLink                string  `pulumi:"mediaLink"`
}





type UserArtifactSourceInput interface {
	pulumi.Input

	ToUserArtifactSourceOutput() UserArtifactSourceOutput
	ToUserArtifactSourceOutputWithContext(context.Context) UserArtifactSourceOutput
}

type UserArtifactSourceArgs struct {
	DefaultConfigurationLink pulumi.StringPtrInput `pulumi:"defaultConfigurationLink"`
	MediaLink                pulumi.StringInput    `pulumi:"mediaLink"`
}

func (UserArtifactSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSource)(nil)).Elem()
}

func (i UserArtifactSourceArgs) ToUserArtifactSourceOutput() UserArtifactSourceOutput {
	return i.ToUserArtifactSourceOutputWithContext(context.Background())
}

func (i UserArtifactSourceArgs) ToUserArtifactSourceOutputWithContext(ctx context.Context) UserArtifactSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArtifactSourceOutput)
}

type UserArtifactSourceOutput struct{ *pulumi.OutputState }

func (UserArtifactSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSource)(nil)).Elem()
}

func (o UserArtifactSourceOutput) ToUserArtifactSourceOutput() UserArtifactSourceOutput {
	return o
}

func (o UserArtifactSourceOutput) ToUserArtifactSourceOutputWithContext(ctx context.Context) UserArtifactSourceOutput {
	return o
}

func (o UserArtifactSourceOutput) DefaultConfigurationLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSource) *string { return v.DefaultConfigurationLink }).(pulumi.StringPtrOutput)
}

func (o UserArtifactSourceOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactSource) string { return v.MediaLink }).(pulumi.StringOutput)
}

type UserArtifactSourceResponse struct {
	DefaultConfigurationLink *string `pulumi:"defaultConfigurationLink"`
	MediaLink                string  `pulumi:"mediaLink"`
}

type UserArtifactSourceResponseOutput struct{ *pulumi.OutputState }

func (UserArtifactSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserArtifactSourceResponse)(nil)).Elem()
}

func (o UserArtifactSourceResponseOutput) ToUserArtifactSourceResponseOutput() UserArtifactSourceResponseOutput {
	return o
}

func (o UserArtifactSourceResponseOutput) ToUserArtifactSourceResponseOutputWithContext(ctx context.Context) UserArtifactSourceResponseOutput {
	return o
}

func (o UserArtifactSourceResponseOutput) DefaultConfigurationLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserArtifactSourceResponse) *string { return v.DefaultConfigurationLink }).(pulumi.StringPtrOutput)
}

func (o UserArtifactSourceResponseOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v UserArtifactSourceResponse) string { return v.MediaLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CommunityGalleryInfoOutput{})
	pulumi.RegisterOutputType(CommunityGalleryInfoPtrOutput{})
	pulumi.RegisterOutputType(CommunityGalleryInfoResponseOutput{})
	pulumi.RegisterOutputType(CommunityGalleryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionArrayOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionResponseOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionResponseArrayOutput{})
	pulumi.RegisterOutputType(DisallowedOutput{})
	pulumi.RegisterOutputType(DisallowedPtrOutput{})
	pulumi.RegisterOutputType(DisallowedResponseOutput{})
	pulumi.RegisterOutputType(DisallowedResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionImagesOutput{})
	pulumi.RegisterOutputType(EncryptionImagesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionImagesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionImagesResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryApplicationVersionPublishingProfileOutput{})
	pulumi.RegisterOutputType(GalleryApplicationVersionPublishingProfileResponseOutput{})
	pulumi.RegisterOutputType(GalleryArtifactVersionSourceOutput{})
	pulumi.RegisterOutputType(GalleryArtifactVersionSourcePtrOutput{})
	pulumi.RegisterOutputType(GalleryArtifactVersionSourceResponseOutput{})
	pulumi.RegisterOutputType(GalleryArtifactVersionSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryDataDiskImageOutput{})
	pulumi.RegisterOutputType(GalleryDataDiskImageArrayOutput{})
	pulumi.RegisterOutputType(GalleryDataDiskImageResponseOutput{})
	pulumi.RegisterOutputType(GalleryDataDiskImageResponseArrayOutput{})
	pulumi.RegisterOutputType(GalleryExtendedLocationOutput{})
	pulumi.RegisterOutputType(GalleryExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(GalleryExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(GalleryExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryIdentifierResponseOutput{})
	pulumi.RegisterOutputType(GalleryIdentifierResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageFeatureOutput{})
	pulumi.RegisterOutputType(GalleryImageFeatureArrayOutput{})
	pulumi.RegisterOutputType(GalleryImageFeatureResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageFeatureResponseArrayOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierOutput{})
	pulumi.RegisterOutputType(GalleryImageIdentifierResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionPublishingProfileOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionPublishingProfilePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionPublishingProfileResponseOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionPublishingProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionStorageProfileOutput{})
	pulumi.RegisterOutputType(GalleryImageVersionStorageProfileResponseOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImageOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImagePtrOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImageResponseOutput{})
	pulumi.RegisterOutputType(GalleryOSDiskImageResponsePtrOutput{})
	pulumi.RegisterOutputType(GalleryTargetExtendedLocationOutput{})
	pulumi.RegisterOutputType(GalleryTargetExtendedLocationArrayOutput{})
	pulumi.RegisterOutputType(GalleryTargetExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(GalleryTargetExtendedLocationResponseArrayOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionPtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionResponseOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageSecurityProfileOutput{})
	pulumi.RegisterOutputType(OSDiskImageSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageSecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(OSDiskImageSecurityProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(RegionalReplicationStatusResponseOutput{})
	pulumi.RegisterOutputType(RegionalReplicationStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(RegionalSharingStatusResponseOutput{})
	pulumi.RegisterOutputType(RegionalSharingStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(ReplicationStatusResponseOutput{})
	pulumi.RegisterOutputType(ResourceRangeOutput{})
	pulumi.RegisterOutputType(ResourceRangePtrOutput{})
	pulumi.RegisterOutputType(ResourceRangeResponseOutput{})
	pulumi.RegisterOutputType(ResourceRangeResponsePtrOutput{})
	pulumi.RegisterOutputType(SharingProfileOutput{})
	pulumi.RegisterOutputType(SharingProfilePtrOutput{})
	pulumi.RegisterOutputType(SharingProfileGroupResponseOutput{})
	pulumi.RegisterOutputType(SharingProfileGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(SharingProfileResponseOutput{})
	pulumi.RegisterOutputType(SharingProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SharingStatusResponseOutput{})
	pulumi.RegisterOutputType(SoftDeletePolicyOutput{})
	pulumi.RegisterOutputType(SoftDeletePolicyPtrOutput{})
	pulumi.RegisterOutputType(SoftDeletePolicyResponseOutput{})
	pulumi.RegisterOutputType(SoftDeletePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetRegionOutput{})
	pulumi.RegisterOutputType(TargetRegionArrayOutput{})
	pulumi.RegisterOutputType(TargetRegionResponseOutput{})
	pulumi.RegisterOutputType(TargetRegionResponseArrayOutput{})
	pulumi.RegisterOutputType(UserArtifactManageOutput{})
	pulumi.RegisterOutputType(UserArtifactManagePtrOutput{})
	pulumi.RegisterOutputType(UserArtifactManageResponseOutput{})
	pulumi.RegisterOutputType(UserArtifactManageResponsePtrOutput{})
	pulumi.RegisterOutputType(UserArtifactSettingsOutput{})
	pulumi.RegisterOutputType(UserArtifactSettingsPtrOutput{})
	pulumi.RegisterOutputType(UserArtifactSettingsResponseOutput{})
	pulumi.RegisterOutputType(UserArtifactSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserArtifactSourceOutput{})
	pulumi.RegisterOutputType(UserArtifactSourceResponseOutput{})
}
