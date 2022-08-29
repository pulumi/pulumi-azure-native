


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiErrorBaseResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
	Target  *string `pulumi:"target"`
}

type ApiErrorBaseResponseOutput struct{ *pulumi.OutputState }

func (ApiErrorBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiErrorBaseResponse)(nil)).Elem()
}

func (o ApiErrorBaseResponseOutput) ToApiErrorBaseResponseOutput() ApiErrorBaseResponseOutput {
	return o
}

func (o ApiErrorBaseResponseOutput) ToApiErrorBaseResponseOutputWithContext(ctx context.Context) ApiErrorBaseResponseOutput {
	return o
}

func (o ApiErrorBaseResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ApiErrorBaseResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ApiErrorBaseResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorBaseResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ApiErrorBaseResponseArrayOutput struct{ *pulumi.OutputState }

func (ApiErrorBaseResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ApiErrorBaseResponse)(nil)).Elem()
}

func (o ApiErrorBaseResponseArrayOutput) ToApiErrorBaseResponseArrayOutput() ApiErrorBaseResponseArrayOutput {
	return o
}

func (o ApiErrorBaseResponseArrayOutput) ToApiErrorBaseResponseArrayOutputWithContext(ctx context.Context) ApiErrorBaseResponseArrayOutput {
	return o
}

func (o ApiErrorBaseResponseArrayOutput) Index(i pulumi.IntInput) ApiErrorBaseResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ApiErrorBaseResponse {
		return vs[0].([]ApiErrorBaseResponse)[vs[1].(int)]
	}).(ApiErrorBaseResponseOutput)
}

type ApiErrorResponse struct {
	Code       *string                `pulumi:"code"`
	Details    []ApiErrorBaseResponse `pulumi:"details"`
	Innererror *InnerErrorResponse    `pulumi:"innererror"`
	Message    *string                `pulumi:"message"`
	Target     *string                `pulumi:"target"`
}

type ApiErrorResponseOutput struct{ *pulumi.OutputState }

func (ApiErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiErrorResponse)(nil)).Elem()
}

func (o ApiErrorResponseOutput) ToApiErrorResponseOutput() ApiErrorResponseOutput {
	return o
}

func (o ApiErrorResponseOutput) ToApiErrorResponseOutputWithContext(ctx context.Context) ApiErrorResponseOutput {
	return o
}

func (o ApiErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponseOutput) Details() ApiErrorBaseResponseArrayOutput {
	return o.ApplyT(func(v ApiErrorResponse) []ApiErrorBaseResponse { return v.Details }).(ApiErrorBaseResponseArrayOutput)
}

func (o ApiErrorResponseOutput) Innererror() InnerErrorResponsePtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *InnerErrorResponse { return v.Innererror }).(InnerErrorResponsePtrOutput)
}

func (o ApiErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ApiErrorResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiErrorResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type CreationData struct {
	CreateOption          string              `pulumi:"createOption"`
	GalleryImageReference *ImageDiskReference `pulumi:"galleryImageReference"`
	ImageReference        *ImageDiskReference `pulumi:"imageReference"`
	LogicalSectorSize     *int                `pulumi:"logicalSectorSize"`
	SecurityDataUri       *string             `pulumi:"securityDataUri"`
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
	SecurityDataUri       pulumi.StringPtrInput      `pulumi:"securityDataUri"`
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

func (o CreationDataOutput) SecurityDataUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationData) *string { return v.SecurityDataUri }).(pulumi.StringPtrOutput)
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
	SecurityDataUri       *string                     `pulumi:"securityDataUri"`
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

func (o CreationDataResponseOutput) SecurityDataUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreationDataResponse) *string { return v.SecurityDataUri }).(pulumi.StringPtrOutput)
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

type DiskSecurityProfile struct {
	SecureVMDiskEncryptionSetId *string `pulumi:"secureVMDiskEncryptionSetId"`
	SecurityType                *string `pulumi:"securityType"`
}





type DiskSecurityProfileInput interface {
	pulumi.Input

	ToDiskSecurityProfileOutput() DiskSecurityProfileOutput
	ToDiskSecurityProfileOutputWithContext(context.Context) DiskSecurityProfileOutput
}

type DiskSecurityProfileArgs struct {
	SecureVMDiskEncryptionSetId pulumi.StringPtrInput `pulumi:"secureVMDiskEncryptionSetId"`
	SecurityType                pulumi.StringPtrInput `pulumi:"securityType"`
}

func (DiskSecurityProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecurityProfile)(nil)).Elem()
}

func (i DiskSecurityProfileArgs) ToDiskSecurityProfileOutput() DiskSecurityProfileOutput {
	return i.ToDiskSecurityProfileOutputWithContext(context.Background())
}

func (i DiskSecurityProfileArgs) ToDiskSecurityProfileOutputWithContext(ctx context.Context) DiskSecurityProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecurityProfileOutput)
}

func (i DiskSecurityProfileArgs) ToDiskSecurityProfilePtrOutput() DiskSecurityProfilePtrOutput {
	return i.ToDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (i DiskSecurityProfileArgs) ToDiskSecurityProfilePtrOutputWithContext(ctx context.Context) DiskSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecurityProfileOutput).ToDiskSecurityProfilePtrOutputWithContext(ctx)
}









type DiskSecurityProfilePtrInput interface {
	pulumi.Input

	ToDiskSecurityProfilePtrOutput() DiskSecurityProfilePtrOutput
	ToDiskSecurityProfilePtrOutputWithContext(context.Context) DiskSecurityProfilePtrOutput
}

type diskSecurityProfilePtrType DiskSecurityProfileArgs

func DiskSecurityProfilePtr(v *DiskSecurityProfileArgs) DiskSecurityProfilePtrInput {
	return (*diskSecurityProfilePtrType)(v)
}

func (*diskSecurityProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSecurityProfile)(nil)).Elem()
}

func (i *diskSecurityProfilePtrType) ToDiskSecurityProfilePtrOutput() DiskSecurityProfilePtrOutput {
	return i.ToDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (i *diskSecurityProfilePtrType) ToDiskSecurityProfilePtrOutputWithContext(ctx context.Context) DiskSecurityProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskSecurityProfilePtrOutput)
}

type DiskSecurityProfileOutput struct{ *pulumi.OutputState }

func (DiskSecurityProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecurityProfile)(nil)).Elem()
}

func (o DiskSecurityProfileOutput) ToDiskSecurityProfileOutput() DiskSecurityProfileOutput {
	return o
}

func (o DiskSecurityProfileOutput) ToDiskSecurityProfileOutputWithContext(ctx context.Context) DiskSecurityProfileOutput {
	return o
}

func (o DiskSecurityProfileOutput) ToDiskSecurityProfilePtrOutput() DiskSecurityProfilePtrOutput {
	return o.ToDiskSecurityProfilePtrOutputWithContext(context.Background())
}

func (o DiskSecurityProfileOutput) ToDiskSecurityProfilePtrOutputWithContext(ctx context.Context) DiskSecurityProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskSecurityProfile) *DiskSecurityProfile {
		return &v
	}).(DiskSecurityProfilePtrOutput)
}

func (o DiskSecurityProfileOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSecurityProfile) *string { return v.SecureVMDiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o DiskSecurityProfileOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSecurityProfile) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

type DiskSecurityProfilePtrOutput struct{ *pulumi.OutputState }

func (DiskSecurityProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSecurityProfile)(nil)).Elem()
}

func (o DiskSecurityProfilePtrOutput) ToDiskSecurityProfilePtrOutput() DiskSecurityProfilePtrOutput {
	return o
}

func (o DiskSecurityProfilePtrOutput) ToDiskSecurityProfilePtrOutputWithContext(ctx context.Context) DiskSecurityProfilePtrOutput {
	return o
}

func (o DiskSecurityProfilePtrOutput) Elem() DiskSecurityProfileOutput {
	return o.ApplyT(func(v *DiskSecurityProfile) DiskSecurityProfile {
		if v != nil {
			return *v
		}
		var ret DiskSecurityProfile
		return ret
	}).(DiskSecurityProfileOutput)
}

func (o DiskSecurityProfilePtrOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.SecureVMDiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o DiskSecurityProfilePtrOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSecurityProfile) *string {
		if v == nil {
			return nil
		}
		return v.SecurityType
	}).(pulumi.StringPtrOutput)
}

type DiskSecurityProfileResponse struct {
	SecureVMDiskEncryptionSetId *string `pulumi:"secureVMDiskEncryptionSetId"`
	SecurityType                *string `pulumi:"securityType"`
}

type DiskSecurityProfileResponseOutput struct{ *pulumi.OutputState }

func (DiskSecurityProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskSecurityProfileResponse)(nil)).Elem()
}

func (o DiskSecurityProfileResponseOutput) ToDiskSecurityProfileResponseOutput() DiskSecurityProfileResponseOutput {
	return o
}

func (o DiskSecurityProfileResponseOutput) ToDiskSecurityProfileResponseOutputWithContext(ctx context.Context) DiskSecurityProfileResponseOutput {
	return o
}

func (o DiskSecurityProfileResponseOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSecurityProfileResponse) *string { return v.SecureVMDiskEncryptionSetId }).(pulumi.StringPtrOutput)
}

func (o DiskSecurityProfileResponseOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskSecurityProfileResponse) *string { return v.SecurityType }).(pulumi.StringPtrOutput)
}

type DiskSecurityProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (DiskSecurityProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskSecurityProfileResponse)(nil)).Elem()
}

func (o DiskSecurityProfileResponsePtrOutput) ToDiskSecurityProfileResponsePtrOutput() DiskSecurityProfileResponsePtrOutput {
	return o
}

func (o DiskSecurityProfileResponsePtrOutput) ToDiskSecurityProfileResponsePtrOutputWithContext(ctx context.Context) DiskSecurityProfileResponsePtrOutput {
	return o
}

func (o DiskSecurityProfileResponsePtrOutput) Elem() DiskSecurityProfileResponseOutput {
	return o.ApplyT(func(v *DiskSecurityProfileResponse) DiskSecurityProfileResponse {
		if v != nil {
			return *v
		}
		var ret DiskSecurityProfileResponse
		return ret
	}).(DiskSecurityProfileResponseOutput)
}

func (o DiskSecurityProfileResponsePtrOutput) SecureVMDiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecureVMDiskEncryptionSetId
	}).(pulumi.StringPtrOutput)
}

func (o DiskSecurityProfileResponsePtrOutput) SecurityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskSecurityProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecurityType
	}).(pulumi.StringPtrOutput)
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

type InnerErrorResponse struct {
	Errordetail   *string `pulumi:"errordetail"`
	Exceptiontype *string `pulumi:"exceptiontype"`
}

type InnerErrorResponseOutput struct{ *pulumi.OutputState }

func (InnerErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutput() InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) ToInnerErrorResponseOutputWithContext(ctx context.Context) InnerErrorResponseOutput {
	return o
}

func (o InnerErrorResponseOutput) Errordetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *string { return v.Errordetail }).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponseOutput) Exceptiontype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerErrorResponse) *string { return v.Exceptiontype }).(pulumi.StringPtrOutput)
}

type InnerErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (InnerErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InnerErrorResponse)(nil)).Elem()
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutput() InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) ToInnerErrorResponsePtrOutputWithContext(ctx context.Context) InnerErrorResponsePtrOutput {
	return o
}

func (o InnerErrorResponsePtrOutput) Elem() InnerErrorResponseOutput {
	return o.ApplyT(func(v *InnerErrorResponse) InnerErrorResponse {
		if v != nil {
			return *v
		}
		var ret InnerErrorResponse
		return ret
	}).(InnerErrorResponseOutput)
}

func (o InnerErrorResponsePtrOutput) Errordetail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Errordetail
	}).(pulumi.StringPtrOutput)
}

func (o InnerErrorResponsePtrOutput) Exceptiontype() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InnerErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exceptiontype
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

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
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

type PropertyUpdatesInProgressResponse struct {
	TargetTier *string `pulumi:"targetTier"`
}

type PropertyUpdatesInProgressResponseOutput struct{ *pulumi.OutputState }

func (PropertyUpdatesInProgressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PropertyUpdatesInProgressResponse)(nil)).Elem()
}

func (o PropertyUpdatesInProgressResponseOutput) ToPropertyUpdatesInProgressResponseOutput() PropertyUpdatesInProgressResponseOutput {
	return o
}

func (o PropertyUpdatesInProgressResponseOutput) ToPropertyUpdatesInProgressResponseOutputWithContext(ctx context.Context) PropertyUpdatesInProgressResponseOutput {
	return o
}

func (o PropertyUpdatesInProgressResponseOutput) TargetTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PropertyUpdatesInProgressResponse) *string { return v.TargetTier }).(pulumi.StringPtrOutput)
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

type SupportedCapabilities struct {
	AcceleratedNetwork *bool `pulumi:"acceleratedNetwork"`
}





type SupportedCapabilitiesInput interface {
	pulumi.Input

	ToSupportedCapabilitiesOutput() SupportedCapabilitiesOutput
	ToSupportedCapabilitiesOutputWithContext(context.Context) SupportedCapabilitiesOutput
}

type SupportedCapabilitiesArgs struct {
	AcceleratedNetwork pulumi.BoolPtrInput `pulumi:"acceleratedNetwork"`
}

func (SupportedCapabilitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCapabilities)(nil)).Elem()
}

func (i SupportedCapabilitiesArgs) ToSupportedCapabilitiesOutput() SupportedCapabilitiesOutput {
	return i.ToSupportedCapabilitiesOutputWithContext(context.Background())
}

func (i SupportedCapabilitiesArgs) ToSupportedCapabilitiesOutputWithContext(ctx context.Context) SupportedCapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportedCapabilitiesOutput)
}

func (i SupportedCapabilitiesArgs) ToSupportedCapabilitiesPtrOutput() SupportedCapabilitiesPtrOutput {
	return i.ToSupportedCapabilitiesPtrOutputWithContext(context.Background())
}

func (i SupportedCapabilitiesArgs) ToSupportedCapabilitiesPtrOutputWithContext(ctx context.Context) SupportedCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportedCapabilitiesOutput).ToSupportedCapabilitiesPtrOutputWithContext(ctx)
}









type SupportedCapabilitiesPtrInput interface {
	pulumi.Input

	ToSupportedCapabilitiesPtrOutput() SupportedCapabilitiesPtrOutput
	ToSupportedCapabilitiesPtrOutputWithContext(context.Context) SupportedCapabilitiesPtrOutput
}

type supportedCapabilitiesPtrType SupportedCapabilitiesArgs

func SupportedCapabilitiesPtr(v *SupportedCapabilitiesArgs) SupportedCapabilitiesPtrInput {
	return (*supportedCapabilitiesPtrType)(v)
}

func (*supportedCapabilitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedCapabilities)(nil)).Elem()
}

func (i *supportedCapabilitiesPtrType) ToSupportedCapabilitiesPtrOutput() SupportedCapabilitiesPtrOutput {
	return i.ToSupportedCapabilitiesPtrOutputWithContext(context.Background())
}

func (i *supportedCapabilitiesPtrType) ToSupportedCapabilitiesPtrOutputWithContext(ctx context.Context) SupportedCapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SupportedCapabilitiesPtrOutput)
}

type SupportedCapabilitiesOutput struct{ *pulumi.OutputState }

func (SupportedCapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCapabilities)(nil)).Elem()
}

func (o SupportedCapabilitiesOutput) ToSupportedCapabilitiesOutput() SupportedCapabilitiesOutput {
	return o
}

func (o SupportedCapabilitiesOutput) ToSupportedCapabilitiesOutputWithContext(ctx context.Context) SupportedCapabilitiesOutput {
	return o
}

func (o SupportedCapabilitiesOutput) ToSupportedCapabilitiesPtrOutput() SupportedCapabilitiesPtrOutput {
	return o.ToSupportedCapabilitiesPtrOutputWithContext(context.Background())
}

func (o SupportedCapabilitiesOutput) ToSupportedCapabilitiesPtrOutputWithContext(ctx context.Context) SupportedCapabilitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedCapabilities) *SupportedCapabilities {
		return &v
	}).(SupportedCapabilitiesPtrOutput)
}

func (o SupportedCapabilitiesOutput) AcceleratedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SupportedCapabilities) *bool { return v.AcceleratedNetwork }).(pulumi.BoolPtrOutput)
}

type SupportedCapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (SupportedCapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedCapabilities)(nil)).Elem()
}

func (o SupportedCapabilitiesPtrOutput) ToSupportedCapabilitiesPtrOutput() SupportedCapabilitiesPtrOutput {
	return o
}

func (o SupportedCapabilitiesPtrOutput) ToSupportedCapabilitiesPtrOutputWithContext(ctx context.Context) SupportedCapabilitiesPtrOutput {
	return o
}

func (o SupportedCapabilitiesPtrOutput) Elem() SupportedCapabilitiesOutput {
	return o.ApplyT(func(v *SupportedCapabilities) SupportedCapabilities {
		if v != nil {
			return *v
		}
		var ret SupportedCapabilities
		return ret
	}).(SupportedCapabilitiesOutput)
}

func (o SupportedCapabilitiesPtrOutput) AcceleratedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SupportedCapabilities) *bool {
		if v == nil {
			return nil
		}
		return v.AcceleratedNetwork
	}).(pulumi.BoolPtrOutput)
}

type SupportedCapabilitiesResponse struct {
	AcceleratedNetwork *bool `pulumi:"acceleratedNetwork"`
}

type SupportedCapabilitiesResponseOutput struct{ *pulumi.OutputState }

func (SupportedCapabilitiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCapabilitiesResponse)(nil)).Elem()
}

func (o SupportedCapabilitiesResponseOutput) ToSupportedCapabilitiesResponseOutput() SupportedCapabilitiesResponseOutput {
	return o
}

func (o SupportedCapabilitiesResponseOutput) ToSupportedCapabilitiesResponseOutputWithContext(ctx context.Context) SupportedCapabilitiesResponseOutput {
	return o
}

func (o SupportedCapabilitiesResponseOutput) AcceleratedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SupportedCapabilitiesResponse) *bool { return v.AcceleratedNetwork }).(pulumi.BoolPtrOutput)
}

type SupportedCapabilitiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SupportedCapabilitiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedCapabilitiesResponse)(nil)).Elem()
}

func (o SupportedCapabilitiesResponsePtrOutput) ToSupportedCapabilitiesResponsePtrOutput() SupportedCapabilitiesResponsePtrOutput {
	return o
}

func (o SupportedCapabilitiesResponsePtrOutput) ToSupportedCapabilitiesResponsePtrOutputWithContext(ctx context.Context) SupportedCapabilitiesResponsePtrOutput {
	return o
}

func (o SupportedCapabilitiesResponsePtrOutput) Elem() SupportedCapabilitiesResponseOutput {
	return o.ApplyT(func(v *SupportedCapabilitiesResponse) SupportedCapabilitiesResponse {
		if v != nil {
			return *v
		}
		var ret SupportedCapabilitiesResponse
		return ret
	}).(SupportedCapabilitiesResponseOutput)
}

func (o SupportedCapabilitiesResponsePtrOutput) AcceleratedNetwork() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SupportedCapabilitiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AcceleratedNetwork
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiErrorBaseResponseOutput{})
	pulumi.RegisterOutputType(ApiErrorBaseResponseArrayOutput{})
	pulumi.RegisterOutputType(ApiErrorResponseOutput{})
	pulumi.RegisterOutputType(CreationDataOutput{})
	pulumi.RegisterOutputType(CreationDataResponseOutput{})
	pulumi.RegisterOutputType(DiskSecurityProfileOutput{})
	pulumi.RegisterOutputType(DiskSecurityProfilePtrOutput{})
	pulumi.RegisterOutputType(DiskSecurityProfileResponseOutput{})
	pulumi.RegisterOutputType(DiskSecurityProfileResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceOutput{})
	pulumi.RegisterOutputType(ImageDiskReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageDiskReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(InnerErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerErrorResponsePtrOutput{})
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
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PropertyUpdatesInProgressResponseOutput{})
	pulumi.RegisterOutputType(PurchasePlanOutput{})
	pulumi.RegisterOutputType(PurchasePlanPtrOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponseOutput{})
	pulumi.RegisterOutputType(PurchasePlanResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SupportedCapabilitiesOutput{})
	pulumi.RegisterOutputType(SupportedCapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(SupportedCapabilitiesResponseOutput{})
	pulumi.RegisterOutputType(SupportedCapabilitiesResponsePtrOutput{})
}
