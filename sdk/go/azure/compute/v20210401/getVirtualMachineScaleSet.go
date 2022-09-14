


package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualMachineScaleSet(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetResult, error) {
	var rv LookupVirtualMachineScaleSetResult
	err := ctx.Invoke("azure-native:compute/v20210401:getVirtualMachineScaleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetArgs struct {
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	VmScaleSetName    string  `pulumi:"vmScaleSetName"`
}


type LookupVirtualMachineScaleSetResult struct {
	AdditionalCapabilities                 *AdditionalCapabilitiesResponse          `pulumi:"additionalCapabilities"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicyResponse          `pulumi:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                                    `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	ExtendedLocation                       *ExtendedLocationResponse                `pulumi:"extendedLocation"`
	HostGroup                              *SubResourceResponse                     `pulumi:"hostGroup"`
	Id                                     string                                   `pulumi:"id"`
	Identity                               *VirtualMachineScaleSetIdentityResponse  `pulumi:"identity"`
	Location                               string                                   `pulumi:"location"`
	Name                                   string                                   `pulumi:"name"`
	OrchestrationMode                      *string                                  `pulumi:"orchestrationMode"`
	Overprovision                          *bool                                    `pulumi:"overprovision"`
	Plan                                   *PlanResponse                            `pulumi:"plan"`
	PlatformFaultDomainCount               *int                                     `pulumi:"platformFaultDomainCount"`
	ProvisioningState                      string                                   `pulumi:"provisioningState"`
	ProximityPlacementGroup                *SubResourceResponse                     `pulumi:"proximityPlacementGroup"`
	ScaleInPolicy                          *ScaleInPolicyResponse                   `pulumi:"scaleInPolicy"`
	SinglePlacementGroup                   *bool                                    `pulumi:"singlePlacementGroup"`
	Sku                                    *SkuResponse                             `pulumi:"sku"`
	SpotRestorePolicy                      *SpotRestorePolicyResponse               `pulumi:"spotRestorePolicy"`
	Tags                                   map[string]string                        `pulumi:"tags"`
	Type                                   string                                   `pulumi:"type"`
	UniqueId                               string                                   `pulumi:"uniqueId"`
	UpgradePolicy                          *UpgradePolicyResponse                   `pulumi:"upgradePolicy"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfileResponse `pulumi:"virtualMachineProfile"`
	ZoneBalance                            *bool                                    `pulumi:"zoneBalance"`
	Zones                                  []string                                 `pulumi:"zones"`
}

func LookupVirtualMachineScaleSetOutput(ctx *pulumi.Context, args LookupVirtualMachineScaleSetOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualMachineScaleSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualMachineScaleSetResult, error) {
			args := v.(LookupVirtualMachineScaleSetArgs)
			r, err := LookupVirtualMachineScaleSet(ctx, &args, opts...)
			var s LookupVirtualMachineScaleSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualMachineScaleSetResultOutput)
}

type LookupVirtualMachineScaleSetOutputArgs struct {
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	VmScaleSetName    pulumi.StringInput    `pulumi:"vmScaleSetName"`
}

func (LookupVirtualMachineScaleSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetArgs)(nil)).Elem()
}


type LookupVirtualMachineScaleSetResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualMachineScaleSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualMachineScaleSetResult)(nil)).Elem()
}

func (o LookupVirtualMachineScaleSetResultOutput) ToLookupVirtualMachineScaleSetResultOutput() LookupVirtualMachineScaleSetResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetResultOutput) ToLookupVirtualMachineScaleSetResultOutputWithContext(ctx context.Context) LookupVirtualMachineScaleSetResultOutput {
	return o
}

func (o LookupVirtualMachineScaleSetResultOutput) AdditionalCapabilities() AdditionalCapabilitiesResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *AdditionalCapabilitiesResponse {
		return v.AdditionalCapabilities
	}).(AdditionalCapabilitiesResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) AutomaticRepairsPolicy() AutomaticRepairsPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *AutomaticRepairsPolicyResponse {
		return v.AutomaticRepairsPolicy
	}).(AutomaticRepairsPolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) DoNotRunExtensionsOnOverprovisionedVMs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *bool { return v.DoNotRunExtensionsOnOverprovisionedVMs }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) HostGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *SubResourceResponse { return v.HostGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Identity() VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *VirtualMachineScaleSetIdentityResponse { return v.Identity }).(VirtualMachineScaleSetIdentityResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) OrchestrationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *string { return v.OrchestrationMode }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Overprovision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *bool { return v.Overprovision }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *int { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *SubResourceResponse { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ScaleInPolicy() ScaleInPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *ScaleInPolicyResponse { return v.ScaleInPolicy }).(ScaleInPolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) SinglePlacementGroup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *bool { return v.SinglePlacementGroup }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) SpotRestorePolicy() SpotRestorePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *SpotRestorePolicyResponse { return v.SpotRestorePolicy }).(SpotRestorePolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) UpgradePolicy() UpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *UpgradePolicyResponse { return v.UpgradePolicy }).(UpgradePolicyResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) VirtualMachineProfile() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *VirtualMachineScaleSetVMProfileResponse {
		return v.VirtualMachineProfile
	}).(VirtualMachineScaleSetVMProfileResponsePtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) ZoneBalance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) *bool { return v.ZoneBalance }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualMachineScaleSetResultOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualMachineScaleSetResult) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualMachineScaleSetResultOutput{})
}
