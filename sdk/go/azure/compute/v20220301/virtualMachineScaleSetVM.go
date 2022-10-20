


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSetVM struct {
	pulumi.CustomResourceState

	AdditionalCapabilities      AdditionalCapabilitiesResponsePtrOutput                              `pulumi:"additionalCapabilities"`
	AvailabilitySet             SubResourceResponsePtrOutput                                         `pulumi:"availabilitySet"`
	DiagnosticsProfile          DiagnosticsProfileResponsePtrOutput                                  `pulumi:"diagnosticsProfile"`
	HardwareProfile             HardwareProfileResponsePtrOutput                                     `pulumi:"hardwareProfile"`
	Identity                    VirtualMachineIdentityResponsePtrOutput                              `pulumi:"identity"`
	InstanceId                  pulumi.StringOutput                                                  `pulumi:"instanceId"`
	InstanceView                VirtualMachineScaleSetVMInstanceViewResponseOutput                   `pulumi:"instanceView"`
	LatestModelApplied          pulumi.BoolOutput                                                    `pulumi:"latestModelApplied"`
	LicenseType                 pulumi.StringPtrOutput                                               `pulumi:"licenseType"`
	Location                    pulumi.StringOutput                                                  `pulumi:"location"`
	ModelDefinitionApplied      pulumi.StringOutput                                                  `pulumi:"modelDefinitionApplied"`
	Name                        pulumi.StringOutput                                                  `pulumi:"name"`
	NetworkProfile              NetworkProfileResponsePtrOutput                                      `pulumi:"networkProfile"`
	NetworkProfileConfiguration VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput `pulumi:"networkProfileConfiguration"`
	OsProfile                   OSProfileResponsePtrOutput                                           `pulumi:"osProfile"`
	Plan                        PlanResponsePtrOutput                                                `pulumi:"plan"`
	ProtectionPolicy            VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput            `pulumi:"protectionPolicy"`
	ProvisioningState           pulumi.StringOutput                                                  `pulumi:"provisioningState"`
	Resources                   VirtualMachineExtensionResponseArrayOutput                           `pulumi:"resources"`
	SecurityProfile             SecurityProfileResponsePtrOutput                                     `pulumi:"securityProfile"`
	Sku                         SkuResponseOutput                                                    `pulumi:"sku"`
	StorageProfile              StorageProfileResponsePtrOutput                                      `pulumi:"storageProfile"`
	Tags                        pulumi.StringMapOutput                                               `pulumi:"tags"`
	Type                        pulumi.StringOutput                                                  `pulumi:"type"`
	UserData                    pulumi.StringPtrOutput                                               `pulumi:"userData"`
	VmId                        pulumi.StringOutput                                                  `pulumi:"vmId"`
	Zones                       pulumi.StringArrayOutput                                             `pulumi:"zones"`
}


func NewVirtualMachineScaleSetVM(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVM, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVM"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetVM"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetVM
	err := ctx.RegisterResource("azure-native:compute/v20220301:VirtualMachineScaleSetVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSetVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVM, error) {
	var resource VirtualMachineScaleSetVM
	err := ctx.ReadResource("azure-native:compute/v20220301:VirtualMachineScaleSetVM", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScaleSetVMState struct {
}

type VirtualMachineScaleSetVMState struct {
}

func (VirtualMachineScaleSetVMState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMState)(nil)).Elem()
}

type virtualMachineScaleSetVMArgs struct {
	AdditionalCapabilities      *AdditionalCapabilities                              `pulumi:"additionalCapabilities"`
	AvailabilitySet             *SubResource                                         `pulumi:"availabilitySet"`
	DiagnosticsProfile          *DiagnosticsProfile                                  `pulumi:"diagnosticsProfile"`
	HardwareProfile             *HardwareProfile                                     `pulumi:"hardwareProfile"`
	Identity                    *VirtualMachineIdentity                              `pulumi:"identity"`
	InstanceId                  *string                                              `pulumi:"instanceId"`
	LicenseType                 *string                                              `pulumi:"licenseType"`
	Location                    *string                                              `pulumi:"location"`
	NetworkProfile              *NetworkProfile                                      `pulumi:"networkProfile"`
	NetworkProfileConfiguration *VirtualMachineScaleSetVMNetworkProfileConfiguration `pulumi:"networkProfileConfiguration"`
	OsProfile                   *OSProfile                                           `pulumi:"osProfile"`
	Plan                        *Plan                                                `pulumi:"plan"`
	ProtectionPolicy            *VirtualMachineScaleSetVMProtectionPolicy            `pulumi:"protectionPolicy"`
	ResourceGroupName           string                                               `pulumi:"resourceGroupName"`
	SecurityProfile             *SecurityProfile                                     `pulumi:"securityProfile"`
	StorageProfile              *StorageProfile                                      `pulumi:"storageProfile"`
	Tags                        map[string]string                                    `pulumi:"tags"`
	UserData                    *string                                              `pulumi:"userData"`
	VmScaleSetName              string                                               `pulumi:"vmScaleSetName"`
}


type VirtualMachineScaleSetVMArgs struct {
	AdditionalCapabilities      AdditionalCapabilitiesPtrInput
	AvailabilitySet             SubResourcePtrInput
	DiagnosticsProfile          DiagnosticsProfilePtrInput
	HardwareProfile             HardwareProfilePtrInput
	Identity                    VirtualMachineIdentityPtrInput
	InstanceId                  pulumi.StringPtrInput
	LicenseType                 pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	NetworkProfile              NetworkProfilePtrInput
	NetworkProfileConfiguration VirtualMachineScaleSetVMNetworkProfileConfigurationPtrInput
	OsProfile                   OSProfilePtrInput
	Plan                        PlanPtrInput
	ProtectionPolicy            VirtualMachineScaleSetVMProtectionPolicyPtrInput
	ResourceGroupName           pulumi.StringInput
	SecurityProfile             SecurityProfilePtrInput
	StorageProfile              StorageProfilePtrInput
	Tags                        pulumi.StringMapInput
	UserData                    pulumi.StringPtrInput
	VmScaleSetName              pulumi.StringInput
}

func (VirtualMachineScaleSetVMArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput
	ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput
}

func (*VirtualMachineScaleSetVM) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVM)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVM) ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput {
	return i.ToVirtualMachineScaleSetVMOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVM) ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMOutput)
}

type VirtualMachineScaleSetVMOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVM)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMOutput) ToVirtualMachineScaleSetVMOutput() VirtualMachineScaleSetVMOutput {
	return o
}

func (o VirtualMachineScaleSetVMOutput) ToVirtualMachineScaleSetVMOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMOutput {
	return o
}

func (o VirtualMachineScaleSetVMOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) AdditionalCapabilitiesResponsePtrOutput {
		return v.AdditionalCapabilities
	}).(AdditionalCapabilitiesResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SubResourceResponsePtrOutput { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) DiagnosticsProfileResponsePtrOutput { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) InstanceView() VirtualMachineScaleSetVMInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMInstanceViewResponseOutput {
		return v.InstanceView
	}).(VirtualMachineScaleSetVMInstanceViewResponseOutput)
}

func (o VirtualMachineScaleSetVMOutput) LatestModelApplied() pulumi.BoolOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.BoolOutput { return v.LatestModelApplied }).(pulumi.BoolOutput)
}

func (o VirtualMachineScaleSetVMOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) ModelDefinitionApplied() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.ModelDefinitionApplied }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) NetworkProfileConfiguration() VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput {
		return v.NetworkProfileConfiguration
	}).(VirtualMachineScaleSetVMNetworkProfileConfigurationResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) ProtectionPolicy() VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput {
		return v.ProtectionPolicy
	}).(VirtualMachineScaleSetVMProtectionPolicyResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) VirtualMachineExtensionResponseArrayOutput { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMOutput) SecurityProfile() SecurityProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SecurityProfileResponsePtrOutput { return v.SecurityProfile }).(SecurityProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o VirtualMachineScaleSetVMOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineScaleSetVMOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVM) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMOutput{})
}
