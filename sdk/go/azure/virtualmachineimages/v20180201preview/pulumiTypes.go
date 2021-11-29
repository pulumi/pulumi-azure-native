


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImageTemplateIsoSource struct {
	Sha256Checksum string `pulumi:"sha256Checksum"`
	SourceURI      string `pulumi:"sourceURI"`
	Type           string `pulumi:"type"`
}





type ImageTemplateIsoSourceInput interface {
	pulumi.Input

	ToImageTemplateIsoSourceOutput() ImageTemplateIsoSourceOutput
	ToImageTemplateIsoSourceOutputWithContext(context.Context) ImageTemplateIsoSourceOutput
}

type ImageTemplateIsoSourceArgs struct {
	Sha256Checksum pulumi.StringInput `pulumi:"sha256Checksum"`
	SourceURI      pulumi.StringInput `pulumi:"sourceURI"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateIsoSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIsoSource)(nil)).Elem()
}

func (i ImageTemplateIsoSourceArgs) ToImageTemplateIsoSourceOutput() ImageTemplateIsoSourceOutput {
	return i.ToImageTemplateIsoSourceOutputWithContext(context.Background())
}

func (i ImageTemplateIsoSourceArgs) ToImageTemplateIsoSourceOutputWithContext(ctx context.Context) ImageTemplateIsoSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIsoSourceOutput)
}

type ImageTemplateIsoSourceOutput struct{ *pulumi.OutputState }

func (ImageTemplateIsoSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIsoSource)(nil)).Elem()
}

func (o ImageTemplateIsoSourceOutput) ToImageTemplateIsoSourceOutput() ImageTemplateIsoSourceOutput {
	return o
}

func (o ImageTemplateIsoSourceOutput) ToImageTemplateIsoSourceOutputWithContext(ctx context.Context) ImageTemplateIsoSourceOutput {
	return o
}

func (o ImageTemplateIsoSourceOutput) Sha256Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSource) string { return v.Sha256Checksum }).(pulumi.StringOutput)
}

func (o ImageTemplateIsoSourceOutput) SourceURI() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSource) string { return v.SourceURI }).(pulumi.StringOutput)
}

func (o ImageTemplateIsoSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSource) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateIsoSourceResponse struct {
	Sha256Checksum string `pulumi:"sha256Checksum"`
	SourceURI      string `pulumi:"sourceURI"`
	Type           string `pulumi:"type"`
}





type ImageTemplateIsoSourceResponseInput interface {
	pulumi.Input

	ToImageTemplateIsoSourceResponseOutput() ImageTemplateIsoSourceResponseOutput
	ToImageTemplateIsoSourceResponseOutputWithContext(context.Context) ImageTemplateIsoSourceResponseOutput
}

type ImageTemplateIsoSourceResponseArgs struct {
	Sha256Checksum pulumi.StringInput `pulumi:"sha256Checksum"`
	SourceURI      pulumi.StringInput `pulumi:"sourceURI"`
	Type           pulumi.StringInput `pulumi:"type"`
}

func (ImageTemplateIsoSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIsoSourceResponse)(nil)).Elem()
}

func (i ImageTemplateIsoSourceResponseArgs) ToImageTemplateIsoSourceResponseOutput() ImageTemplateIsoSourceResponseOutput {
	return i.ToImageTemplateIsoSourceResponseOutputWithContext(context.Background())
}

func (i ImageTemplateIsoSourceResponseArgs) ToImageTemplateIsoSourceResponseOutputWithContext(ctx context.Context) ImageTemplateIsoSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIsoSourceResponseOutput)
}

type ImageTemplateIsoSourceResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateIsoSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIsoSourceResponse)(nil)).Elem()
}

func (o ImageTemplateIsoSourceResponseOutput) ToImageTemplateIsoSourceResponseOutput() ImageTemplateIsoSourceResponseOutput {
	return o
}

func (o ImageTemplateIsoSourceResponseOutput) ToImageTemplateIsoSourceResponseOutputWithContext(ctx context.Context) ImageTemplateIsoSourceResponseOutput {
	return o
}

func (o ImageTemplateIsoSourceResponseOutput) Sha256Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSourceResponse) string { return v.Sha256Checksum }).(pulumi.StringOutput)
}

func (o ImageTemplateIsoSourceResponseOutput) SourceURI() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSourceResponse) string { return v.SourceURI }).(pulumi.StringOutput)
}

func (o ImageTemplateIsoSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIsoSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateLastRunStatusResponse struct {
	EndTime     *string `pulumi:"endTime"`
	Message     *string `pulumi:"message"`
	RunState    *string `pulumi:"runState"`
	RunSubState *string `pulumi:"runSubState"`
	StartTime   *string `pulumi:"startTime"`
}





type ImageTemplateLastRunStatusResponseInput interface {
	pulumi.Input

	ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput
	ToImageTemplateLastRunStatusResponseOutputWithContext(context.Context) ImageTemplateLastRunStatusResponseOutput
}

type ImageTemplateLastRunStatusResponseArgs struct {
	EndTime     pulumi.StringPtrInput `pulumi:"endTime"`
	Message     pulumi.StringPtrInput `pulumi:"message"`
	RunState    pulumi.StringPtrInput `pulumi:"runState"`
	RunSubState pulumi.StringPtrInput `pulumi:"runSubState"`
	StartTime   pulumi.StringPtrInput `pulumi:"startTime"`
}

func (ImageTemplateLastRunStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput {
	return i.ToImageTemplateLastRunStatusResponseOutputWithContext(context.Background())
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponseOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponseOutput)
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return i.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (i ImageTemplateLastRunStatusResponseArgs) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponseOutput).ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx)
}









type ImageTemplateLastRunStatusResponsePtrInput interface {
	pulumi.Input

	ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput
	ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Context) ImageTemplateLastRunStatusResponsePtrOutput
}

type imageTemplateLastRunStatusResponsePtrType ImageTemplateLastRunStatusResponseArgs

func ImageTemplateLastRunStatusResponsePtr(v *ImageTemplateLastRunStatusResponseArgs) ImageTemplateLastRunStatusResponsePtrInput {
	return (*imageTemplateLastRunStatusResponsePtrType)(v)
}

func (*imageTemplateLastRunStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (i *imageTemplateLastRunStatusResponsePtrType) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return i.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (i *imageTemplateLastRunStatusResponsePtrType) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateLastRunStatusResponsePtrOutput)
}

type ImageTemplateLastRunStatusResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateLastRunStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutput() ImageTemplateLastRunStatusResponseOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponseOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponseOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return o.ToImageTemplateLastRunStatusResponsePtrOutputWithContext(context.Background())
}

func (o ImageTemplateLastRunStatusResponseOutput) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateLastRunStatusResponse) *ImageTemplateLastRunStatusResponse {
		return &v
	}).(ImageTemplateLastRunStatusResponsePtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) RunState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunState }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) RunSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.RunSubState }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateLastRunStatusResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type ImageTemplateLastRunStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateLastRunStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateLastRunStatusResponse)(nil)).Elem()
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) ToImageTemplateLastRunStatusResponsePtrOutput() ImageTemplateLastRunStatusResponsePtrOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) ToImageTemplateLastRunStatusResponsePtrOutputWithContext(ctx context.Context) ImageTemplateLastRunStatusResponsePtrOutput {
	return o
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) Elem() ImageTemplateLastRunStatusResponseOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) ImageTemplateLastRunStatusResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateLastRunStatusResponse
		return ret
	}).(ImageTemplateLastRunStatusResponseOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) RunState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunState
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) RunSubState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunSubState
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateLastRunStatusResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateLastRunStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type ImageTemplateManagedImageDistributor struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateManagedImageDistributorInput interface {
	pulumi.Input

	ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput
	ToImageTemplateManagedImageDistributorOutputWithContext(context.Context) ImageTemplateManagedImageDistributorOutput
}

type ImageTemplateManagedImageDistributorArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	ImageId       pulumi.StringInput    `pulumi:"imageId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateManagedImageDistributorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributor)(nil)).Elem()
}

func (i ImageTemplateManagedImageDistributorArgs) ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput {
	return i.ToImageTemplateManagedImageDistributorOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageDistributorArgs) ToImageTemplateManagedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageDistributorOutput)
}

type ImageTemplateManagedImageDistributorOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageDistributorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributor)(nil)).Elem()
}

func (o ImageTemplateManagedImageDistributorOutput) ToImageTemplateManagedImageDistributorOutput() ImageTemplateManagedImageDistributorOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorOutput) ToImageTemplateManagedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.Location }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributor) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateManagedImageDistributorResponse struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}





type ImageTemplateManagedImageDistributorResponseInput interface {
	pulumi.Input

	ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput
	ToImageTemplateManagedImageDistributorResponseOutputWithContext(context.Context) ImageTemplateManagedImageDistributorResponseOutput
}

type ImageTemplateManagedImageDistributorResponseArgs struct {
	ArtifactTags  pulumi.StringMapInput `pulumi:"artifactTags"`
	ImageId       pulumi.StringInput    `pulumi:"imageId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	RunOutputName pulumi.StringInput    `pulumi:"runOutputName"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateManagedImageDistributorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributorResponse)(nil)).Elem()
}

func (i ImageTemplateManagedImageDistributorResponseArgs) ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput {
	return i.ToImageTemplateManagedImageDistributorResponseOutputWithContext(context.Background())
}

func (i ImageTemplateManagedImageDistributorResponseArgs) ToImageTemplateManagedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateManagedImageDistributorResponseOutput)
}

type ImageTemplateManagedImageDistributorResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateManagedImageDistributorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateManagedImageDistributorResponse)(nil)).Elem()
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ToImageTemplateManagedImageDistributorResponseOutput() ImageTemplateManagedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ToImageTemplateManagedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateManagedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.ImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateManagedImageDistributorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateManagedImageDistributorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplatePlatformImageSource struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Type      string  `pulumi:"type"`
	Version   *string `pulumi:"version"`
}





type ImageTemplatePlatformImageSourceInput interface {
	pulumi.Input

	ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput
	ToImageTemplatePlatformImageSourceOutputWithContext(context.Context) ImageTemplatePlatformImageSourceOutput
}

type ImageTemplatePlatformImageSourceArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Type      pulumi.StringInput    `pulumi:"type"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageTemplatePlatformImageSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSource)(nil)).Elem()
}

func (i ImageTemplatePlatformImageSourceArgs) ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput {
	return i.ToImageTemplatePlatformImageSourceOutputWithContext(context.Background())
}

func (i ImageTemplatePlatformImageSourceArgs) ToImageTemplatePlatformImageSourceOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePlatformImageSourceOutput)
}

type ImageTemplatePlatformImageSourceOutput struct{ *pulumi.OutputState }

func (ImageTemplatePlatformImageSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSource)(nil)).Elem()
}

func (o ImageTemplatePlatformImageSourceOutput) ToImageTemplatePlatformImageSourceOutput() ImageTemplatePlatformImageSourceOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceOutput) ToImageTemplatePlatformImageSourceOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePlatformImageSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageTemplatePlatformImageSourceResponse struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Type      string  `pulumi:"type"`
	Version   *string `pulumi:"version"`
}





type ImageTemplatePlatformImageSourceResponseInput interface {
	pulumi.Input

	ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput
	ToImageTemplatePlatformImageSourceResponseOutputWithContext(context.Context) ImageTemplatePlatformImageSourceResponseOutput
}

type ImageTemplatePlatformImageSourceResponseArgs struct {
	Offer     pulumi.StringPtrInput `pulumi:"offer"`
	Publisher pulumi.StringPtrInput `pulumi:"publisher"`
	Sku       pulumi.StringPtrInput `pulumi:"sku"`
	Type      pulumi.StringInput    `pulumi:"type"`
	Version   pulumi.StringPtrInput `pulumi:"version"`
}

func (ImageTemplatePlatformImageSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSourceResponse)(nil)).Elem()
}

func (i ImageTemplatePlatformImageSourceResponseArgs) ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput {
	return i.ToImageTemplatePlatformImageSourceResponseOutputWithContext(context.Background())
}

func (i ImageTemplatePlatformImageSourceResponseArgs) ToImageTemplatePlatformImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplatePlatformImageSourceResponseOutput)
}

type ImageTemplatePlatformImageSourceResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplatePlatformImageSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplatePlatformImageSourceResponse)(nil)).Elem()
}

func (o ImageTemplatePlatformImageSourceResponseOutput) ToImageTemplatePlatformImageSourceResponseOutput() ImageTemplatePlatformImageSourceResponseOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceResponseOutput) ToImageTemplatePlatformImageSourceResponseOutputWithContext(ctx context.Context) ImageTemplatePlatformImageSourceResponseOutput {
	return o
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ImageTemplatePlatformImageSourceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplatePlatformImageSourceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ImageTemplateSharedImageDistributor struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	Type               string            `pulumi:"type"`
}





type ImageTemplateSharedImageDistributorInput interface {
	pulumi.Input

	ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput
	ToImageTemplateSharedImageDistributorOutputWithContext(context.Context) ImageTemplateSharedImageDistributorOutput
}

type ImageTemplateSharedImageDistributorArgs struct {
	ArtifactTags       pulumi.StringMapInput   `pulumi:"artifactTags"`
	GalleryImageId     pulumi.StringInput      `pulumi:"galleryImageId"`
	ReplicationRegions pulumi.StringArrayInput `pulumi:"replicationRegions"`
	RunOutputName      pulumi.StringInput      `pulumi:"runOutputName"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateSharedImageDistributorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributor)(nil)).Elem()
}

func (i ImageTemplateSharedImageDistributorArgs) ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput {
	return i.ToImageTemplateSharedImageDistributorOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageDistributorArgs) ToImageTemplateSharedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageDistributorOutput)
}

type ImageTemplateSharedImageDistributorOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageDistributorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributor)(nil)).Elem()
}

func (o ImageTemplateSharedImageDistributorOutput) ToImageTemplateSharedImageDistributorOutput() ImageTemplateSharedImageDistributorOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorOutput) ToImageTemplateSharedImageDistributorOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) GalleryImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.GalleryImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) ReplicationRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) []string { return v.ReplicationRegions }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributor) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateSharedImageDistributorResponse struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	Type               string            `pulumi:"type"`
}





type ImageTemplateSharedImageDistributorResponseInput interface {
	pulumi.Input

	ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput
	ToImageTemplateSharedImageDistributorResponseOutputWithContext(context.Context) ImageTemplateSharedImageDistributorResponseOutput
}

type ImageTemplateSharedImageDistributorResponseArgs struct {
	ArtifactTags       pulumi.StringMapInput   `pulumi:"artifactTags"`
	GalleryImageId     pulumi.StringInput      `pulumi:"galleryImageId"`
	ReplicationRegions pulumi.StringArrayInput `pulumi:"replicationRegions"`
	RunOutputName      pulumi.StringInput      `pulumi:"runOutputName"`
	Type               pulumi.StringInput      `pulumi:"type"`
}

func (ImageTemplateSharedImageDistributorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributorResponse)(nil)).Elem()
}

func (i ImageTemplateSharedImageDistributorResponseArgs) ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput {
	return i.ToImageTemplateSharedImageDistributorResponseOutputWithContext(context.Background())
}

func (i ImageTemplateSharedImageDistributorResponseArgs) ToImageTemplateSharedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateSharedImageDistributorResponseOutput)
}

type ImageTemplateSharedImageDistributorResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateSharedImageDistributorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateSharedImageDistributorResponse)(nil)).Elem()
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ToImageTemplateSharedImageDistributorResponseOutput() ImageTemplateSharedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ToImageTemplateSharedImageDistributorResponseOutputWithContext(ctx context.Context) ImageTemplateSharedImageDistributorResponseOutput {
	return o
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ArtifactTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) map[string]string { return v.ArtifactTags }).(pulumi.StringMapOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) GalleryImageId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.GalleryImageId }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) ReplicationRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) []string { return v.ReplicationRegions }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) RunOutputName() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.RunOutputName }).(pulumi.StringOutput)
}

func (o ImageTemplateSharedImageDistributorResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateSharedImageDistributorResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateShellCustomizer struct {
	Name   *string `pulumi:"name"`
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}





type ImageTemplateShellCustomizerInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput
	ToImageTemplateShellCustomizerOutputWithContext(context.Context) ImageTemplateShellCustomizerOutput
}

type ImageTemplateShellCustomizerArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Script pulumi.StringPtrInput `pulumi:"script"`
	Type   pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateShellCustomizerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizer)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerArgs) ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput {
	return i.ToImageTemplateShellCustomizerOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerArgs) ToImageTemplateShellCustomizerOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerOutput)
}





type ImageTemplateShellCustomizerArrayInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerArrayOutput() ImageTemplateShellCustomizerArrayOutput
	ToImageTemplateShellCustomizerArrayOutputWithContext(context.Context) ImageTemplateShellCustomizerArrayOutput
}

type ImageTemplateShellCustomizerArray []ImageTemplateShellCustomizerInput

func (ImageTemplateShellCustomizerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageTemplateShellCustomizer)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerArray) ToImageTemplateShellCustomizerArrayOutput() ImageTemplateShellCustomizerArrayOutput {
	return i.ToImageTemplateShellCustomizerArrayOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerArray) ToImageTemplateShellCustomizerArrayOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerArrayOutput)
}

type ImageTemplateShellCustomizerOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizer)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerOutput) ToImageTemplateShellCustomizerOutput() ImageTemplateShellCustomizerOutput {
	return o
}

func (o ImageTemplateShellCustomizerOutput) ToImageTemplateShellCustomizerOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerOutput {
	return o
}

func (o ImageTemplateShellCustomizerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizer) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateShellCustomizerArrayOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageTemplateShellCustomizer)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerArrayOutput) ToImageTemplateShellCustomizerArrayOutput() ImageTemplateShellCustomizerArrayOutput {
	return o
}

func (o ImageTemplateShellCustomizerArrayOutput) ToImageTemplateShellCustomizerArrayOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerArrayOutput {
	return o
}

func (o ImageTemplateShellCustomizerArrayOutput) Index(i pulumi.IntInput) ImageTemplateShellCustomizerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageTemplateShellCustomizer {
		return vs[0].([]ImageTemplateShellCustomizer)[vs[1].(int)]
	}).(ImageTemplateShellCustomizerOutput)
}

type ImageTemplateShellCustomizerResponse struct {
	Name   *string `pulumi:"name"`
	Script *string `pulumi:"script"`
	Type   string  `pulumi:"type"`
}





type ImageTemplateShellCustomizerResponseInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput
	ToImageTemplateShellCustomizerResponseOutputWithContext(context.Context) ImageTemplateShellCustomizerResponseOutput
}

type ImageTemplateShellCustomizerResponseArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Script pulumi.StringPtrInput `pulumi:"script"`
	Type   pulumi.StringInput    `pulumi:"type"`
}

func (ImageTemplateShellCustomizerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerResponseArgs) ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput {
	return i.ToImageTemplateShellCustomizerResponseOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerResponseArgs) ToImageTemplateShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerResponseOutput)
}





type ImageTemplateShellCustomizerResponseArrayInput interface {
	pulumi.Input

	ToImageTemplateShellCustomizerResponseArrayOutput() ImageTemplateShellCustomizerResponseArrayOutput
	ToImageTemplateShellCustomizerResponseArrayOutputWithContext(context.Context) ImageTemplateShellCustomizerResponseArrayOutput
}

type ImageTemplateShellCustomizerResponseArray []ImageTemplateShellCustomizerResponseInput

func (ImageTemplateShellCustomizerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (i ImageTemplateShellCustomizerResponseArray) ToImageTemplateShellCustomizerResponseArrayOutput() ImageTemplateShellCustomizerResponseArrayOutput {
	return i.ToImageTemplateShellCustomizerResponseArrayOutputWithContext(context.Background())
}

func (i ImageTemplateShellCustomizerResponseArray) ToImageTemplateShellCustomizerResponseArrayOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateShellCustomizerResponseArrayOutput)
}

type ImageTemplateShellCustomizerResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerResponseOutput) ToImageTemplateShellCustomizerResponseOutput() ImageTemplateShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseOutput) ToImageTemplateShellCustomizerResponseOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) *string { return v.Script }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateShellCustomizerResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateShellCustomizerResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ImageTemplateShellCustomizerResponseArrayOutput struct{ *pulumi.OutputState }

func (ImageTemplateShellCustomizerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ImageTemplateShellCustomizerResponse)(nil)).Elem()
}

func (o ImageTemplateShellCustomizerResponseArrayOutput) ToImageTemplateShellCustomizerResponseArrayOutput() ImageTemplateShellCustomizerResponseArrayOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseArrayOutput) ToImageTemplateShellCustomizerResponseArrayOutputWithContext(ctx context.Context) ImageTemplateShellCustomizerResponseArrayOutput {
	return o
}

func (o ImageTemplateShellCustomizerResponseArrayOutput) Index(i pulumi.IntInput) ImageTemplateShellCustomizerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ImageTemplateShellCustomizerResponse {
		return vs[0].([]ImageTemplateShellCustomizerResponse)[vs[1].(int)]
	}).(ImageTemplateShellCustomizerResponseOutput)
}

type ProvisioningErrorResponse struct {
	Message               *string `pulumi:"message"`
	ProvisioningErrorCode *string `pulumi:"provisioningErrorCode"`
}





type ProvisioningErrorResponseInput interface {
	pulumi.Input

	ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput
	ToProvisioningErrorResponseOutputWithContext(context.Context) ProvisioningErrorResponseOutput
}

type ProvisioningErrorResponseArgs struct {
	Message               pulumi.StringPtrInput `pulumi:"message"`
	ProvisioningErrorCode pulumi.StringPtrInput `pulumi:"provisioningErrorCode"`
}

func (ProvisioningErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningErrorResponse)(nil)).Elem()
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput {
	return i.ToProvisioningErrorResponseOutputWithContext(context.Background())
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponseOutputWithContext(ctx context.Context) ProvisioningErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponseOutput)
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return i.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (i ProvisioningErrorResponseArgs) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponseOutput).ToProvisioningErrorResponsePtrOutputWithContext(ctx)
}









type ProvisioningErrorResponsePtrInput interface {
	pulumi.Input

	ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput
	ToProvisioningErrorResponsePtrOutputWithContext(context.Context) ProvisioningErrorResponsePtrOutput
}

type provisioningErrorResponsePtrType ProvisioningErrorResponseArgs

func ProvisioningErrorResponsePtr(v *ProvisioningErrorResponseArgs) ProvisioningErrorResponsePtrInput {
	return (*provisioningErrorResponsePtrType)(v)
}

func (*provisioningErrorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningErrorResponse)(nil)).Elem()
}

func (i *provisioningErrorResponsePtrType) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return i.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (i *provisioningErrorResponsePtrType) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProvisioningErrorResponsePtrOutput)
}

type ProvisioningErrorResponseOutput struct{ *pulumi.OutputState }

func (ProvisioningErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningErrorResponse)(nil)).Elem()
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutput() ProvisioningErrorResponseOutput {
	return o
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponseOutputWithContext(ctx context.Context) ProvisioningErrorResponseOutput {
	return o
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return o.ToProvisioningErrorResponsePtrOutputWithContext(context.Background())
}

func (o ProvisioningErrorResponseOutput) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningErrorResponse) *ProvisioningErrorResponse {
		return &v
	}).(ProvisioningErrorResponsePtrOutput)
}

func (o ProvisioningErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ProvisioningErrorResponseOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.ProvisioningErrorCode }).(pulumi.StringPtrOutput)
}

type ProvisioningErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningErrorResponse)(nil)).Elem()
}

func (o ProvisioningErrorResponsePtrOutput) ToProvisioningErrorResponsePtrOutput() ProvisioningErrorResponsePtrOutput {
	return o
}

func (o ProvisioningErrorResponsePtrOutput) ToProvisioningErrorResponsePtrOutputWithContext(ctx context.Context) ProvisioningErrorResponsePtrOutput {
	return o
}

func (o ProvisioningErrorResponsePtrOutput) Elem() ProvisioningErrorResponseOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) ProvisioningErrorResponse {
		if v != nil {
			return *v
		}
		var ret ProvisioningErrorResponse
		return ret
	}).(ProvisioningErrorResponseOutput)
}

func (o ProvisioningErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ProvisioningErrorResponsePtrOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProvisioningErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningErrorCode
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageTemplateIsoSourceOutput{})
	pulumi.RegisterOutputType(ImageTemplateIsoSourceResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageDistributorOutput{})
	pulumi.RegisterOutputType(ImageTemplateManagedImageDistributorResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplatePlatformImageSourceOutput{})
	pulumi.RegisterOutputType(ImageTemplatePlatformImageSourceResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageDistributorOutput{})
	pulumi.RegisterOutputType(ImageTemplateSharedImageDistributorResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerArrayOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerResponseArrayOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponsePtrOutput{})
}
