


package v20200930

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreationData struct {
	CreateOption          string              `pulumi:"createOption"`
	GalleryImageReference *ImageDiskReference `pulumi:"galleryImageReference"`
	ImageReference        *ImageDiskReference `pulumi:"imageReference"`
	LogicalSectorSize     *int                `pulumi:"logicalSectorSize"`
	SourceResourceId      *string             `pulumi:"sourceResourceId"`
	SourceUri             *string             `pulumi:"sourceUri"`
	StorageAccountId      *string             `pulumi:"storageAccountId"`
	UploadSizeBytes       *float64            `pulumi:"uploadSizeBytes"`
}





type CreationDataInput interface {
	pulumi.Input

	ToCreationDataOutput() CreationDataOutput
	ToCreationDataOutputWithContext(context.Context) CreationDataOutput
}

type CreationDataArgs struct {
	CreateOption          pulumi.StringInput         `pulumi:"createOption"`
	GalleryImageReference ImageDiskReferencePtrInput `pulumi:"galleryImageReference"`
	ImageReference        ImageDiskReferencePtrInput `pulumi:"imageReference"`
	LogicalSectorSize     pulumi.IntPtrInput         `pulumi:"logicalSectorSize"`
	SourceResourceId      pulumi.StringPtrInput      `pulumi:"sourceResourceId"`
	SourceUri             pulumi.StringPtrInput      `pulumi:"sourceUri"`
	StorageAccountId      pulumi.StringPtrInput      `pulumi:"storageAccountId"`
	UploadSizeBytes       pulumi.Float64PtrInput     `pulumi:"uploadSizeBytes"`
}

func (CreationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationData)(nil)).Elem()
}

func (i CreationDataArgs) ToCreationDataOutput() CreationDataOutput {
	return i.ToCreationDataOutputWithContext(context.Background())
}

func (i CreationDataArgs) ToCreationDataOutputWithContext(ctx context.Context) CreationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataOutput)
}

type CreationDataOutput struct{ *pulumi.OutputState }

func (CreationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationData)(nil)).Elem()
}

func (o CreationDataOutput) ToCreationDataOutput() CreationDataOutput {
	return o
}

func (o CreationDataOutput) ToCreationDataOutputWithContext(ctx context.Context) CreationDataOutput {
	return o
}

func (o CreationDataOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v CreationData) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o CreationDataOutput) GalleryImageReference() ImageDiskReferencePtrOutput {
	return o.ApplyT(func(v CreationData) *ImageDiskReference { return v.GalleryImageReference }).(ImageDiskReferencePtrOutput)
}

func (o CreationDataOutput) ImageReference() ImageDiskReferencePtrOutput {
	return o.ApplyT(func(v CreationData) *ImageDiskReference { return v.ImageReference }).(ImageDiskReferencePtrOutput)
}

func (o CreationDataOutput) LogicalSectorSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreationData) *int { return v.LogicalSectorSize }).(pulumi.IntPtrOutput)
}

func (o CreationDataOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o CreationDataOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o CreationDataOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o CreationDataOutput) UploadSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CreationData) *float64 { return v.UploadSizeBytes }).(pulumi.Float64PtrOutput)
}

type CreationDataResponse struct {
	CreateOption          string                      `pulumi:"createOption"`
	GalleryImageReference *ImageDiskReferenceResponse `pulumi:"galleryImageReference"`
	ImageReference        *ImageDiskReferenceResponse `pulumi:"imageReference"`
	LogicalSectorSize     *int                        `pulumi:"logicalSectorSize"`
	SourceResourceId      *string                     `pulumi:"sourceResourceId"`
	SourceUniqueId        string                      `pulumi:"sourceUniqueId"`
	SourceUri             *string                     `pulumi:"sourceUri"`
	StorageAccountId      *string                     `pulumi:"storageAccountId"`
	UploadSizeBytes       *float64                    `pulumi:"uploadSizeBytes"`
}

type CreationDataResponseOutput struct{ *pulumi.OutputState }

func (CreationDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationDataResponse)(nil)).Elem()
}

func (o CreationDataResponseOutput) ToCreationDataResponseOutput() CreationDataResponseOutput {
	return o
}

func (o CreationDataResponseOutput) ToCreationDataResponseOutputWithContext(ctx context.Context) CreationDataResponseOutput {
	return o
}

func (o CreationDataResponseOutput) CreateOption() pulumi.StringOutput {
	return o.ApplyT(func(v CreationDataResponse) string { return v.CreateOption }).(pulumi.StringOutput)
}

func (o CreationDataResponseOutput) GalleryImageReference() ImageDiskReferenceResponsePtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *ImageDiskReferenceResponse { return v.GalleryImageReference }).(ImageDiskReferenceResponsePtrOutput)
}

func (o CreationDataResponseOutput) ImageReference() ImageDiskReferenceResponsePtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *ImageDiskReferenceResponse { return v.ImageReference }).(ImageDiskReferenceResponsePtrOutput)
}

func (o CreationDataResponseOutput) LogicalSectorSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *int { return v.LogicalSectorSize }).(pulumi.IntPtrOutput)
}

func (o CreationDataResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o CreationDataResponseOutput) SourceUniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v CreationDataResponse) string { return v.SourceUniqueId }).(pulumi.StringOutput)
}

func (o CreationDataResponseOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.SourceUri }).(pulumi.StringPtrOutput)
}

func (o CreationDataResponseOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o CreationDataResponseOutput) UploadSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *float64 { return v.UploadSizeBytes }).(pulumi.Float64PtrOutput)
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

type DiskSku struct {
	Name *string `pulumi:"name"`
}





type DiskSkuInput interface {
	pulumi.Input

	ToDiskSkuOutput() DiskSkuOutput
	ToDiskSkuOutputWithContext(context.Context) DiskSkuOutput
}

type DiskSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DiskSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSku)(nil)).Elem()
}

func (i DiskSkuArgs) ToDiskSkuOutput() DiskSkuOutput {
	return i.ToDiskSkuOutputWithContext(context.Background())
}

func (i DiskSkuArgs) ToDiskSkuOutputWithContext(ctx context.Context) DiskSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuOutput)
}

func (i DiskSkuArgs) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return i.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (i DiskSkuArgs) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuOutput).ToDiskSkuPtrOutputWithContext(ctx)
}









type DiskSkuPtrInput interface {
	pulumi.Input

	ToDiskSkuPtrOutput() DiskSkuPtrOutput
	ToDiskSkuPtrOutputWithContext(context.Context) DiskSkuPtrOutput
}

type diskSkuPtrType DiskSkuArgs

func DiskSkuPtr(v *DiskSkuArgs) DiskSkuPtrInput {
	return (*diskSkuPtrType)(v)
}

func (*diskSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSku)(nil)).Elem()
}

func (i *diskSkuPtrType) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return i.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (i *diskSkuPtrType) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuPtrOutput)
}

type DiskSkuOutput struct{ *pulumi.OutputState }

func (DiskSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSku)(nil)).Elem()
}

func (o DiskSkuOutput) ToDiskSkuOutput() DiskSkuOutput {
	return o
}

func (o DiskSkuOutput) ToDiskSkuOutputWithContext(ctx context.Context) DiskSkuOutput {
	return o
}

func (o DiskSkuOutput) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return o.ToDiskSkuPtrOutputWithContext(context.Background())
}

func (o DiskSkuOutput) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskSku) *DiskSku {
		return &v
	}).(DiskSkuPtrOutput)
}

func (o DiskSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DiskSkuPtrOutput struct{ *pulumi.OutputState }

func (DiskSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSku)(nil)).Elem()
}

func (o DiskSkuPtrOutput) ToDiskSkuPtrOutput() DiskSkuPtrOutput {
	return o
}

func (o DiskSkuPtrOutput) ToDiskSkuPtrOutputWithContext(ctx context.Context) DiskSkuPtrOutput {
	return o
}

func (o DiskSkuPtrOutput) Elem() DiskSkuOutput {
	return o.ApplyT(func(v *DiskSku) DiskSku {
		if v != nil {
			return *v
		}
		var ret DiskSku
		return ret
	}).(DiskSkuOutput)
}

func (o DiskSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DiskSkuResponse struct {
	Name *string `pulumi:"name"`
	Tier string  `pulumi:"tier"`
}

type DiskSkuResponseOutput struct{ *pulumi.OutputState }

func (DiskSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutput() DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) ToDiskSkuResponseOutputWithContext(ctx context.Context) DiskSkuResponseOutput {
	return o
}

func (o DiskSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DiskSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v DiskSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type DiskSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSkuResponse)(nil)).Elem()
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return o
}

func (o DiskSkuResponsePtrOutput) Elem() DiskSkuResponseOutput {
	return o.ApplyT(func(v *DiskSkuResponse) DiskSkuResponse {
		if v != nil {
			return *v
		}
		var ret DiskSkuResponse
		return ret
	}).(DiskSkuResponseOutput)
}

func (o DiskSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DiskSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type Encryption struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	Type                *string `pulumi:"type"`
}





type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

type EncryptionArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
	Type                pulumi.StringPtrInput `pulumi:"type"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}









type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

type EncryptionOutput struct{ *pulumi.OutputState }

func (EncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (o EncryptionOutput) ToEncryptionOutput() EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return o
}

func (o EncryptionOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o.ToEncryptionPtrOutputWithContext(context.Background())
}

func (o EncryptionOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encryption) *Encryption {
		return &v
	}).(EncryptionPtrOutput)
}

func (o EncryptionOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o EncryptionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Encryption) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EncryptionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return o
}

func (o EncryptionPtrOutput) Elem() EncryptionOutput {
	return o.ApplyT(func(v *Encryption) Encryption {
		if v != nil {
			return *v
		}
		var ret Encryption
		return ret
	}).(EncryptionOutput)
}

func (o EncryptionPtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Encryption) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
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

type EncryptionResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	Type                *string `pulumi:"type"`
}

type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o EncryptionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionResponse
		return ret
	}).(EncryptionResponseOutput)
}

func (o EncryptionResponsePtrOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type EncryptionSetIdentity struct {
	Type *string `pulumi:"type"`
}





type EncryptionSetIdentityInput interface {
	pulumi.Input

	ToEncryptionSetIdentityOutput() EncryptionSetIdentityOutput
	ToEncryptionSetIdentityOutputWithContext(context.Context) EncryptionSetIdentityOutput
}

type EncryptionSetIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (EncryptionSetIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSetIdentity)(nil)).Elem()
}

func (i EncryptionSetIdentityArgs) ToEncryptionSetIdentityOutput() EncryptionSetIdentityOutput {
	return i.ToEncryptionSetIdentityOutputWithContext(context.Background())
}

func (i EncryptionSetIdentityArgs) ToEncryptionSetIdentityOutputWithContext(ctx context.Context) EncryptionSetIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityOutput)
}

func (i EncryptionSetIdentityArgs) ToEncryptionSetIdentityPtrOutput() EncryptionSetIdentityPtrOutput {
	return i.ToEncryptionSetIdentityPtrOutputWithContext(context.Background())
}

func (i EncryptionSetIdentityArgs) ToEncryptionSetIdentityPtrOutputWithContext(ctx context.Context) EncryptionSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityOutput).ToEncryptionSetIdentityPtrOutputWithContext(ctx)
}









type EncryptionSetIdentityPtrInput interface {
	pulumi.Input

	ToEncryptionSetIdentityPtrOutput() EncryptionSetIdentityPtrOutput
	ToEncryptionSetIdentityPtrOutputWithContext(context.Context) EncryptionSetIdentityPtrOutput
}

type encryptionSetIdentityPtrType EncryptionSetIdentityArgs

func EncryptionSetIdentityPtr(v *EncryptionSetIdentityArgs) EncryptionSetIdentityPtrInput {
	return (*encryptionSetIdentityPtrType)(v)
}

func (*encryptionSetIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSetIdentity)(nil)).Elem()
}

func (i *encryptionSetIdentityPtrType) ToEncryptionSetIdentityPtrOutput() EncryptionSetIdentityPtrOutput {
	return i.ToEncryptionSetIdentityPtrOutputWithContext(context.Background())
}

func (i *encryptionSetIdentityPtrType) ToEncryptionSetIdentityPtrOutputWithContext(ctx context.Context) EncryptionSetIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityPtrOutput)
}

type EncryptionSetIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionSetIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSetIdentity)(nil)).Elem()
}

func (o EncryptionSetIdentityOutput) ToEncryptionSetIdentityOutput() EncryptionSetIdentityOutput {
	return o
}

func (o EncryptionSetIdentityOutput) ToEncryptionSetIdentityOutputWithContext(ctx context.Context) EncryptionSetIdentityOutput {
	return o
}

func (o EncryptionSetIdentityOutput) ToEncryptionSetIdentityPtrOutput() EncryptionSetIdentityPtrOutput {
	return o.ToEncryptionSetIdentityPtrOutputWithContext(context.Background())
}

func (o EncryptionSetIdentityOutput) ToEncryptionSetIdentityPtrOutputWithContext(ctx context.Context) EncryptionSetIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionSetIdentity) *EncryptionSetIdentity {
		return &v
	}).(EncryptionSetIdentityPtrOutput)
}

func (o EncryptionSetIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionSetIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EncryptionSetIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionSetIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSetIdentity)(nil)).Elem()
}

func (o EncryptionSetIdentityPtrOutput) ToEncryptionSetIdentityPtrOutput() EncryptionSetIdentityPtrOutput {
	return o
}

func (o EncryptionSetIdentityPtrOutput) ToEncryptionSetIdentityPtrOutputWithContext(ctx context.Context) EncryptionSetIdentityPtrOutput {
	return o
}

func (o EncryptionSetIdentityPtrOutput) Elem() EncryptionSetIdentityOutput {
	return o.ApplyT(func(v *EncryptionSetIdentity) EncryptionSetIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionSetIdentity
		return ret
	}).(EncryptionSetIdentityOutput)
}

func (o EncryptionSetIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSetIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type EncryptionSetIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type EncryptionSetIdentityResponseOutput struct{ *pulumi.OutputState }

func (EncryptionSetIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSetIdentityResponse)(nil)).Elem()
}

func (o EncryptionSetIdentityResponseOutput) ToEncryptionSetIdentityResponseOutput() EncryptionSetIdentityResponseOutput {
	return o
}

func (o EncryptionSetIdentityResponseOutput) ToEncryptionSetIdentityResponseOutputWithContext(ctx context.Context) EncryptionSetIdentityResponseOutput {
	return o
}

func (o EncryptionSetIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionSetIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o EncryptionSetIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionSetIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o EncryptionSetIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionSetIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type EncryptionSetIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionSetIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSetIdentityResponse)(nil)).Elem()
}

func (o EncryptionSetIdentityResponsePtrOutput) ToEncryptionSetIdentityResponsePtrOutput() EncryptionSetIdentityResponsePtrOutput {
	return o
}

func (o EncryptionSetIdentityResponsePtrOutput) ToEncryptionSetIdentityResponsePtrOutputWithContext(ctx context.Context) EncryptionSetIdentityResponsePtrOutput {
	return o
}

func (o EncryptionSetIdentityResponsePtrOutput) Elem() EncryptionSetIdentityResponseOutput {
	return o.ApplyT(func(v *EncryptionSetIdentityResponse) EncryptionSetIdentityResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionSetIdentityResponse
		return ret
	}).(EncryptionSetIdentityResponseOutput)
}

func (o EncryptionSetIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionSetIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionSetIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSetIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type EncryptionSettingsCollection struct {
	Enabled                   bool                        `pulumi:"enabled"`
	EncryptionSettings        []EncryptionSettingsElement `pulumi:"encryptionSettings"`
	EncryptionSettingsVersion *string                     `pulumi:"encryptionSettingsVersion"`
}





type EncryptionSettingsCollectionInput interface {
	pulumi.Input

	ToEncryptionSettingsCollectionOutput() EncryptionSettingsCollectionOutput
	ToEncryptionSettingsCollectionOutputWithContext(context.Context) EncryptionSettingsCollectionOutput
}

type EncryptionSettingsCollectionArgs struct {
	Enabled                   pulumi.BoolInput                    `pulumi:"enabled"`
	EncryptionSettings        EncryptionSettingsElementArrayInput `pulumi:"encryptionSettings"`
	EncryptionSettingsVersion pulumi.StringPtrInput               `pulumi:"encryptionSettingsVersion"`
}

func (EncryptionSettingsCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsCollection)(nil)).Elem()
}

func (i EncryptionSettingsCollectionArgs) ToEncryptionSettingsCollectionOutput() EncryptionSettingsCollectionOutput {
	return i.ToEncryptionSettingsCollectionOutputWithContext(context.Background())
}

func (i EncryptionSettingsCollectionArgs) ToEncryptionSettingsCollectionOutputWithContext(ctx context.Context) EncryptionSettingsCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionOutput)
}

func (i EncryptionSettingsCollectionArgs) ToEncryptionSettingsCollectionPtrOutput() EncryptionSettingsCollectionPtrOutput {
	return i.ToEncryptionSettingsCollectionPtrOutputWithContext(context.Background())
}

func (i EncryptionSettingsCollectionArgs) ToEncryptionSettingsCollectionPtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionOutput).ToEncryptionSettingsCollectionPtrOutputWithContext(ctx)
}









type EncryptionSettingsCollectionPtrInput interface {
	pulumi.Input

	ToEncryptionSettingsCollectionPtrOutput() EncryptionSettingsCollectionPtrOutput
	ToEncryptionSettingsCollectionPtrOutputWithContext(context.Context) EncryptionSettingsCollectionPtrOutput
}

type encryptionSettingsCollectionPtrType EncryptionSettingsCollectionArgs

func EncryptionSettingsCollectionPtr(v *EncryptionSettingsCollectionArgs) EncryptionSettingsCollectionPtrInput {
	return (*encryptionSettingsCollectionPtrType)(v)
}

func (*encryptionSettingsCollectionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettingsCollection)(nil)).Elem()
}

func (i *encryptionSettingsCollectionPtrType) ToEncryptionSettingsCollectionPtrOutput() EncryptionSettingsCollectionPtrOutput {
	return i.ToEncryptionSettingsCollectionPtrOutputWithContext(context.Background())
}

func (i *encryptionSettingsCollectionPtrType) ToEncryptionSettingsCollectionPtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionPtrOutput)
}

type EncryptionSettingsCollectionOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsCollection)(nil)).Elem()
}

func (o EncryptionSettingsCollectionOutput) ToEncryptionSettingsCollectionOutput() EncryptionSettingsCollectionOutput {
	return o
}

func (o EncryptionSettingsCollectionOutput) ToEncryptionSettingsCollectionOutputWithContext(ctx context.Context) EncryptionSettingsCollectionOutput {
	return o
}

func (o EncryptionSettingsCollectionOutput) ToEncryptionSettingsCollectionPtrOutput() EncryptionSettingsCollectionPtrOutput {
	return o.ToEncryptionSettingsCollectionPtrOutputWithContext(context.Background())
}

func (o EncryptionSettingsCollectionOutput) ToEncryptionSettingsCollectionPtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionSettingsCollection) *EncryptionSettingsCollection {
		return &v
	}).(EncryptionSettingsCollectionPtrOutput)
}

func (o EncryptionSettingsCollectionOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v EncryptionSettingsCollection) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o EncryptionSettingsCollectionOutput) EncryptionSettings() EncryptionSettingsElementArrayOutput {
	return o.ApplyT(func(v EncryptionSettingsCollection) []EncryptionSettingsElement { return v.EncryptionSettings }).(EncryptionSettingsElementArrayOutput)
}

func (o EncryptionSettingsCollectionOutput) EncryptionSettingsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionSettingsCollection) *string { return v.EncryptionSettingsVersion }).(pulumi.StringPtrOutput)
}

type EncryptionSettingsCollectionPtrOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsCollectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettingsCollection)(nil)).Elem()
}

func (o EncryptionSettingsCollectionPtrOutput) ToEncryptionSettingsCollectionPtrOutput() EncryptionSettingsCollectionPtrOutput {
	return o
}

func (o EncryptionSettingsCollectionPtrOutput) ToEncryptionSettingsCollectionPtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionPtrOutput {
	return o
}

func (o EncryptionSettingsCollectionPtrOutput) Elem() EncryptionSettingsCollectionOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollection) EncryptionSettingsCollection {
		if v != nil {
			return *v
		}
		var ret EncryptionSettingsCollection
		return ret
	}).(EncryptionSettingsCollectionOutput)
}

func (o EncryptionSettingsCollectionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollection) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsCollectionPtrOutput) EncryptionSettings() EncryptionSettingsElementArrayOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollection) []EncryptionSettingsElement {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(EncryptionSettingsElementArrayOutput)
}

func (o EncryptionSettingsCollectionPtrOutput) EncryptionSettingsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollection) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionSettingsVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionSettingsCollectionResponse struct {
	Enabled                   bool                                `pulumi:"enabled"`
	EncryptionSettings        []EncryptionSettingsElementResponse `pulumi:"encryptionSettings"`
	EncryptionSettingsVersion *string                             `pulumi:"encryptionSettingsVersion"`
}

type EncryptionSettingsCollectionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsCollectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsCollectionResponse)(nil)).Elem()
}

func (o EncryptionSettingsCollectionResponseOutput) ToEncryptionSettingsCollectionResponseOutput() EncryptionSettingsCollectionResponseOutput {
	return o
}

func (o EncryptionSettingsCollectionResponseOutput) ToEncryptionSettingsCollectionResponseOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponseOutput {
	return o
}

func (o EncryptionSettingsCollectionResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v EncryptionSettingsCollectionResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o EncryptionSettingsCollectionResponseOutput) EncryptionSettings() EncryptionSettingsElementResponseArrayOutput {
	return o.ApplyT(func(v EncryptionSettingsCollectionResponse) []EncryptionSettingsElementResponse {
		return v.EncryptionSettings
	}).(EncryptionSettingsElementResponseArrayOutput)
}

func (o EncryptionSettingsCollectionResponseOutput) EncryptionSettingsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionSettingsCollectionResponse) *string { return v.EncryptionSettingsVersion }).(pulumi.StringPtrOutput)
}

type EncryptionSettingsCollectionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsCollectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettingsCollectionResponse)(nil)).Elem()
}

func (o EncryptionSettingsCollectionResponsePtrOutput) ToEncryptionSettingsCollectionResponsePtrOutput() EncryptionSettingsCollectionResponsePtrOutput {
	return o
}

func (o EncryptionSettingsCollectionResponsePtrOutput) ToEncryptionSettingsCollectionResponsePtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponsePtrOutput {
	return o
}

func (o EncryptionSettingsCollectionResponsePtrOutput) Elem() EncryptionSettingsCollectionResponseOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollectionResponse) EncryptionSettingsCollectionResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionSettingsCollectionResponse
		return ret
	}).(EncryptionSettingsCollectionResponseOutput)
}

func (o EncryptionSettingsCollectionResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollectionResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o EncryptionSettingsCollectionResponsePtrOutput) EncryptionSettings() EncryptionSettingsElementResponseArrayOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollectionResponse) []EncryptionSettingsElementResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionSettings
	}).(EncryptionSettingsElementResponseArrayOutput)
}

func (o EncryptionSettingsCollectionResponsePtrOutput) EncryptionSettingsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionSettingsCollectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionSettingsVersion
	}).(pulumi.StringPtrOutput)
}

type EncryptionSettingsElement struct {
	DiskEncryptionKey *KeyVaultAndSecretReference `pulumi:"diskEncryptionKey"`
	KeyEncryptionKey  *KeyVaultAndKeyReference    `pulumi:"keyEncryptionKey"`
}





type EncryptionSettingsElementInput interface {
	pulumi.Input

	ToEncryptionSettingsElementOutput() EncryptionSettingsElementOutput
	ToEncryptionSettingsElementOutputWithContext(context.Context) EncryptionSettingsElementOutput
}

type EncryptionSettingsElementArgs struct {
	DiskEncryptionKey KeyVaultAndSecretReferencePtrInput `pulumi:"diskEncryptionKey"`
	KeyEncryptionKey  KeyVaultAndKeyReferencePtrInput    `pulumi:"keyEncryptionKey"`
}

func (EncryptionSettingsElementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsElement)(nil)).Elem()
}

func (i EncryptionSettingsElementArgs) ToEncryptionSettingsElementOutput() EncryptionSettingsElementOutput {
	return i.ToEncryptionSettingsElementOutputWithContext(context.Background())
}

func (i EncryptionSettingsElementArgs) ToEncryptionSettingsElementOutputWithContext(ctx context.Context) EncryptionSettingsElementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsElementOutput)
}





type EncryptionSettingsElementArrayInput interface {
	pulumi.Input

	ToEncryptionSettingsElementArrayOutput() EncryptionSettingsElementArrayOutput
	ToEncryptionSettingsElementArrayOutputWithContext(context.Context) EncryptionSettingsElementArrayOutput
}

type EncryptionSettingsElementArray []EncryptionSettingsElementInput

func (EncryptionSettingsElementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncryptionSettingsElement)(nil)).Elem()
}

func (i EncryptionSettingsElementArray) ToEncryptionSettingsElementArrayOutput() EncryptionSettingsElementArrayOutput {
	return i.ToEncryptionSettingsElementArrayOutputWithContext(context.Background())
}

func (i EncryptionSettingsElementArray) ToEncryptionSettingsElementArrayOutputWithContext(ctx context.Context) EncryptionSettingsElementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsElementArrayOutput)
}

type EncryptionSettingsElementOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsElementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsElement)(nil)).Elem()
}

func (o EncryptionSettingsElementOutput) ToEncryptionSettingsElementOutput() EncryptionSettingsElementOutput {
	return o
}

func (o EncryptionSettingsElementOutput) ToEncryptionSettingsElementOutputWithContext(ctx context.Context) EncryptionSettingsElementOutput {
	return o
}

func (o EncryptionSettingsElementOutput) DiskEncryptionKey() KeyVaultAndSecretReferencePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsElement) *KeyVaultAndSecretReference { return v.DiskEncryptionKey }).(KeyVaultAndSecretReferencePtrOutput)
}

func (o EncryptionSettingsElementOutput) KeyEncryptionKey() KeyVaultAndKeyReferencePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsElement) *KeyVaultAndKeyReference { return v.KeyEncryptionKey }).(KeyVaultAndKeyReferencePtrOutput)
}

type EncryptionSettingsElementArrayOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsElementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncryptionSettingsElement)(nil)).Elem()
}

func (o EncryptionSettingsElementArrayOutput) ToEncryptionSettingsElementArrayOutput() EncryptionSettingsElementArrayOutput {
	return o
}

func (o EncryptionSettingsElementArrayOutput) ToEncryptionSettingsElementArrayOutputWithContext(ctx context.Context) EncryptionSettingsElementArrayOutput {
	return o
}

func (o EncryptionSettingsElementArrayOutput) Index(i pulumi.IntInput) EncryptionSettingsElementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncryptionSettingsElement {
		return vs[0].([]EncryptionSettingsElement)[vs[1].(int)]
	}).(EncryptionSettingsElementOutput)
}

type EncryptionSettingsElementResponse struct {
	DiskEncryptionKey *KeyVaultAndSecretReferenceResponse `pulumi:"diskEncryptionKey"`
	KeyEncryptionKey  *KeyVaultAndKeyReferenceResponse    `pulumi:"keyEncryptionKey"`
}

type EncryptionSettingsElementResponseOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsElementResponse)(nil)).Elem()
}

func (o EncryptionSettingsElementResponseOutput) ToEncryptionSettingsElementResponseOutput() EncryptionSettingsElementResponseOutput {
	return o
}

func (o EncryptionSettingsElementResponseOutput) ToEncryptionSettingsElementResponseOutputWithContext(ctx context.Context) EncryptionSettingsElementResponseOutput {
	return o
}

func (o EncryptionSettingsElementResponseOutput) DiskEncryptionKey() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsElementResponse) *KeyVaultAndSecretReferenceResponse {
		return v.DiskEncryptionKey
	}).(KeyVaultAndSecretReferenceResponsePtrOutput)
}

func (o EncryptionSettingsElementResponseOutput) KeyEncryptionKey() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionSettingsElementResponse) *KeyVaultAndKeyReferenceResponse { return v.KeyEncryptionKey }).(KeyVaultAndKeyReferenceResponsePtrOutput)
}

type EncryptionSettingsElementResponseArrayOutput struct{ *pulumi.OutputState }

func (EncryptionSettingsElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncryptionSettingsElementResponse)(nil)).Elem()
}

func (o EncryptionSettingsElementResponseArrayOutput) ToEncryptionSettingsElementResponseArrayOutput() EncryptionSettingsElementResponseArrayOutput {
	return o
}

func (o EncryptionSettingsElementResponseArrayOutput) ToEncryptionSettingsElementResponseArrayOutputWithContext(ctx context.Context) EncryptionSettingsElementResponseArrayOutput {
	return o
}

func (o EncryptionSettingsElementResponseArrayOutput) Index(i pulumi.IntInput) EncryptionSettingsElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EncryptionSettingsElementResponse {
		return vs[0].([]EncryptionSettingsElementResponse)[vs[1].(int)]
	}).(EncryptionSettingsElementResponseOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type GalleryApplicationVersionPublishingProfile struct {
	EnableHealthCheck  *bool               `pulumi:"enableHealthCheck"`
	EndOfLifeDate      *string             `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  *bool               `pulumi:"excludeFromLatest"`
	ManageActions      *UserArtifactManage `pulumi:"manageActions"`
	ReplicaCount       *int                `pulumi:"replicaCount"`
	Source             UserArtifactSource  `pulumi:"source"`
	StorageAccountType *string             `pulumi:"storageAccountType"`
	TargetRegions      []TargetRegion      `pulumi:"targetRegions"`
}





type GalleryApplicationVersionPublishingProfileInput interface {
	pulumi.Input

	ToGalleryApplicationVersionPublishingProfileOutput() GalleryApplicationVersionPublishingProfileOutput
	ToGalleryApplicationVersionPublishingProfileOutputWithContext(context.Context) GalleryApplicationVersionPublishingProfileOutput
}

type GalleryApplicationVersionPublishingProfileArgs struct {
	EnableHealthCheck  pulumi.BoolPtrInput        `pulumi:"enableHealthCheck"`
	EndOfLifeDate      pulumi.StringPtrInput      `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  pulumi.BoolPtrInput        `pulumi:"excludeFromLatest"`
	ManageActions      UserArtifactManagePtrInput `pulumi:"manageActions"`
	ReplicaCount       pulumi.IntPtrInput         `pulumi:"replicaCount"`
	Source             UserArtifactSourceInput    `pulumi:"source"`
	StorageAccountType pulumi.StringPtrInput      `pulumi:"storageAccountType"`
	TargetRegions      TargetRegionArrayInput     `pulumi:"targetRegions"`
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

func (o GalleryApplicationVersionPublishingProfileOutput) Source() UserArtifactSourceOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) UserArtifactSource { return v.Source }).(UserArtifactSourceOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
}

func (o GalleryApplicationVersionPublishingProfileOutput) TargetRegions() TargetRegionArrayOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfile) []TargetRegion { return v.TargetRegions }).(TargetRegionArrayOutput)
}

type GalleryApplicationVersionPublishingProfileResponse struct {
	EnableHealthCheck  *bool                       `pulumi:"enableHealthCheck"`
	EndOfLifeDate      *string                     `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  *bool                       `pulumi:"excludeFromLatest"`
	ManageActions      *UserArtifactManageResponse `pulumi:"manageActions"`
	PublishedDate      string                      `pulumi:"publishedDate"`
	ReplicaCount       *int                        `pulumi:"replicaCount"`
	Source             UserArtifactSourceResponse  `pulumi:"source"`
	StorageAccountType *string                     `pulumi:"storageAccountType"`
	TargetRegions      []TargetRegionResponse      `pulumi:"targetRegions"`
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

func (o GalleryApplicationVersionPublishingProfileResponseOutput) Source() UserArtifactSourceResponseOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) UserArtifactSourceResponse { return v.Source }).(UserArtifactSourceResponseOutput)
}

func (o GalleryApplicationVersionPublishingProfileResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryApplicationVersionPublishingProfileResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
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
	EndOfLifeDate      *string        `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  *bool          `pulumi:"excludeFromLatest"`
	ReplicaCount       *int           `pulumi:"replicaCount"`
	StorageAccountType *string        `pulumi:"storageAccountType"`
	TargetRegions      []TargetRegion `pulumi:"targetRegions"`
}





type GalleryImageVersionPublishingProfileInput interface {
	pulumi.Input

	ToGalleryImageVersionPublishingProfileOutput() GalleryImageVersionPublishingProfileOutput
	ToGalleryImageVersionPublishingProfileOutputWithContext(context.Context) GalleryImageVersionPublishingProfileOutput
}

type GalleryImageVersionPublishingProfileArgs struct {
	EndOfLifeDate      pulumi.StringPtrInput  `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  pulumi.BoolPtrInput    `pulumi:"excludeFromLatest"`
	ReplicaCount       pulumi.IntPtrInput     `pulumi:"replicaCount"`
	StorageAccountType pulumi.StringPtrInput  `pulumi:"storageAccountType"`
	TargetRegions      TargetRegionArrayInput `pulumi:"targetRegions"`
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

func (o GalleryImageVersionPublishingProfileOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfile) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
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

func (o GalleryImageVersionPublishingProfilePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
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
	EndOfLifeDate      *string                `pulumi:"endOfLifeDate"`
	ExcludeFromLatest  *bool                  `pulumi:"excludeFromLatest"`
	PublishedDate      string                 `pulumi:"publishedDate"`
	ReplicaCount       *int                   `pulumi:"replicaCount"`
	StorageAccountType *string                `pulumi:"storageAccountType"`
	TargetRegions      []TargetRegionResponse `pulumi:"targetRegions"`
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

func (o GalleryImageVersionPublishingProfileResponseOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageVersionPublishingProfileResponse) *string { return v.StorageAccountType }).(pulumi.StringPtrOutput)
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

func (o GalleryImageVersionPublishingProfileResponsePtrOutput) StorageAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GalleryImageVersionPublishingProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountType
	}).(pulumi.StringPtrOutput)
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

type ImageDiskReference struct {
	Id  string `pulumi:"id"`
	Lun *int   `pulumi:"lun"`
}





type ImageDiskReferenceInput interface {
	pulumi.Input

	ToImageDiskReferenceOutput() ImageDiskReferenceOutput
	ToImageDiskReferenceOutputWithContext(context.Context) ImageDiskReferenceOutput
}

type ImageDiskReferenceArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Lun pulumi.IntPtrInput `pulumi:"lun"`
}

func (ImageDiskReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReference)(nil)).Elem()
}

func (i ImageDiskReferenceArgs) ToImageDiskReferenceOutput() ImageDiskReferenceOutput {
	return i.ToImageDiskReferenceOutputWithContext(context.Background())
}

func (i ImageDiskReferenceArgs) ToImageDiskReferenceOutputWithContext(ctx context.Context) ImageDiskReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceOutput)
}

func (i ImageDiskReferenceArgs) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return i.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (i ImageDiskReferenceArgs) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceOutput).ToImageDiskReferencePtrOutputWithContext(ctx)
}









type ImageDiskReferencePtrInput interface {
	pulumi.Input

	ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput
	ToImageDiskReferencePtrOutputWithContext(context.Context) ImageDiskReferencePtrOutput
}

type imageDiskReferencePtrType ImageDiskReferenceArgs

func ImageDiskReferencePtr(v *ImageDiskReferenceArgs) ImageDiskReferencePtrInput {
	return (*imageDiskReferencePtrType)(v)
}

func (*imageDiskReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReference)(nil)).Elem()
}

func (i *imageDiskReferencePtrType) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return i.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (i *imageDiskReferencePtrType) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferencePtrOutput)
}

type ImageDiskReferenceOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReference)(nil)).Elem()
}

func (o ImageDiskReferenceOutput) ToImageDiskReferenceOutput() ImageDiskReferenceOutput {
	return o
}

func (o ImageDiskReferenceOutput) ToImageDiskReferenceOutputWithContext(ctx context.Context) ImageDiskReferenceOutput {
	return o
}

func (o ImageDiskReferenceOutput) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return o.ToImageDiskReferencePtrOutputWithContext(context.Background())
}

func (o ImageDiskReferenceOutput) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageDiskReference) *ImageDiskReference {
		return &v
	}).(ImageDiskReferencePtrOutput)
}

func (o ImageDiskReferenceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ImageDiskReference) string { return v.Id }).(pulumi.StringOutput)
}

func (o ImageDiskReferenceOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDiskReference) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

type ImageDiskReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageDiskReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReference)(nil)).Elem()
}

func (o ImageDiskReferencePtrOutput) ToImageDiskReferencePtrOutput() ImageDiskReferencePtrOutput {
	return o
}

func (o ImageDiskReferencePtrOutput) ToImageDiskReferencePtrOutputWithContext(ctx context.Context) ImageDiskReferencePtrOutput {
	return o
}

func (o ImageDiskReferencePtrOutput) Elem() ImageDiskReferenceOutput {
	return o.ApplyT(func(v *ImageDiskReference) ImageDiskReference {
		if v != nil {
			return *v
		}
		var ret ImageDiskReference
		return ret
	}).(ImageDiskReferenceOutput)
}

func (o ImageDiskReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDiskReference) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageDiskReferencePtrOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageDiskReference) *int {
		if v == nil {
			return nil
		}
		return v.Lun
	}).(pulumi.IntPtrOutput)
}

type ImageDiskReferenceResponse struct {
	Id  string `pulumi:"id"`
	Lun *int   `pulumi:"lun"`
}

type ImageDiskReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReferenceResponse)(nil)).Elem()
}

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponseOutput() ImageDiskReferenceResponseOutput {
	return o
}

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponseOutputWithContext(ctx context.Context) ImageDiskReferenceResponseOutput {
	return o
}

func (o ImageDiskReferenceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ImageDiskReferenceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ImageDiskReferenceResponseOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageDiskReferenceResponse) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

type ImageDiskReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageDiskReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReferenceResponse)(nil)).Elem()
}

func (o ImageDiskReferenceResponsePtrOutput) ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput {
	return o
}

func (o ImageDiskReferenceResponsePtrOutput) ToImageDiskReferenceResponsePtrOutputWithContext(ctx context.Context) ImageDiskReferenceResponsePtrOutput {
	return o
}

func (o ImageDiskReferenceResponsePtrOutput) Elem() ImageDiskReferenceResponseOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) ImageDiskReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageDiskReferenceResponse
		return ret
	}).(ImageDiskReferenceResponseOutput)
}

func (o ImageDiskReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ImageDiskReferenceResponsePtrOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageDiskReferenceResponse) *int {
		if v == nil {
			return nil
		}
		return v.Lun
	}).(pulumi.IntPtrOutput)
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

type KeyForDiskEncryptionSet struct {
	KeyUrl      string       `pulumi:"keyUrl"`
	SourceVault *SourceVault `pulumi:"sourceVault"`
}





type KeyForDiskEncryptionSetInput interface {
	pulumi.Input

	ToKeyForDiskEncryptionSetOutput() KeyForDiskEncryptionSetOutput
	ToKeyForDiskEncryptionSetOutputWithContext(context.Context) KeyForDiskEncryptionSetOutput
}

type KeyForDiskEncryptionSetArgs struct {
	KeyUrl      pulumi.StringInput  `pulumi:"keyUrl"`
	SourceVault SourceVaultPtrInput `pulumi:"sourceVault"`
}

func (KeyForDiskEncryptionSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyForDiskEncryptionSet)(nil)).Elem()
}

func (i KeyForDiskEncryptionSetArgs) ToKeyForDiskEncryptionSetOutput() KeyForDiskEncryptionSetOutput {
	return i.ToKeyForDiskEncryptionSetOutputWithContext(context.Background())
}

func (i KeyForDiskEncryptionSetArgs) ToKeyForDiskEncryptionSetOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyForDiskEncryptionSetOutput)
}

func (i KeyForDiskEncryptionSetArgs) ToKeyForDiskEncryptionSetPtrOutput() KeyForDiskEncryptionSetPtrOutput {
	return i.ToKeyForDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (i KeyForDiskEncryptionSetArgs) ToKeyForDiskEncryptionSetPtrOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyForDiskEncryptionSetOutput).ToKeyForDiskEncryptionSetPtrOutputWithContext(ctx)
}









type KeyForDiskEncryptionSetPtrInput interface {
	pulumi.Input

	ToKeyForDiskEncryptionSetPtrOutput() KeyForDiskEncryptionSetPtrOutput
	ToKeyForDiskEncryptionSetPtrOutputWithContext(context.Context) KeyForDiskEncryptionSetPtrOutput
}

type keyForDiskEncryptionSetPtrType KeyForDiskEncryptionSetArgs

func KeyForDiskEncryptionSetPtr(v *KeyForDiskEncryptionSetArgs) KeyForDiskEncryptionSetPtrInput {
	return (*keyForDiskEncryptionSetPtrType)(v)
}

func (*keyForDiskEncryptionSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyForDiskEncryptionSet)(nil)).Elem()
}

func (i *keyForDiskEncryptionSetPtrType) ToKeyForDiskEncryptionSetPtrOutput() KeyForDiskEncryptionSetPtrOutput {
	return i.ToKeyForDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (i *keyForDiskEncryptionSetPtrType) ToKeyForDiskEncryptionSetPtrOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyForDiskEncryptionSetPtrOutput)
}

type KeyForDiskEncryptionSetOutput struct{ *pulumi.OutputState }

func (KeyForDiskEncryptionSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyForDiskEncryptionSet)(nil)).Elem()
}

func (o KeyForDiskEncryptionSetOutput) ToKeyForDiskEncryptionSetOutput() KeyForDiskEncryptionSetOutput {
	return o
}

func (o KeyForDiskEncryptionSetOutput) ToKeyForDiskEncryptionSetOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetOutput {
	return o
}

func (o KeyForDiskEncryptionSetOutput) ToKeyForDiskEncryptionSetPtrOutput() KeyForDiskEncryptionSetPtrOutput {
	return o.ToKeyForDiskEncryptionSetPtrOutputWithContext(context.Background())
}

func (o KeyForDiskEncryptionSetOutput) ToKeyForDiskEncryptionSetPtrOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyForDiskEncryptionSet) *KeyForDiskEncryptionSet {
		return &v
	}).(KeyForDiskEncryptionSetPtrOutput)
}

func (o KeyForDiskEncryptionSetOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyForDiskEncryptionSet) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyForDiskEncryptionSetOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v KeyForDiskEncryptionSet) *SourceVault { return v.SourceVault }).(SourceVaultPtrOutput)
}

type KeyForDiskEncryptionSetPtrOutput struct{ *pulumi.OutputState }

func (KeyForDiskEncryptionSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyForDiskEncryptionSet)(nil)).Elem()
}

func (o KeyForDiskEncryptionSetPtrOutput) ToKeyForDiskEncryptionSetPtrOutput() KeyForDiskEncryptionSetPtrOutput {
	return o
}

func (o KeyForDiskEncryptionSetPtrOutput) ToKeyForDiskEncryptionSetPtrOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetPtrOutput {
	return o
}

func (o KeyForDiskEncryptionSetPtrOutput) Elem() KeyForDiskEncryptionSetOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSet) KeyForDiskEncryptionSet {
		if v != nil {
			return *v
		}
		var ret KeyForDiskEncryptionSet
		return ret
	}).(KeyForDiskEncryptionSetOutput)
}

func (o KeyForDiskEncryptionSetPtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSet) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyForDiskEncryptionSetPtrOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSet) *SourceVault {
		if v == nil {
			return nil
		}
		return v.SourceVault
	}).(SourceVaultPtrOutput)
}

type KeyForDiskEncryptionSetResponse struct {
	KeyUrl      string               `pulumi:"keyUrl"`
	SourceVault *SourceVaultResponse `pulumi:"sourceVault"`
}

type KeyForDiskEncryptionSetResponseOutput struct{ *pulumi.OutputState }

func (KeyForDiskEncryptionSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyForDiskEncryptionSetResponse)(nil)).Elem()
}

func (o KeyForDiskEncryptionSetResponseOutput) ToKeyForDiskEncryptionSetResponseOutput() KeyForDiskEncryptionSetResponseOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponseOutput) ToKeyForDiskEncryptionSetResponseOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetResponseOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyForDiskEncryptionSetResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyForDiskEncryptionSetResponseOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v KeyForDiskEncryptionSetResponse) *SourceVaultResponse { return v.SourceVault }).(SourceVaultResponsePtrOutput)
}

type KeyForDiskEncryptionSetResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyForDiskEncryptionSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyForDiskEncryptionSetResponse)(nil)).Elem()
}

func (o KeyForDiskEncryptionSetResponsePtrOutput) ToKeyForDiskEncryptionSetResponsePtrOutput() KeyForDiskEncryptionSetResponsePtrOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponsePtrOutput) ToKeyForDiskEncryptionSetResponsePtrOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetResponsePtrOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponsePtrOutput) Elem() KeyForDiskEncryptionSetResponseOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSetResponse) KeyForDiskEncryptionSetResponse {
		if v != nil {
			return *v
		}
		var ret KeyForDiskEncryptionSetResponse
		return ret
	}).(KeyForDiskEncryptionSetResponseOutput)
}

func (o KeyForDiskEncryptionSetResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyForDiskEncryptionSetResponsePtrOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v *KeyForDiskEncryptionSetResponse) *SourceVaultResponse {
		if v == nil {
			return nil
		}
		return v.SourceVault
	}).(SourceVaultResponsePtrOutput)
}

type KeyForDiskEncryptionSetResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyForDiskEncryptionSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyForDiskEncryptionSetResponse)(nil)).Elem()
}

func (o KeyForDiskEncryptionSetResponseArrayOutput) ToKeyForDiskEncryptionSetResponseArrayOutput() KeyForDiskEncryptionSetResponseArrayOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponseArrayOutput) ToKeyForDiskEncryptionSetResponseArrayOutputWithContext(ctx context.Context) KeyForDiskEncryptionSetResponseArrayOutput {
	return o
}

func (o KeyForDiskEncryptionSetResponseArrayOutput) Index(i pulumi.IntInput) KeyForDiskEncryptionSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyForDiskEncryptionSetResponse {
		return vs[0].([]KeyForDiskEncryptionSetResponse)[vs[1].(int)]
	}).(KeyForDiskEncryptionSetResponseOutput)
}

type KeyVaultAndKeyReference struct {
	KeyUrl      string      `pulumi:"keyUrl"`
	SourceVault SourceVault `pulumi:"sourceVault"`
}





type KeyVaultAndKeyReferenceInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput
	ToKeyVaultAndKeyReferenceOutputWithContext(context.Context) KeyVaultAndKeyReferenceOutput
}

type KeyVaultAndKeyReferenceArgs struct {
	KeyUrl      pulumi.StringInput `pulumi:"keyUrl"`
	SourceVault SourceVaultInput   `pulumi:"sourceVault"`
}

func (KeyVaultAndKeyReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReference)(nil)).Elem()
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput {
	return i.ToKeyVaultAndKeyReferenceOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferenceOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceOutput)
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return i.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceArgs) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceOutput).ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx)
}









type KeyVaultAndKeyReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput
	ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Context) KeyVaultAndKeyReferencePtrOutput
}

type keyVaultAndKeyReferencePtrType KeyVaultAndKeyReferenceArgs

func KeyVaultAndKeyReferencePtr(v *KeyVaultAndKeyReferenceArgs) KeyVaultAndKeyReferencePtrInput {
	return (*keyVaultAndKeyReferencePtrType)(v)
}

func (*keyVaultAndKeyReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReference)(nil)).Elem()
}

func (i *keyVaultAndKeyReferencePtrType) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return i.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndKeyReferencePtrType) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferencePtrOutput)
}

type KeyVaultAndKeyReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReference)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferenceOutput() KeyVaultAndKeyReferenceOutput {
	return o
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferenceOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceOutput {
	return o
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return o.ToKeyVaultAndKeyReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndKeyReferenceOutput) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndKeyReference) *KeyVaultAndKeyReference {
		return &v
	}).(KeyVaultAndKeyReferencePtrOutput)
}

func (o KeyVaultAndKeyReferenceOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReference) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndKeyReferenceOutput) SourceVault() SourceVaultOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReference) SourceVault { return v.SourceVault }).(SourceVaultOutput)
}

type KeyVaultAndKeyReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReference)(nil)).Elem()
}

func (o KeyVaultAndKeyReferencePtrOutput) ToKeyVaultAndKeyReferencePtrOutput() KeyVaultAndKeyReferencePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferencePtrOutput) ToKeyVaultAndKeyReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferencePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferencePtrOutput) Elem() KeyVaultAndKeyReferenceOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) KeyVaultAndKeyReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndKeyReference
		return ret
	}).(KeyVaultAndKeyReferenceOutput)
}

func (o KeyVaultAndKeyReferencePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndKeyReferencePtrOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReference) *SourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultPtrOutput)
}

type KeyVaultAndKeyReferenceResponse struct {
	KeyUrl      string              `pulumi:"keyUrl"`
	SourceVault SourceVaultResponse `pulumi:"sourceVault"`
}

type KeyVaultAndKeyReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponseOutput() KeyVaultAndKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponseOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseOutput) KeyUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReferenceResponse) string { return v.KeyUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndKeyReferenceResponseOutput) SourceVault() SourceVaultResponseOutput {
	return o.ApplyT(func(v KeyVaultAndKeyReferenceResponse) SourceVaultResponse { return v.SourceVault }).(SourceVaultResponseOutput)
}

type KeyVaultAndKeyReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) Elem() KeyVaultAndKeyReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) KeyVaultAndKeyReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndKeyReferenceResponse
		return ret
	}).(KeyVaultAndKeyReferenceResponseOutput)
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) KeyUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndKeyReferenceResponsePtrOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultAndKeyReferenceResponse) *SourceVaultResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultResponsePtrOutput)
}

type KeyVaultAndSecretReference struct {
	SecretUrl   string      `pulumi:"secretUrl"`
	SourceVault SourceVault `pulumi:"sourceVault"`
}





type KeyVaultAndSecretReferenceInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput
	ToKeyVaultAndSecretReferenceOutputWithContext(context.Context) KeyVaultAndSecretReferenceOutput
}

type KeyVaultAndSecretReferenceArgs struct {
	SecretUrl   pulumi.StringInput `pulumi:"secretUrl"`
	SourceVault SourceVaultInput   `pulumi:"sourceVault"`
}

func (KeyVaultAndSecretReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReference)(nil)).Elem()
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput {
	return i.ToKeyVaultAndSecretReferenceOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferenceOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceOutput)
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return i.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceArgs) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceOutput).ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx)
}









type KeyVaultAndSecretReferencePtrInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput
	ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Context) KeyVaultAndSecretReferencePtrOutput
}

type keyVaultAndSecretReferencePtrType KeyVaultAndSecretReferenceArgs

func KeyVaultAndSecretReferencePtr(v *KeyVaultAndSecretReferenceArgs) KeyVaultAndSecretReferencePtrInput {
	return (*keyVaultAndSecretReferencePtrType)(v)
}

func (*keyVaultAndSecretReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReference)(nil)).Elem()
}

func (i *keyVaultAndSecretReferencePtrType) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return i.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndSecretReferencePtrType) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferencePtrOutput)
}

type KeyVaultAndSecretReferenceOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReference)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferenceOutput() KeyVaultAndSecretReferenceOutput {
	return o
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferenceOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceOutput {
	return o
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return o.ToKeyVaultAndSecretReferencePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndSecretReferenceOutput) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndSecretReference) *KeyVaultAndSecretReference {
		return &v
	}).(KeyVaultAndSecretReferencePtrOutput)
}

func (o KeyVaultAndSecretReferenceOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReference) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndSecretReferenceOutput) SourceVault() SourceVaultOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReference) SourceVault { return v.SourceVault }).(SourceVaultOutput)
}

type KeyVaultAndSecretReferencePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReference)(nil)).Elem()
}

func (o KeyVaultAndSecretReferencePtrOutput) ToKeyVaultAndSecretReferencePtrOutput() KeyVaultAndSecretReferencePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferencePtrOutput) ToKeyVaultAndSecretReferencePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferencePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferencePtrOutput) Elem() KeyVaultAndSecretReferenceOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) KeyVaultAndSecretReference {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndSecretReference
		return ret
	}).(KeyVaultAndSecretReferenceOutput)
}

func (o KeyVaultAndSecretReferencePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndSecretReferencePtrOutput) SourceVault() SourceVaultPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReference) *SourceVault {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultPtrOutput)
}

type KeyVaultAndSecretReferenceResponse struct {
	SecretUrl   string              `pulumi:"secretUrl"`
	SourceVault SourceVaultResponse `pulumi:"sourceVault"`
}

type KeyVaultAndSecretReferenceResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponseOutput() KeyVaultAndSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponseOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponseOutput) SecretUrl() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReferenceResponse) string { return v.SecretUrl }).(pulumi.StringOutput)
}

func (o KeyVaultAndSecretReferenceResponseOutput) SourceVault() SourceVaultResponseOutput {
	return o.ApplyT(func(v KeyVaultAndSecretReferenceResponse) SourceVaultResponse { return v.SourceVault }).(SourceVaultResponseOutput)
}

type KeyVaultAndSecretReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultAndSecretReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponsePtrOutput {
	return o
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) Elem() KeyVaultAndSecretReferenceResponseOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) KeyVaultAndSecretReferenceResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultAndSecretReferenceResponse
		return ret
	}).(KeyVaultAndSecretReferenceResponseOutput)
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) SecretUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecretUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultAndSecretReferenceResponsePtrOutput) SourceVault() SourceVaultResponsePtrOutput {
	return o.ApplyT(func(v *KeyVaultAndSecretReferenceResponse) *SourceVaultResponse {
		if v == nil {
			return nil
		}
		return &v.SourceVault
	}).(SourceVaultResponsePtrOutput)
}

type OSDiskImageEncryption struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
}





type OSDiskImageEncryptionInput interface {
	pulumi.Input

	ToOSDiskImageEncryptionOutput() OSDiskImageEncryptionOutput
	ToOSDiskImageEncryptionOutputWithContext(context.Context) OSDiskImageEncryptionOutput
}

type OSDiskImageEncryptionArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
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

type OSDiskImageEncryptionResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
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

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PurchasePlan struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
}





type PurchasePlanInput interface {
	pulumi.Input

	ToPurchasePlanOutput() PurchasePlanOutput
	ToPurchasePlanOutputWithContext(context.Context) PurchasePlanOutput
}

type PurchasePlanArgs struct {
	Name          pulumi.StringInput    `pulumi:"name"`
	Product       pulumi.StringInput    `pulumi:"product"`
	PromotionCode pulumi.StringPtrInput `pulumi:"promotionCode"`
	Publisher     pulumi.StringInput    `pulumi:"publisher"`
}

func (PurchasePlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlan)(nil)).Elem()
}

func (i PurchasePlanArgs) ToPurchasePlanOutput() PurchasePlanOutput {
	return i.ToPurchasePlanOutputWithContext(context.Background())
}

func (i PurchasePlanArgs) ToPurchasePlanOutputWithContext(ctx context.Context) PurchasePlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanOutput)
}

func (i PurchasePlanArgs) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return i.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (i PurchasePlanArgs) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanOutput).ToPurchasePlanPtrOutputWithContext(ctx)
}









type PurchasePlanPtrInput interface {
	pulumi.Input

	ToPurchasePlanPtrOutput() PurchasePlanPtrOutput
	ToPurchasePlanPtrOutputWithContext(context.Context) PurchasePlanPtrOutput
}

type purchasePlanPtrType PurchasePlanArgs

func PurchasePlanPtr(v *PurchasePlanArgs) PurchasePlanPtrInput {
	return (*purchasePlanPtrType)(v)
}

func (*purchasePlanPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlan)(nil)).Elem()
}

func (i *purchasePlanPtrType) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return i.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (i *purchasePlanPtrType) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PurchasePlanPtrOutput)
}

type PurchasePlanOutput struct{ *pulumi.OutputState }

func (PurchasePlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlan)(nil)).Elem()
}

func (o PurchasePlanOutput) ToPurchasePlanOutput() PurchasePlanOutput {
	return o
}

func (o PurchasePlanOutput) ToPurchasePlanOutputWithContext(ctx context.Context) PurchasePlanOutput {
	return o
}

func (o PurchasePlanOutput) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return o.ToPurchasePlanPtrOutputWithContext(context.Background())
}

func (o PurchasePlanOutput) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PurchasePlan) *PurchasePlan {
		return &v
	}).(PurchasePlanPtrOutput)
}

func (o PurchasePlanOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlan) string { return v.Name }).(pulumi.StringOutput)
}

func (o PurchasePlanOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlan) string { return v.Product }).(pulumi.StringOutput)
}

func (o PurchasePlanOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlan) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlan) string { return v.Publisher }).(pulumi.StringOutput)
}

type PurchasePlanPtrOutput struct{ *pulumi.OutputState }

func (PurchasePlanPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlan)(nil)).Elem()
}

func (o PurchasePlanPtrOutput) ToPurchasePlanPtrOutput() PurchasePlanPtrOutput {
	return o
}

func (o PurchasePlanPtrOutput) ToPurchasePlanPtrOutputWithContext(ctx context.Context) PurchasePlanPtrOutput {
	return o
}

func (o PurchasePlanPtrOutput) Elem() PurchasePlanOutput {
	return o.ApplyT(func(v *PurchasePlan) PurchasePlan {
		if v != nil {
			return *v
		}
		var ret PurchasePlan
		return ret
	}).(PurchasePlanOutput)
}

func (o PurchasePlanPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlan) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

type PurchasePlanResponse struct {
	Name          string  `pulumi:"name"`
	Product       string  `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     string  `pulumi:"publisher"`
}

type PurchasePlanResponseOutput struct{ *pulumi.OutputState }

func (PurchasePlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PurchasePlanResponse)(nil)).Elem()
}

func (o PurchasePlanResponseOutput) ToPurchasePlanResponseOutput() PurchasePlanResponseOutput {
	return o
}

func (o PurchasePlanResponseOutput) ToPurchasePlanResponseOutputWithContext(ctx context.Context) PurchasePlanResponseOutput {
	return o
}

func (o PurchasePlanResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlanResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PurchasePlanResponseOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlanResponse) string { return v.Product }).(pulumi.StringOutput)
}

func (o PurchasePlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PurchasePlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v PurchasePlanResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

type PurchasePlanResponsePtrOutput struct{ *pulumi.OutputState }

func (PurchasePlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PurchasePlanResponse)(nil)).Elem()
}

func (o PurchasePlanResponsePtrOutput) ToPurchasePlanResponsePtrOutput() PurchasePlanResponsePtrOutput {
	return o
}

func (o PurchasePlanResponsePtrOutput) ToPurchasePlanResponsePtrOutputWithContext(ctx context.Context) PurchasePlanResponsePtrOutput {
	return o
}

func (o PurchasePlanResponsePtrOutput) Elem() PurchasePlanResponseOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) PurchasePlanResponse {
		if v != nil {
			return *v
		}
		var ret PurchasePlanResponse
		return ret
	}).(PurchasePlanResponseOutput)
}

func (o PurchasePlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Product
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o PurchasePlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PurchasePlanResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
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

type ShareInfoElementResponse struct {
	VmUri string `pulumi:"vmUri"`
}

type ShareInfoElementResponseOutput struct{ *pulumi.OutputState }

func (ShareInfoElementResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareInfoElementResponse)(nil)).Elem()
}

func (o ShareInfoElementResponseOutput) ToShareInfoElementResponseOutput() ShareInfoElementResponseOutput {
	return o
}

func (o ShareInfoElementResponseOutput) ToShareInfoElementResponseOutputWithContext(ctx context.Context) ShareInfoElementResponseOutput {
	return o
}

func (o ShareInfoElementResponseOutput) VmUri() pulumi.StringOutput {
	return o.ApplyT(func(v ShareInfoElementResponse) string { return v.VmUri }).(pulumi.StringOutput)
}

type ShareInfoElementResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareInfoElementResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareInfoElementResponse)(nil)).Elem()
}

func (o ShareInfoElementResponseArrayOutput) ToShareInfoElementResponseArrayOutput() ShareInfoElementResponseArrayOutput {
	return o
}

func (o ShareInfoElementResponseArrayOutput) ToShareInfoElementResponseArrayOutputWithContext(ctx context.Context) ShareInfoElementResponseArrayOutput {
	return o
}

func (o ShareInfoElementResponseArrayOutput) Index(i pulumi.IntInput) ShareInfoElementResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareInfoElementResponse {
		return vs[0].([]ShareInfoElementResponse)[vs[1].(int)]
	}).(ShareInfoElementResponseOutput)
}

type SharingProfile struct {
	Permissions *string `pulumi:"permissions"`
}





type SharingProfileInput interface {
	pulumi.Input

	ToSharingProfileOutput() SharingProfileOutput
	ToSharingProfileOutputWithContext(context.Context) SharingProfileOutput
}

type SharingProfileArgs struct {
	Permissions pulumi.StringPtrInput `pulumi:"permissions"`
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
	Groups      []SharingProfileGroupResponse `pulumi:"groups"`
	Permissions *string                       `pulumi:"permissions"`
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

type SnapshotSku struct {
	Name *string `pulumi:"name"`
}





type SnapshotSkuInput interface {
	pulumi.Input

	ToSnapshotSkuOutput() SnapshotSkuOutput
	ToSnapshotSkuOutputWithContext(context.Context) SnapshotSkuOutput
}

type SnapshotSkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SnapshotSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSku)(nil)).Elem()
}

func (i SnapshotSkuArgs) ToSnapshotSkuOutput() SnapshotSkuOutput {
	return i.ToSnapshotSkuOutputWithContext(context.Background())
}

func (i SnapshotSkuArgs) ToSnapshotSkuOutputWithContext(ctx context.Context) SnapshotSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuOutput)
}

func (i SnapshotSkuArgs) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return i.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (i SnapshotSkuArgs) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuOutput).ToSnapshotSkuPtrOutputWithContext(ctx)
}









type SnapshotSkuPtrInput interface {
	pulumi.Input

	ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput
	ToSnapshotSkuPtrOutputWithContext(context.Context) SnapshotSkuPtrOutput
}

type snapshotSkuPtrType SnapshotSkuArgs

func SnapshotSkuPtr(v *SnapshotSkuArgs) SnapshotSkuPtrInput {
	return (*snapshotSkuPtrType)(v)
}

func (*snapshotSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSku)(nil)).Elem()
}

func (i *snapshotSkuPtrType) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return i.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (i *snapshotSkuPtrType) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuPtrOutput)
}

type SnapshotSkuOutput struct{ *pulumi.OutputState }

func (SnapshotSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSku)(nil)).Elem()
}

func (o SnapshotSkuOutput) ToSnapshotSkuOutput() SnapshotSkuOutput {
	return o
}

func (o SnapshotSkuOutput) ToSnapshotSkuOutputWithContext(ctx context.Context) SnapshotSkuOutput {
	return o
}

func (o SnapshotSkuOutput) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return o.ToSnapshotSkuPtrOutputWithContext(context.Background())
}

func (o SnapshotSkuOutput) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotSku) *SnapshotSku {
		return &v
	}).(SnapshotSkuPtrOutput)
}

func (o SnapshotSkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotSku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SnapshotSkuPtrOutput struct{ *pulumi.OutputState }

func (SnapshotSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSku)(nil)).Elem()
}

func (o SnapshotSkuPtrOutput) ToSnapshotSkuPtrOutput() SnapshotSkuPtrOutput {
	return o
}

func (o SnapshotSkuPtrOutput) ToSnapshotSkuPtrOutputWithContext(ctx context.Context) SnapshotSkuPtrOutput {
	return o
}

func (o SnapshotSkuPtrOutput) Elem() SnapshotSkuOutput {
	return o.ApplyT(func(v *SnapshotSku) SnapshotSku {
		if v != nil {
			return *v
		}
		var ret SnapshotSku
		return ret
	}).(SnapshotSkuOutput)
}

func (o SnapshotSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SnapshotSkuResponse struct {
	Name *string `pulumi:"name"`
	Tier string  `pulumi:"tier"`
}

type SnapshotSkuResponseOutput struct{ *pulumi.OutputState }

func (SnapshotSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSkuResponse)(nil)).Elem()
}

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponseOutput() SnapshotSkuResponseOutput {
	return o
}

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponseOutputWithContext(ctx context.Context) SnapshotSkuResponseOutput {
	return o
}

func (o SnapshotSkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SnapshotSkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SnapshotSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SnapshotSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SnapshotSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SnapshotSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSkuResponse)(nil)).Elem()
}

func (o SnapshotSkuResponsePtrOutput) ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput {
	return o
}

func (o SnapshotSkuResponsePtrOutput) ToSnapshotSkuResponsePtrOutputWithContext(ctx context.Context) SnapshotSkuResponsePtrOutput {
	return o
}

func (o SnapshotSkuResponsePtrOutput) Elem() SnapshotSkuResponseOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) SnapshotSkuResponse {
		if v != nil {
			return *v
		}
		var ret SnapshotSkuResponse
		return ret
	}).(SnapshotSkuResponseOutput)
}

func (o SnapshotSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SnapshotSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnapshotSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SourceVault struct {
	Id *string `pulumi:"id"`
}





type SourceVaultInput interface {
	pulumi.Input

	ToSourceVaultOutput() SourceVaultOutput
	ToSourceVaultOutputWithContext(context.Context) SourceVaultOutput
}

type SourceVaultArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SourceVaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVault)(nil)).Elem()
}

func (i SourceVaultArgs) ToSourceVaultOutput() SourceVaultOutput {
	return i.ToSourceVaultOutputWithContext(context.Background())
}

func (i SourceVaultArgs) ToSourceVaultOutputWithContext(ctx context.Context) SourceVaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultOutput)
}

func (i SourceVaultArgs) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return i.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (i SourceVaultArgs) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultOutput).ToSourceVaultPtrOutputWithContext(ctx)
}









type SourceVaultPtrInput interface {
	pulumi.Input

	ToSourceVaultPtrOutput() SourceVaultPtrOutput
	ToSourceVaultPtrOutputWithContext(context.Context) SourceVaultPtrOutput
}

type sourceVaultPtrType SourceVaultArgs

func SourceVaultPtr(v *SourceVaultArgs) SourceVaultPtrInput {
	return (*sourceVaultPtrType)(v)
}

func (*sourceVaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVault)(nil)).Elem()
}

func (i *sourceVaultPtrType) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return i.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (i *sourceVaultPtrType) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultPtrOutput)
}

type SourceVaultOutput struct{ *pulumi.OutputState }

func (SourceVaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVault)(nil)).Elem()
}

func (o SourceVaultOutput) ToSourceVaultOutput() SourceVaultOutput {
	return o
}

func (o SourceVaultOutput) ToSourceVaultOutputWithContext(ctx context.Context) SourceVaultOutput {
	return o
}

func (o SourceVaultOutput) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return o.ToSourceVaultPtrOutputWithContext(context.Background())
}

func (o SourceVaultOutput) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceVault) *SourceVault {
		return &v
	}).(SourceVaultPtrOutput)
}

func (o SourceVaultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceVault) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SourceVaultPtrOutput struct{ *pulumi.OutputState }

func (SourceVaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVault)(nil)).Elem()
}

func (o SourceVaultPtrOutput) ToSourceVaultPtrOutput() SourceVaultPtrOutput {
	return o
}

func (o SourceVaultPtrOutput) ToSourceVaultPtrOutputWithContext(ctx context.Context) SourceVaultPtrOutput {
	return o
}

func (o SourceVaultPtrOutput) Elem() SourceVaultOutput {
	return o.ApplyT(func(v *SourceVault) SourceVault {
		if v != nil {
			return *v
		}
		var ret SourceVault
		return ret
	}).(SourceVaultOutput)
}

func (o SourceVaultPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceVault) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SourceVaultResponse struct {
	Id *string `pulumi:"id"`
}

type SourceVaultResponseOutput struct{ *pulumi.OutputState }

func (SourceVaultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVaultResponse)(nil)).Elem()
}

func (o SourceVaultResponseOutput) ToSourceVaultResponseOutput() SourceVaultResponseOutput {
	return o
}

func (o SourceVaultResponseOutput) ToSourceVaultResponseOutputWithContext(ctx context.Context) SourceVaultResponseOutput {
	return o
}

func (o SourceVaultResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceVaultResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SourceVaultResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceVaultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVaultResponse)(nil)).Elem()
}

func (o SourceVaultResponsePtrOutput) ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput {
	return o
}

func (o SourceVaultResponsePtrOutput) ToSourceVaultResponsePtrOutputWithContext(ctx context.Context) SourceVaultResponsePtrOutput {
	return o
}

func (o SourceVaultResponsePtrOutput) Elem() SourceVaultResponseOutput {
	return o.ApplyT(func(v *SourceVaultResponse) SourceVaultResponse {
		if v != nil {
			return *v
		}
		var ret SourceVaultResponse
		return ret
	}).(SourceVaultResponseOutput)
}

func (o SourceVaultResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceVaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(CreationDataOutput{})
	pulumi.RegisterOutputType(CreationDataResponseOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionArrayOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionResponseOutput{})
	pulumi.RegisterOutputType(DataDiskImageEncryptionResponseArrayOutput{})
	pulumi.RegisterOutputType(DisallowedOutput{})
	pulumi.RegisterOutputType(DisallowedPtrOutput{})
	pulumi.RegisterOutputType(DisallowedResponseOutput{})
	pulumi.RegisterOutputType(DisallowedResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskSkuOutput{})
	pulumi.RegisterOutputType(DiskSkuPtrOutput{})
	pulumi.RegisterOutputType(DiskSkuResponseOutput{})
	pulumi.RegisterOutputType(DiskSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionImagesOutput{})
	pulumi.RegisterOutputType(EncryptionImagesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionImagesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionImagesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionSetIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionSetIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionSetIdentityResponseOutput{})
	pulumi.RegisterOutputType(EncryptionSetIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsCollectionOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsCollectionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsCollectionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsCollectionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsElementOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsElementArrayOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsElementResponseOutput{})
	pulumi.RegisterOutputType(EncryptionSettingsElementResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ImageDiskReferenceOutput{})
	pulumi.RegisterOutputType(ImageDiskReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(ImagePurchasePlanResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyForDiskEncryptionSetOutput{})
	pulumi.RegisterOutputType(KeyForDiskEncryptionSetPtrOutput{})
	pulumi.RegisterOutputType(KeyForDiskEncryptionSetResponseOutput{})
	pulumi.RegisterOutputType(KeyForDiskEncryptionSetResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyForDiskEncryptionSetResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionPtrOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionResponseOutput{})
	pulumi.RegisterOutputType(OSDiskImageEncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PurchasePlanOutput{})
	pulumi.RegisterOutputType(PurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponsePtrOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationResponseOutput{})
	pulumi.RegisterOutputType(RecommendedMachineConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(RegionalReplicationStatusResponseOutput{})
	pulumi.RegisterOutputType(RegionalReplicationStatusResponseArrayOutput{})
	pulumi.RegisterOutputType(ReplicationStatusResponseOutput{})
	pulumi.RegisterOutputType(ResourceRangeOutput{})
	pulumi.RegisterOutputType(ResourceRangePtrOutput{})
	pulumi.RegisterOutputType(ResourceRangeResponseOutput{})
	pulumi.RegisterOutputType(ResourceRangeResponsePtrOutput{})
	pulumi.RegisterOutputType(ShareInfoElementResponseOutput{})
	pulumi.RegisterOutputType(ShareInfoElementResponseArrayOutput{})
	pulumi.RegisterOutputType(SharingProfileOutput{})
	pulumi.RegisterOutputType(SharingProfilePtrOutput{})
	pulumi.RegisterOutputType(SharingProfileGroupResponseOutput{})
	pulumi.RegisterOutputType(SharingProfileGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(SharingProfileResponseOutput{})
	pulumi.RegisterOutputType(SharingProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SnapshotSkuOutput{})
	pulumi.RegisterOutputType(SnapshotSkuPtrOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponseOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceVaultOutput{})
	pulumi.RegisterOutputType(SourceVaultPtrOutput{})
	pulumi.RegisterOutputType(SourceVaultResponseOutput{})
	pulumi.RegisterOutputType(SourceVaultResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetRegionOutput{})
	pulumi.RegisterOutputType(TargetRegionArrayOutput{})
	pulumi.RegisterOutputType(TargetRegionResponseOutput{})
	pulumi.RegisterOutputType(TargetRegionResponseArrayOutput{})
	pulumi.RegisterOutputType(UserArtifactManageOutput{})
	pulumi.RegisterOutputType(UserArtifactManagePtrOutput{})
	pulumi.RegisterOutputType(UserArtifactManageResponseOutput{})
	pulumi.RegisterOutputType(UserArtifactManageResponsePtrOutput{})
	pulumi.RegisterOutputType(UserArtifactSourceOutput{})
	pulumi.RegisterOutputType(UserArtifactSourceResponseOutput{})
}
