


package v20190501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImageTemplateFileCustomizer struct {
	Destination    *string `pulumi:"destination"`
	Name           *string `pulumi:"name"`
	Sha256Checksum *string `pulumi:"sha256Checksum"`
	SourceUri      *string `pulumi:"sourceUri"`
	Type           string  `pulumi:"type"`
}

type ImageTemplateFileCustomizerResponse struct {
	Destination    *string `pulumi:"destination"`
	Name           *string `pulumi:"name"`
	Sha256Checksum *string `pulumi:"sha256Checksum"`
	SourceUri      *string `pulumi:"sourceUri"`
	Type           string  `pulumi:"type"`
}

type ImageTemplateIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ImageTemplateIdentityInput interface {
	pulumi.Input

	ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput
	ToImageTemplateIdentityOutputWithContext(context.Context) ImageTemplateIdentityOutput
}

type ImageTemplateIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ImageTemplateIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentity)(nil)).Elem()
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput {
	return i.ToImageTemplateIdentityOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityOutputWithContext(ctx context.Context) ImageTemplateIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityOutput)
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return i.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (i ImageTemplateIdentityArgs) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityOutput).ToImageTemplateIdentityPtrOutputWithContext(ctx)
}









type ImageTemplateIdentityPtrInput interface {
	pulumi.Input

	ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput
	ToImageTemplateIdentityPtrOutputWithContext(context.Context) ImageTemplateIdentityPtrOutput
}

type imageTemplateIdentityPtrType ImageTemplateIdentityArgs

func ImageTemplateIdentityPtr(v *ImageTemplateIdentityArgs) ImageTemplateIdentityPtrInput {
	return (*imageTemplateIdentityPtrType)(v)
}

func (*imageTemplateIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentity)(nil)).Elem()
}

func (i *imageTemplateIdentityPtrType) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return i.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (i *imageTemplateIdentityPtrType) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateIdentityPtrOutput)
}

type ImageTemplateIdentityOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentity)(nil)).Elem()
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityOutput() ImageTemplateIdentityOutput {
	return o
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityOutputWithContext(ctx context.Context) ImageTemplateIdentityOutput {
	return o
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return o.ToImageTemplateIdentityPtrOutputWithContext(context.Background())
}

func (o ImageTemplateIdentityOutput) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateIdentity) *ImageTemplateIdentity {
		return &v
	}).(ImageTemplateIdentityPtrOutput)
}

func (o ImageTemplateIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ImageTemplateIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ImageTemplateIdentityPtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentity)(nil)).Elem()
}

func (o ImageTemplateIdentityPtrOutput) ToImageTemplateIdentityPtrOutput() ImageTemplateIdentityPtrOutput {
	return o
}

func (o ImageTemplateIdentityPtrOutput) ToImageTemplateIdentityPtrOutputWithContext(ctx context.Context) ImageTemplateIdentityPtrOutput {
	return o
}

func (o ImageTemplateIdentityPtrOutput) Elem() ImageTemplateIdentityOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) ImageTemplateIdentity {
		if v != nil {
			return *v
		}
		var ret ImageTemplateIdentity
		return ret
	}).(ImageTemplateIdentityOutput)
}

func (o ImageTemplateIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ImageTemplateIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ImageTemplateIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ImageTemplateIdentityResponse struct {
	Type                   *string                                                        `pulumi:"type"`
	UserAssignedIdentities map[string]ImageTemplateIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}

type ImageTemplateIdentityResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponse)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponseOutput() ImageTemplateIdentityResponseOutput {
	return o
}

func (o ImageTemplateIdentityResponseOutput) ToImageTemplateIdentityResponseOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseOutput {
	return o
}

func (o ImageTemplateIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateIdentityResponseOutput) UserAssignedIdentities() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponse) map[string]ImageTemplateIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ImageTemplateIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateIdentityResponse)(nil)).Elem()
}

func (o ImageTemplateIdentityResponsePtrOutput) ToImageTemplateIdentityResponsePtrOutput() ImageTemplateIdentityResponsePtrOutput {
	return o
}

func (o ImageTemplateIdentityResponsePtrOutput) ToImageTemplateIdentityResponsePtrOutputWithContext(ctx context.Context) ImageTemplateIdentityResponsePtrOutput {
	return o
}

func (o ImageTemplateIdentityResponsePtrOutput) Elem() ImageTemplateIdentityResponseOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) ImageTemplateIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateIdentityResponse
		return ret
	}).(ImageTemplateIdentityResponseOutput)
}

func (o ImageTemplateIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateIdentityResponsePtrOutput) UserAssignedIdentities() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ImageTemplateIdentityResponse) map[string]ImageTemplateIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ImageTemplateIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ImageTemplateIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ImageTemplateIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput() ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) ToImageTemplateIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ImageTemplateIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ImageTemplateIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ImageTemplateIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput)
}

type ImageTemplateIsoSource struct {
	Sha256Checksum string `pulumi:"sha256Checksum"`
	SourceUri      string `pulumi:"sourceUri"`
	Type           string `pulumi:"type"`
}

type ImageTemplateIsoSourceResponse struct {
	Sha256Checksum string `pulumi:"sha256Checksum"`
	SourceUri      string `pulumi:"sourceUri"`
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
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}

type ImageTemplatePowerShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
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

type ImageTemplateSharedImageVersionSource struct {
	ImageVersionId string `pulumi:"imageVersionId"`
	Type           string `pulumi:"type"`
}

type ImageTemplateSharedImageVersionSourceResponse struct {
	ImageVersionId string `pulumi:"imageVersionId"`
	Type           string `pulumi:"type"`
}

type ImageTemplateShellCustomizer struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
}

type ImageTemplateShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
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

type ImageTemplateVmProfile struct {
	VmSize *string `pulumi:"vmSize"`
}





type ImageTemplateVmProfileInput interface {
	pulumi.Input

	ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput
	ToImageTemplateVmProfileOutputWithContext(context.Context) ImageTemplateVmProfileOutput
}

type ImageTemplateVmProfileArgs struct {
	VmSize pulumi.StringPtrInput `pulumi:"vmSize"`
}

func (ImageTemplateVmProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfile)(nil)).Elem()
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput {
	return i.ToImageTemplateVmProfileOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfileOutputWithContext(ctx context.Context) ImageTemplateVmProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileOutput)
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return i.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (i ImageTemplateVmProfileArgs) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfileOutput).ToImageTemplateVmProfilePtrOutputWithContext(ctx)
}









type ImageTemplateVmProfilePtrInput interface {
	pulumi.Input

	ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput
	ToImageTemplateVmProfilePtrOutputWithContext(context.Context) ImageTemplateVmProfilePtrOutput
}

type imageTemplateVmProfilePtrType ImageTemplateVmProfileArgs

func ImageTemplateVmProfilePtr(v *ImageTemplateVmProfileArgs) ImageTemplateVmProfilePtrInput {
	return (*imageTemplateVmProfilePtrType)(v)
}

func (*imageTemplateVmProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfile)(nil)).Elem()
}

func (i *imageTemplateVmProfilePtrType) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return i.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (i *imageTemplateVmProfilePtrType) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageTemplateVmProfilePtrOutput)
}

type ImageTemplateVmProfileOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfile)(nil)).Elem()
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput {
	return o
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfileOutputWithContext(ctx context.Context) ImageTemplateVmProfileOutput {
	return o
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return o.ToImageTemplateVmProfilePtrOutputWithContext(context.Background())
}

func (o ImageTemplateVmProfileOutput) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageTemplateVmProfile) *ImageTemplateVmProfile {
		return &v
	}).(ImageTemplateVmProfilePtrOutput)
}

func (o ImageTemplateVmProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ImageTemplateVmProfilePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfile)(nil)).Elem()
}

func (o ImageTemplateVmProfilePtrOutput) ToImageTemplateVmProfilePtrOutput() ImageTemplateVmProfilePtrOutput {
	return o
}

func (o ImageTemplateVmProfilePtrOutput) ToImageTemplateVmProfilePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfilePtrOutput {
	return o
}

func (o ImageTemplateVmProfilePtrOutput) Elem() ImageTemplateVmProfileOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) ImageTemplateVmProfile {
		if v != nil {
			return *v
		}
		var ret ImageTemplateVmProfile
		return ret
	}).(ImageTemplateVmProfileOutput)
}

func (o ImageTemplateVmProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

type ImageTemplateVmProfileResponse struct {
	VmSize *string `pulumi:"vmSize"`
}

type ImageTemplateVmProfileResponseOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponseOutput() ImageTemplateVmProfileResponseOutput {
	return o
}

func (o ImageTemplateVmProfileResponseOutput) ToImageTemplateVmProfileResponseOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponseOutput {
	return o
}

func (o ImageTemplateVmProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

type ImageTemplateVmProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageTemplateVmProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageTemplateVmProfileResponse)(nil)).Elem()
}

func (o ImageTemplateVmProfileResponsePtrOutput) ToImageTemplateVmProfileResponsePtrOutput() ImageTemplateVmProfileResponsePtrOutput {
	return o
}

func (o ImageTemplateVmProfileResponsePtrOutput) ToImageTemplateVmProfileResponsePtrOutputWithContext(ctx context.Context) ImageTemplateVmProfileResponsePtrOutput {
	return o
}

func (o ImageTemplateVmProfileResponsePtrOutput) Elem() ImageTemplateVmProfileResponseOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) ImageTemplateVmProfileResponse {
		if v != nil {
			return *v
		}
		var ret ImageTemplateVmProfileResponse
		return ret
	}).(ImageTemplateVmProfileResponseOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(ImageTemplateIdentityOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityPtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfilePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
}
