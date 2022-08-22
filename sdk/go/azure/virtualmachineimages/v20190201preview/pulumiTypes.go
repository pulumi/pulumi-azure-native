


package v20190201preview

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

type ImageTemplateManagedImageSource struct {
	ImageId string `pulumi:"imageId"`
	Type    string `pulumi:"type"`
}

type ImageTemplateManagedImageSourceResponse struct {
	ImageId string `pulumi:"imageId"`
	Type    string `pulumi:"type"`
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

type ImageTemplatePowerShellCustomizer struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	Script         *string  `pulumi:"script"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}

type ImageTemplatePowerShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	Script         *string  `pulumi:"script"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}

type ImageTemplateRestartCustomizer struct {
	Name                *string `pulumi:"name"`
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	RestartCommand      *string `pulumi:"restartCommand"`
	RestartTimeout      *string `pulumi:"restartTimeout"`
	Type                string  `pulumi:"type"`
}

type ImageTemplateRestartCustomizerResponse struct {
	Name                *string `pulumi:"name"`
	RestartCheckCommand *string `pulumi:"restartCheckCommand"`
	RestartCommand      *string `pulumi:"restartCommand"`
	RestartTimeout      *string `pulumi:"restartTimeout"`
	Type                string  `pulumi:"type"`
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
	Inline []string `pulumi:"inline"`
	Name   *string  `pulumi:"name"`
	Script *string  `pulumi:"script"`
	Type   string   `pulumi:"type"`
}

type ImageTemplateShellCustomizerResponse struct {
	Inline []string `pulumi:"inline"`
	Name   *string  `pulumi:"name"`
	Script *string  `pulumi:"script"`
	Type   string   `pulumi:"type"`
}

type ImageTemplateVhdDistributor struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
}

type ImageTemplateVhdDistributorResponse struct {
	ArtifactTags  map[string]string `pulumi:"artifactTags"`
	RunOutputName string            `pulumi:"runOutputName"`
	Type          string            `pulumi:"type"`
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
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
}
