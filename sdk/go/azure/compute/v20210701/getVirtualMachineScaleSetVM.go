


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVM(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMResult, error) {
	var rv LookupVirtualMachineScaleSetVMResult
	err := ctx.Invoke("azure-native:compute/v20210701:getVirtualMachineScaleSetVM", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetVMArgs struct {
	Expand            *string `pulumi:"expand"`
	InstanceId        string  `pulumi:"instanceId"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmScaleSetName    string  `pulumi:"vmScaleSetName"`
}


type LookupVirtualMachineScaleSetVMResult struct {
	AdditionalCapabilities      *AdditionalCapabilitiesResponse                              `pulumi:"additionalCapabilities"`
	AvailabilitySet             *SubResourceResponse                                         `pulumi:"availabilitySet"`
	DiagnosticsProfile          *DiagnosticsProfileResponse                                  `pulumi:"diagnosticsProfile"`
	HardwareProfile             *HardwareProfileResponse                                     `pulumi:"hardwareProfile"`
	Id                          string                                                       `pulumi:"id"`
	InstanceId                  string                                                       `pulumi:"instanceId"`
	InstanceView                VirtualMachineScaleSetVMInstanceViewResponse                 `pulumi:"instanceView"`
	LatestModelApplied          bool                                                         `pulumi:"latestModelApplied"`
	LicenseType                 *string                                                      `pulumi:"licenseType"`
	Location                    string                                                       `pulumi:"location"`
	ModelDefinitionApplied      string                                                       `pulumi:"modelDefinitionApplied"`
	Name                        string                                                       `pulumi:"name"`
	NetworkProfile              *NetworkProfileResponse                                      `pulumi:"networkProfile"`
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfigurationResponse `pulumi:"networkProfileConfiguration"`
	OsProfile                   *OSProfileResponse                                           `pulumi:"osProfile"`
	Plan                        *PlanResponse                                                `pulumi:"plan"`
	ProtectionPolicy            *VirtualMachineScaleSetVMProtectionPolicyResponse            `pulumi:"protectionPolicy"`
	ProvisioningState           string                                                       `pulumi:"provisioningState"`
	Resources                   []VirtualMachineExtensionResponse                            `pulumi:"resources"`
	SecurityProfile             *SecurityProfileResponse                                     `pulumi:"securityProfile"`
	Sku                         SkuResponse                                                  `pulumi:"sku"`
	StorageProfile              *StorageProfileResponse                                      `pulumi:"storageProfile"`
	Tags                        map[string]string                                            `pulumi:"tags"`
	Type                        string                                                       `pulumi:"type"`
	UserData                    *string                                                      `pulumi:"userData"`
	VmId                        string                                                       `pulumi:"vmId"`
	Zones                       []string                                                     `pulumi:"zones"`
}

func LookupVirtualMachineScaleSetVMOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetVMOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetVMResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetVMResult, error) {
			args := v.(LookupVirtualMachineScaleSetVMArgs)
			r, err := LookupVirtualMachineScaleSetVM(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetVMResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetVMResultOutput)
}

type LookupVirtualMachineScaleSetVMOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	InstanceId        pulumi.StringInput    `pulumi:"instanceId"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmScaleSetName    pulumi.StringInput    `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetVMOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetVMResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetVMResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetVMResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetVMResultOutput) ToLookupVirtualMachineScaleSetVMResultOutput() LookupVirtualMachineScaleSetVMResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMResultOutput) ToLookupVirtualMachineScaleSetVMResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetVMResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetVMResultOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *AdditionalCapabilitiesResponse {
		return v.AdditionalCapabilities
	}).(AdditionalCapabilitiesResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *SubResourceResponse { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *DiagnosticsProfileResponse { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *HardwareProfileResponse { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) InstanceView() VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) VirtualMachineScaleSetVMInstanceViewResponse {
		return v.InstanceView
	}).(VirtualMachineScaleSetVMInstanceViewResponseOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) LatestModelApplied() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) bool { return v.LatestModelApplied }).(pulumi.BoolOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) ModelDefinitionApplied() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.ModelDefinitionApplied }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *NetworkProfileResponse { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) NetworkProfileConfiguration() VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *VirtualMachineScaleSetVMNetworkProfileConfigurationResponse {
		return v.NetworkProfileConfiguration
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *OSProfileResponse { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) ProtectionPolicy() VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *VirtualMachineScaleSetVMProtectionPolicyResponse {
		return v.ProtectionPolicy
	}).(VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) []VirtualMachineExtensionResponse { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *SecurityProfileResponse { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *StorageProfileResponse { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) *string { return v.UserData }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) string { return v.VmId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetVMResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetVMResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetVMResultOutput{})
}
