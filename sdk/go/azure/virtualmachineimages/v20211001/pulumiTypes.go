


package v20211001

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


func (val *ImageTemplateFileCustomizer) Defaults() *ImageTemplateFileCustomizer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
}

type ImageTemplateFileCustomizerResponse struct {
	Destination    *string `pulumi:"destination"`
	Name           *string `pulumi:"name"`
	Sha256Checksum *string `pulumi:"sha256Checksum"`
	SourceUri      *string `pulumi:"sourceUri"`
	Type           string  `pulumi:"type"`
}


func (val *ImageTemplateFileCustomizerResponse) Defaults() *ImageTemplateFileCustomizerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
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

func (o ImageTemplateIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ImageTemplateIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ImageTemplateIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
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
	Offer     *string                    `pulumi:"offer"`
	PlanInfo  *PlatformImagePurchasePlan `pulumi:"planInfo"`
	Publisher *string                    `pulumi:"publisher"`
	Sku       *string                    `pulumi:"sku"`
	Type      string                     `pulumi:"type"`
	Version   *string                    `pulumi:"version"`
}

type ImageTemplatePlatformImageSourceResponse struct {
	ExactVersion string                             `pulumi:"exactVersion"`
	Offer        *string                            `pulumi:"offer"`
	PlanInfo     *PlatformImagePurchasePlanResponse `pulumi:"planInfo"`
	Publisher    *string                            `pulumi:"publisher"`
	Sku          *string                            `pulumi:"sku"`
	Type         string                             `pulumi:"type"`
	Version      *string                            `pulumi:"version"`
}

type ImageTemplatePowerShellCustomizer struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	RunAsSystem    *bool    `pulumi:"runAsSystem"`
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}


func (val *ImageTemplatePowerShellCustomizer) Defaults() *ImageTemplatePowerShellCustomizer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RunAsSystem) {
		runAsSystem_ := false
		tmp.RunAsSystem = &runAsSystem_
	}
	if isZero(tmp.RunElevated) {
		runElevated_ := false
		tmp.RunElevated = &runElevated_
	}
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
}

type ImageTemplatePowerShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	RunAsSystem    *bool    `pulumi:"runAsSystem"`
	RunElevated    *bool    `pulumi:"runElevated"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
	ValidExitCodes []int    `pulumi:"validExitCodes"`
}


func (val *ImageTemplatePowerShellCustomizerResponse) Defaults() *ImageTemplatePowerShellCustomizerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.RunAsSystem) {
		runAsSystem_ := false
		tmp.RunAsSystem = &runAsSystem_
	}
	if isZero(tmp.RunElevated) {
		runElevated_ := false
		tmp.RunElevated = &runElevated_
	}
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
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
	ExcludeFromLatest  *bool             `pulumi:"excludeFromLatest"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	StorageAccountType *string           `pulumi:"storageAccountType"`
	Type               string            `pulumi:"type"`
}


func (val *ImageTemplateSharedImageDistributor) Defaults() *ImageTemplateSharedImageDistributor {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExcludeFromLatest) {
		excludeFromLatest_ := false
		tmp.ExcludeFromLatest = &excludeFromLatest_
	}
	return &tmp
}

type ImageTemplateSharedImageDistributorResponse struct {
	ArtifactTags       map[string]string `pulumi:"artifactTags"`
	ExcludeFromLatest  *bool             `pulumi:"excludeFromLatest"`
	GalleryImageId     string            `pulumi:"galleryImageId"`
	ReplicationRegions []string          `pulumi:"replicationRegions"`
	RunOutputName      string            `pulumi:"runOutputName"`
	StorageAccountType *string           `pulumi:"storageAccountType"`
	Type               string            `pulumi:"type"`
}


func (val *ImageTemplateSharedImageDistributorResponse) Defaults() *ImageTemplateSharedImageDistributorResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExcludeFromLatest) {
		excludeFromLatest_ := false
		tmp.ExcludeFromLatest = &excludeFromLatest_
	}
	return &tmp
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


func (val *ImageTemplateShellCustomizer) Defaults() *ImageTemplateShellCustomizer {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
}

type ImageTemplateShellCustomizerResponse struct {
	Inline         []string `pulumi:"inline"`
	Name           *string  `pulumi:"name"`
	ScriptUri      *string  `pulumi:"scriptUri"`
	Sha256Checksum *string  `pulumi:"sha256Checksum"`
	Type           string   `pulumi:"type"`
}


func (val *ImageTemplateShellCustomizerResponse) Defaults() *ImageTemplateShellCustomizerResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Sha256Checksum) {
		sha256Checksum_ := ""
		tmp.Sha256Checksum = &sha256Checksum_
	}
	return &tmp
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
	OsDiskSizeGB           *int                  `pulumi:"osDiskSizeGB"`
	UserAssignedIdentities []string              `pulumi:"userAssignedIdentities"`
	VmSize                 *string               `pulumi:"vmSize"`
	VnetConfig             *VirtualNetworkConfig `pulumi:"vnetConfig"`
}


func (val *ImageTemplateVmProfile) Defaults() *ImageTemplateVmProfile {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OsDiskSizeGB) {
		osDiskSizeGB_ := 0
		tmp.OsDiskSizeGB = &osDiskSizeGB_
	}
	if isZero(tmp.VmSize) {
		vmSize_ := ""
		tmp.VmSize = &vmSize_
	}
	tmp.VnetConfig = tmp.VnetConfig.Defaults()

	return &tmp
}





type ImageTemplateVmProfileInput interface {
	pulumi.Input

	ToImageTemplateVmProfileOutput() ImageTemplateVmProfileOutput
	ToImageTemplateVmProfileOutputWithContext(context.Context) ImageTemplateVmProfileOutput
}

type ImageTemplateVmProfileArgs struct {
	OsDiskSizeGB           pulumi.IntPtrInput           `pulumi:"osDiskSizeGB"`
	UserAssignedIdentities pulumi.StringArrayInput      `pulumi:"userAssignedIdentities"`
	VmSize                 pulumi.StringPtrInput        `pulumi:"vmSize"`
	VnetConfig             VirtualNetworkConfigPtrInput `pulumi:"vnetConfig"`
}


func (val *ImageTemplateVmProfileArgs) Defaults() *ImageTemplateVmProfileArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OsDiskSizeGB) {
		tmp.OsDiskSizeGB = pulumi.IntPtr(0)
	}
	if isZero(tmp.VmSize) {
		tmp.VmSize = pulumi.StringPtr("")
	}

	return &tmp
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

func (o ImageTemplateVmProfileOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateVmProfileOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileOutput) VnetConfig() VirtualNetworkConfigPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfile) *VirtualNetworkConfig { return v.VnetConfig }).(VirtualNetworkConfigPtrOutput)
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

func (o ImageTemplateVmProfilePtrOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *int {
		if v == nil {
			return nil
		}
		return v.OsDiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfilePtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

func (o ImageTemplateVmProfilePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfilePtrOutput) VnetConfig() VirtualNetworkConfigPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfile) *VirtualNetworkConfig {
		if v == nil {
			return nil
		}
		return v.VnetConfig
	}).(VirtualNetworkConfigPtrOutput)
}

type ImageTemplateVmProfileResponse struct {
	OsDiskSizeGB           *int                          `pulumi:"osDiskSizeGB"`
	UserAssignedIdentities []string                      `pulumi:"userAssignedIdentities"`
	VmSize                 *string                       `pulumi:"vmSize"`
	VnetConfig             *VirtualNetworkConfigResponse `pulumi:"vnetConfig"`
}


func (val *ImageTemplateVmProfileResponse) Defaults() *ImageTemplateVmProfileResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.OsDiskSizeGB) {
		osDiskSizeGB_ := 0
		tmp.OsDiskSizeGB = &osDiskSizeGB_
	}
	if isZero(tmp.VmSize) {
		vmSize_ := ""
		tmp.VmSize = &vmSize_
	}
	tmp.VnetConfig = tmp.VnetConfig.Defaults()

	return &tmp
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

func (o ImageTemplateVmProfileResponseOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *int { return v.OsDiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileResponseOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) []string { return v.UserAssignedIdentities }).(pulumi.StringArrayOutput)
}

func (o ImageTemplateVmProfileResponseOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *string { return v.VmSize }).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileResponseOutput) VnetConfig() VirtualNetworkConfigResponsePtrOutput {
	return o.ApplyT(func(v ImageTemplateVmProfileResponse) *VirtualNetworkConfigResponse { return v.VnetConfig }).(VirtualNetworkConfigResponsePtrOutput)
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

func (o ImageTemplateVmProfileResponsePtrOutput) OsDiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.OsDiskSizeGB
	}).(pulumi.IntPtrOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) UserAssignedIdentities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) []string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.StringArrayOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) VmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmSize
	}).(pulumi.StringPtrOutput)
}

func (o ImageTemplateVmProfileResponsePtrOutput) VnetConfig() VirtualNetworkConfigResponsePtrOutput {
	return o.ApplyT(func(v *ImageTemplateVmProfileResponse) *VirtualNetworkConfigResponse {
		if v == nil {
			return nil
		}
		return v.VnetConfig
	}).(VirtualNetworkConfigResponsePtrOutput)
}

type ImageTemplateWindowsUpdateCustomizer struct {
	Filters        []string `pulumi:"filters"`
	Name           *string  `pulumi:"name"`
	SearchCriteria *string  `pulumi:"searchCriteria"`
	Type           string   `pulumi:"type"`
	UpdateLimit    *int     `pulumi:"updateLimit"`
}

type ImageTemplateWindowsUpdateCustomizerResponse struct {
	Filters        []string `pulumi:"filters"`
	Name           *string  `pulumi:"name"`
	SearchCriteria *string  `pulumi:"searchCriteria"`
	Type           string   `pulumi:"type"`
	UpdateLimit    *int     `pulumi:"updateLimit"`
}

type PlatformImagePurchasePlan struct {
	PlanName      string `pulumi:"planName"`
	PlanProduct   string `pulumi:"planProduct"`
	PlanPublisher string `pulumi:"planPublisher"`
}

type PlatformImagePurchasePlanResponse struct {
	PlanName      string `pulumi:"planName"`
	PlanProduct   string `pulumi:"planProduct"`
	PlanPublisher string `pulumi:"planPublisher"`
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfig struct {
	ProxyVmSize *string `pulumi:"proxyVmSize"`
	SubnetId    *string `pulumi:"subnetId"`
}


func (val *VirtualNetworkConfig) Defaults() *VirtualNetworkConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ProxyVmSize) {
		proxyVmSize_ := ""
		tmp.ProxyVmSize = &proxyVmSize_
	}
	return &tmp
}





type VirtualNetworkConfigInput interface {
	pulumi.Input

	ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput
	ToVirtualNetworkConfigOutputWithContext(context.Context) VirtualNetworkConfigOutput
}

type VirtualNetworkConfigArgs struct {
	ProxyVmSize pulumi.StringPtrInput `pulumi:"proxyVmSize"`
	SubnetId    pulumi.StringPtrInput `pulumi:"subnetId"`
}


func (val *VirtualNetworkConfigArgs) Defaults() *VirtualNetworkConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ProxyVmSize) {
		tmp.ProxyVmSize = pulumi.StringPtr("")
	}
	return &tmp
}
func (VirtualNetworkConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfig)(nil)).Elem()
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput {
	return i.ToVirtualNetworkConfigOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigOutputWithContext(ctx context.Context) VirtualNetworkConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigOutput)
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return i.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (i VirtualNetworkConfigArgs) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigOutput).ToVirtualNetworkConfigPtrOutputWithContext(ctx)
}









type VirtualNetworkConfigPtrInput interface {
	pulumi.Input

	ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput
	ToVirtualNetworkConfigPtrOutputWithContext(context.Context) VirtualNetworkConfigPtrOutput
}

type virtualNetworkConfigPtrType VirtualNetworkConfigArgs

func VirtualNetworkConfigPtr(v *VirtualNetworkConfigArgs) VirtualNetworkConfigPtrInput {
	return (*virtualNetworkConfigPtrType)(v)
}

func (*virtualNetworkConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfig)(nil)).Elem()
}

func (i *virtualNetworkConfigPtrType) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return i.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (i *virtualNetworkConfigPtrType) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkConfigPtrOutput)
}

type VirtualNetworkConfigOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfig)(nil)).Elem()
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigOutput() VirtualNetworkConfigOutput {
	return o
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigOutputWithContext(ctx context.Context) VirtualNetworkConfigOutput {
	return o
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return o.ToVirtualNetworkConfigPtrOutputWithContext(context.Background())
}

func (o VirtualNetworkConfigOutput) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkConfig) *VirtualNetworkConfig {
		return &v
	}).(VirtualNetworkConfigPtrOutput)
}

func (o VirtualNetworkConfigOutput) ProxyVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfig) *string { return v.ProxyVmSize }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfig) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigPtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfig)(nil)).Elem()
}

func (o VirtualNetworkConfigPtrOutput) ToVirtualNetworkConfigPtrOutput() VirtualNetworkConfigPtrOutput {
	return o
}

func (o VirtualNetworkConfigPtrOutput) ToVirtualNetworkConfigPtrOutputWithContext(ctx context.Context) VirtualNetworkConfigPtrOutput {
	return o
}

func (o VirtualNetworkConfigPtrOutput) Elem() VirtualNetworkConfigOutput {
	return o.ApplyT(func(v *VirtualNetworkConfig) VirtualNetworkConfig {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfig
		return ret
	}).(VirtualNetworkConfigOutput)
}

func (o VirtualNetworkConfigPtrOutput) ProxyVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfig) *string {
		if v == nil {
			return nil
		}
		return v.ProxyVmSize
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigPtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfig) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigResponse struct {
	ProxyVmSize *string `pulumi:"proxyVmSize"`
	SubnetId    *string `pulumi:"subnetId"`
}


func (val *VirtualNetworkConfigResponse) Defaults() *VirtualNetworkConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ProxyVmSize) {
		proxyVmSize_ := ""
		tmp.ProxyVmSize = &proxyVmSize_
	}
	return &tmp
}

type VirtualNetworkConfigResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkConfigResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponseOutput() VirtualNetworkConfigResponseOutput {
	return o
}

func (o VirtualNetworkConfigResponseOutput) ToVirtualNetworkConfigResponseOutputWithContext(ctx context.Context) VirtualNetworkConfigResponseOutput {
	return o
}

func (o VirtualNetworkConfigResponseOutput) ProxyVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfigResponse) *string { return v.ProxyVmSize }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigResponseOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkConfigResponse) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type VirtualNetworkConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkConfigResponse)(nil)).Elem()
}

func (o VirtualNetworkConfigResponsePtrOutput) ToVirtualNetworkConfigResponsePtrOutput() VirtualNetworkConfigResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigResponsePtrOutput) ToVirtualNetworkConfigResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkConfigResponsePtrOutput {
	return o
}

func (o VirtualNetworkConfigResponsePtrOutput) Elem() VirtualNetworkConfigResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigResponse) VirtualNetworkConfigResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkConfigResponse
		return ret
	}).(VirtualNetworkConfigResponseOutput)
}

func (o VirtualNetworkConfigResponsePtrOutput) ProxyVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProxyVmSize
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkConfigResponsePtrOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImageTemplateIdentityOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ImageTemplateIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(ImageTemplateLastRunStatusResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfilePtrOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponseOutput{})
	pulumi.RegisterOutputType(ImageTemplateVmProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(ProvisioningErrorResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigPtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkConfigResponsePtrOutput{})
}
