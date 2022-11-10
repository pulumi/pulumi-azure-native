


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

type ImageTemplateIsoSourceResponse struct {
	Sha256Checksum string `pulumi:"sha256Checksum"`
	SourceURI      string `pulumi:"sourceURI"`
	Type           string `pulumi:"type"`
}

type ImageTemplateLastRunStatusResponse struct {
	EndTime     *string `pulumi:"endTime"`
	Message     *string `pulumi:"message"`
	RunState    *string `pulumi:"runState"`
	RunSubState *string `pulumi:"runSubState"`
	StartTime   *string `pulumi:"startTime"`
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

type ImageTemplateManagedImageDistributor struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}

type ImageTemplateManagedImageDistributorResponse struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	ImageId       string            `pulumi:"imageId"`
	Location      string            `pulumi:"location"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}

type ImageTemplatePlatformImageSource struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Type      string  `pulumi:"type"`
	Version   *string `pulumi:"version"`
}

type ImageTemplatePlatformImageSourceResponse struct {
	Offer     *string `pulumi:"offer"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Type      string  `pulumi:"type"`
	Version   *string `pulumi:"version"`
}

type ImageTemplateSharedImageDistributor struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	Type               string            `pulumi:"type"`
}

type ImageTemplateSharedImageDistributorResponse struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	Type               string            `pulumi:"type"`
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

func (o ProvisioningErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ProvisioningErrorResponseOutput) ProvisioningErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProvisioningErrorResponse) *string { return v.ProvisioningErrorCode }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerArrayOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateShellCustomizerResponseArrayOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
}
