


package hybridnetwork

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomProfile struct {
	MetadataConfigurationPath *string `pulumi:"metadataConfigurationPath"`
}





type CustomProfileInput interface {
	pulumi.Input

	ToCustomProfileOutput() CustomProfileOutput
	ToCustomProfileOutputWithContext(context.Context) CustomProfileOutput
}

type CustomProfileArgs struct {
	MetadataConfigurationPath pulumi.StringPtrInput `pulumi:"metadataConfigurationPath"`
}

func (CustomProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProfile)(nil)).Elem()
}

func (i CustomProfileArgs) ToCustomProfileOutput() CustomProfileOutput {
	return i.ToCustomProfileOutputWithContext(context.Background())
}

func (i CustomProfileArgs) ToCustomProfileOutputWithContext(ctx context.Context) CustomProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfileOutput)
}

func (i CustomProfileArgs) ToCustomProfilePtrOutput() CustomProfilePtrOutput {
	return i.ToCustomProfilePtrOutputWithContext(context.Background())
}

func (i CustomProfileArgs) ToCustomProfilePtrOutputWithContext(ctx context.Context) CustomProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfileOutput).ToCustomProfilePtrOutputWithContext(ctx)
}









type CustomProfilePtrInput interface {
	pulumi.Input

	ToCustomProfilePtrOutput() CustomProfilePtrOutput
	ToCustomProfilePtrOutputWithContext(context.Context) CustomProfilePtrOutput
}

type customProfilePtrType CustomProfileArgs

func CustomProfilePtr(v *CustomProfileArgs) CustomProfilePtrInput {
	return (*customProfilePtrType)(v)
}

func (*customProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomProfile)(nil)).Elem()
}

func (i *customProfilePtrType) ToCustomProfilePtrOutput() CustomProfilePtrOutput {
	return i.ToCustomProfilePtrOutputWithContext(context.Background())
}

func (i *customProfilePtrType) ToCustomProfilePtrOutputWithContext(ctx context.Context) CustomProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfilePtrOutput)
}

type CustomProfileOutput struct{ *pulumi.OutputState }

func (CustomProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProfile)(nil)).Elem()
}

func (o CustomProfileOutput) ToCustomProfileOutput() CustomProfileOutput {
	return o
}

func (o CustomProfileOutput) ToCustomProfileOutputWithContext(ctx context.Context) CustomProfileOutput {
	return o
}

func (o CustomProfileOutput) ToCustomProfilePtrOutput() CustomProfilePtrOutput {
	return o.ToCustomProfilePtrOutputWithContext(context.Background())
}

func (o CustomProfileOutput) ToCustomProfilePtrOutputWithContext(ctx context.Context) CustomProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomProfile) *CustomProfile {
		return &v
	}).(CustomProfilePtrOutput)
}

func (o CustomProfileOutput) MetadataConfigurationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomProfile) *string { return v.MetadataConfigurationPath }).(pulumi.StringPtrOutput)
}

type CustomProfilePtrOutput struct{ *pulumi.OutputState }

func (CustomProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomProfile)(nil)).Elem()
}

func (o CustomProfilePtrOutput) ToCustomProfilePtrOutput() CustomProfilePtrOutput {
	return o
}

func (o CustomProfilePtrOutput) ToCustomProfilePtrOutputWithContext(ctx context.Context) CustomProfilePtrOutput {
	return o
}

func (o CustomProfilePtrOutput) Elem() CustomProfileOutput {
	return o.ApplyT(func(v *CustomProfile) CustomProfile {
		if v != nil {
			return *v
		}
		var ret CustomProfile
		return ret
	}).(CustomProfileOutput)
}

func (o CustomProfilePtrOutput) MetadataConfigurationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomProfile) *string {
		if v == nil {
			return nil
		}
		return v.MetadataConfigurationPath
	}).(pulumi.StringPtrOutput)
}

type CustomProfileResponse struct {
	MetadataConfigurationPath *string `pulumi:"metadataConfigurationPath"`
}





type CustomProfileResponseInput interface {
	pulumi.Input

	ToCustomProfileResponseOutput() CustomProfileResponseOutput
	ToCustomProfileResponseOutputWithContext(context.Context) CustomProfileResponseOutput
}

type CustomProfileResponseArgs struct {
	MetadataConfigurationPath pulumi.StringPtrInput `pulumi:"metadataConfigurationPath"`
}

func (CustomProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProfileResponse)(nil)).Elem()
}

func (i CustomProfileResponseArgs) ToCustomProfileResponseOutput() CustomProfileResponseOutput {
	return i.ToCustomProfileResponseOutputWithContext(context.Background())
}

func (i CustomProfileResponseArgs) ToCustomProfileResponseOutputWithContext(ctx context.Context) CustomProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfileResponseOutput)
}

func (i CustomProfileResponseArgs) ToCustomProfileResponsePtrOutput() CustomProfileResponsePtrOutput {
	return i.ToCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (i CustomProfileResponseArgs) ToCustomProfileResponsePtrOutputWithContext(ctx context.Context) CustomProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfileResponseOutput).ToCustomProfileResponsePtrOutputWithContext(ctx)
}









type CustomProfileResponsePtrInput interface {
	pulumi.Input

	ToCustomProfileResponsePtrOutput() CustomProfileResponsePtrOutput
	ToCustomProfileResponsePtrOutputWithContext(context.Context) CustomProfileResponsePtrOutput
}

type customProfileResponsePtrType CustomProfileResponseArgs

func CustomProfileResponsePtr(v *CustomProfileResponseArgs) CustomProfileResponsePtrInput {
	return (*customProfileResponsePtrType)(v)
}

func (*customProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomProfileResponse)(nil)).Elem()
}

func (i *customProfileResponsePtrType) ToCustomProfileResponsePtrOutput() CustomProfileResponsePtrOutput {
	return i.ToCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (i *customProfileResponsePtrType) ToCustomProfileResponsePtrOutputWithContext(ctx context.Context) CustomProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomProfileResponsePtrOutput)
}

type CustomProfileResponseOutput struct{ *pulumi.OutputState }

func (CustomProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomProfileResponse)(nil)).Elem()
}

func (o CustomProfileResponseOutput) ToCustomProfileResponseOutput() CustomProfileResponseOutput {
	return o
}

func (o CustomProfileResponseOutput) ToCustomProfileResponseOutputWithContext(ctx context.Context) CustomProfileResponseOutput {
	return o
}

func (o CustomProfileResponseOutput) ToCustomProfileResponsePtrOutput() CustomProfileResponsePtrOutput {
	return o.ToCustomProfileResponsePtrOutputWithContext(context.Background())
}

func (o CustomProfileResponseOutput) ToCustomProfileResponsePtrOutputWithContext(ctx context.Context) CustomProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomProfileResponse) *CustomProfileResponse {
		return &v
	}).(CustomProfileResponsePtrOutput)
}

func (o CustomProfileResponseOutput) MetadataConfigurationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomProfileResponse) *string { return v.MetadataConfigurationPath }).(pulumi.StringPtrOutput)
}

type CustomProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomProfileResponse)(nil)).Elem()
}

func (o CustomProfileResponsePtrOutput) ToCustomProfileResponsePtrOutput() CustomProfileResponsePtrOutput {
	return o
}

func (o CustomProfileResponsePtrOutput) ToCustomProfileResponsePtrOutputWithContext(ctx context.Context) CustomProfileResponsePtrOutput {
	return o
}

func (o CustomProfileResponsePtrOutput) Elem() CustomProfileResponseOutput {
	return o.ApplyT(func(v *CustomProfileResponse) CustomProfileResponse {
		if v != nil {
			return *v
		}
		var ret CustomProfileResponse
		return ret
	}).(CustomProfileResponseOutput)
}

func (o CustomProfileResponsePtrOutput) MetadataConfigurationPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.MetadataConfigurationPath
	}).(pulumi.StringPtrOutput)
}

type DataDisk struct {
	CreateOption *string `pulumi:"createOption"`
	DiskSizeGB   *int    `pulumi:"diskSizeGB"`
	Name         *string `pulumi:"name"`
}





type DataDiskInput interface {
	pulumi.Input

	ToDataDiskOutput() DataDiskOutput
	ToDataDiskOutputWithContext(context.Context) DataDiskOutput
}

type DataDiskArgs struct {
	CreateOption pulumi.StringPtrInput `pulumi:"createOption"`
	DiskSizeGB   pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (DataDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (i DataDiskArgs) ToDataDiskOutput() DataDiskOutput {
	return i.ToDataDiskOutputWithContext(context.Background())
}

func (i DataDiskArgs) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskOutput)
}





type DataDiskArrayInput interface {
	pulumi.Input

	ToDataDiskArrayOutput() DataDiskArrayOutput
	ToDataDiskArrayOutputWithContext(context.Context) DataDiskArrayOutput
}

type DataDiskArray []DataDiskInput

func (DataDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (i DataDiskArray) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return i.ToDataDiskArrayOutputWithContext(context.Background())
}

func (i DataDiskArray) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskArrayOutput)
}

type DataDiskOutput struct{ *pulumi.OutputState }

func (DataDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDisk)(nil)).Elem()
}

func (o DataDiskOutput) ToDataDiskOutput() DataDiskOutput {
	return o
}

func (o DataDiskOutput) ToDataDiskOutputWithContext(ctx context.Context) DataDiskOutput {
	return o
}

func (o DataDiskOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.CreateOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DataDiskArrayOutput struct{ *pulumi.OutputState }

func (DataDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDisk)(nil)).Elem()
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutput() DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) ToDataDiskArrayOutputWithContext(ctx context.Context) DataDiskArrayOutput {
	return o
}

func (o DataDiskArrayOutput) Index(i pulumi.IntInput) DataDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDisk {
		return vs[0].([]DataDisk)[vs[1].(int)]
	}).(DataDiskOutput)
}

type DataDiskResponse struct {
	CreateOption *string `pulumi:"createOption"`
	DiskSizeGB   *int    `pulumi:"diskSizeGB"`
	Name         *string `pulumi:"name"`
}





type DataDiskResponseInput interface {
	pulumi.Input

	ToDataDiskResponseOutput() DataDiskResponseOutput
	ToDataDiskResponseOutputWithContext(context.Context) DataDiskResponseOutput
}

type DataDiskResponseArgs struct {
	CreateOption pulumi.StringPtrInput `pulumi:"createOption"`
	DiskSizeGB   pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (DataDiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (i DataDiskResponseArgs) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return i.ToDataDiskResponseOutputWithContext(context.Background())
}

func (i DataDiskResponseArgs) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskResponseOutput)
}





type DataDiskResponseArrayInput interface {
	pulumi.Input

	ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput
	ToDataDiskResponseArrayOutputWithContext(context.Context) DataDiskResponseArrayOutput
}

type DataDiskResponseArray []DataDiskResponseInput

func (DataDiskResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (i DataDiskResponseArray) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return i.ToDataDiskResponseArrayOutputWithContext(context.Background())
}

func (i DataDiskResponseArray) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskResponseArrayOutput)
}

type DataDiskResponseOutput struct{ *pulumi.OutputState }

func (DataDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutput() DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) ToDataDiskResponseOutputWithContext(ctx context.Context) DataDiskResponseOutput {
	return o
}

func (o DataDiskResponseOutput) CreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.CreateOption }).(pulumi.StringPtrOutput)
}

func (o DataDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o DataDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DataDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskResponse)(nil)).Elem()
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutput() DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) ToDataDiskResponseArrayOutputWithContext(ctx context.Context) DataDiskResponseArrayOutput {
	return o
}

func (o DataDiskResponseArrayOutput) Index(i pulumi.IntInput) DataDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskResponse {
		return vs[0].([]DataDiskResponse)[vs[1].(int)]
	}).(DataDiskResponseOutput)
}

type ImageReference struct {
	ExactVersion *string `pulumi:"exactVersion"`
	Offer        *string `pulumi:"offer"`
	Publisher    *string `pulumi:"publisher"`
	Sku          *string `pulumi:"sku"`
	Version      *string `pulumi:"version"`
}





type ImageReferenceInput interface {
	pulumi.Input

	ToImageReferenceOutput() ImageReferenceOutput
	ToImageReferenceOutputWithContext(context.Context) ImageReferenceOutput
}

type ImageReferenceArgs struct {
	ExactVersion pulumi.StringPtrInput `pulumi:"exactVersion"`
	Offer        pulumi.StringPtrInput `pulumi:"offer"`
	Publisher    pulumi.StringPtrInput `pulumi:"publisher"`
	Sku          pulumi.StringPtrInput `pulumi:"sku"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (i ImageReferenceArgs) ToImageReferenceOutput() ImageReferenceOutput {
	return i.ToImageReferenceOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput)
}

func (i ImageReferenceArgs) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i ImageReferenceArgs) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceOutput).ToImageReferencePtrOutputWithContext(ctx)
}









type ImageReferencePtrInput interface {
	pulumi.Input

	ToImageReferencePtrOutput() ImageReferencePtrOutput
	ToImageReferencePtrOutputWithContext(context.Context) ImageReferencePtrOutput
}

type imageReferencePtrType ImageReferenceArgs

func ImageReferencePtr(v *ImageReferenceArgs) ImageReferencePtrInput {
	return (*imageReferencePtrType)(v)
}

func (*imageReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (i *imageReferencePtrType) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return i.ToImageReferencePtrOutputWithContext(context.Background())
}

func (i *imageReferencePtrType) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferencePtrOutput)
}

type ImageReferenceOutput struct{ *pulumi.OutputState }

func (ImageReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReference)(nil)).Elem()
}

func (o ImageReferenceOutput) ToImageReferenceOutput() ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferenceOutputWithContext(ctx context.Context) ImageReferenceOutput {
	return o
}

func (o ImageReferenceOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o.ToImageReferencePtrOutputWithContext(context.Background())
}

func (o ImageReferenceOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReference) *ImageReference {
		return &v
	}).(ImageReferencePtrOutput)
}

func (o ImageReferenceOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.ExactVersion }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReference) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferencePtrOutput struct{ *pulumi.OutputState }

func (ImageReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReference)(nil)).Elem()
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutput() ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) ToImageReferencePtrOutputWithContext(ctx context.Context) ImageReferencePtrOutput {
	return o
}

func (o ImageReferencePtrOutput) Elem() ImageReferenceOutput {
	return o.ApplyT(func(v *ImageReference) ImageReference {
		if v != nil {
			return *v
		}
		var ret ImageReference
		return ret
	}).(ImageReferenceOutput)
}

func (o ImageReferencePtrOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.ExactVersion
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferencePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReference) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ImageReferenceResponse struct {
	ExactVersion *string `pulumi:"exactVersion"`
	Offer        *string `pulumi:"offer"`
	Publisher    *string `pulumi:"publisher"`
	Sku          *string `pulumi:"sku"`
	Version      *string `pulumi:"version"`
}





type ImageReferenceResponseInput interface {
	pulumi.Input

	ToImageReferenceResponseOutput() ImageReferenceResponseOutput
	ToImageReferenceResponseOutputWithContext(context.Context) ImageReferenceResponseOutput
}

type ImageReferenceResponseArgs struct {
	ExactVersion pulumi.StringPtrInput `pulumi:"exactVersion"`
	Offer        pulumi.StringPtrInput `pulumi:"offer"`
	Publisher    pulumi.StringPtrInput `pulumi:"publisher"`
	Sku          pulumi.StringPtrInput `pulumi:"sku"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return i.ToImageReferenceResponseOutputWithContext(context.Background())
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponseOutput)
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return i.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ImageReferenceResponseArgs) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponseOutput).ToImageReferenceResponsePtrOutputWithContext(ctx)
}









type ImageReferenceResponsePtrInput interface {
	pulumi.Input

	ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput
	ToImageReferenceResponsePtrOutputWithContext(context.Context) ImageReferenceResponsePtrOutput
}

type imageReferenceResponsePtrType ImageReferenceResponseArgs

func ImageReferenceResponsePtr(v *ImageReferenceResponseArgs) ImageReferenceResponsePtrInput {
	return (*imageReferenceResponsePtrType)(v)
}

func (*imageReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (i *imageReferenceResponsePtrType) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return i.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *imageReferenceResponsePtrType) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageReferenceResponsePtrOutput)
}

type ImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutput() ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponseOutputWithContext(ctx context.Context) ImageReferenceResponseOutput {
	return o
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o.ToImageReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ImageReferenceResponseOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageReferenceResponse) *ImageReferenceResponse {
		return &v
	}).(ImageReferenceResponsePtrOutput)
}

func (o ImageReferenceResponseOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.ExactVersion }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageReferenceResponse)(nil)).Elem()
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutput() ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) ToImageReferenceResponsePtrOutputWithContext(ctx context.Context) ImageReferenceResponsePtrOutput {
	return o
}

func (o ImageReferenceResponsePtrOutput) Elem() ImageReferenceResponseOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) ImageReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ImageReferenceResponse
		return ret
	}).(ImageReferenceResponseOutput)
}

func (o ImageReferenceResponsePtrOutput) ExactVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExactVersion
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Offer
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

func (o ImageReferenceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type LinuxConfiguration struct {
	Ssh *SshConfiguration `pulumi:"ssh"`
}





type LinuxConfigurationInput interface {
	pulumi.Input

	ToLinuxConfigurationOutput() LinuxConfigurationOutput
	ToLinuxConfigurationOutputWithContext(context.Context) LinuxConfigurationOutput
}

type LinuxConfigurationArgs struct {
	Ssh SshConfigurationPtrInput `pulumi:"ssh"`
}

func (LinuxConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return i.ToLinuxConfigurationOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput)
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i LinuxConfigurationArgs) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationOutput).ToLinuxConfigurationPtrOutputWithContext(ctx)
}









type LinuxConfigurationPtrInput interface {
	pulumi.Input

	ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput
	ToLinuxConfigurationPtrOutputWithContext(context.Context) LinuxConfigurationPtrOutput
}

type linuxConfigurationPtrType LinuxConfigurationArgs

func LinuxConfigurationPtr(v *LinuxConfigurationArgs) LinuxConfigurationPtrInput {
	return (*linuxConfigurationPtrType)(v)
}

func (*linuxConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return i.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (i *linuxConfigurationPtrType) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationPtrOutput)
}

type LinuxConfigurationOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutput() LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationOutputWithContext(ctx context.Context) LinuxConfigurationOutput {
	return o
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o.ToLinuxConfigurationPtrOutputWithContext(context.Background())
}

func (o LinuxConfigurationOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxConfiguration) *LinuxConfiguration {
		return &v
	}).(LinuxConfigurationPtrOutput)
}

func (o LinuxConfigurationOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v LinuxConfiguration) *SshConfiguration { return v.Ssh }).(SshConfigurationPtrOutput)
}

type LinuxConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfiguration)(nil)).Elem()
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutput() LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) ToLinuxConfigurationPtrOutputWithContext(ctx context.Context) LinuxConfigurationPtrOutput {
	return o
}

func (o LinuxConfigurationPtrOutput) Elem() LinuxConfigurationOutput {
	return o.ApplyT(func(v *LinuxConfiguration) LinuxConfiguration {
		if v != nil {
			return *v
		}
		var ret LinuxConfiguration
		return ret
	}).(LinuxConfigurationOutput)
}

func (o LinuxConfigurationPtrOutput) Ssh() SshConfigurationPtrOutput {
	return o.ApplyT(func(v *LinuxConfiguration) *SshConfiguration {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationPtrOutput)
}

type LinuxConfigurationResponse struct {
	Ssh *SshConfigurationResponse `pulumi:"ssh"`
}





type LinuxConfigurationResponseInput interface {
	pulumi.Input

	ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput
	ToLinuxConfigurationResponseOutputWithContext(context.Context) LinuxConfigurationResponseOutput
}

type LinuxConfigurationResponseArgs struct {
	Ssh SshConfigurationResponsePtrInput `pulumi:"ssh"`
}

func (LinuxConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfigurationResponse)(nil)).Elem()
}

func (i LinuxConfigurationResponseArgs) ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput {
	return i.ToLinuxConfigurationResponseOutputWithContext(context.Background())
}

func (i LinuxConfigurationResponseArgs) ToLinuxConfigurationResponseOutputWithContext(ctx context.Context) LinuxConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationResponseOutput)
}

func (i LinuxConfigurationResponseArgs) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return i.ToLinuxConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i LinuxConfigurationResponseArgs) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationResponseOutput).ToLinuxConfigurationResponsePtrOutputWithContext(ctx)
}









type LinuxConfigurationResponsePtrInput interface {
	pulumi.Input

	ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput
	ToLinuxConfigurationResponsePtrOutputWithContext(context.Context) LinuxConfigurationResponsePtrOutput
}

type linuxConfigurationResponsePtrType LinuxConfigurationResponseArgs

func LinuxConfigurationResponsePtr(v *LinuxConfigurationResponseArgs) LinuxConfigurationResponsePtrInput {
	return (*linuxConfigurationResponsePtrType)(v)
}

func (*linuxConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfigurationResponse)(nil)).Elem()
}

func (i *linuxConfigurationResponsePtrType) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return i.ToLinuxConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *linuxConfigurationResponsePtrType) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxConfigurationResponsePtrOutput)
}

type LinuxConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutput() LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponseOutputWithContext(ctx context.Context) LinuxConfigurationResponseOutput {
	return o
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return o.ToLinuxConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o LinuxConfigurationResponseOutput) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxConfigurationResponse) *LinuxConfigurationResponse {
		return &v
	}).(LinuxConfigurationResponsePtrOutput)
}

func (o LinuxConfigurationResponseOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LinuxConfigurationResponse) *SshConfigurationResponse { return v.Ssh }).(SshConfigurationResponsePtrOutput)
}

type LinuxConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxConfigurationResponse)(nil)).Elem()
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutput() LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) ToLinuxConfigurationResponsePtrOutputWithContext(ctx context.Context) LinuxConfigurationResponsePtrOutput {
	return o
}

func (o LinuxConfigurationResponsePtrOutput) Elem() LinuxConfigurationResponseOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) LinuxConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LinuxConfigurationResponse
		return ret
	}).(LinuxConfigurationResponseOutput)
}

func (o LinuxConfigurationResponsePtrOutput) Ssh() SshConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *LinuxConfigurationResponse) *SshConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Ssh
	}).(SshConfigurationResponsePtrOutput)
}

type NetworkFunctionRoleConfiguration struct {
	CustomProfile      *CustomProfile     `pulumi:"customProfile"`
	NetworkInterfaces  []NetworkInterface `pulumi:"networkInterfaces"`
	OsProfile          *OsProfile         `pulumi:"osProfile"`
	RoleName           *string            `pulumi:"roleName"`
	RoleType           *string            `pulumi:"roleType"`
	StorageProfile     *StorageProfile    `pulumi:"storageProfile"`
	UserDataParameters interface{}        `pulumi:"userDataParameters"`
	UserDataTemplate   interface{}        `pulumi:"userDataTemplate"`
	VirtualMachineSize *string            `pulumi:"virtualMachineSize"`
}


func (val *NetworkFunctionRoleConfiguration) Defaults() *NetworkFunctionRoleConfiguration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.OsProfile = tmp.OsProfile.Defaults()

	return &tmp
}





type NetworkFunctionRoleConfigurationInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationOutput() NetworkFunctionRoleConfigurationOutput
	ToNetworkFunctionRoleConfigurationOutputWithContext(context.Context) NetworkFunctionRoleConfigurationOutput
}

type NetworkFunctionRoleConfigurationArgs struct {
	CustomProfile      CustomProfilePtrInput      `pulumi:"customProfile"`
	NetworkInterfaces  NetworkInterfaceArrayInput `pulumi:"networkInterfaces"`
	OsProfile          OsProfilePtrInput          `pulumi:"osProfile"`
	RoleName           pulumi.StringPtrInput      `pulumi:"roleName"`
	RoleType           pulumi.StringPtrInput      `pulumi:"roleType"`
	StorageProfile     StorageProfilePtrInput     `pulumi:"storageProfile"`
	UserDataParameters pulumi.Input               `pulumi:"userDataParameters"`
	UserDataTemplate   pulumi.Input               `pulumi:"userDataTemplate"`
	VirtualMachineSize pulumi.StringPtrInput      `pulumi:"virtualMachineSize"`
}

func (NetworkFunctionRoleConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfiguration)(nil)).Elem()
}

func (i NetworkFunctionRoleConfigurationArgs) ToNetworkFunctionRoleConfigurationOutput() NetworkFunctionRoleConfigurationOutput {
	return i.ToNetworkFunctionRoleConfigurationOutputWithContext(context.Background())
}

func (i NetworkFunctionRoleConfigurationArgs) ToNetworkFunctionRoleConfigurationOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionRoleConfigurationOutput)
}





type NetworkFunctionRoleConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationArrayOutput() NetworkFunctionRoleConfigurationArrayOutput
	ToNetworkFunctionRoleConfigurationArrayOutputWithContext(context.Context) NetworkFunctionRoleConfigurationArrayOutput
}

type NetworkFunctionRoleConfigurationArray []NetworkFunctionRoleConfigurationInput

func (NetworkFunctionRoleConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionRoleConfiguration)(nil)).Elem()
}

func (i NetworkFunctionRoleConfigurationArray) ToNetworkFunctionRoleConfigurationArrayOutput() NetworkFunctionRoleConfigurationArrayOutput {
	return i.ToNetworkFunctionRoleConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkFunctionRoleConfigurationArray) ToNetworkFunctionRoleConfigurationArrayOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionRoleConfigurationArrayOutput)
}

type NetworkFunctionRoleConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfiguration)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationOutput) ToNetworkFunctionRoleConfigurationOutput() NetworkFunctionRoleConfigurationOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationOutput) ToNetworkFunctionRoleConfigurationOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationOutput) CustomProfile() CustomProfilePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *CustomProfile { return v.CustomProfile }).(CustomProfilePtrOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) []NetworkInterface { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) OsProfile() OsProfilePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *OsProfile { return v.OsProfile }).(OsProfilePtrOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) RoleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *string { return v.RoleType }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) UserDataParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) interface{} { return v.UserDataParameters }).(pulumi.AnyOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) UserDataTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) interface{} { return v.UserDataTemplate }).(pulumi.AnyOutput)
}

func (o NetworkFunctionRoleConfigurationOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfiguration) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type NetworkFunctionRoleConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionRoleConfiguration)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationArrayOutput) ToNetworkFunctionRoleConfigurationArrayOutput() NetworkFunctionRoleConfigurationArrayOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationArrayOutput) ToNetworkFunctionRoleConfigurationArrayOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationArrayOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkFunctionRoleConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkFunctionRoleConfiguration {
		return vs[0].([]NetworkFunctionRoleConfiguration)[vs[1].(int)]
	}).(NetworkFunctionRoleConfigurationOutput)
}

type NetworkFunctionRoleConfigurationResponse struct {
	CustomProfile      *CustomProfileResponse     `pulumi:"customProfile"`
	NetworkInterfaces  []NetworkInterfaceResponse `pulumi:"networkInterfaces"`
	OsProfile          *OsProfileResponse         `pulumi:"osProfile"`
	RoleName           *string                    `pulumi:"roleName"`
	RoleType           *string                    `pulumi:"roleType"`
	StorageProfile     *StorageProfileResponse    `pulumi:"storageProfile"`
	UserDataParameters interface{}                `pulumi:"userDataParameters"`
	UserDataTemplate   interface{}                `pulumi:"userDataTemplate"`
	VirtualMachineSize *string                    `pulumi:"virtualMachineSize"`
}


func (val *NetworkFunctionRoleConfigurationResponse) Defaults() *NetworkFunctionRoleConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.OsProfile = tmp.OsProfile.Defaults()

	return &tmp
}





type NetworkFunctionRoleConfigurationResponseInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationResponseOutput() NetworkFunctionRoleConfigurationResponseOutput
	ToNetworkFunctionRoleConfigurationResponseOutputWithContext(context.Context) NetworkFunctionRoleConfigurationResponseOutput
}

type NetworkFunctionRoleConfigurationResponseArgs struct {
	CustomProfile      CustomProfileResponsePtrInput      `pulumi:"customProfile"`
	NetworkInterfaces  NetworkInterfaceResponseArrayInput `pulumi:"networkInterfaces"`
	OsProfile          OsProfileResponsePtrInput          `pulumi:"osProfile"`
	RoleName           pulumi.StringPtrInput              `pulumi:"roleName"`
	RoleType           pulumi.StringPtrInput              `pulumi:"roleType"`
	StorageProfile     StorageProfileResponsePtrInput     `pulumi:"storageProfile"`
	UserDataParameters pulumi.Input                       `pulumi:"userDataParameters"`
	UserDataTemplate   pulumi.Input                       `pulumi:"userDataTemplate"`
	VirtualMachineSize pulumi.StringPtrInput              `pulumi:"virtualMachineSize"`
}

func (NetworkFunctionRoleConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfigurationResponse)(nil)).Elem()
}

func (i NetworkFunctionRoleConfigurationResponseArgs) ToNetworkFunctionRoleConfigurationResponseOutput() NetworkFunctionRoleConfigurationResponseOutput {
	return i.ToNetworkFunctionRoleConfigurationResponseOutputWithContext(context.Background())
}

func (i NetworkFunctionRoleConfigurationResponseArgs) ToNetworkFunctionRoleConfigurationResponseOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionRoleConfigurationResponseOutput)
}





type NetworkFunctionRoleConfigurationResponseArrayInput interface {
	pulumi.Input

	ToNetworkFunctionRoleConfigurationResponseArrayOutput() NetworkFunctionRoleConfigurationResponseArrayOutput
	ToNetworkFunctionRoleConfigurationResponseArrayOutputWithContext(context.Context) NetworkFunctionRoleConfigurationResponseArrayOutput
}

type NetworkFunctionRoleConfigurationResponseArray []NetworkFunctionRoleConfigurationResponseInput

func (NetworkFunctionRoleConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionRoleConfigurationResponse)(nil)).Elem()
}

func (i NetworkFunctionRoleConfigurationResponseArray) ToNetworkFunctionRoleConfigurationResponseArrayOutput() NetworkFunctionRoleConfigurationResponseArrayOutput {
	return i.ToNetworkFunctionRoleConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i NetworkFunctionRoleConfigurationResponseArray) ToNetworkFunctionRoleConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionRoleConfigurationResponseArrayOutput)
}

type NetworkFunctionRoleConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionRoleConfigurationResponse)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationResponseOutput) ToNetworkFunctionRoleConfigurationResponseOutput() NetworkFunctionRoleConfigurationResponseOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationResponseOutput) ToNetworkFunctionRoleConfigurationResponseOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationResponseOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationResponseOutput) CustomProfile() CustomProfileResponsePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *CustomProfileResponse { return v.CustomProfile }).(CustomProfileResponsePtrOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) []NetworkInterfaceResponse {
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *OsProfileResponse { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) RoleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *string { return v.RoleType }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) UserDataParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) interface{} { return v.UserDataParameters }).(pulumi.AnyOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) UserDataTemplate() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) interface{} { return v.UserDataTemplate }).(pulumi.AnyOutput)
}

func (o NetworkFunctionRoleConfigurationResponseOutput) VirtualMachineSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionRoleConfigurationResponse) *string { return v.VirtualMachineSize }).(pulumi.StringPtrOutput)
}

type NetworkFunctionRoleConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkFunctionRoleConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionRoleConfigurationResponse)(nil)).Elem()
}

func (o NetworkFunctionRoleConfigurationResponseArrayOutput) ToNetworkFunctionRoleConfigurationResponseArrayOutput() NetworkFunctionRoleConfigurationResponseArrayOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationResponseArrayOutput) ToNetworkFunctionRoleConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkFunctionRoleConfigurationResponseArrayOutput {
	return o
}

func (o NetworkFunctionRoleConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkFunctionRoleConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkFunctionRoleConfigurationResponse {
		return vs[0].([]NetworkFunctionRoleConfigurationResponse)[vs[1].(int)]
	}).(NetworkFunctionRoleConfigurationResponseOutput)
}

type NetworkFunctionTemplate struct {
	NetworkFunctionRoleConfigurations []NetworkFunctionRoleConfiguration `pulumi:"networkFunctionRoleConfigurations"`
}





type NetworkFunctionTemplateInput interface {
	pulumi.Input

	ToNetworkFunctionTemplateOutput() NetworkFunctionTemplateOutput
	ToNetworkFunctionTemplateOutputWithContext(context.Context) NetworkFunctionTemplateOutput
}

type NetworkFunctionTemplateArgs struct {
	NetworkFunctionRoleConfigurations NetworkFunctionRoleConfigurationArrayInput `pulumi:"networkFunctionRoleConfigurations"`
}

func (NetworkFunctionTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionTemplate)(nil)).Elem()
}

func (i NetworkFunctionTemplateArgs) ToNetworkFunctionTemplateOutput() NetworkFunctionTemplateOutput {
	return i.ToNetworkFunctionTemplateOutputWithContext(context.Background())
}

func (i NetworkFunctionTemplateArgs) ToNetworkFunctionTemplateOutputWithContext(ctx context.Context) NetworkFunctionTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplateOutput)
}

func (i NetworkFunctionTemplateArgs) ToNetworkFunctionTemplatePtrOutput() NetworkFunctionTemplatePtrOutput {
	return i.ToNetworkFunctionTemplatePtrOutputWithContext(context.Background())
}

func (i NetworkFunctionTemplateArgs) ToNetworkFunctionTemplatePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplateOutput).ToNetworkFunctionTemplatePtrOutputWithContext(ctx)
}









type NetworkFunctionTemplatePtrInput interface {
	pulumi.Input

	ToNetworkFunctionTemplatePtrOutput() NetworkFunctionTemplatePtrOutput
	ToNetworkFunctionTemplatePtrOutputWithContext(context.Context) NetworkFunctionTemplatePtrOutput
}

type networkFunctionTemplatePtrType NetworkFunctionTemplateArgs

func NetworkFunctionTemplatePtr(v *NetworkFunctionTemplateArgs) NetworkFunctionTemplatePtrInput {
	return (*networkFunctionTemplatePtrType)(v)
}

func (*networkFunctionTemplatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionTemplate)(nil)).Elem()
}

func (i *networkFunctionTemplatePtrType) ToNetworkFunctionTemplatePtrOutput() NetworkFunctionTemplatePtrOutput {
	return i.ToNetworkFunctionTemplatePtrOutputWithContext(context.Background())
}

func (i *networkFunctionTemplatePtrType) ToNetworkFunctionTemplatePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplatePtrOutput)
}

type NetworkFunctionTemplateOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionTemplate)(nil)).Elem()
}

func (o NetworkFunctionTemplateOutput) ToNetworkFunctionTemplateOutput() NetworkFunctionTemplateOutput {
	return o
}

func (o NetworkFunctionTemplateOutput) ToNetworkFunctionTemplateOutputWithContext(ctx context.Context) NetworkFunctionTemplateOutput {
	return o
}

func (o NetworkFunctionTemplateOutput) ToNetworkFunctionTemplatePtrOutput() NetworkFunctionTemplatePtrOutput {
	return o.ToNetworkFunctionTemplatePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionTemplateOutput) ToNetworkFunctionTemplatePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionTemplate) *NetworkFunctionTemplate {
		return &v
	}).(NetworkFunctionTemplatePtrOutput)
}

func (o NetworkFunctionTemplateOutput) NetworkFunctionRoleConfigurations() NetworkFunctionRoleConfigurationArrayOutput {
	return o.ApplyT(func(v NetworkFunctionTemplate) []NetworkFunctionRoleConfiguration {
		return v.NetworkFunctionRoleConfigurations
	}).(NetworkFunctionRoleConfigurationArrayOutput)
}

type NetworkFunctionTemplatePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionTemplate)(nil)).Elem()
}

func (o NetworkFunctionTemplatePtrOutput) ToNetworkFunctionTemplatePtrOutput() NetworkFunctionTemplatePtrOutput {
	return o
}

func (o NetworkFunctionTemplatePtrOutput) ToNetworkFunctionTemplatePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplatePtrOutput {
	return o
}

func (o NetworkFunctionTemplatePtrOutput) Elem() NetworkFunctionTemplateOutput {
	return o.ApplyT(func(v *NetworkFunctionTemplate) NetworkFunctionTemplate {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionTemplate
		return ret
	}).(NetworkFunctionTemplateOutput)
}

func (o NetworkFunctionTemplatePtrOutput) NetworkFunctionRoleConfigurations() NetworkFunctionRoleConfigurationArrayOutput {
	return o.ApplyT(func(v *NetworkFunctionTemplate) []NetworkFunctionRoleConfiguration {
		if v == nil {
			return nil
		}
		return v.NetworkFunctionRoleConfigurations
	}).(NetworkFunctionRoleConfigurationArrayOutput)
}

type NetworkFunctionTemplateResponse struct {
	NetworkFunctionRoleConfigurations []NetworkFunctionRoleConfigurationResponse `pulumi:"networkFunctionRoleConfigurations"`
}





type NetworkFunctionTemplateResponseInput interface {
	pulumi.Input

	ToNetworkFunctionTemplateResponseOutput() NetworkFunctionTemplateResponseOutput
	ToNetworkFunctionTemplateResponseOutputWithContext(context.Context) NetworkFunctionTemplateResponseOutput
}

type NetworkFunctionTemplateResponseArgs struct {
	NetworkFunctionRoleConfigurations NetworkFunctionRoleConfigurationResponseArrayInput `pulumi:"networkFunctionRoleConfigurations"`
}

func (NetworkFunctionTemplateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionTemplateResponse)(nil)).Elem()
}

func (i NetworkFunctionTemplateResponseArgs) ToNetworkFunctionTemplateResponseOutput() NetworkFunctionTemplateResponseOutput {
	return i.ToNetworkFunctionTemplateResponseOutputWithContext(context.Background())
}

func (i NetworkFunctionTemplateResponseArgs) ToNetworkFunctionTemplateResponseOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplateResponseOutput)
}

func (i NetworkFunctionTemplateResponseArgs) ToNetworkFunctionTemplateResponsePtrOutput() NetworkFunctionTemplateResponsePtrOutput {
	return i.ToNetworkFunctionTemplateResponsePtrOutputWithContext(context.Background())
}

func (i NetworkFunctionTemplateResponseArgs) ToNetworkFunctionTemplateResponsePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplateResponseOutput).ToNetworkFunctionTemplateResponsePtrOutputWithContext(ctx)
}









type NetworkFunctionTemplateResponsePtrInput interface {
	pulumi.Input

	ToNetworkFunctionTemplateResponsePtrOutput() NetworkFunctionTemplateResponsePtrOutput
	ToNetworkFunctionTemplateResponsePtrOutputWithContext(context.Context) NetworkFunctionTemplateResponsePtrOutput
}

type networkFunctionTemplateResponsePtrType NetworkFunctionTemplateResponseArgs

func NetworkFunctionTemplateResponsePtr(v *NetworkFunctionTemplateResponseArgs) NetworkFunctionTemplateResponsePtrInput {
	return (*networkFunctionTemplateResponsePtrType)(v)
}

func (*networkFunctionTemplateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionTemplateResponse)(nil)).Elem()
}

func (i *networkFunctionTemplateResponsePtrType) ToNetworkFunctionTemplateResponsePtrOutput() NetworkFunctionTemplateResponsePtrOutput {
	return i.ToNetworkFunctionTemplateResponsePtrOutputWithContext(context.Background())
}

func (i *networkFunctionTemplateResponsePtrType) ToNetworkFunctionTemplateResponsePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionTemplateResponsePtrOutput)
}

type NetworkFunctionTemplateResponseOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionTemplateResponse)(nil)).Elem()
}

func (o NetworkFunctionTemplateResponseOutput) ToNetworkFunctionTemplateResponseOutput() NetworkFunctionTemplateResponseOutput {
	return o
}

func (o NetworkFunctionTemplateResponseOutput) ToNetworkFunctionTemplateResponseOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponseOutput {
	return o
}

func (o NetworkFunctionTemplateResponseOutput) ToNetworkFunctionTemplateResponsePtrOutput() NetworkFunctionTemplateResponsePtrOutput {
	return o.ToNetworkFunctionTemplateResponsePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionTemplateResponseOutput) ToNetworkFunctionTemplateResponsePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionTemplateResponse) *NetworkFunctionTemplateResponse {
		return &v
	}).(NetworkFunctionTemplateResponsePtrOutput)
}

func (o NetworkFunctionTemplateResponseOutput) NetworkFunctionRoleConfigurations() NetworkFunctionRoleConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NetworkFunctionTemplateResponse) []NetworkFunctionRoleConfigurationResponse {
		return v.NetworkFunctionRoleConfigurations
	}).(NetworkFunctionRoleConfigurationResponseArrayOutput)
}

type NetworkFunctionTemplateResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionTemplateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionTemplateResponse)(nil)).Elem()
}

func (o NetworkFunctionTemplateResponsePtrOutput) ToNetworkFunctionTemplateResponsePtrOutput() NetworkFunctionTemplateResponsePtrOutput {
	return o
}

func (o NetworkFunctionTemplateResponsePtrOutput) ToNetworkFunctionTemplateResponsePtrOutputWithContext(ctx context.Context) NetworkFunctionTemplateResponsePtrOutput {
	return o
}

func (o NetworkFunctionTemplateResponsePtrOutput) Elem() NetworkFunctionTemplateResponseOutput {
	return o.ApplyT(func(v *NetworkFunctionTemplateResponse) NetworkFunctionTemplateResponse {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionTemplateResponse
		return ret
	}).(NetworkFunctionTemplateResponseOutput)
}

func (o NetworkFunctionTemplateResponsePtrOutput) NetworkFunctionRoleConfigurations() NetworkFunctionRoleConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *NetworkFunctionTemplateResponse) []NetworkFunctionRoleConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.NetworkFunctionRoleConfigurations
	}).(NetworkFunctionRoleConfigurationResponseArrayOutput)
}

type NetworkFunctionUserConfiguration struct {
	NetworkInterfaces  []NetworkInterface                         `pulumi:"networkInterfaces"`
	OsProfile          *NetworkFunctionUserConfigurationOsProfile `pulumi:"osProfile"`
	RoleName           *string                                    `pulumi:"roleName"`
	UserDataParameters interface{}                                `pulumi:"userDataParameters"`
}





type NetworkFunctionUserConfigurationInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationOutput() NetworkFunctionUserConfigurationOutput
	ToNetworkFunctionUserConfigurationOutputWithContext(context.Context) NetworkFunctionUserConfigurationOutput
}

type NetworkFunctionUserConfigurationArgs struct {
	NetworkInterfaces  NetworkInterfaceArrayInput                        `pulumi:"networkInterfaces"`
	OsProfile          NetworkFunctionUserConfigurationOsProfilePtrInput `pulumi:"osProfile"`
	RoleName           pulumi.StringPtrInput                             `pulumi:"roleName"`
	UserDataParameters pulumi.Input                                      `pulumi:"userDataParameters"`
}

func (NetworkFunctionUserConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfiguration)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationArgs) ToNetworkFunctionUserConfigurationOutput() NetworkFunctionUserConfigurationOutput {
	return i.ToNetworkFunctionUserConfigurationOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationArgs) ToNetworkFunctionUserConfigurationOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationOutput)
}





type NetworkFunctionUserConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationArrayOutput() NetworkFunctionUserConfigurationArrayOutput
	ToNetworkFunctionUserConfigurationArrayOutputWithContext(context.Context) NetworkFunctionUserConfigurationArrayOutput
}

type NetworkFunctionUserConfigurationArray []NetworkFunctionUserConfigurationInput

func (NetworkFunctionUserConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionUserConfiguration)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationArray) ToNetworkFunctionUserConfigurationArrayOutput() NetworkFunctionUserConfigurationArrayOutput {
	return i.ToNetworkFunctionUserConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationArray) ToNetworkFunctionUserConfigurationArrayOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationArrayOutput)
}

type NetworkFunctionUserConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfiguration)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationOutput) ToNetworkFunctionUserConfigurationOutput() NetworkFunctionUserConfigurationOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOutput) ToNetworkFunctionUserConfigurationOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfiguration) []NetworkInterface { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
}

func (o NetworkFunctionUserConfigurationOutput) OsProfile() NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfiguration) *NetworkFunctionUserConfigurationOsProfile {
		return v.OsProfile
	}).(NetworkFunctionUserConfigurationOsProfilePtrOutput)
}

func (o NetworkFunctionUserConfigurationOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfiguration) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionUserConfigurationOutput) UserDataParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfiguration) interface{} { return v.UserDataParameters }).(pulumi.AnyOutput)
}

type NetworkFunctionUserConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionUserConfiguration)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationArrayOutput) ToNetworkFunctionUserConfigurationArrayOutput() NetworkFunctionUserConfigurationArrayOutput {
	return o
}

func (o NetworkFunctionUserConfigurationArrayOutput) ToNetworkFunctionUserConfigurationArrayOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationArrayOutput {
	return o
}

func (o NetworkFunctionUserConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkFunctionUserConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkFunctionUserConfiguration {
		return vs[0].([]NetworkFunctionUserConfiguration)[vs[1].(int)]
	}).(NetworkFunctionUserConfigurationOutput)
}

type NetworkFunctionUserConfigurationOsProfile struct {
	CustomData *string `pulumi:"customData"`
}





type NetworkFunctionUserConfigurationOsProfileInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationOsProfileOutput() NetworkFunctionUserConfigurationOsProfileOutput
	ToNetworkFunctionUserConfigurationOsProfileOutputWithContext(context.Context) NetworkFunctionUserConfigurationOsProfileOutput
}

type NetworkFunctionUserConfigurationOsProfileArgs struct {
	CustomData pulumi.StringPtrInput `pulumi:"customData"`
}

func (NetworkFunctionUserConfigurationOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationOsProfile)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationOsProfileArgs) ToNetworkFunctionUserConfigurationOsProfileOutput() NetworkFunctionUserConfigurationOsProfileOutput {
	return i.ToNetworkFunctionUserConfigurationOsProfileOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationOsProfileArgs) ToNetworkFunctionUserConfigurationOsProfileOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationOsProfileOutput)
}

func (i NetworkFunctionUserConfigurationOsProfileArgs) ToNetworkFunctionUserConfigurationOsProfilePtrOutput() NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return i.ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationOsProfileArgs) ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationOsProfileOutput).ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(ctx)
}









type NetworkFunctionUserConfigurationOsProfilePtrInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationOsProfilePtrOutput() NetworkFunctionUserConfigurationOsProfilePtrOutput
	ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(context.Context) NetworkFunctionUserConfigurationOsProfilePtrOutput
}

type networkFunctionUserConfigurationOsProfilePtrType NetworkFunctionUserConfigurationOsProfileArgs

func NetworkFunctionUserConfigurationOsProfilePtr(v *NetworkFunctionUserConfigurationOsProfileArgs) NetworkFunctionUserConfigurationOsProfilePtrInput {
	return (*networkFunctionUserConfigurationOsProfilePtrType)(v)
}

func (*networkFunctionUserConfigurationOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionUserConfigurationOsProfile)(nil)).Elem()
}

func (i *networkFunctionUserConfigurationOsProfilePtrType) ToNetworkFunctionUserConfigurationOsProfilePtrOutput() NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return i.ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(context.Background())
}

func (i *networkFunctionUserConfigurationOsProfilePtrType) ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationOsProfilePtrOutput)
}

type NetworkFunctionUserConfigurationOsProfileOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationOsProfile)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationOsProfileOutput) ToNetworkFunctionUserConfigurationOsProfileOutput() NetworkFunctionUserConfigurationOsProfileOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOsProfileOutput) ToNetworkFunctionUserConfigurationOsProfileOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfileOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOsProfileOutput) ToNetworkFunctionUserConfigurationOsProfilePtrOutput() NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return o.ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionUserConfigurationOsProfileOutput) ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionUserConfigurationOsProfile) *NetworkFunctionUserConfigurationOsProfile {
		return &v
	}).(NetworkFunctionUserConfigurationOsProfilePtrOutput)
}

func (o NetworkFunctionUserConfigurationOsProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationOsProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

type NetworkFunctionUserConfigurationOsProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionUserConfigurationOsProfile)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationOsProfilePtrOutput) ToNetworkFunctionUserConfigurationOsProfilePtrOutput() NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOsProfilePtrOutput) ToNetworkFunctionUserConfigurationOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationOsProfilePtrOutput {
	return o
}

func (o NetworkFunctionUserConfigurationOsProfilePtrOutput) Elem() NetworkFunctionUserConfigurationOsProfileOutput {
	return o.ApplyT(func(v *NetworkFunctionUserConfigurationOsProfile) NetworkFunctionUserConfigurationOsProfile {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionUserConfigurationOsProfile
		return ret
	}).(NetworkFunctionUserConfigurationOsProfileOutput)
}

func (o NetworkFunctionUserConfigurationOsProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkFunctionUserConfigurationOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

type NetworkFunctionUserConfigurationResponse struct {
	NetworkInterfaces  []NetworkInterfaceResponse                         `pulumi:"networkInterfaces"`
	OsProfile          *NetworkFunctionUserConfigurationResponseOsProfile `pulumi:"osProfile"`
	RoleName           *string                                            `pulumi:"roleName"`
	UserDataParameters interface{}                                        `pulumi:"userDataParameters"`
}





type NetworkFunctionUserConfigurationResponseInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationResponseOutput() NetworkFunctionUserConfigurationResponseOutput
	ToNetworkFunctionUserConfigurationResponseOutputWithContext(context.Context) NetworkFunctionUserConfigurationResponseOutput
}

type NetworkFunctionUserConfigurationResponseArgs struct {
	NetworkInterfaces  NetworkInterfaceResponseArrayInput                        `pulumi:"networkInterfaces"`
	OsProfile          NetworkFunctionUserConfigurationResponseOsProfilePtrInput `pulumi:"osProfile"`
	RoleName           pulumi.StringPtrInput                                     `pulumi:"roleName"`
	UserDataParameters pulumi.Input                                              `pulumi:"userDataParameters"`
}

func (NetworkFunctionUserConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationResponse)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationResponseArgs) ToNetworkFunctionUserConfigurationResponseOutput() NetworkFunctionUserConfigurationResponseOutput {
	return i.ToNetworkFunctionUserConfigurationResponseOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationResponseArgs) ToNetworkFunctionUserConfigurationResponseOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationResponseOutput)
}





type NetworkFunctionUserConfigurationResponseArrayInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationResponseArrayOutput() NetworkFunctionUserConfigurationResponseArrayOutput
	ToNetworkFunctionUserConfigurationResponseArrayOutputWithContext(context.Context) NetworkFunctionUserConfigurationResponseArrayOutput
}

type NetworkFunctionUserConfigurationResponseArray []NetworkFunctionUserConfigurationResponseInput

func (NetworkFunctionUserConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionUserConfigurationResponse)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationResponseArray) ToNetworkFunctionUserConfigurationResponseArrayOutput() NetworkFunctionUserConfigurationResponseArrayOutput {
	return i.ToNetworkFunctionUserConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationResponseArray) ToNetworkFunctionUserConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationResponseArrayOutput)
}

type NetworkFunctionUserConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationResponse)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationResponseOutput) ToNetworkFunctionUserConfigurationResponseOutput() NetworkFunctionUserConfigurationResponseOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOutput) ToNetworkFunctionUserConfigurationResponseOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationResponse) []NetworkInterfaceResponse {
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkFunctionUserConfigurationResponseOutput) OsProfile() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationResponse) *NetworkFunctionUserConfigurationResponseOsProfile {
		return v.OsProfile
	}).(NetworkFunctionUserConfigurationResponseOsProfilePtrOutput)
}

func (o NetworkFunctionUserConfigurationResponseOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationResponse) *string { return v.RoleName }).(pulumi.StringPtrOutput)
}

func (o NetworkFunctionUserConfigurationResponseOutput) UserDataParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationResponse) interface{} { return v.UserDataParameters }).(pulumi.AnyOutput)
}

type NetworkFunctionUserConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkFunctionUserConfigurationResponse)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationResponseArrayOutput) ToNetworkFunctionUserConfigurationResponseArrayOutput() NetworkFunctionUserConfigurationResponseArrayOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseArrayOutput) ToNetworkFunctionUserConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseArrayOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkFunctionUserConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkFunctionUserConfigurationResponse {
		return vs[0].([]NetworkFunctionUserConfigurationResponse)[vs[1].(int)]
	}).(NetworkFunctionUserConfigurationResponseOutput)
}

type NetworkFunctionUserConfigurationResponseOsProfile struct {
	CustomData *string `pulumi:"customData"`
}





type NetworkFunctionUserConfigurationResponseOsProfileInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationResponseOsProfileOutput() NetworkFunctionUserConfigurationResponseOsProfileOutput
	ToNetworkFunctionUserConfigurationResponseOsProfileOutputWithContext(context.Context) NetworkFunctionUserConfigurationResponseOsProfileOutput
}

type NetworkFunctionUserConfigurationResponseOsProfileArgs struct {
	CustomData pulumi.StringPtrInput `pulumi:"customData"`
}

func (NetworkFunctionUserConfigurationResponseOsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationResponseOsProfile)(nil)).Elem()
}

func (i NetworkFunctionUserConfigurationResponseOsProfileArgs) ToNetworkFunctionUserConfigurationResponseOsProfileOutput() NetworkFunctionUserConfigurationResponseOsProfileOutput {
	return i.ToNetworkFunctionUserConfigurationResponseOsProfileOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationResponseOsProfileArgs) ToNetworkFunctionUserConfigurationResponseOsProfileOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationResponseOsProfileOutput)
}

func (i NetworkFunctionUserConfigurationResponseOsProfileArgs) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutput() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return i.ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(context.Background())
}

func (i NetworkFunctionUserConfigurationResponseOsProfileArgs) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationResponseOsProfileOutput).ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(ctx)
}









type NetworkFunctionUserConfigurationResponseOsProfilePtrInput interface {
	pulumi.Input

	ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutput() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput
	ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(context.Context) NetworkFunctionUserConfigurationResponseOsProfilePtrOutput
}

type networkFunctionUserConfigurationResponseOsProfilePtrType NetworkFunctionUserConfigurationResponseOsProfileArgs

func NetworkFunctionUserConfigurationResponseOsProfilePtr(v *NetworkFunctionUserConfigurationResponseOsProfileArgs) NetworkFunctionUserConfigurationResponseOsProfilePtrInput {
	return (*networkFunctionUserConfigurationResponseOsProfilePtrType)(v)
}

func (*networkFunctionUserConfigurationResponseOsProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionUserConfigurationResponseOsProfile)(nil)).Elem()
}

func (i *networkFunctionUserConfigurationResponseOsProfilePtrType) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutput() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return i.ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(context.Background())
}

func (i *networkFunctionUserConfigurationResponseOsProfilePtrType) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFunctionUserConfigurationResponseOsProfilePtrOutput)
}

type NetworkFunctionUserConfigurationResponseOsProfileOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationResponseOsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkFunctionUserConfigurationResponseOsProfile)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationResponseOsProfileOutput) ToNetworkFunctionUserConfigurationResponseOsProfileOutput() NetworkFunctionUserConfigurationResponseOsProfileOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOsProfileOutput) ToNetworkFunctionUserConfigurationResponseOsProfileOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfileOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOsProfileOutput) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutput() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return o.ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(context.Background())
}

func (o NetworkFunctionUserConfigurationResponseOsProfileOutput) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkFunctionUserConfigurationResponseOsProfile) *NetworkFunctionUserConfigurationResponseOsProfile {
		return &v
	}).(NetworkFunctionUserConfigurationResponseOsProfilePtrOutput)
}

func (o NetworkFunctionUserConfigurationResponseOsProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkFunctionUserConfigurationResponseOsProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

type NetworkFunctionUserConfigurationResponseOsProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkFunctionUserConfigurationResponseOsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFunctionUserConfigurationResponseOsProfile)(nil)).Elem()
}

func (o NetworkFunctionUserConfigurationResponseOsProfilePtrOutput) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutput() NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOsProfilePtrOutput) ToNetworkFunctionUserConfigurationResponseOsProfilePtrOutputWithContext(ctx context.Context) NetworkFunctionUserConfigurationResponseOsProfilePtrOutput {
	return o
}

func (o NetworkFunctionUserConfigurationResponseOsProfilePtrOutput) Elem() NetworkFunctionUserConfigurationResponseOsProfileOutput {
	return o.ApplyT(func(v *NetworkFunctionUserConfigurationResponseOsProfile) NetworkFunctionUserConfigurationResponseOsProfile {
		if v != nil {
			return *v
		}
		var ret NetworkFunctionUserConfigurationResponseOsProfile
		return ret
	}).(NetworkFunctionUserConfigurationResponseOsProfileOutput)
}

func (o NetworkFunctionUserConfigurationResponseOsProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkFunctionUserConfigurationResponseOsProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

type NetworkInterface struct {
	IpConfigurations     []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	MacAddress           *string                           `pulumi:"macAddress"`
	NetworkInterfaceName *string                           `pulumi:"networkInterfaceName"`
	VmSwitchType         *string                           `pulumi:"vmSwitchType"`
}





type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(context.Context) NetworkInterfaceOutput
}

type NetworkInterfaceArgs struct {
	IpConfigurations     NetworkInterfaceIPConfigurationArrayInput `pulumi:"ipConfigurations"`
	MacAddress           pulumi.StringPtrInput                     `pulumi:"macAddress"`
	NetworkInterfaceName pulumi.StringPtrInput                     `pulumi:"networkInterfaceName"`
	VmSwitchType         pulumi.StringPtrInput                     `pulumi:"vmSwitchType"`
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}





type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) IpConfigurations() NetworkInterfaceIPConfigurationArrayOutput {
	return o.ApplyT(func(v NetworkInterface) []NetworkInterfaceIPConfiguration { return v.IpConfigurations }).(NetworkInterfaceIPConfigurationArrayOutput)
}

func (o NetworkInterfaceOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) NetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.NetworkInterfaceName }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceOutput) VmSwitchType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.VmSwitchType }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceIPConfiguration struct {
	DnsServers         []string `pulumi:"dnsServers"`
	Gateway            *string  `pulumi:"gateway"`
	IpAddress          *string  `pulumi:"ipAddress"`
	IpAllocationMethod *string  `pulumi:"ipAllocationMethod"`
	IpVersion          *string  `pulumi:"ipVersion"`
	Subnet             *string  `pulumi:"subnet"`
}





type NetworkInterfaceIPConfigurationInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput
	ToNetworkInterfaceIPConfigurationOutputWithContext(context.Context) NetworkInterfaceIPConfigurationOutput
}

type NetworkInterfaceIPConfigurationArgs struct {
	DnsServers         pulumi.StringArrayInput `pulumi:"dnsServers"`
	Gateway            pulumi.StringPtrInput   `pulumi:"gateway"`
	IpAddress          pulumi.StringPtrInput   `pulumi:"ipAddress"`
	IpAllocationMethod pulumi.StringPtrInput   `pulumi:"ipAllocationMethod"`
	IpVersion          pulumi.StringPtrInput   `pulumi:"ipVersion"`
	Subnet             pulumi.StringPtrInput   `pulumi:"subnet"`
}

func (NetworkInterfaceIPConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return i.ToNetworkInterfaceIPConfigurationOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArgs) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationOutput)
}





type NetworkInterfaceIPConfigurationArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput
	ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Context) NetworkInterfaceIPConfigurationArrayOutput
}

type NetworkInterfaceIPConfigurationArray []NetworkInterfaceIPConfigurationInput

func (NetworkInterfaceIPConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return i.ToNetworkInterfaceIPConfigurationArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationArray) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationArrayOutput)
}

type NetworkInterfaceIPConfigurationOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutput() NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) ToNetworkInterfaceIPConfigurationOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Gateway }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.IpVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfiguration) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceIPConfigurationArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfiguration)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutput() NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) ToNetworkInterfaceIPConfigurationArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfiguration {
		return vs[0].([]NetworkInterfaceIPConfiguration)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationOutput)
}

type NetworkInterfaceIPConfigurationResponse struct {
	DnsServers         []string `pulumi:"dnsServers"`
	Gateway            *string  `pulumi:"gateway"`
	IpAddress          *string  `pulumi:"ipAddress"`
	IpAllocationMethod *string  `pulumi:"ipAllocationMethod"`
	IpVersion          *string  `pulumi:"ipVersion"`
	Subnet             *string  `pulumi:"subnet"`
}





type NetworkInterfaceIPConfigurationResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput
	ToNetworkInterfaceIPConfigurationResponseOutputWithContext(context.Context) NetworkInterfaceIPConfigurationResponseOutput
}

type NetworkInterfaceIPConfigurationResponseArgs struct {
	DnsServers         pulumi.StringArrayInput `pulumi:"dnsServers"`
	Gateway            pulumi.StringPtrInput   `pulumi:"gateway"`
	IpAddress          pulumi.StringPtrInput   `pulumi:"ipAddress"`
	IpAllocationMethod pulumi.StringPtrInput   `pulumi:"ipAllocationMethod"`
	IpVersion          pulumi.StringPtrInput   `pulumi:"ipVersion"`
	Subnet             pulumi.StringPtrInput   `pulumi:"subnet"`
}

func (NetworkInterfaceIPConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput {
	return i.ToNetworkInterfaceIPConfigurationResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationResponseArgs) ToNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponseOutput)
}





type NetworkInterfaceIPConfigurationResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput
	ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput
}

type NetworkInterfaceIPConfigurationResponseArray []NetworkInterfaceIPConfigurationResponseInput

func (NetworkInterfaceIPConfigurationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (i NetworkInterfaceIPConfigurationResponseArray) ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return i.ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceIPConfigurationResponseArray) ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

type NetworkInterfaceIPConfigurationResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutput() NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) ToNetworkInterfaceIPConfigurationResponseOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Gateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Gateway }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) IpAllocationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.IpAllocationMethod }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) IpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.IpVersion }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceIPConfigurationResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceIPConfigurationResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceIPConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceIPConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceIPConfigurationResponse)(nil)).Elem()
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutput() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) ToNetworkInterfaceIPConfigurationResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o
}

func (o NetworkInterfaceIPConfigurationResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceIPConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceIPConfigurationResponse {
		return vs[0].([]NetworkInterfaceIPConfigurationResponse)[vs[1].(int)]
	}).(NetworkInterfaceIPConfigurationResponseOutput)
}

type NetworkInterfaceResponse struct {
	IpConfigurations     []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	MacAddress           *string                                   `pulumi:"macAddress"`
	NetworkInterfaceName *string                                   `pulumi:"networkInterfaceName"`
	VmSwitchType         *string                                   `pulumi:"vmSwitchType"`
}





type NetworkInterfaceResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput
	ToNetworkInterfaceResponseOutputWithContext(context.Context) NetworkInterfaceResponseOutput
}

type NetworkInterfaceResponseArgs struct {
	IpConfigurations     NetworkInterfaceIPConfigurationResponseArrayInput `pulumi:"ipConfigurations"`
	MacAddress           pulumi.StringPtrInput                             `pulumi:"macAddress"`
	NetworkInterfaceName pulumi.StringPtrInput                             `pulumi:"networkInterfaceName"`
	VmSwitchType         pulumi.StringPtrInput                             `pulumi:"vmSwitchType"`
}

func (NetworkInterfaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return i.ToNetworkInterfaceResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseOutput)
}





type NetworkInterfaceResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput
	ToNetworkInterfaceResponseArrayOutputWithContext(context.Context) NetworkInterfaceResponseArrayOutput
}

type NetworkInterfaceResponseArray []NetworkInterfaceResponseInput

func (NetworkInterfaceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return i.ToNetworkInterfaceResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseArrayOutput)
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) IpConfigurations() NetworkInterfaceIPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) []NetworkInterfaceIPConfigurationResponse { return v.IpConfigurations }).(NetworkInterfaceIPConfigurationResponseArrayOutput)
}

func (o NetworkInterfaceResponseOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) NetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.NetworkInterfaceName }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfaceResponseOutput) VmSwitchType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.VmSwitchType }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceResponse {
		return vs[0].([]NetworkInterfaceResponse)[vs[1].(int)]
	}).(NetworkInterfaceResponseOutput)
}

type OsDisk struct {
	DiskSizeGB *int             `pulumi:"diskSizeGB"`
	Name       *string          `pulumi:"name"`
	OsType     *string          `pulumi:"osType"`
	Vhd        *VirtualHardDisk `pulumi:"vhd"`
}





type OsDiskInput interface {
	pulumi.Input

	ToOsDiskOutput() OsDiskOutput
	ToOsDiskOutputWithContext(context.Context) OsDiskOutput
}

type OsDiskArgs struct {
	DiskSizeGB pulumi.IntPtrInput      `pulumi:"diskSizeGB"`
	Name       pulumi.StringPtrInput   `pulumi:"name"`
	OsType     pulumi.StringPtrInput   `pulumi:"osType"`
	Vhd        VirtualHardDiskPtrInput `pulumi:"vhd"`
}

func (OsDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDisk)(nil)).Elem()
}

func (i OsDiskArgs) ToOsDiskOutput() OsDiskOutput {
	return i.ToOsDiskOutputWithContext(context.Background())
}

func (i OsDiskArgs) ToOsDiskOutputWithContext(ctx context.Context) OsDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskOutput)
}

func (i OsDiskArgs) ToOsDiskPtrOutput() OsDiskPtrOutput {
	return i.ToOsDiskPtrOutputWithContext(context.Background())
}

func (i OsDiskArgs) ToOsDiskPtrOutputWithContext(ctx context.Context) OsDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskOutput).ToOsDiskPtrOutputWithContext(ctx)
}









type OsDiskPtrInput interface {
	pulumi.Input

	ToOsDiskPtrOutput() OsDiskPtrOutput
	ToOsDiskPtrOutputWithContext(context.Context) OsDiskPtrOutput
}

type osDiskPtrType OsDiskArgs

func OsDiskPtr(v *OsDiskArgs) OsDiskPtrInput {
	return (*osDiskPtrType)(v)
}

func (*osDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsDisk)(nil)).Elem()
}

func (i *osDiskPtrType) ToOsDiskPtrOutput() OsDiskPtrOutput {
	return i.ToOsDiskPtrOutputWithContext(context.Background())
}

func (i *osDiskPtrType) ToOsDiskPtrOutputWithContext(ctx context.Context) OsDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskPtrOutput)
}

type OsDiskOutput struct{ *pulumi.OutputState }

func (OsDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDisk)(nil)).Elem()
}

func (o OsDiskOutput) ToOsDiskOutput() OsDiskOutput {
	return o
}

func (o OsDiskOutput) ToOsDiskOutputWithContext(ctx context.Context) OsDiskOutput {
	return o
}

func (o OsDiskOutput) ToOsDiskPtrOutput() OsDiskPtrOutput {
	return o.ToOsDiskPtrOutputWithContext(context.Background())
}

func (o OsDiskOutput) ToOsDiskPtrOutputWithContext(ctx context.Context) OsDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsDisk) *OsDisk {
		return &v
	}).(OsDiskPtrOutput)
}

func (o OsDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OsDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OsDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OsDiskOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsDisk) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OsDiskOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v OsDisk) *VirtualHardDisk { return v.Vhd }).(VirtualHardDiskPtrOutput)
}

type OsDiskPtrOutput struct{ *pulumi.OutputState }

func (OsDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsDisk)(nil)).Elem()
}

func (o OsDiskPtrOutput) ToOsDiskPtrOutput() OsDiskPtrOutput {
	return o
}

func (o OsDiskPtrOutput) ToOsDiskPtrOutputWithContext(ctx context.Context) OsDiskPtrOutput {
	return o
}

func (o OsDiskPtrOutput) Elem() OsDiskOutput {
	return o.ApplyT(func(v *OsDisk) OsDisk {
		if v != nil {
			return *v
		}
		var ret OsDisk
		return ret
	}).(OsDiskOutput)
}

func (o OsDiskPtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OsDisk) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OsDiskPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsDisk) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OsDiskPtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsDisk) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OsDiskPtrOutput) Vhd() VirtualHardDiskPtrOutput {
	return o.ApplyT(func(v *OsDisk) *VirtualHardDisk {
		if v == nil {
			return nil
		}
		return v.Vhd
	}).(VirtualHardDiskPtrOutput)
}

type OsDiskResponse struct {
	DiskSizeGB *int    `pulumi:"diskSizeGB"`
	Name       *string `pulumi:"name"`
	OsType     *string `pulumi:"osType"`
}





type OsDiskResponseInput interface {
	pulumi.Input

	ToOsDiskResponseOutput() OsDiskResponseOutput
	ToOsDiskResponseOutputWithContext(context.Context) OsDiskResponseOutput
}

type OsDiskResponseArgs struct {
	DiskSizeGB pulumi.IntPtrInput    `pulumi:"diskSizeGB"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	OsType     pulumi.StringPtrInput `pulumi:"osType"`
}

func (OsDiskResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDiskResponse)(nil)).Elem()
}

func (i OsDiskResponseArgs) ToOsDiskResponseOutput() OsDiskResponseOutput {
	return i.ToOsDiskResponseOutputWithContext(context.Background())
}

func (i OsDiskResponseArgs) ToOsDiskResponseOutputWithContext(ctx context.Context) OsDiskResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskResponseOutput)
}

func (i OsDiskResponseArgs) ToOsDiskResponsePtrOutput() OsDiskResponsePtrOutput {
	return i.ToOsDiskResponsePtrOutputWithContext(context.Background())
}

func (i OsDiskResponseArgs) ToOsDiskResponsePtrOutputWithContext(ctx context.Context) OsDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskResponseOutput).ToOsDiskResponsePtrOutputWithContext(ctx)
}









type OsDiskResponsePtrInput interface {
	pulumi.Input

	ToOsDiskResponsePtrOutput() OsDiskResponsePtrOutput
	ToOsDiskResponsePtrOutputWithContext(context.Context) OsDiskResponsePtrOutput
}

type osDiskResponsePtrType OsDiskResponseArgs

func OsDiskResponsePtr(v *OsDiskResponseArgs) OsDiskResponsePtrInput {
	return (*osDiskResponsePtrType)(v)
}

func (*osDiskResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsDiskResponse)(nil)).Elem()
}

func (i *osDiskResponsePtrType) ToOsDiskResponsePtrOutput() OsDiskResponsePtrOutput {
	return i.ToOsDiskResponsePtrOutputWithContext(context.Background())
}

func (i *osDiskResponsePtrType) ToOsDiskResponsePtrOutputWithContext(ctx context.Context) OsDiskResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskResponsePtrOutput)
}

type OsDiskResponseOutput struct{ *pulumi.OutputState }

func (OsDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDiskResponse)(nil)).Elem()
}

func (o OsDiskResponseOutput) ToOsDiskResponseOutput() OsDiskResponseOutput {
	return o
}

func (o OsDiskResponseOutput) ToOsDiskResponseOutputWithContext(ctx context.Context) OsDiskResponseOutput {
	return o
}

func (o OsDiskResponseOutput) ToOsDiskResponsePtrOutput() OsDiskResponsePtrOutput {
	return o.ToOsDiskResponsePtrOutputWithContext(context.Background())
}

func (o OsDiskResponseOutput) ToOsDiskResponsePtrOutputWithContext(ctx context.Context) OsDiskResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsDiskResponse) *OsDiskResponse {
		return &v
	}).(OsDiskResponsePtrOutput)
}

func (o OsDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v OsDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o OsDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o OsDiskResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsDiskResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

type OsDiskResponsePtrOutput struct{ *pulumi.OutputState }

func (OsDiskResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsDiskResponse)(nil)).Elem()
}

func (o OsDiskResponsePtrOutput) ToOsDiskResponsePtrOutput() OsDiskResponsePtrOutput {
	return o
}

func (o OsDiskResponsePtrOutput) ToOsDiskResponsePtrOutputWithContext(ctx context.Context) OsDiskResponsePtrOutput {
	return o
}

func (o OsDiskResponsePtrOutput) Elem() OsDiskResponseOutput {
	return o.ApplyT(func(v *OsDiskResponse) OsDiskResponse {
		if v != nil {
			return *v
		}
		var ret OsDiskResponse
		return ret
	}).(OsDiskResponseOutput)
}

func (o OsDiskResponsePtrOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OsDiskResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o OsDiskResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o OsDiskResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsDiskResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

type OsProfile struct {
	AdminUsername      *string             `pulumi:"adminUsername"`
	CustomData         *string             `pulumi:"customData"`
	CustomDataRequired *bool               `pulumi:"customDataRequired"`
	LinuxConfiguration *LinuxConfiguration `pulumi:"linuxConfiguration"`
}


func (val *OsProfile) Defaults() *OsProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CustomDataRequired) {
		customDataRequired_ := true
		tmp.CustomDataRequired = &customDataRequired_
	}
	return &tmp
}





type OsProfileInput interface {
	pulumi.Input

	ToOsProfileOutput() OsProfileOutput
	ToOsProfileOutputWithContext(context.Context) OsProfileOutput
}

type OsProfileArgs struct {
	AdminUsername      pulumi.StringPtrInput      `pulumi:"adminUsername"`
	CustomData         pulumi.StringPtrInput      `pulumi:"customData"`
	CustomDataRequired pulumi.BoolPtrInput        `pulumi:"customDataRequired"`
	LinuxConfiguration LinuxConfigurationPtrInput `pulumi:"linuxConfiguration"`
}

func (OsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (i OsProfileArgs) ToOsProfileOutput() OsProfileOutput {
	return i.ToOsProfileOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput)
}

func (i OsProfileArgs) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput).ToOsProfilePtrOutputWithContext(ctx)
}









type OsProfilePtrInput interface {
	pulumi.Input

	ToOsProfilePtrOutput() OsProfilePtrOutput
	ToOsProfilePtrOutputWithContext(context.Context) OsProfilePtrOutput
}

type osProfilePtrType OsProfileArgs

func OsProfilePtr(v *OsProfileArgs) OsProfilePtrInput {
	return (*osProfilePtrType)(v)
}

func (*osProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (i *osProfilePtrType) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i *osProfilePtrType) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfilePtrOutput)
}

type OsProfileOutput struct{ *pulumi.OutputState }

func (OsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (o OsProfileOutput) ToOsProfileOutput() OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o.ToOsProfilePtrOutputWithContext(context.Background())
}

func (o OsProfileOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsProfile) *OsProfile {
		return &v
	}).(OsProfilePtrOutput)
}

func (o OsProfileOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) CustomDataRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OsProfile) *bool { return v.CustomDataRequired }).(pulumi.BoolPtrOutput)
}

func (o OsProfileOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v OsProfile) *LinuxConfiguration { return v.LinuxConfiguration }).(LinuxConfigurationPtrOutput)
}

type OsProfilePtrOutput struct{ *pulumi.OutputState }

func (OsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) Elem() OsProfileOutput {
	return o.ApplyT(func(v *OsProfile) OsProfile {
		if v != nil {
			return *v
		}
		var ret OsProfile
		return ret
	}).(OsProfileOutput)
}

func (o OsProfilePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) CustomDataRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OsProfile) *bool {
		if v == nil {
			return nil
		}
		return v.CustomDataRequired
	}).(pulumi.BoolPtrOutput)
}

func (o OsProfilePtrOutput) LinuxConfiguration() LinuxConfigurationPtrOutput {
	return o.ApplyT(func(v *OsProfile) *LinuxConfiguration {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationPtrOutput)
}

type OsProfileResponse struct {
	AdminUsername      *string                     `pulumi:"adminUsername"`
	CustomData         *string                     `pulumi:"customData"`
	CustomDataRequired *bool                       `pulumi:"customDataRequired"`
	LinuxConfiguration *LinuxConfigurationResponse `pulumi:"linuxConfiguration"`
}


func (val *OsProfileResponse) Defaults() *OsProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.CustomDataRequired) {
		customDataRequired_ := true
		tmp.CustomDataRequired = &customDataRequired_
	}
	return &tmp
}





type OsProfileResponseInput interface {
	pulumi.Input

	ToOsProfileResponseOutput() OsProfileResponseOutput
	ToOsProfileResponseOutputWithContext(context.Context) OsProfileResponseOutput
}

type OsProfileResponseArgs struct {
	AdminUsername      pulumi.StringPtrInput              `pulumi:"adminUsername"`
	CustomData         pulumi.StringPtrInput              `pulumi:"customData"`
	CustomDataRequired pulumi.BoolPtrInput                `pulumi:"customDataRequired"`
	LinuxConfiguration LinuxConfigurationResponsePtrInput `pulumi:"linuxConfiguration"`
}

func (OsProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfileResponse)(nil)).Elem()
}

func (i OsProfileResponseArgs) ToOsProfileResponseOutput() OsProfileResponseOutput {
	return i.ToOsProfileResponseOutputWithContext(context.Background())
}

func (i OsProfileResponseArgs) ToOsProfileResponseOutputWithContext(ctx context.Context) OsProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileResponseOutput)
}

func (i OsProfileResponseArgs) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return i.ToOsProfileResponsePtrOutputWithContext(context.Background())
}

func (i OsProfileResponseArgs) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileResponseOutput).ToOsProfileResponsePtrOutputWithContext(ctx)
}









type OsProfileResponsePtrInput interface {
	pulumi.Input

	ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput
	ToOsProfileResponsePtrOutputWithContext(context.Context) OsProfileResponsePtrOutput
}

type osProfileResponsePtrType OsProfileResponseArgs

func OsProfileResponsePtr(v *OsProfileResponseArgs) OsProfileResponsePtrInput {
	return (*osProfileResponsePtrType)(v)
}

func (*osProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfileResponse)(nil)).Elem()
}

func (i *osProfileResponsePtrType) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return i.ToOsProfileResponsePtrOutputWithContext(context.Background())
}

func (i *osProfileResponsePtrType) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileResponsePtrOutput)
}

type OsProfileResponseOutput struct{ *pulumi.OutputState }

func (OsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutput() OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutputWithContext(ctx context.Context) OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return o.ToOsProfileResponsePtrOutputWithContext(context.Background())
}

func (o OsProfileResponseOutput) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsProfileResponse) *OsProfileResponse {
		return &v
	}).(OsProfileResponsePtrOutput)
}

func (o OsProfileResponseOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.AdminUsername }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.CustomData }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) CustomDataRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *bool { return v.CustomDataRequired }).(pulumi.BoolPtrOutput)
}

func (o OsProfileResponseOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *LinuxConfigurationResponse { return v.LinuxConfiguration }).(LinuxConfigurationResponsePtrOutput)
}

type OsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) Elem() OsProfileResponseOutput {
	return o.ApplyT(func(v *OsProfileResponse) OsProfileResponse {
		if v != nil {
			return *v
		}
		var ret OsProfileResponse
		return ret
	}).(OsProfileResponseOutput)
}

func (o OsProfileResponsePtrOutput) AdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AdminUsername
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) CustomData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomData
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) CustomDataRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *bool {
		if v == nil {
			return nil
		}
		return v.CustomDataRequired
	}).(pulumi.BoolPtrOutput)
}

func (o OsProfileResponsePtrOutput) LinuxConfiguration() LinuxConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *LinuxConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LinuxConfiguration
	}).(LinuxConfigurationResponsePtrOutput)
}

type SshConfiguration struct {
	PublicKeys []SshPublicKey `pulumi:"publicKeys"`
}





type SshConfigurationInput interface {
	pulumi.Input

	ToSshConfigurationOutput() SshConfigurationOutput
	ToSshConfigurationOutputWithContext(context.Context) SshConfigurationOutput
}

type SshConfigurationArgs struct {
	PublicKeys SshPublicKeyArrayInput `pulumi:"publicKeys"`
}

func (SshConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (i SshConfigurationArgs) ToSshConfigurationOutput() SshConfigurationOutput {
	return i.ToSshConfigurationOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput)
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i SshConfigurationArgs) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationOutput).ToSshConfigurationPtrOutputWithContext(ctx)
}









type SshConfigurationPtrInput interface {
	pulumi.Input

	ToSshConfigurationPtrOutput() SshConfigurationPtrOutput
	ToSshConfigurationPtrOutputWithContext(context.Context) SshConfigurationPtrOutput
}

type sshConfigurationPtrType SshConfigurationArgs

func SshConfigurationPtr(v *SshConfigurationArgs) SshConfigurationPtrInput {
	return (*sshConfigurationPtrType)(v)
}

func (*sshConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return i.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (i *sshConfigurationPtrType) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationPtrOutput)
}

type SshConfigurationOutput struct{ *pulumi.OutputState }

func (SshConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationOutput) ToSshConfigurationOutput() SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationOutputWithContext(ctx context.Context) SshConfigurationOutput {
	return o
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o.ToSshConfigurationPtrOutputWithContext(context.Background())
}

func (o SshConfigurationOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshConfiguration) *SshConfiguration {
		return &v
	}).(SshConfigurationPtrOutput)
}

func (o SshConfigurationOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v SshConfiguration) []SshPublicKey { return v.PublicKeys }).(SshPublicKeyArrayOutput)
}

type SshConfigurationPtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfiguration)(nil)).Elem()
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutput() SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) ToSshConfigurationPtrOutputWithContext(ctx context.Context) SshConfigurationPtrOutput {
	return o
}

func (o SshConfigurationPtrOutput) Elem() SshConfigurationOutput {
	return o.ApplyT(func(v *SshConfiguration) SshConfiguration {
		if v != nil {
			return *v
		}
		var ret SshConfiguration
		return ret
	}).(SshConfigurationOutput)
}

func (o SshConfigurationPtrOutput) PublicKeys() SshPublicKeyArrayOutput {
	return o.ApplyT(func(v *SshConfiguration) []SshPublicKey {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyArrayOutput)
}

type SshConfigurationResponse struct {
	PublicKeys []SshPublicKeyResponse `pulumi:"publicKeys"`
}





type SshConfigurationResponseInput interface {
	pulumi.Input

	ToSshConfigurationResponseOutput() SshConfigurationResponseOutput
	ToSshConfigurationResponseOutputWithContext(context.Context) SshConfigurationResponseOutput
}

type SshConfigurationResponseArgs struct {
	PublicKeys SshPublicKeyResponseArrayInput `pulumi:"publicKeys"`
}

func (SshConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfigurationResponse)(nil)).Elem()
}

func (i SshConfigurationResponseArgs) ToSshConfigurationResponseOutput() SshConfigurationResponseOutput {
	return i.ToSshConfigurationResponseOutputWithContext(context.Background())
}

func (i SshConfigurationResponseArgs) ToSshConfigurationResponseOutputWithContext(ctx context.Context) SshConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationResponseOutput)
}

func (i SshConfigurationResponseArgs) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return i.ToSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i SshConfigurationResponseArgs) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationResponseOutput).ToSshConfigurationResponsePtrOutputWithContext(ctx)
}









type SshConfigurationResponsePtrInput interface {
	pulumi.Input

	ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput
	ToSshConfigurationResponsePtrOutputWithContext(context.Context) SshConfigurationResponsePtrOutput
}

type sshConfigurationResponsePtrType SshConfigurationResponseArgs

func SshConfigurationResponsePtr(v *SshConfigurationResponseArgs) SshConfigurationResponsePtrInput {
	return (*sshConfigurationResponsePtrType)(v)
}

func (*sshConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfigurationResponse)(nil)).Elem()
}

func (i *sshConfigurationResponsePtrType) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return i.ToSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *sshConfigurationResponsePtrType) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshConfigurationResponsePtrOutput)
}

type SshConfigurationResponseOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutput() SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponseOutputWithContext(ctx context.Context) SshConfigurationResponseOutput {
	return o
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return o.ToSshConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o SshConfigurationResponseOutput) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SshConfigurationResponse) *SshConfigurationResponse {
		return &v
	}).(SshConfigurationResponsePtrOutput)
}

func (o SshConfigurationResponseOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v SshConfigurationResponse) []SshPublicKeyResponse { return v.PublicKeys }).(SshPublicKeyResponseArrayOutput)
}

type SshConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (SshConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SshConfigurationResponse)(nil)).Elem()
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutput() SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) ToSshConfigurationResponsePtrOutputWithContext(ctx context.Context) SshConfigurationResponsePtrOutput {
	return o
}

func (o SshConfigurationResponsePtrOutput) Elem() SshConfigurationResponseOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) SshConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret SshConfigurationResponse
		return ret
	}).(SshConfigurationResponseOutput)
}

func (o SshConfigurationResponsePtrOutput) PublicKeys() SshPublicKeyResponseArrayOutput {
	return o.ApplyT(func(v *SshConfigurationResponse) []SshPublicKeyResponse {
		if v == nil {
			return nil
		}
		return v.PublicKeys
	}).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKey struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type SshPublicKeyInput interface {
	pulumi.Input

	ToSshPublicKeyOutput() SshPublicKeyOutput
	ToSshPublicKeyOutputWithContext(context.Context) SshPublicKeyOutput
}

type SshPublicKeyArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (SshPublicKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return i.ToSshPublicKeyOutputWithContext(context.Background())
}

func (i SshPublicKeyArgs) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyOutput)
}





type SshPublicKeyArrayInput interface {
	pulumi.Input

	ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput
	ToSshPublicKeyArrayOutputWithContext(context.Context) SshPublicKeyArrayOutput
}

type SshPublicKeyArray []SshPublicKeyInput

func (SshPublicKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return i.ToSshPublicKeyArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyArray) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyArrayOutput)
}

type SshPublicKeyOutput struct{ *pulumi.OutputState }

func (SshPublicKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutput() SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) ToSshPublicKeyOutputWithContext(ctx context.Context) SshPublicKeyOutput {
	return o
}

func (o SshPublicKeyOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKey) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKey)(nil)).Elem()
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutput() SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) ToSshPublicKeyArrayOutputWithContext(ctx context.Context) SshPublicKeyArrayOutput {
	return o
}

func (o SshPublicKeyArrayOutput) Index(i pulumi.IntInput) SshPublicKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKey {
		return vs[0].([]SshPublicKey)[vs[1].(int)]
	}).(SshPublicKeyOutput)
}

type SshPublicKeyResponse struct {
	KeyData *string `pulumi:"keyData"`
	Path    *string `pulumi:"path"`
}





type SshPublicKeyResponseInput interface {
	pulumi.Input

	ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput
	ToSshPublicKeyResponseOutputWithContext(context.Context) SshPublicKeyResponseOutput
}

type SshPublicKeyResponseArgs struct {
	KeyData pulumi.StringPtrInput `pulumi:"keyData"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (SshPublicKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (i SshPublicKeyResponseArgs) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return i.ToSshPublicKeyResponseOutputWithContext(context.Background())
}

func (i SshPublicKeyResponseArgs) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyResponseOutput)
}





type SshPublicKeyResponseArrayInput interface {
	pulumi.Input

	ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput
	ToSshPublicKeyResponseArrayOutputWithContext(context.Context) SshPublicKeyResponseArrayOutput
}

type SshPublicKeyResponseArray []SshPublicKeyResponseInput

func (SshPublicKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (i SshPublicKeyResponseArray) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return i.ToSshPublicKeyResponseArrayOutputWithContext(context.Background())
}

func (i SshPublicKeyResponseArray) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SshPublicKeyResponseArrayOutput)
}

type SshPublicKeyResponseOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutput() SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) ToSshPublicKeyResponseOutputWithContext(ctx context.Context) SshPublicKeyResponseOutput {
	return o
}

func (o SshPublicKeyResponseOutput) KeyData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.KeyData }).(pulumi.StringPtrOutput)
}

func (o SshPublicKeyResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SshPublicKeyResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type SshPublicKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (SshPublicKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SshPublicKeyResponse)(nil)).Elem()
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutput() SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) ToSshPublicKeyResponseArrayOutputWithContext(ctx context.Context) SshPublicKeyResponseArrayOutput {
	return o
}

func (o SshPublicKeyResponseArrayOutput) Index(i pulumi.IntInput) SshPublicKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SshPublicKeyResponse {
		return vs[0].([]SshPublicKeyResponse)[vs[1].(int)]
	}).(SshPublicKeyResponseOutput)
}

type StorageProfile struct {
	DataDisks      []DataDisk      `pulumi:"dataDisks"`
	ImageReference *ImageReference `pulumi:"imageReference"`
	OsDisk         *OsDisk         `pulumi:"osDisk"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	DataDisks      DataDiskArrayInput     `pulumi:"dataDisks"`
	ImageReference ImageReferencePtrInput `pulumi:"imageReference"`
	OsDisk         OsDiskPtrInput         `pulumi:"osDisk"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []DataDisk { return v.DataDisks }).(DataDiskArrayOutput)
}

func (o StorageProfileOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v StorageProfile) *ImageReference { return v.ImageReference }).(ImageReferencePtrOutput)
}

func (o StorageProfileOutput) OsDisk() OsDiskPtrOutput {
	return o.ApplyT(func(v StorageProfile) *OsDisk { return v.OsDisk }).(OsDiskPtrOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) DataDisks() DataDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []DataDisk {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskArrayOutput)
}

func (o StorageProfilePtrOutput) ImageReference() ImageReferencePtrOutput {
	return o.ApplyT(func(v *StorageProfile) *ImageReference {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferencePtrOutput)
}

func (o StorageProfilePtrOutput) OsDisk() OsDiskPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *OsDisk {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OsDiskPtrOutput)
}

type StorageProfileResponse struct {
	DataDisks      []DataDiskResponse      `pulumi:"dataDisks"`
	ImageReference *ImageReferenceResponse `pulumi:"imageReference"`
	OsDisk         *OsDiskResponse         `pulumi:"osDisk"`
}





type StorageProfileResponseInput interface {
	pulumi.Input

	ToStorageProfileResponseOutput() StorageProfileResponseOutput
	ToStorageProfileResponseOutputWithContext(context.Context) StorageProfileResponseOutput
}

type StorageProfileResponseArgs struct {
	DataDisks      DataDiskResponseArrayInput     `pulumi:"dataDisks"`
	ImageReference ImageReferenceResponsePtrInput `pulumi:"imageReference"`
	OsDisk         OsDiskResponsePtrInput         `pulumi:"osDisk"`
}

func (StorageProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return i.ToStorageProfileResponseOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput)
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput).ToStorageProfileResponsePtrOutputWithContext(ctx)
}









type StorageProfileResponsePtrInput interface {
	pulumi.Input

	ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput
	ToStorageProfileResponsePtrOutputWithContext(context.Context) StorageProfileResponsePtrOutput
}

type storageProfileResponsePtrType StorageProfileResponseArgs

func StorageProfileResponsePtr(v *StorageProfileResponseArgs) StorageProfileResponsePtrInput {
	return (*storageProfileResponsePtrType)(v)
}

func (*storageProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponsePtrOutput)
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfileResponse) *StorageProfileResponse {
		return &v
	}).(StorageProfileResponsePtrOutput)
}

func (o StorageProfileResponseOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []DataDiskResponse { return v.DataDisks }).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponseOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *ImageReferenceResponse { return v.ImageReference }).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponseOutput) OsDisk() OsDiskResponsePtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *OsDiskResponse { return v.OsDisk }).(OsDiskResponsePtrOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) DataDisks() DataDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []DataDiskResponse {
		if v == nil {
			return nil
		}
		return v.DataDisks
	}).(DataDiskResponseArrayOutput)
}

func (o StorageProfileResponsePtrOutput) ImageReference() ImageReferenceResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *ImageReferenceResponse {
		if v == nil {
			return nil
		}
		return v.ImageReference
	}).(ImageReferenceResponsePtrOutput)
}

func (o StorageProfileResponsePtrOutput) OsDisk() OsDiskResponsePtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *OsDiskResponse {
		if v == nil {
			return nil
		}
		return v.OsDisk
	}).(OsDiskResponsePtrOutput)
}

type SubResource struct {
	Id *string `pulumi:"id"`
}





type SubResourceInput interface {
	pulumi.Input

	ToSubResourceOutput() SubResourceOutput
	ToSubResourceOutputWithContext(context.Context) SubResourceOutput
}

type SubResourceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (i SubResourceArgs) ToSubResourceOutput() SubResourceOutput {
	return i.ToSubResourceOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput)
}

func (i SubResourceArgs) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i SubResourceArgs) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceOutput).ToSubResourcePtrOutputWithContext(ctx)
}









type SubResourcePtrInput interface {
	pulumi.Input

	ToSubResourcePtrOutput() SubResourcePtrOutput
	ToSubResourcePtrOutputWithContext(context.Context) SubResourcePtrOutput
}

type subResourcePtrType SubResourceArgs

func SubResourcePtr(v *SubResourceArgs) SubResourcePtrInput {
	return (*subResourcePtrType)(v)
}

func (*subResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (i *subResourcePtrType) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return i.ToSubResourcePtrOutputWithContext(context.Background())
}

func (i *subResourcePtrType) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourcePtrOutput)
}

type SubResourceOutput struct{ *pulumi.OutputState }

func (SubResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResource)(nil)).Elem()
}

func (o SubResourceOutput) ToSubResourceOutput() SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourceOutputWithContext(ctx context.Context) SubResourceOutput {
	return o
}

func (o SubResourceOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o.ToSubResourcePtrOutputWithContext(context.Background())
}

func (o SubResourceOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResource) *SubResource {
		return &v
	}).(SubResourcePtrOutput)
}

func (o SubResourceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResource) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourcePtrOutput struct{ *pulumi.OutputState }

func (SubResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResource)(nil)).Elem()
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutput() SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) ToSubResourcePtrOutputWithContext(ctx context.Context) SubResourcePtrOutput {
	return o
}

func (o SubResourcePtrOutput) Elem() SubResourceOutput {
	return o.ApplyT(func(v *SubResource) SubResource {
		if v != nil {
			return *v
		}
		var ret SubResource
		return ret
	}).(SubResourceOutput)
}

func (o SubResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResource) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponse struct {
	Id *string `pulumi:"id"`
}





type SubResourceResponseInput interface {
	pulumi.Input

	ToSubResourceResponseOutput() SubResourceResponseOutput
	ToSubResourceResponseOutputWithContext(context.Context) SubResourceResponseOutput
}

type SubResourceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (SubResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return i.ToSubResourceResponseOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput)
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i SubResourceResponseArgs) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseOutput).ToSubResourceResponsePtrOutputWithContext(ctx)
}









type SubResourceResponsePtrInput interface {
	pulumi.Input

	ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput
	ToSubResourceResponsePtrOutputWithContext(context.Context) SubResourceResponsePtrOutput
}

type subResourceResponsePtrType SubResourceResponseArgs

func SubResourceResponsePtr(v *SubResourceResponseArgs) SubResourceResponsePtrInput {
	return (*subResourceResponsePtrType)(v)
}

func (*subResourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return i.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (i *subResourceResponsePtrType) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponsePtrOutput)
}





type SubResourceResponseArrayInput interface {
	pulumi.Input

	ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput
	ToSubResourceResponseArrayOutputWithContext(context.Context) SubResourceResponseArrayOutput
}

type SubResourceResponseArray []SubResourceResponseInput

func (SubResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (i SubResourceResponseArray) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return i.ToSubResourceResponseArrayOutputWithContext(context.Background())
}

func (i SubResourceResponseArray) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubResourceResponseArrayOutput)
}

type SubResourceResponseOutput struct{ *pulumi.OutputState }

func (SubResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutput() SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponseOutputWithContext(ctx context.Context) SubResourceResponseOutput {
	return o
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o.ToSubResourceResponsePtrOutputWithContext(context.Background())
}

func (o SubResourceResponseOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubResourceResponse) *SubResourceResponse {
		return &v
	}).(SubResourceResponsePtrOutput)
}

func (o SubResourceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubResourceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type SubResourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SubResourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutput() SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) ToSubResourceResponsePtrOutputWithContext(ctx context.Context) SubResourceResponsePtrOutput {
	return o
}

func (o SubResourceResponsePtrOutput) Elem() SubResourceResponseOutput {
	return o.ApplyT(func(v *SubResourceResponse) SubResourceResponse {
		if v != nil {
			return *v
		}
		var ret SubResourceResponse
		return ret
	}).(SubResourceResponseOutput)
}

func (o SubResourceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubResourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type SubResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SubResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubResourceResponse)(nil)).Elem()
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutput() SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) ToSubResourceResponseArrayOutputWithContext(ctx context.Context) SubResourceResponseArrayOutput {
	return o
}

func (o SubResourceResponseArrayOutput) Index(i pulumi.IntInput) SubResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubResourceResponse {
		return vs[0].([]SubResourceResponse)[vs[1].(int)]
	}).(SubResourceResponseOutput)
}

type VirtualHardDisk struct {
	Uri *string `pulumi:"uri"`
}





type VirtualHardDiskInput interface {
	pulumi.Input

	ToVirtualHardDiskOutput() VirtualHardDiskOutput
	ToVirtualHardDiskOutputWithContext(context.Context) VirtualHardDiskOutput
}

type VirtualHardDiskArgs struct {
	Uri pulumi.StringPtrInput `pulumi:"uri"`
}

func (VirtualHardDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return i.ToVirtualHardDiskOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput)
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i VirtualHardDiskArgs) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskOutput).ToVirtualHardDiskPtrOutputWithContext(ctx)
}









type VirtualHardDiskPtrInput interface {
	pulumi.Input

	ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput
	ToVirtualHardDiskPtrOutputWithContext(context.Context) VirtualHardDiskPtrOutput
}

type virtualHardDiskPtrType VirtualHardDiskArgs

func VirtualHardDiskPtr(v *VirtualHardDiskArgs) VirtualHardDiskPtrInput {
	return (*virtualHardDiskPtrType)(v)
}

func (*virtualHardDiskPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return i.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (i *virtualHardDiskPtrType) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHardDiskPtrOutput)
}

type VirtualHardDiskOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutput() VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskOutputWithContext(ctx context.Context) VirtualHardDiskOutput {
	return o
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o.ToVirtualHardDiskPtrOutputWithContext(context.Background())
}

func (o VirtualHardDiskOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualHardDisk) *VirtualHardDisk {
		return &v
	}).(VirtualHardDiskPtrOutput)
}

func (o VirtualHardDiskOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualHardDisk) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type VirtualHardDiskPtrOutput struct{ *pulumi.OutputState }

func (VirtualHardDiskPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHardDisk)(nil)).Elem()
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutput() VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) ToVirtualHardDiskPtrOutputWithContext(ctx context.Context) VirtualHardDiskPtrOutput {
	return o
}

func (o VirtualHardDiskPtrOutput) Elem() VirtualHardDiskOutput {
	return o.ApplyT(func(v *VirtualHardDisk) VirtualHardDisk {
		if v != nil {
			return *v
		}
		var ret VirtualHardDisk
		return ret
	}).(VirtualHardDiskOutput)
}

func (o VirtualHardDiskPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHardDisk) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomProfileOutput{})
	pulumi.RegisterOutputType(CustomProfilePtrOutput{})
	pulumi.RegisterOutputType(CustomProfileResponseOutput{})
	pulumi.RegisterOutputType(CustomProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskOutput{})
	pulumi.RegisterOutputType(DataDiskArrayOutput{})
	pulumi.RegisterOutputType(DataDiskResponseOutput{})
	pulumi.RegisterOutputType(DataDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageReferenceOutput{})
	pulumi.RegisterOutputType(ImageReferencePtrOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(ImageReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LinuxConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkFunctionRoleConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTemplateOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTemplatePtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTemplateResponseOutput{})
	pulumi.RegisterOutputType(NetworkFunctionTemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationOsProfileOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationOsProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationResponseOsProfileOutput{})
	pulumi.RegisterOutputType(NetworkFunctionUserConfigurationResponseOsProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceIPConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseArrayOutput{})
	pulumi.RegisterOutputType(OsDiskOutput{})
	pulumi.RegisterOutputType(OsDiskPtrOutput{})
	pulumi.RegisterOutputType(OsDiskResponseOutput{})
	pulumi.RegisterOutputType(OsDiskResponsePtrOutput{})
	pulumi.RegisterOutputType(OsProfileOutput{})
	pulumi.RegisterOutputType(OsProfilePtrOutput{})
	pulumi.RegisterOutputType(OsProfileResponseOutput{})
	pulumi.RegisterOutputType(OsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationOutput{})
	pulumi.RegisterOutputType(SshConfigurationPtrOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponseOutput{})
	pulumi.RegisterOutputType(SshConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SshPublicKeyOutput{})
	pulumi.RegisterOutputType(SshPublicKeyArrayOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseOutput{})
	pulumi.RegisterOutputType(SshPublicKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceOutput{})
	pulumi.RegisterOutputType(SubResourcePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseOutput{})
	pulumi.RegisterOutputType(SubResourceResponsePtrOutput{})
	pulumi.RegisterOutputType(SubResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskOutput{})
	pulumi.RegisterOutputType(VirtualHardDiskPtrOutput{})
}
