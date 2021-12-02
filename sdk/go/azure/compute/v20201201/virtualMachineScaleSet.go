


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	AdditionalCapabilities                 AdditionalCapabilitiesResponsePtrOutput          `pulumi:"additionalCapabilities"`
	AutomaticRepairsPolicy                 AutomaticRepairsPolicyResponsePtrOutput          `pulumi:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrOutput                             `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	ExtendedLocation                       ExtendedLocationResponsePtrOutput                `pulumi:"extendedLocation"`
	HostGroup                              SubResourceResponsePtrOutput                     `pulumi:"hostGroup"`
	Identity                               VirtualMachineScaleSetIdentityResponsePtrOutput  `pulumi:"identity"`
	Location                               pulumi.StringOutput                              `pulumi:"location"`
	Name                                   pulumi.StringOutput                              `pulumi:"name"`
	OrchestrationMode                      pulumi.StringPtrOutput                           `pulumi:"orchestrationMode"`
	Overprovision                          pulumi.BoolPtrOutput                             `pulumi:"overprovision"`
	Plan                                   PlanResponsePtrOutput                            `pulumi:"plan"`
	PlatformFaultDomainCount               pulumi.IntPtrOutput                              `pulumi:"platformFaultDomainCount"`
	ProvisioningState                      pulumi.StringOutput                              `pulumi:"provisioningState"`
	ProximityPlacementGroup                SubResourceResponsePtrOutput                     `pulumi:"proximityPlacementGroup"`
	ScaleInPolicy                          ScaleInPolicyResponsePtrOutput                   `pulumi:"scaleInPolicy"`
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
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSet"),
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
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSet"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSet
	err := ctx.RegisterResource("azure-native:compute/v20201201:VirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	var resource VirtualMachineScaleSet
	err := ctx.ReadResource("azure-native:compute/v20201201:VirtualMachineScaleSet", name, id, state, &resource, opts...)
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
	AdditionalCapabilities                 *AdditionalCapabilities          `pulumi:"additionalCapabilities"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy          `pulumi:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                            `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	ExtendedLocation                       *ExtendedLocation                `pulumi:"extendedLocation"`
	HostGroup                              *SubResource                     `pulumi:"hostGroup"`
	Identity                               *VirtualMachineScaleSetIdentity  `pulumi:"identity"`
	Location                               *string                          `pulumi:"location"`
	OrchestrationMode                      *string                          `pulumi:"orchestrationMode"`
	Overprovision                          *bool                            `pulumi:"overprovision"`
	Plan                                   *Plan                            `pulumi:"plan"`
	PlatformFaultDomainCount               *int                             `pulumi:"platformFaultDomainCount"`
	ProximityPlacementGroup                *SubResource                     `pulumi:"proximityPlacementGroup"`
	ResourceGroupName                      string                           `pulumi:"resourceGroupName"`
	ScaleInPolicy                          *ScaleInPolicy                   `pulumi:"scaleInPolicy"`
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
	AdditionalCapabilities                 AdditionalCapabilitiesPtrInput
	AutomaticRepairsPolicy                 AutomaticRepairsPolicyPtrInput
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrInput
	ExtendedLocation                       ExtendedLocationPtrInput
	HostGroup                              SubResourcePtrInput
	Identity                               VirtualMachineScaleSetIdentityPtrInput
	Location                               pulumi.StringPtrInput
	OrchestrationMode                      pulumi.StringPtrInput
	Overprovision                          pulumi.BoolPtrInput
	Plan                                   PlanPtrInput
	PlatformFaultDomainCount               pulumi.IntPtrInput
	ProximityPlacementGroup                SubResourcePtrInput
	ResourceGroupName                      pulumi.StringInput
	ScaleInPolicy                          ScaleInPolicyPtrInput
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
	return reflect.TypeOf((*VirtualMachineScaleSet)(nil))
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return i.ToVirtualMachineScaleSetOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSet) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetOutput)
}

type VirtualMachineScaleSetOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineScaleSet)(nil))
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutput() VirtualMachineScaleSetOutput {
	return o
}

func (o VirtualMachineScaleSetOutput) ToVirtualMachineScaleSetOutputWithContext(ctx context.Context) VirtualMachineScaleSetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetOutput{})
}
