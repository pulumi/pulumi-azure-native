


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSetVM(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetVMArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetVMResult, error) {
	var rv LookupVirtualMachineScaleSetVMResult
	err := ctx.Invoke("azure-native:compute/v20200601:getVirtualMachineScaleSetVM", args, &rv, opts...)
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
	VmId                        string                                                       `pulumi:"vmId"`
	Zones                       []string                                                     `pulumi:"zones"`
}
