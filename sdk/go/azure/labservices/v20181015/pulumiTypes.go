


package v20181015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentDetailsResponse struct {
	Description           string                        `pulumi:"description"`
	EnvironmentState      string                        `pulumi:"environmentState"`
	Id                    string                        `pulumi:"id"`
	LatestOperationResult LatestOperationResultResponse `pulumi:"latestOperationResult"`
	Name                  string                        `pulumi:"name"`
	PasswordLastReset     string                        `pulumi:"passwordLastReset"`
	ProvisioningState     string                        `pulumi:"provisioningState"`
	TotalUsage            string                        `pulumi:"totalUsage"`
	VirtualMachineDetails VirtualMachineDetailsResponse `pulumi:"virtualMachineDetails"`
}

type EnvironmentSizeResponse struct {
	MaxPrice         float64            `pulumi:"maxPrice"`
	MinMemory        float64            `pulumi:"minMemory"`
	MinNumberOfCores int                `pulumi:"minNumberOfCores"`
	Name             *string            `pulumi:"name"`
	VmSizes          []SizeInfoResponse `pulumi:"vmSizes"`
}

type EnvironmentSizeResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentSizeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentSizeResponse)(nil)).Elem()
}

func (o EnvironmentSizeResponseOutput) ToEnvironmentSizeResponseOutput() EnvironmentSizeResponseOutput {
	return o
}

func (o EnvironmentSizeResponseOutput) ToEnvironmentSizeResponseOutputWithContext(ctx context.Context) EnvironmentSizeResponseOutput {
	return o
}

func (o EnvironmentSizeResponseOutput) MaxPrice() pulumi.Float64Output {
	return o.ApplyT(func(v EnvironmentSizeResponse) float64 { return v.MaxPrice }).(pulumi.Float64Output)
}

func (o EnvironmentSizeResponseOutput) MinMemory() pulumi.Float64Output {
	return o.ApplyT(func(v EnvironmentSizeResponse) float64 { return v.MinMemory }).(pulumi.Float64Output)
}

func (o EnvironmentSizeResponseOutput) MinNumberOfCores() pulumi.IntOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) int { return v.MinNumberOfCores }).(pulumi.IntOutput)
}

func (o EnvironmentSizeResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentSizeResponseOutput) VmSizes() SizeInfoResponseArrayOutput {
	return o.ApplyT(func(v EnvironmentSizeResponse) []SizeInfoResponse { return v.VmSizes }).(SizeInfoResponseArrayOutput)
}

type EnvironmentSizeResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentSizeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentSizeResponse)(nil)).Elem()
}

func (o EnvironmentSizeResponseArrayOutput) ToEnvironmentSizeResponseArrayOutput() EnvironmentSizeResponseArrayOutput {
	return o
}

func (o EnvironmentSizeResponseArrayOutput) ToEnvironmentSizeResponseArrayOutputWithContext(ctx context.Context) EnvironmentSizeResponseArrayOutput {
	return o
}

func (o EnvironmentSizeResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentSizeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentSizeResponse {
		return vs[0].([]EnvironmentSizeResponse)[vs[1].(int)]
	}).(EnvironmentSizeResponseOutput)
}

type GalleryImageReferenceResponse struct {
	Offer     *string `pulumi:"offer"`
	OsType    *string `pulumi:"osType"`
	Publisher *string `pulumi:"publisher"`
	Sku       *string `pulumi:"sku"`
	Version   *string `pulumi:"version"`
}

type GalleryImageReferenceResponseOutput struct{ *pulumi.OutputState }

func (GalleryImageReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryImageReferenceResponse)(nil)).Elem()
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutput() GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) ToGalleryImageReferenceResponseOutputWithContext(ctx context.Context) GalleryImageReferenceResponseOutput {
	return o
}

func (o GalleryImageReferenceResponseOutput) Offer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Offer }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o GalleryImageReferenceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GalleryImageReferenceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type LabDetailsResponse struct {
	Id                *string `pulumi:"id"`
	Name              *string `pulumi:"name"`
	ProvisioningState *string `pulumi:"provisioningState"`
	UsageQuota        string  `pulumi:"usageQuota"`
}

type LatestOperationResultResponse struct {
	ErrorCode    string `pulumi:"errorCode"`
	ErrorMessage string `pulumi:"errorMessage"`
	HttpMethod   string `pulumi:"httpMethod"`
	OperationUrl string `pulumi:"operationUrl"`
	RequestUri   string `pulumi:"requestUri"`
	Status       string `pulumi:"status"`
}

type LatestOperationResultResponseOutput struct{ *pulumi.OutputState }

func (LatestOperationResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LatestOperationResultResponse)(nil)).Elem()
}

func (o LatestOperationResultResponseOutput) ToLatestOperationResultResponseOutput() LatestOperationResultResponseOutput {
	return o
}

func (o LatestOperationResultResponseOutput) ToLatestOperationResultResponseOutputWithContext(ctx context.Context) LatestOperationResultResponseOutput {
	return o
}

func (o LatestOperationResultResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.HttpMethod }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) OperationUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.OperationUrl }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) RequestUri() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.RequestUri }).(pulumi.StringOutput)
}

func (o LatestOperationResultResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LatestOperationResultResponse) string { return v.Status }).(pulumi.StringOutput)
}

type NetworkInterfaceResponse struct {
	PrivateIpAddress string `pulumi:"privateIpAddress"`
	RdpAuthority     string `pulumi:"rdpAuthority"`
	SshAuthority     string `pulumi:"sshAuthority"`
	Username         string `pulumi:"username"`
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

func (o NetworkInterfaceResponseOutput) PrivateIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.PrivateIpAddress }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) RdpAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.RdpAuthority }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) SshAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.SshAuthority }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Username }).(pulumi.StringOutput)
}

type OperationBatchStatusResponseItemResponse struct {
	OperationUrl string `pulumi:"operationUrl"`
	Status       string `pulumi:"status"`
}

type ReferenceVm struct {
	Password *string `pulumi:"password"`
	UserName string  `pulumi:"userName"`
}





type ReferenceVmInput interface {
	pulumi.Input

	ToReferenceVmOutput() ReferenceVmOutput
	ToReferenceVmOutputWithContext(context.Context) ReferenceVmOutput
}

type ReferenceVmArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	UserName pulumi.StringInput    `pulumi:"userName"`
}

func (ReferenceVmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVm)(nil)).Elem()
}

func (i ReferenceVmArgs) ToReferenceVmOutput() ReferenceVmOutput {
	return i.ToReferenceVmOutputWithContext(context.Background())
}

func (i ReferenceVmArgs) ToReferenceVmOutputWithContext(ctx context.Context) ReferenceVmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceVmOutput)
}

type ReferenceVmOutput struct{ *pulumi.OutputState }

func (ReferenceVmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVm)(nil)).Elem()
}

func (o ReferenceVmOutput) ToReferenceVmOutput() ReferenceVmOutput {
	return o
}

func (o ReferenceVmOutput) ToReferenceVmOutputWithContext(ctx context.Context) ReferenceVmOutput {
	return o
}

func (o ReferenceVmOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceVm) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ReferenceVmOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVm) string { return v.UserName }).(pulumi.StringOutput)
}

type ReferenceVmResponse struct {
	Password       *string                `pulumi:"password"`
	UserName       string                 `pulumi:"userName"`
	VmResourceId   string                 `pulumi:"vmResourceId"`
	VmStateDetails VmStateDetailsResponse `pulumi:"vmStateDetails"`
}

type ReferenceVmResponseOutput struct{ *pulumi.OutputState }

func (ReferenceVmResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceVmResponse)(nil)).Elem()
}

func (o ReferenceVmResponseOutput) ToReferenceVmResponseOutput() ReferenceVmResponseOutput {
	return o
}

func (o ReferenceVmResponseOutput) ToReferenceVmResponseOutputWithContext(ctx context.Context) ReferenceVmResponseOutput {
	return o
}

func (o ReferenceVmResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceVmResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ReferenceVmResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVmResponse) string { return v.UserName }).(pulumi.StringOutput)
}

func (o ReferenceVmResponseOutput) VmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReferenceVmResponse) string { return v.VmResourceId }).(pulumi.StringOutput)
}

func (o ReferenceVmResponseOutput) VmStateDetails() VmStateDetailsResponseOutput {
	return o.ApplyT(func(v ReferenceVmResponse) VmStateDetailsResponse { return v.VmStateDetails }).(VmStateDetailsResponseOutput)
}

type RegionalAvailabilityResponse struct {
	Region             *string                    `pulumi:"region"`
	SizeAvailabilities []SizeAvailabilityResponse `pulumi:"sizeAvailabilities"`
}

type ResourceSet struct {
	ResourceSettingId *string `pulumi:"resourceSettingId"`
	VmResourceId      *string `pulumi:"vmResourceId"`
}





type ResourceSetInput interface {
	pulumi.Input

	ToResourceSetOutput() ResourceSetOutput
	ToResourceSetOutputWithContext(context.Context) ResourceSetOutput
}

type ResourceSetArgs struct {
	ResourceSettingId pulumi.StringPtrInput `pulumi:"resourceSettingId"`
	VmResourceId      pulumi.StringPtrInput `pulumi:"vmResourceId"`
}

func (ResourceSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSet)(nil)).Elem()
}

func (i ResourceSetArgs) ToResourceSetOutput() ResourceSetOutput {
	return i.ToResourceSetOutputWithContext(context.Background())
}

func (i ResourceSetArgs) ToResourceSetOutputWithContext(ctx context.Context) ResourceSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetOutput)
}

func (i ResourceSetArgs) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return i.ToResourceSetPtrOutputWithContext(context.Background())
}

func (i ResourceSetArgs) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetOutput).ToResourceSetPtrOutputWithContext(ctx)
}









type ResourceSetPtrInput interface {
	pulumi.Input

	ToResourceSetPtrOutput() ResourceSetPtrOutput
	ToResourceSetPtrOutputWithContext(context.Context) ResourceSetPtrOutput
}

type resourceSetPtrType ResourceSetArgs

func ResourceSetPtr(v *ResourceSetArgs) ResourceSetPtrInput {
	return (*resourceSetPtrType)(v)
}

func (*resourceSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSet)(nil)).Elem()
}

func (i *resourceSetPtrType) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return i.ToResourceSetPtrOutputWithContext(context.Background())
}

func (i *resourceSetPtrType) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSetPtrOutput)
}

type ResourceSetOutput struct{ *pulumi.OutputState }

func (ResourceSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSet)(nil)).Elem()
}

func (o ResourceSetOutput) ToResourceSetOutput() ResourceSetOutput {
	return o
}

func (o ResourceSetOutput) ToResourceSetOutputWithContext(ctx context.Context) ResourceSetOutput {
	return o
}

func (o ResourceSetOutput) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return o.ToResourceSetPtrOutputWithContext(context.Background())
}

func (o ResourceSetOutput) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSet) *ResourceSet {
		return &v
	}).(ResourceSetPtrOutput)
}

func (o ResourceSetOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSet) *string { return v.ResourceSettingId }).(pulumi.StringPtrOutput)
}

func (o ResourceSetOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSet) *string { return v.VmResourceId }).(pulumi.StringPtrOutput)
}

type ResourceSetPtrOutput struct{ *pulumi.OutputState }

func (ResourceSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSet)(nil)).Elem()
}

func (o ResourceSetPtrOutput) ToResourceSetPtrOutput() ResourceSetPtrOutput {
	return o
}

func (o ResourceSetPtrOutput) ToResourceSetPtrOutputWithContext(ctx context.Context) ResourceSetPtrOutput {
	return o
}

func (o ResourceSetPtrOutput) Elem() ResourceSetOutput {
	return o.ApplyT(func(v *ResourceSet) ResourceSet {
		if v != nil {
			return *v
		}
		var ret ResourceSet
		return ret
	}).(ResourceSetOutput)
}

func (o ResourceSetPtrOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSet) *string {
		if v == nil {
			return nil
		}
		return v.ResourceSettingId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSetPtrOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSet) *string {
		if v == nil {
			return nil
		}
		return v.VmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceSetResponse struct {
	ResourceSettingId *string `pulumi:"resourceSettingId"`
	VmResourceId      *string `pulumi:"vmResourceId"`
}

type ResourceSetResponseOutput struct{ *pulumi.OutputState }

func (ResourceSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSetResponse)(nil)).Elem()
}

func (o ResourceSetResponseOutput) ToResourceSetResponseOutput() ResourceSetResponseOutput {
	return o
}

func (o ResourceSetResponseOutput) ToResourceSetResponseOutputWithContext(ctx context.Context) ResourceSetResponseOutput {
	return o
}

func (o ResourceSetResponseOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSetResponse) *string { return v.ResourceSettingId }).(pulumi.StringPtrOutput)
}

func (o ResourceSetResponseOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSetResponse) *string { return v.VmResourceId }).(pulumi.StringPtrOutput)
}

type ResourceSetResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSetResponse)(nil)).Elem()
}

func (o ResourceSetResponsePtrOutput) ToResourceSetResponsePtrOutput() ResourceSetResponsePtrOutput {
	return o
}

func (o ResourceSetResponsePtrOutput) ToResourceSetResponsePtrOutputWithContext(ctx context.Context) ResourceSetResponsePtrOutput {
	return o
}

func (o ResourceSetResponsePtrOutput) Elem() ResourceSetResponseOutput {
	return o.ApplyT(func(v *ResourceSetResponse) ResourceSetResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSetResponse
		return ret
	}).(ResourceSetResponseOutput)
}

func (o ResourceSetResponsePtrOutput) ResourceSettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceSettingId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSetResponsePtrOutput) VmResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.VmResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceSettings struct {
	GalleryImageResourceId *string     `pulumi:"galleryImageResourceId"`
	ReferenceVm            ReferenceVm `pulumi:"referenceVm"`
	Size                   *string     `pulumi:"size"`
}





type ResourceSettingsInput interface {
	pulumi.Input

	ToResourceSettingsOutput() ResourceSettingsOutput
	ToResourceSettingsOutputWithContext(context.Context) ResourceSettingsOutput
}

type ResourceSettingsArgs struct {
	GalleryImageResourceId pulumi.StringPtrInput `pulumi:"galleryImageResourceId"`
	ReferenceVm            ReferenceVmInput      `pulumi:"referenceVm"`
	Size                   pulumi.StringPtrInput `pulumi:"size"`
}

func (ResourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettings)(nil)).Elem()
}

func (i ResourceSettingsArgs) ToResourceSettingsOutput() ResourceSettingsOutput {
	return i.ToResourceSettingsOutputWithContext(context.Background())
}

func (i ResourceSettingsArgs) ToResourceSettingsOutputWithContext(ctx context.Context) ResourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSettingsOutput)
}

type ResourceSettingsOutput struct{ *pulumi.OutputState }

func (ResourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettings)(nil)).Elem()
}

func (o ResourceSettingsOutput) ToResourceSettingsOutput() ResourceSettingsOutput {
	return o
}

func (o ResourceSettingsOutput) ToResourceSettingsOutputWithContext(ctx context.Context) ResourceSettingsOutput {
	return o
}

func (o ResourceSettingsOutput) GalleryImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettings) *string { return v.GalleryImageResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceSettingsOutput) ReferenceVm() ReferenceVmOutput {
	return o.ApplyT(func(v ResourceSettings) ReferenceVm { return v.ReferenceVm }).(ReferenceVmOutput)
}

func (o ResourceSettingsOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettings) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type ResourceSettingsResponse struct {
	Cores                  int                 `pulumi:"cores"`
	GalleryImageResourceId *string             `pulumi:"galleryImageResourceId"`
	Id                     string              `pulumi:"id"`
	ImageName              string              `pulumi:"imageName"`
	ReferenceVm            ReferenceVmResponse `pulumi:"referenceVm"`
	Size                   *string             `pulumi:"size"`
}

type ResourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (ResourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSettingsResponse)(nil)).Elem()
}

func (o ResourceSettingsResponseOutput) ToResourceSettingsResponseOutput() ResourceSettingsResponseOutput {
	return o
}

func (o ResourceSettingsResponseOutput) ToResourceSettingsResponseOutputWithContext(ctx context.Context) ResourceSettingsResponseOutput {
	return o
}

func (o ResourceSettingsResponseOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) int { return v.Cores }).(pulumi.IntOutput)
}

func (o ResourceSettingsResponseOutput) GalleryImageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) *string { return v.GalleryImageResourceId }).(pulumi.StringPtrOutput)
}

func (o ResourceSettingsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ResourceSettingsResponseOutput) ImageName() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) string { return v.ImageName }).(pulumi.StringOutput)
}

func (o ResourceSettingsResponseOutput) ReferenceVm() ReferenceVmResponseOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) ReferenceVmResponse { return v.ReferenceVm }).(ReferenceVmResponseOutput)
}

func (o ResourceSettingsResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSettingsResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

type SizeAvailabilityResponse struct {
	IsAvailable  *bool   `pulumi:"isAvailable"`
	SizeCategory *string `pulumi:"sizeCategory"`
}

type SizeConfigurationPropertiesResponse struct {
	EnvironmentSizes []EnvironmentSizeResponse `pulumi:"environmentSizes"`
}

type SizeConfigurationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SizeConfigurationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeConfigurationPropertiesResponse)(nil)).Elem()
}

func (o SizeConfigurationPropertiesResponseOutput) ToSizeConfigurationPropertiesResponseOutput() SizeConfigurationPropertiesResponseOutput {
	return o
}

func (o SizeConfigurationPropertiesResponseOutput) ToSizeConfigurationPropertiesResponseOutputWithContext(ctx context.Context) SizeConfigurationPropertiesResponseOutput {
	return o
}

func (o SizeConfigurationPropertiesResponseOutput) EnvironmentSizes() EnvironmentSizeResponseArrayOutput {
	return o.ApplyT(func(v SizeConfigurationPropertiesResponse) []EnvironmentSizeResponse { return v.EnvironmentSizes }).(EnvironmentSizeResponseArrayOutput)
}

type SizeInfoResponse struct {
	ComputeSize   *string  `pulumi:"computeSize"`
	Memory        *float64 `pulumi:"memory"`
	NumberOfCores *int     `pulumi:"numberOfCores"`
	Price         *float64 `pulumi:"price"`
}

type SizeInfoResponseOutput struct{ *pulumi.OutputState }

func (SizeInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SizeInfoResponse)(nil)).Elem()
}

func (o SizeInfoResponseOutput) ToSizeInfoResponseOutput() SizeInfoResponseOutput {
	return o
}

func (o SizeInfoResponseOutput) ToSizeInfoResponseOutputWithContext(ctx context.Context) SizeInfoResponseOutput {
	return o
}

func (o SizeInfoResponseOutput) ComputeSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *string { return v.ComputeSize }).(pulumi.StringPtrOutput)
}

func (o SizeInfoResponseOutput) Memory() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *float64 { return v.Memory }).(pulumi.Float64PtrOutput)
}

func (o SizeInfoResponseOutput) NumberOfCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *int { return v.NumberOfCores }).(pulumi.IntPtrOutput)
}

func (o SizeInfoResponseOutput) Price() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SizeInfoResponse) *float64 { return v.Price }).(pulumi.Float64PtrOutput)
}

type SizeInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (SizeInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SizeInfoResponse)(nil)).Elem()
}

func (o SizeInfoResponseArrayOutput) ToSizeInfoResponseArrayOutput() SizeInfoResponseArrayOutput {
	return o
}

func (o SizeInfoResponseArrayOutput) ToSizeInfoResponseArrayOutputWithContext(ctx context.Context) SizeInfoResponseArrayOutput {
	return o
}

func (o SizeInfoResponseArrayOutput) Index(i pulumi.IntInput) SizeInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SizeInfoResponse {
		return vs[0].([]SizeInfoResponse)[vs[1].(int)]
	}).(SizeInfoResponseOutput)
}

type VirtualMachineDetailsResponse struct {
	LastKnownPowerState string `pulumi:"lastKnownPowerState"`
	PrivateIpAddress    string `pulumi:"privateIpAddress"`
	ProvisioningState   string `pulumi:"provisioningState"`
	RdpAuthority        string `pulumi:"rdpAuthority"`
	SshAuthority        string `pulumi:"sshAuthority"`
	UserName            string `pulumi:"userName"`
}

type VmStateDetailsResponse struct {
	LastKnownPowerState string `pulumi:"lastKnownPowerState"`
	PowerState          string `pulumi:"powerState"`
	RdpAuthority        string `pulumi:"rdpAuthority"`
	SshAuthority        string `pulumi:"sshAuthority"`
}

type VmStateDetailsResponseOutput struct{ *pulumi.OutputState }

func (VmStateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmStateDetailsResponse)(nil)).Elem()
}

func (o VmStateDetailsResponseOutput) ToVmStateDetailsResponseOutput() VmStateDetailsResponseOutput {
	return o
}

func (o VmStateDetailsResponseOutput) ToVmStateDetailsResponseOutputWithContext(ctx context.Context) VmStateDetailsResponseOutput {
	return o
}

func (o VmStateDetailsResponseOutput) LastKnownPowerState() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.LastKnownPowerState }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.PowerState }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) RdpAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.RdpAuthority }).(pulumi.StringOutput)
}

func (o VmStateDetailsResponseOutput) SshAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v VmStateDetailsResponse) string { return v.SshAuthority }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnvironmentSizeResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentSizeResponseArrayOutput{})
	pulumi.RegisterOutputType(GalleryImageReferenceResponseOutput{})
	pulumi.RegisterOutputType(LatestOperationResultResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(ReferenceVmOutput{})
	pulumi.RegisterOutputType(ReferenceVmResponseOutput{})
	pulumi.RegisterOutputType(ResourceSetOutput{})
	pulumi.RegisterOutputType(ResourceSetPtrOutput{})
	pulumi.RegisterOutputType(ResourceSetResponseOutput{})
	pulumi.RegisterOutputType(ResourceSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceSettingsOutput{})
	pulumi.RegisterOutputType(ResourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SizeConfigurationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SizeInfoResponseOutput{})
	pulumi.RegisterOutputType(SizeInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(VmStateDetailsResponseOutput{})
}
