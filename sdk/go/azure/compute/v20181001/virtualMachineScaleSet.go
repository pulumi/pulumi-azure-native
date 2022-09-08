


package v20181001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	AutomaticRepairsPolicy                 AutomaticRepairsPolicyResponsePtrOutput          `pulumi:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrOutput                             `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	Identity                               VirtualMachineScaleSetIdentityResponsePtrOutput  `pulumi:"identity"`
	Location                               pulumi.StringOutput                              `pulumi:"location"`
	Name                                   pulumi.StringOutput                              `pulumi:"name"`
	Overprovision                          pulumi.BoolPtrOutput                             `pulumi:"overprovision"`
	Plan                                   PlanResponsePtrOutput                            `pulumi:"plan"`
	PlatformFaultDomainCount               pulumi.IntPtrOutput                              `pulumi:"platformFaultDomainCount"`
	ProvisioningState                      pulumi.StringOutput                              `pulumi:"provisioningState"`
	ProximityPlacementGroup                SubResourceResponsePtrOutput                     `pulumi:"proximityPlacementGroup"`
	SinglePlacementGroup                   pulumi.BoolPtrOutput                             `pulumi:"singlePlacementGroup"`
	Sku                                    SkuResponsePtrOutput                             `pulumi:"sku"`
	Tags                                   pulumi.StringMapOutput                           `pulumi:"tags"`
	Type                                   pulumi.StringOutput                              `pulumi:"type"`
	UniqueId                               pulumi.StringOutput                              `pulumi:"uniqueId"`
	UpgradePolicy                          UpgradePolicyResponsePtrOutput                   `pulumi:"upgradePolicy"`
	VirtualMachineProfile                  VirtualMachineScaleSetVMProfileResponsePtrOutput `pulumi:"virtualMachineProfile"`
	ZoneBalance                            pulumi.BoolPtrOutput                             `pulumi:"zoneBalance"`
	Zones                                  pulumi.StringArrayOutput                         `pulumi:"zones"`
}


func NewVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSet
	err := ctx.RegisterResource("azure-native:compute/v20181001:VirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	var resource VirtualMachineScaleSet
	err := ctx.ReadResource("azure-native:compute/v20181001:VirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScaleSetState struct {
}

type VirtualMachineScaleSetState struct {
}

func (VirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetState)(nil)).Elem()
}

type virtualMachineScaleSetArgs struct {
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy          `pulumi:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                            `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	Identity                               *VirtualMachineScaleSetIdentity  `pulumi:"identity"`
	Location                               *string                          `pulumi:"location"`
	Overprovision                          *bool                            `pulumi:"overprovision"`
	Plan                                   *Plan                            `pulumi:"plan"`
	PlatformFaultDomainCount               *int                             `pulumi:"platformFaultDomainCount"`
	ProximityPlacementGroup                *SubResource                     `pulumi:"proximityPlacementGroup"`
	ResourceGroupName                      string                           `pulumi:"resourceGroupName"`
	SinglePlacementGroup                   *bool                            `pulumi:"singlePlacementGroup"`
	Sku                                    *Sku                             `pulumi:"sku"`
	Tags                                   map[string]string                `pulumi:"tags"`
	UpgradePolicy                          *UpgradePolicy                   `pulumi:"upgradePolicy"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfile `pulumi:"virtualMachineProfile"`
	VmScaleSetName                         *string                          `pulumi:"vmScaleSetName"`
	ZoneBalance                            *bool                            `pulumi:"zoneBalance"`
	Zones                                  []string                         `pulumi:"zones"`
}


type VirtualMachineScaleSetArgs struct {
	AutomaticRepairsPolicy                 AutomaticRepairsPolicyPtrInput
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrInput
	Identity                               VirtualMachineScaleSetIdentityPtrInput
	Location                               pulumi.StringPtrInput
	Overprovision                          pulumi.BoolPtrInput
	Plan                                   PlanPtrInput
	PlatformFaultDomainCount               pulumi.IntPtrInput
	ProximityPlacementGroup                SubResourcePtrInput
	ResourceGroupName                      pulumi.StringInput
	SinglePlacementGroup                   pulumi.BoolPtrInput
	Sku                                    SkuPtrInput
	Tags                                   pulumi.StringMapInput
	UpgradePolicy                          UpgradePolicyPtrInput
	VirtualMachineProfile                  VirtualMachineScaleSetVMProfilePtrInput
	VmScaleSetName                         pulumi.StringPtrInput
	ZoneBalance                            pulumi.BoolPtrInput
	Zones                                  pulumi.StringArrayInput
}

func (VirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetArgs)(nil)).Elem()
}

type VirtualMachineScaleSetInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput
	ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput
}

func (*VirtualMachineScaleSet) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSet)(nil)).Elem()
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return i.ToVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOutput)
}

type VirtualMachineScaleSetOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSet)(nil)).Elem()
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return o
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return o
}

func (o VirtualMachineScaleSetOutput) AutomaticRepairsPolicy() AutomaticRepairsPolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) AutomaticRepairsPolicyResponsePtrOutput {
		return v.AutomaticRepairsPolicy
	}).(AutomaticRepairsPolicyResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) DoNotRunExtensionsOnOverprovisionedVMs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.DoNotRunExtensionsOnOverprovisionedVMs }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOutput) Identity() VirtualMachineScaleSetIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) VirtualMachineScaleSetIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineScaleSetIdentityResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOutput) Overprovision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.Overprovision }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.IntPtrOutput { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) SubResourceResponsePtrOutput { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) SinglePlacementGroup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.SinglePlacementGroup }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineScaleSetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetOutput) UpgradePolicy() UpgradePolicyResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) UpgradePolicyResponsePtrOutput { return v.UpgradePolicy }).(UpgradePolicyResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) VirtualMachineProfile() VirtualMachineScaleSetVMProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) VirtualMachineScaleSetVMProfileResponsePtrOutput {
		return v.VirtualMachineProfile
	}).(VirtualMachineScaleSetVMProfileResponsePtrOutput)
}

func (o VirtualMachineScaleSetOutput) ZoneBalance() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.BoolPtrOutput { return v.ZoneBalance }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSet) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetOutput{})
}
