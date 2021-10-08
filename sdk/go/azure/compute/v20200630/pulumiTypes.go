


package v20200630

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

func (i CreationDataArgs) ToCreationDataPtrOutput() CreationDataPtrOutput {
	return i.ToCreationDataPtrOutputWithContext(context.Background())
}

func (i CreationDataArgs) ToCreationDataPtrOutputWithContext(ctx context.Context) CreationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataOutput).ToCreationDataPtrOutputWithContext(ctx)
}









type CreationDataPtrInput interface {
	pulumi.Input

	ToCreationDataPtrOutput() CreationDataPtrOutput
	ToCreationDataPtrOutputWithContext(context.Context) CreationDataPtrOutput
}

type creationDataPtrType CreationDataArgs

func CreationDataPtr(v *CreationDataArgs) CreationDataPtrInput {
	return (*creationDataPtrType)(v)
}

func (*creationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreationData)(nil)).Elem()
}

func (i *creationDataPtrType) ToCreationDataPtrOutput() CreationDataPtrOutput {
	return i.ToCreationDataPtrOutputWithContext(context.Background())
}

func (i *creationDataPtrType) ToCreationDataPtrOutputWithContext(ctx context.Context) CreationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataPtrOutput)
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

func (o CreationDataOutput) ToCreationDataPtrOutput() CreationDataPtrOutput {
	return o.ToCreationDataPtrOutputWithContext(context.Background())
}

func (o CreationDataOutput) ToCreationDataPtrOutputWithContext(ctx context.Context) CreationDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreationData) *CreationData {
		return &v
	}).(CreationDataPtrOutput)
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

type CreationDataPtrOutput struct{ *pulumi.OutputState }

func (CreationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreationData)(nil)).Elem()
}

func (o CreationDataPtrOutput) ToCreationDataPtrOutput() CreationDataPtrOutput {
	return o
}

func (o CreationDataPtrOutput) ToCreationDataPtrOutputWithContext(ctx context.Context) CreationDataPtrOutput {
	return o
}

func (o CreationDataPtrOutput) Elem() CreationDataOutput {
	return o.ApplyT(func(v *CreationData) CreationData {
		if v != nil {
			return *v
		}
		var ret CreationData
		return ret
	}).(CreationDataOutput)
}

func (o CreationDataPtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationData) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataPtrOutput) GalleryImageReference() ImageDiskReferencePtrOutput {
	return o.ApplyT(func(v *CreationData) *ImageDiskReference {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(ImageDiskReferencePtrOutput)
}

func (o CreationDataPtrOutput) ImageReference() ImageDiskReferencePtrOutput {
	return o.ApplyT(func(v *CreationData) *ImageDiskReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageDiskReferencePtrOutput)
}

func (o CreationDataPtrOutput) LogicalSectorSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreationData) *int {
		if v == nil {
			return nil
		}
		return v.LogicalSectorSize
	}).(pulumi.IntPtrOutput)
}

func (o CreationDataPtrOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationData) *string {
		if v == nil {
			return nil
		}
		return v.SourceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataPtrOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationData) *string {
		if v == nil {
			return nil
		}
		return v.SourceUri
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataPtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationData) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataPtrOutput) UploadSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CreationData) *float64 {
		if v == nil {
			return nil
		}
		return v.UploadSizeBytes
	}).(pulumi.Float64PtrOutput)
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





type CreationDataResponseInput interface {
	pulumi.Input

	ToCreationDataResponseOutput() CreationDataResponseOutput
	ToCreationDataResponseOutputWithContext(context.Context) CreationDataResponseOutput
}

type CreationDataResponseArgs struct {
	CreateOption          pulumi.StringInput                 `pulumi:"createOption"`
	GalleryImageReference ImageDiskReferenceResponsePtrInput `pulumi:"galleryImageReference"`
	ImageReference        ImageDiskReferenceResponsePtrInput `pulumi:"imageReference"`
	LogicalSectorSize     pulumi.IntPtrInput                 `pulumi:"logicalSectorSize"`
	SourceResourceId      pulumi.StringPtrInput              `pulumi:"sourceResourceId"`
	SourceUniqueId        pulumi.StringInput                 `pulumi:"sourceUniqueId"`
	SourceUri             pulumi.StringPtrInput              `pulumi:"sourceUri"`
	StorageAccountId      pulumi.StringPtrInput              `pulumi:"storageAccountId"`
	UploadSizeBytes       pulumi.Float64PtrInput             `pulumi:"uploadSizeBytes"`
}

func (CreationDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreationDataResponse)(nil)).Elem()
}

func (i CreationDataResponseArgs) ToCreationDataResponseOutput() CreationDataResponseOutput {
	return i.ToCreationDataResponseOutputWithContext(context.Background())
}

func (i CreationDataResponseArgs) ToCreationDataResponseOutputWithContext(ctx context.Context) CreationDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataResponseOutput)
}

func (i CreationDataResponseArgs) ToCreationDataResponsePtrOutput() CreationDataResponsePtrOutput {
	return i.ToCreationDataResponsePtrOutputWithContext(context.Background())
}

func (i CreationDataResponseArgs) ToCreationDataResponsePtrOutputWithContext(ctx context.Context) CreationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataResponseOutput).ToCreationDataResponsePtrOutputWithContext(ctx)
}









type CreationDataResponsePtrInput interface {
	pulumi.Input

	ToCreationDataResponsePtrOutput() CreationDataResponsePtrOutput
	ToCreationDataResponsePtrOutputWithContext(context.Context) CreationDataResponsePtrOutput
}

type creationDataResponsePtrType CreationDataResponseArgs

func CreationDataResponsePtr(v *CreationDataResponseArgs) CreationDataResponsePtrInput {
	return (*creationDataResponsePtrType)(v)
}

func (*creationDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreationDataResponse)(nil)).Elem()
}

func (i *creationDataResponsePtrType) ToCreationDataResponsePtrOutput() CreationDataResponsePtrOutput {
	return i.ToCreationDataResponsePtrOutputWithContext(context.Background())
}

func (i *creationDataResponsePtrType) ToCreationDataResponsePtrOutputWithContext(ctx context.Context) CreationDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreationDataResponsePtrOutput)
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

func (o CreationDataResponseOutput) ToCreationDataResponsePtrOutput() CreationDataResponsePtrOutput {
	return o.ToCreationDataResponsePtrOutputWithContext(context.Background())
}

func (o CreationDataResponseOutput) ToCreationDataResponsePtrOutputWithContext(ctx context.Context) CreationDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreationDataResponse) *CreationDataResponse {
		return &v
	}).(CreationDataResponsePtrOutput)
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

type CreationDataResponsePtrOutput struct{ *pulumi.OutputState }

func (CreationDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreationDataResponse)(nil)).Elem()
}

func (o CreationDataResponsePtrOutput) ToCreationDataResponsePtrOutput() CreationDataResponsePtrOutput {
	return o
}

func (o CreationDataResponsePtrOutput) ToCreationDataResponsePtrOutputWithContext(ctx context.Context) CreationDataResponsePtrOutput {
	return o
}

func (o CreationDataResponsePtrOutput) Elem() CreationDataResponseOutput {
	return o.ApplyT(func(v *CreationDataResponse) CreationDataResponse {
		if v != nil {
			return *v
		}
		var ret CreationDataResponse
		return ret
	}).(CreationDataResponseOutput)
}

func (o CreationDataResponsePtrOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreateOption
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataResponsePtrOutput) GalleryImageReference() ImageDiskReferenceResponsePtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *ImageDiskReferenceResponse {
		if v == nil {
			return nil
		}
		return v.GalleryImageReference
	}).(ImageDiskReferenceResponsePtrOutput)
}

func (o CreationDataResponsePtrOutput) ImageReference() ImageDiskReferenceResponsePtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *ImageDiskReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageDiskReferenceResponsePtrOutput)
}

func (o CreationDataResponsePtrOutput) LogicalSectorSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogicalSectorSize
	}).(pulumi.IntPtrOutput)
}

func (o CreationDataResponsePtrOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataResponsePtrOutput) SourceUniqueId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceUniqueId
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataResponsePtrOutput) SourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceUri
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountId
	}).(pulumi.StringPtrOutput)
}

func (o CreationDataResponsePtrOutput) UploadSizeBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CreationDataResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.UploadSizeBytes
	}).(pulumi.Float64PtrOutput)
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





type DiskSkuResponseInput interface {
	pulumi.Input

	ToDiskSkuResponseOutput() DiskSkuResponseOutput
	ToDiskSkuResponseOutputWithContext(context.Context) DiskSkuResponseOutput
}

type DiskSkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringInput    `pulumi:"tier"`
}

func (DiskSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSkuResponse)(nil)).Elem()
}

func (i DiskSkuResponseArgs) ToDiskSkuResponseOutput() DiskSkuResponseOutput {
	return i.ToDiskSkuResponseOutputWithContext(context.Background())
}

func (i DiskSkuResponseArgs) ToDiskSkuResponseOutputWithContext(ctx context.Context) DiskSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuResponseOutput)
}

func (i DiskSkuResponseArgs) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return i.ToDiskSkuResponsePtrOutputWithContext(context.Background())
}

func (i DiskSkuResponseArgs) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuResponseOutput).ToDiskSkuResponsePtrOutputWithContext(ctx)
}









type DiskSkuResponsePtrInput interface {
	pulumi.Input

	ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput
	ToDiskSkuResponsePtrOutputWithContext(context.Context) DiskSkuResponsePtrOutput
}

type diskSkuResponsePtrType DiskSkuResponseArgs

func DiskSkuResponsePtr(v *DiskSkuResponseArgs) DiskSkuResponsePtrInput {
	return (*diskSkuResponsePtrType)(v)
}

func (*diskSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSkuResponse)(nil)).Elem()
}

func (i *diskSkuResponsePtrType) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return i.ToDiskSkuResponsePtrOutputWithContext(context.Background())
}

func (i *diskSkuResponsePtrType) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSkuResponsePtrOutput)
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

func (o DiskSkuResponseOutput) ToDiskSkuResponsePtrOutput() DiskSkuResponsePtrOutput {
	return o.ToDiskSkuResponsePtrOutputWithContext(context.Background())
}

func (o DiskSkuResponseOutput) ToDiskSkuResponsePtrOutputWithContext(ctx context.Context) DiskSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskSkuResponse) *DiskSkuResponse {
		return &v
	}).(DiskSkuResponsePtrOutput)
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

type EncryptionResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	Type                *string `pulumi:"type"`
}





type EncryptionResponseInput interface {
	pulumi.Input

	ToEncryptionResponseOutput() EncryptionResponseOutput
	ToEncryptionResponseOutputWithContext(context.Context) EncryptionResponseOutput
}

type EncryptionResponseArgs struct {
	DiskEncryptionSetId pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
	Type                pulumi.StringPtrInput `pulumi:"type"`
}

func (EncryptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return i.ToEncryptionResponseOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput)
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionResponseArgs) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponseOutput).ToEncryptionResponsePtrOutputWithContext(ctx)
}









type EncryptionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput
	ToEncryptionResponsePtrOutputWithContext(context.Context) EncryptionResponsePtrOutput
}

type encryptionResponsePtrType EncryptionResponseArgs

func EncryptionResponsePtr(v *EncryptionResponseArgs) EncryptionResponsePtrInput {
	return (*encryptionResponsePtrType)(v)
}

func (*encryptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return i.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionResponsePtrType) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionResponsePtrOutput)
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

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
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





type EncryptionSetIdentityResponseInput interface {
	pulumi.Input

	ToEncryptionSetIdentityResponseOutput() EncryptionSetIdentityResponseOutput
	ToEncryptionSetIdentityResponseOutputWithContext(context.Context) EncryptionSetIdentityResponseOutput
}

type EncryptionSetIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (EncryptionSetIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSetIdentityResponse)(nil)).Elem()
}

func (i EncryptionSetIdentityResponseArgs) ToEncryptionSetIdentityResponseOutput() EncryptionSetIdentityResponseOutput {
	return i.ToEncryptionSetIdentityResponseOutputWithContext(context.Background())
}

func (i EncryptionSetIdentityResponseArgs) ToEncryptionSetIdentityResponseOutputWithContext(ctx context.Context) EncryptionSetIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityResponseOutput)
}

func (i EncryptionSetIdentityResponseArgs) ToEncryptionSetIdentityResponsePtrOutput() EncryptionSetIdentityResponsePtrOutput {
	return i.ToEncryptionSetIdentityResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionSetIdentityResponseArgs) ToEncryptionSetIdentityResponsePtrOutputWithContext(ctx context.Context) EncryptionSetIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityResponseOutput).ToEncryptionSetIdentityResponsePtrOutputWithContext(ctx)
}









type EncryptionSetIdentityResponsePtrInput interface {
	pulumi.Input

	ToEncryptionSetIdentityResponsePtrOutput() EncryptionSetIdentityResponsePtrOutput
	ToEncryptionSetIdentityResponsePtrOutputWithContext(context.Context) EncryptionSetIdentityResponsePtrOutput
}

type encryptionSetIdentityResponsePtrType EncryptionSetIdentityResponseArgs

func EncryptionSetIdentityResponsePtr(v *EncryptionSetIdentityResponseArgs) EncryptionSetIdentityResponsePtrInput {
	return (*encryptionSetIdentityResponsePtrType)(v)
}

func (*encryptionSetIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSetIdentityResponse)(nil)).Elem()
}

func (i *encryptionSetIdentityResponsePtrType) ToEncryptionSetIdentityResponsePtrOutput() EncryptionSetIdentityResponsePtrOutput {
	return i.ToEncryptionSetIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionSetIdentityResponsePtrType) ToEncryptionSetIdentityResponsePtrOutputWithContext(ctx context.Context) EncryptionSetIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSetIdentityResponsePtrOutput)
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

func (o EncryptionSetIdentityResponseOutput) ToEncryptionSetIdentityResponsePtrOutput() EncryptionSetIdentityResponsePtrOutput {
	return o.ToEncryptionSetIdentityResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionSetIdentityResponseOutput) ToEncryptionSetIdentityResponsePtrOutputWithContext(ctx context.Context) EncryptionSetIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionSetIdentityResponse) *EncryptionSetIdentityResponse {
		return &v
	}).(EncryptionSetIdentityResponsePtrOutput)
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





type EncryptionSettingsCollectionResponseInput interface {
	pulumi.Input

	ToEncryptionSettingsCollectionResponseOutput() EncryptionSettingsCollectionResponseOutput
	ToEncryptionSettingsCollectionResponseOutputWithContext(context.Context) EncryptionSettingsCollectionResponseOutput
}

type EncryptionSettingsCollectionResponseArgs struct {
	Enabled                   pulumi.BoolInput                            `pulumi:"enabled"`
	EncryptionSettings        EncryptionSettingsElementResponseArrayInput `pulumi:"encryptionSettings"`
	EncryptionSettingsVersion pulumi.StringPtrInput                       `pulumi:"encryptionSettingsVersion"`
}

func (EncryptionSettingsCollectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsCollectionResponse)(nil)).Elem()
}

func (i EncryptionSettingsCollectionResponseArgs) ToEncryptionSettingsCollectionResponseOutput() EncryptionSettingsCollectionResponseOutput {
	return i.ToEncryptionSettingsCollectionResponseOutputWithContext(context.Background())
}

func (i EncryptionSettingsCollectionResponseArgs) ToEncryptionSettingsCollectionResponseOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionResponseOutput)
}

func (i EncryptionSettingsCollectionResponseArgs) ToEncryptionSettingsCollectionResponsePtrOutput() EncryptionSettingsCollectionResponsePtrOutput {
	return i.ToEncryptionSettingsCollectionResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionSettingsCollectionResponseArgs) ToEncryptionSettingsCollectionResponsePtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionResponseOutput).ToEncryptionSettingsCollectionResponsePtrOutputWithContext(ctx)
}









type EncryptionSettingsCollectionResponsePtrInput interface {
	pulumi.Input

	ToEncryptionSettingsCollectionResponsePtrOutput() EncryptionSettingsCollectionResponsePtrOutput
	ToEncryptionSettingsCollectionResponsePtrOutputWithContext(context.Context) EncryptionSettingsCollectionResponsePtrOutput
}

type encryptionSettingsCollectionResponsePtrType EncryptionSettingsCollectionResponseArgs

func EncryptionSettingsCollectionResponsePtr(v *EncryptionSettingsCollectionResponseArgs) EncryptionSettingsCollectionResponsePtrInput {
	return (*encryptionSettingsCollectionResponsePtrType)(v)
}

func (*encryptionSettingsCollectionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionSettingsCollectionResponse)(nil)).Elem()
}

func (i *encryptionSettingsCollectionResponsePtrType) ToEncryptionSettingsCollectionResponsePtrOutput() EncryptionSettingsCollectionResponsePtrOutput {
	return i.ToEncryptionSettingsCollectionResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionSettingsCollectionResponsePtrType) ToEncryptionSettingsCollectionResponsePtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsCollectionResponsePtrOutput)
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

func (o EncryptionSettingsCollectionResponseOutput) ToEncryptionSettingsCollectionResponsePtrOutput() EncryptionSettingsCollectionResponsePtrOutput {
	return o.ToEncryptionSettingsCollectionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionSettingsCollectionResponseOutput) ToEncryptionSettingsCollectionResponsePtrOutputWithContext(ctx context.Context) EncryptionSettingsCollectionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionSettingsCollectionResponse) *EncryptionSettingsCollectionResponse {
		return &v
	}).(EncryptionSettingsCollectionResponsePtrOutput)
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





type EncryptionSettingsElementResponseInput interface {
	pulumi.Input

	ToEncryptionSettingsElementResponseOutput() EncryptionSettingsElementResponseOutput
	ToEncryptionSettingsElementResponseOutputWithContext(context.Context) EncryptionSettingsElementResponseOutput
}

type EncryptionSettingsElementResponseArgs struct {
	DiskEncryptionKey KeyVaultAndSecretReferenceResponsePtrInput `pulumi:"diskEncryptionKey"`
	KeyEncryptionKey  KeyVaultAndKeyReferenceResponsePtrInput    `pulumi:"keyEncryptionKey"`
}

func (EncryptionSettingsElementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionSettingsElementResponse)(nil)).Elem()
}

func (i EncryptionSettingsElementResponseArgs) ToEncryptionSettingsElementResponseOutput() EncryptionSettingsElementResponseOutput {
	return i.ToEncryptionSettingsElementResponseOutputWithContext(context.Background())
}

func (i EncryptionSettingsElementResponseArgs) ToEncryptionSettingsElementResponseOutputWithContext(ctx context.Context) EncryptionSettingsElementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsElementResponseOutput)
}





type EncryptionSettingsElementResponseArrayInput interface {
	pulumi.Input

	ToEncryptionSettingsElementResponseArrayOutput() EncryptionSettingsElementResponseArrayOutput
	ToEncryptionSettingsElementResponseArrayOutputWithContext(context.Context) EncryptionSettingsElementResponseArrayOutput
}

type EncryptionSettingsElementResponseArray []EncryptionSettingsElementResponseInput

func (EncryptionSettingsElementResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EncryptionSettingsElementResponse)(nil)).Elem()
}

func (i EncryptionSettingsElementResponseArray) ToEncryptionSettingsElementResponseArrayOutput() EncryptionSettingsElementResponseArrayOutput {
	return i.ToEncryptionSettingsElementResponseArrayOutputWithContext(context.Background())
}

func (i EncryptionSettingsElementResponseArray) ToEncryptionSettingsElementResponseArrayOutputWithContext(ctx context.Context) EncryptionSettingsElementResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionSettingsElementResponseArrayOutput)
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





type ImageDiskReferenceResponseInput interface {
	pulumi.Input

	ToImageDiskReferenceResponseOutput() ImageDiskReferenceResponseOutput
	ToImageDiskReferenceResponseOutputWithContext(context.Context) ImageDiskReferenceResponseOutput
}

type ImageDiskReferenceResponseArgs struct {
	Id  pulumi.StringInput `pulumi:"id"`
	Lun pulumi.IntPtrInput `pulumi:"lun"`
}

func (ImageDiskReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageDiskReferenceResponse)(nil)).Elem()
}

func (i ImageDiskReferenceResponseArgs) ToImageDiskReferenceResponseOutput() ImageDiskReferenceResponseOutput {
	return i.ToImageDiskReferenceResponseOutputWithContext(context.Background())
}

func (i ImageDiskReferenceResponseArgs) ToImageDiskReferenceResponseOutputWithContext(ctx context.Context) ImageDiskReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceResponseOutput)
}

func (i ImageDiskReferenceResponseArgs) ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput {
	return i.ToImageDiskReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ImageDiskReferenceResponseArgs) ToImageDiskReferenceResponsePtrOutputWithContext(ctx context.Context) ImageDiskReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceResponseOutput).ToImageDiskReferenceResponsePtrOutputWithContext(ctx)
}









type ImageDiskReferenceResponsePtrInput interface {
	pulumi.Input

	ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput
	ToImageDiskReferenceResponsePtrOutputWithContext(context.Context) ImageDiskReferenceResponsePtrOutput
}

type imageDiskReferenceResponsePtrType ImageDiskReferenceResponseArgs

func ImageDiskReferenceResponsePtr(v *ImageDiskReferenceResponseArgs) ImageDiskReferenceResponsePtrInput {
	return (*imageDiskReferenceResponsePtrType)(v)
}

func (*imageDiskReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageDiskReferenceResponse)(nil)).Elem()
}

func (i *imageDiskReferenceResponsePtrType) ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput {
	return i.ToImageDiskReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *imageDiskReferenceResponsePtrType) ToImageDiskReferenceResponsePtrOutputWithContext(ctx context.Context) ImageDiskReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageDiskReferenceResponsePtrOutput)
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

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponsePtrOutput() ImageDiskReferenceResponsePtrOutput {
	return o.ToImageDiskReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ImageDiskReferenceResponseOutput) ToImageDiskReferenceResponsePtrOutputWithContext(ctx context.Context) ImageDiskReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageDiskReferenceResponse) *ImageDiskReferenceResponse {
		return &v
	}).(ImageDiskReferenceResponsePtrOutput)
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





type KeyVaultAndKeyReferenceResponseInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferenceResponseOutput() KeyVaultAndKeyReferenceResponseOutput
	ToKeyVaultAndKeyReferenceResponseOutputWithContext(context.Context) KeyVaultAndKeyReferenceResponseOutput
}

type KeyVaultAndKeyReferenceResponseArgs struct {
	KeyUrl      pulumi.StringInput       `pulumi:"keyUrl"`
	SourceVault SourceVaultResponseInput `pulumi:"sourceVault"`
}

func (KeyVaultAndKeyReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (i KeyVaultAndKeyReferenceResponseArgs) ToKeyVaultAndKeyReferenceResponseOutput() KeyVaultAndKeyReferenceResponseOutput {
	return i.ToKeyVaultAndKeyReferenceResponseOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceResponseArgs) ToKeyVaultAndKeyReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceResponseOutput)
}

func (i KeyVaultAndKeyReferenceResponseArgs) ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput {
	return i.ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceResponseArgs) ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceResponseOutput).ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx)
}









type KeyVaultAndKeyReferenceResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput
	ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(context.Context) KeyVaultAndKeyReferenceResponsePtrOutput
}

type keyVaultAndKeyReferenceResponsePtrType KeyVaultAndKeyReferenceResponseArgs

func KeyVaultAndKeyReferenceResponsePtr(v *KeyVaultAndKeyReferenceResponseArgs) KeyVaultAndKeyReferenceResponsePtrInput {
	return (*keyVaultAndKeyReferenceResponsePtrType)(v)
}

func (*keyVaultAndKeyReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (i *keyVaultAndKeyReferenceResponsePtrType) ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput {
	return i.ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndKeyReferenceResponsePtrType) ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceResponsePtrOutput)
}





type KeyVaultAndKeyReferenceResponseArrayInput interface {
	pulumi.Input

	ToKeyVaultAndKeyReferenceResponseArrayOutput() KeyVaultAndKeyReferenceResponseArrayOutput
	ToKeyVaultAndKeyReferenceResponseArrayOutputWithContext(context.Context) KeyVaultAndKeyReferenceResponseArrayOutput
}

type KeyVaultAndKeyReferenceResponseArray []KeyVaultAndKeyReferenceResponseInput

func (KeyVaultAndKeyReferenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (i KeyVaultAndKeyReferenceResponseArray) ToKeyVaultAndKeyReferenceResponseArrayOutput() KeyVaultAndKeyReferenceResponseArrayOutput {
	return i.ToKeyVaultAndKeyReferenceResponseArrayOutputWithContext(context.Background())
}

func (i KeyVaultAndKeyReferenceResponseArray) ToKeyVaultAndKeyReferenceResponseArrayOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndKeyReferenceResponseArrayOutput)
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

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponsePtrOutput() KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndKeyReferenceResponseOutput) ToKeyVaultAndKeyReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndKeyReferenceResponse) *KeyVaultAndKeyReferenceResponse {
		return &v
	}).(KeyVaultAndKeyReferenceResponsePtrOutput)
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

type KeyVaultAndKeyReferenceResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyVaultAndKeyReferenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyVaultAndKeyReferenceResponse)(nil)).Elem()
}

func (o KeyVaultAndKeyReferenceResponseArrayOutput) ToKeyVaultAndKeyReferenceResponseArrayOutput() KeyVaultAndKeyReferenceResponseArrayOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseArrayOutput) ToKeyVaultAndKeyReferenceResponseArrayOutputWithContext(ctx context.Context) KeyVaultAndKeyReferenceResponseArrayOutput {
	return o
}

func (o KeyVaultAndKeyReferenceResponseArrayOutput) Index(i pulumi.IntInput) KeyVaultAndKeyReferenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyVaultAndKeyReferenceResponse {
		return vs[0].([]KeyVaultAndKeyReferenceResponse)[vs[1].(int)]
	}).(KeyVaultAndKeyReferenceResponseOutput)
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





type KeyVaultAndSecretReferenceResponseInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferenceResponseOutput() KeyVaultAndSecretReferenceResponseOutput
	ToKeyVaultAndSecretReferenceResponseOutputWithContext(context.Context) KeyVaultAndSecretReferenceResponseOutput
}

type KeyVaultAndSecretReferenceResponseArgs struct {
	SecretUrl   pulumi.StringInput       `pulumi:"secretUrl"`
	SourceVault SourceVaultResponseInput `pulumi:"sourceVault"`
}

func (KeyVaultAndSecretReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (i KeyVaultAndSecretReferenceResponseArgs) ToKeyVaultAndSecretReferenceResponseOutput() KeyVaultAndSecretReferenceResponseOutput {
	return i.ToKeyVaultAndSecretReferenceResponseOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceResponseArgs) ToKeyVaultAndSecretReferenceResponseOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceResponseOutput)
}

func (i KeyVaultAndSecretReferenceResponseArgs) ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput {
	return i.ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultAndSecretReferenceResponseArgs) ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceResponseOutput).ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx)
}









type KeyVaultAndSecretReferenceResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput
	ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(context.Context) KeyVaultAndSecretReferenceResponsePtrOutput
}

type keyVaultAndSecretReferenceResponsePtrType KeyVaultAndSecretReferenceResponseArgs

func KeyVaultAndSecretReferenceResponsePtr(v *KeyVaultAndSecretReferenceResponseArgs) KeyVaultAndSecretReferenceResponsePtrInput {
	return (*keyVaultAndSecretReferenceResponsePtrType)(v)
}

func (*keyVaultAndSecretReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultAndSecretReferenceResponse)(nil)).Elem()
}

func (i *keyVaultAndSecretReferenceResponsePtrType) ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput {
	return i.ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultAndSecretReferenceResponsePtrType) ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultAndSecretReferenceResponsePtrOutput)
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

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponsePtrOutput() KeyVaultAndSecretReferenceResponsePtrOutput {
	return o.ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultAndSecretReferenceResponseOutput) ToKeyVaultAndSecretReferenceResponsePtrOutputWithContext(ctx context.Context) KeyVaultAndSecretReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultAndSecretReferenceResponse) *KeyVaultAndSecretReferenceResponse {
		return &v
	}).(KeyVaultAndSecretReferenceResponsePtrOutput)
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

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
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





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
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

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
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

type ShareInfoElementResponse struct {
	VmUri string `pulumi:"vmUri"`
}





type ShareInfoElementResponseInput interface {
	pulumi.Input

	ToShareInfoElementResponseOutput() ShareInfoElementResponseOutput
	ToShareInfoElementResponseOutputWithContext(context.Context) ShareInfoElementResponseOutput
}

type ShareInfoElementResponseArgs struct {
	VmUri pulumi.StringInput `pulumi:"vmUri"`
}

func (ShareInfoElementResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareInfoElementResponse)(nil)).Elem()
}

func (i ShareInfoElementResponseArgs) ToShareInfoElementResponseOutput() ShareInfoElementResponseOutput {
	return i.ToShareInfoElementResponseOutputWithContext(context.Background())
}

func (i ShareInfoElementResponseArgs) ToShareInfoElementResponseOutputWithContext(ctx context.Context) ShareInfoElementResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareInfoElementResponseOutput)
}





type ShareInfoElementResponseArrayInput interface {
	pulumi.Input

	ToShareInfoElementResponseArrayOutput() ShareInfoElementResponseArrayOutput
	ToShareInfoElementResponseArrayOutputWithContext(context.Context) ShareInfoElementResponseArrayOutput
}

type ShareInfoElementResponseArray []ShareInfoElementResponseInput

func (ShareInfoElementResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareInfoElementResponse)(nil)).Elem()
}

func (i ShareInfoElementResponseArray) ToShareInfoElementResponseArrayOutput() ShareInfoElementResponseArrayOutput {
	return i.ToShareInfoElementResponseArrayOutputWithContext(context.Background())
}

func (i ShareInfoElementResponseArray) ToShareInfoElementResponseArrayOutputWithContext(ctx context.Context) ShareInfoElementResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareInfoElementResponseArrayOutput)
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





type SnapshotSkuResponseInput interface {
	pulumi.Input

	ToSnapshotSkuResponseOutput() SnapshotSkuResponseOutput
	ToSnapshotSkuResponseOutputWithContext(context.Context) SnapshotSkuResponseOutput
}

type SnapshotSkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Tier pulumi.StringInput    `pulumi:"tier"`
}

func (SnapshotSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SnapshotSkuResponse)(nil)).Elem()
}

func (i SnapshotSkuResponseArgs) ToSnapshotSkuResponseOutput() SnapshotSkuResponseOutput {
	return i.ToSnapshotSkuResponseOutputWithContext(context.Background())
}

func (i SnapshotSkuResponseArgs) ToSnapshotSkuResponseOutputWithContext(ctx context.Context) SnapshotSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuResponseOutput)
}

func (i SnapshotSkuResponseArgs) ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput {
	return i.ToSnapshotSkuResponsePtrOutputWithContext(context.Background())
}

func (i SnapshotSkuResponseArgs) ToSnapshotSkuResponsePtrOutputWithContext(ctx context.Context) SnapshotSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuResponseOutput).ToSnapshotSkuResponsePtrOutputWithContext(ctx)
}









type SnapshotSkuResponsePtrInput interface {
	pulumi.Input

	ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput
	ToSnapshotSkuResponsePtrOutputWithContext(context.Context) SnapshotSkuResponsePtrOutput
}

type snapshotSkuResponsePtrType SnapshotSkuResponseArgs

func SnapshotSkuResponsePtr(v *SnapshotSkuResponseArgs) SnapshotSkuResponsePtrInput {
	return (*snapshotSkuResponsePtrType)(v)
}

func (*snapshotSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SnapshotSkuResponse)(nil)).Elem()
}

func (i *snapshotSkuResponsePtrType) ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput {
	return i.ToSnapshotSkuResponsePtrOutputWithContext(context.Background())
}

func (i *snapshotSkuResponsePtrType) ToSnapshotSkuResponsePtrOutputWithContext(ctx context.Context) SnapshotSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnapshotSkuResponsePtrOutput)
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

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponsePtrOutput() SnapshotSkuResponsePtrOutput {
	return o.ToSnapshotSkuResponsePtrOutputWithContext(context.Background())
}

func (o SnapshotSkuResponseOutput) ToSnapshotSkuResponsePtrOutputWithContext(ctx context.Context) SnapshotSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SnapshotSkuResponse) *SnapshotSkuResponse {
		return &v
	}).(SnapshotSkuResponsePtrOutput)
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





type SourceVaultResponseInput interface {
	pulumi.Input

	ToSourceVaultResponseOutput() SourceVaultResponseOutput
	ToSourceVaultResponseOutputWithContext(context.Context) SourceVaultResponseOutput
}

type SourceVaultResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SourceVaultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceVaultResponse)(nil)).Elem()
}

func (i SourceVaultResponseArgs) ToSourceVaultResponseOutput() SourceVaultResponseOutput {
	return i.ToSourceVaultResponseOutputWithContext(context.Background())
}

func (i SourceVaultResponseArgs) ToSourceVaultResponseOutputWithContext(ctx context.Context) SourceVaultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultResponseOutput)
}

func (i SourceVaultResponseArgs) ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput {
	return i.ToSourceVaultResponsePtrOutputWithContext(context.Background())
}

func (i SourceVaultResponseArgs) ToSourceVaultResponsePtrOutputWithContext(ctx context.Context) SourceVaultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultResponseOutput).ToSourceVaultResponsePtrOutputWithContext(ctx)
}









type SourceVaultResponsePtrInput interface {
	pulumi.Input

	ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput
	ToSourceVaultResponsePtrOutputWithContext(context.Context) SourceVaultResponsePtrOutput
}

type sourceVaultResponsePtrType SourceVaultResponseArgs

func SourceVaultResponsePtr(v *SourceVaultResponseArgs) SourceVaultResponsePtrInput {
	return (*sourceVaultResponsePtrType)(v)
}

func (*sourceVaultResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceVaultResponse)(nil)).Elem()
}

func (i *sourceVaultResponsePtrType) ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput {
	return i.ToSourceVaultResponsePtrOutputWithContext(context.Background())
}

func (i *sourceVaultResponsePtrType) ToSourceVaultResponsePtrOutputWithContext(ctx context.Context) SourceVaultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceVaultResponsePtrOutput)
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

func (o SourceVaultResponseOutput) ToSourceVaultResponsePtrOutput() SourceVaultResponsePtrOutput {
	return o.ToSourceVaultResponsePtrOutputWithContext(context.Background())
}

func (o SourceVaultResponseOutput) ToSourceVaultResponsePtrOutputWithContext(ctx context.Context) SourceVaultResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceVaultResponse) *SourceVaultResponse {
		return &v
	}).(SourceVaultResponsePtrOutput)
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CreationDataInput)(nil)).Elem(), CreationDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreationDataPtrInput)(nil)).Elem(), CreationDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreationDataResponseInput)(nil)).Elem(), CreationDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreationDataResponsePtrInput)(nil)).Elem(), CreationDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskSkuInput)(nil)).Elem(), DiskSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskSkuPtrInput)(nil)).Elem(), DiskSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskSkuResponseInput)(nil)).Elem(), DiskSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskSkuResponsePtrInput)(nil)).Elem(), DiskSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionInput)(nil)).Elem(), EncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionPtrInput)(nil)).Elem(), EncryptionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionResponseInput)(nil)).Elem(), EncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionResponsePtrInput)(nil)).Elem(), EncryptionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSetIdentityInput)(nil)).Elem(), EncryptionSetIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSetIdentityPtrInput)(nil)).Elem(), EncryptionSetIdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSetIdentityResponseInput)(nil)).Elem(), EncryptionSetIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSetIdentityResponsePtrInput)(nil)).Elem(), EncryptionSetIdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsCollectionInput)(nil)).Elem(), EncryptionSettingsCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsCollectionPtrInput)(nil)).Elem(), EncryptionSettingsCollectionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsCollectionResponseInput)(nil)).Elem(), EncryptionSettingsCollectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsCollectionResponsePtrInput)(nil)).Elem(), EncryptionSettingsCollectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsElementInput)(nil)).Elem(), EncryptionSettingsElementArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsElementArrayInput)(nil)).Elem(), EncryptionSettingsElementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsElementResponseInput)(nil)).Elem(), EncryptionSettingsElementResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionSettingsElementResponseArrayInput)(nil)).Elem(), EncryptionSettingsElementResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageDiskReferenceInput)(nil)).Elem(), ImageDiskReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageDiskReferencePtrInput)(nil)).Elem(), ImageDiskReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageDiskReferenceResponseInput)(nil)).Elem(), ImageDiskReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageDiskReferenceResponsePtrInput)(nil)).Elem(), ImageDiskReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndKeyReferenceInput)(nil)).Elem(), KeyVaultAndKeyReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndKeyReferencePtrInput)(nil)).Elem(), KeyVaultAndKeyReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndKeyReferenceResponseInput)(nil)).Elem(), KeyVaultAndKeyReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndKeyReferenceResponsePtrInput)(nil)).Elem(), KeyVaultAndKeyReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndKeyReferenceResponseArrayInput)(nil)).Elem(), KeyVaultAndKeyReferenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndSecretReferenceInput)(nil)).Elem(), KeyVaultAndSecretReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndSecretReferencePtrInput)(nil)).Elem(), KeyVaultAndSecretReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndSecretReferenceResponseInput)(nil)).Elem(), KeyVaultAndSecretReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyVaultAndSecretReferenceResponsePtrInput)(nil)).Elem(), KeyVaultAndSecretReferenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareInfoElementResponseInput)(nil)).Elem(), ShareInfoElementResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareInfoElementResponseArrayInput)(nil)).Elem(), ShareInfoElementResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotSkuInput)(nil)).Elem(), SnapshotSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotSkuPtrInput)(nil)).Elem(), SnapshotSkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotSkuResponseInput)(nil)).Elem(), SnapshotSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnapshotSkuResponsePtrInput)(nil)).Elem(), SnapshotSkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceVaultInput)(nil)).Elem(), SourceVaultArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceVaultPtrInput)(nil)).Elem(), SourceVaultArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceVaultResponseInput)(nil)).Elem(), SourceVaultResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SourceVaultResponsePtrInput)(nil)).Elem(), SourceVaultResponseArgs{})
	pulumi.RegisterOutputType(CreationDataOutput{})
	pulumi.RegisterOutputType(CreationDataPtrOutput{})
	pulumi.RegisterOutputType(CreationDataResponseOutput{})
	pulumi.RegisterOutputType(CreationDataResponsePtrOutput{})
	pulumi.RegisterOutputType(DiskSkuOutput{})
	pulumi.RegisterOutputType(DiskSkuPtrOutput{})
	pulumi.RegisterOutputType(DiskSkuResponseOutput{})
	pulumi.RegisterOutputType(DiskSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
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
	pulumi.RegisterOutputType(ImageDiskReferenceOutput{})
	pulumi.RegisterOutputType(ImageDiskReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndKeyReferenceResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferencePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultAndSecretReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ShareInfoElementResponseOutput{})
	pulumi.RegisterOutputType(ShareInfoElementResponseArrayOutput{})
	pulumi.RegisterOutputType(SnapshotSkuOutput{})
	pulumi.RegisterOutputType(SnapshotSkuPtrOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponseOutput{})
	pulumi.RegisterOutputType(SnapshotSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceVaultOutput{})
	pulumi.RegisterOutputType(SourceVaultPtrOutput{})
	pulumi.RegisterOutputType(SourceVaultResponseOutput{})
	pulumi.RegisterOutputType(SourceVaultResponsePtrOutput{})
}
