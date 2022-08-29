


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualMachine struct {
	pulumi.CustomResourceState

	AvailabilitySet         SubResourceResponsePtrOutput               `pulumi:"availabilitySet"`
	DiagnosticsProfile      DiagnosticsProfileResponsePtrOutput        `pulumi:"diagnosticsProfile"`
	HardwareProfile         HardwareProfileResponsePtrOutput           `pulumi:"hardwareProfile"`
	Identity                VirtualMachineIdentityResponsePtrOutput    `pulumi:"identity"`
	InstanceView            VirtualMachineInstanceViewResponseOutput   `pulumi:"instanceView"`
	LicenseType             pulumi.StringPtrOutput                     `pulumi:"licenseType"`
	Location                pulumi.StringOutput                        `pulumi:"location"`
	Name                    pulumi.StringOutput                        `pulumi:"name"`
	NetworkProfile          NetworkProfileResponsePtrOutput            `pulumi:"networkProfile"`
	OsProfile               OSProfileResponsePtrOutput                 `pulumi:"osProfile"`
	Plan                    PlanResponsePtrOutput                      `pulumi:"plan"`
	ProvisioningState       pulumi.StringOutput                        `pulumi:"provisioningState"`
	ProximityPlacementGroup SubResourceResponsePtrOutput               `pulumi:"proximityPlacementGroup"`
	Resources               VirtualMachineExtensionResponseArrayOutput `pulumi:"resources"`
	StorageProfile          StorageProfileResponsePtrOutput            `pulumi:"storageProfile"`
	Tags                    pulumi.StringMapOutput                     `pulumi:"tags"`
	Type                    pulumi.StringOutput                        `pulumi:"type"`
	VmId                    pulumi.StringOutput                        `pulumi:"vmId"`
	Zones                   pulumi.StringArrayOutput                   `pulumi:"zones"`
}


func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:compute/v20180401:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:compute/v20180401:VirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineState struct {
}

type VirtualMachineState struct {
}

func (VirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineState)(nil)).Elem()
}

type virtualMachineArgs struct {
	AvailabilitySet         *SubResource            `pulumi:"availabilitySet"`
	DiagnosticsProfile      *DiagnosticsProfile     `pulumi:"diagnosticsProfile"`
	HardwareProfile         *HardwareProfile        `pulumi:"hardwareProfile"`
	Identity                *VirtualMachineIdentity `pulumi:"identity"`
	LicenseType             *string                 `pulumi:"licenseType"`
	Location                *string                 `pulumi:"location"`
	NetworkProfile          *NetworkProfile         `pulumi:"networkProfile"`
	OsProfile               *OSProfile              `pulumi:"osProfile"`
	Plan                    *Plan                   `pulumi:"plan"`
	ProximityPlacementGroup *SubResource            `pulumi:"proximityPlacementGroup"`
	ResourceGroupName       string                  `pulumi:"resourceGroupName"`
	StorageProfile          *StorageProfile         `pulumi:"storageProfile"`
	Tags                    map[string]string       `pulumi:"tags"`
	VmName                  *string                 `pulumi:"vmName"`
	Zones                   []string                `pulumi:"zones"`
}


type VirtualMachineArgs struct {
	AvailabilitySet         SubResourcePtrInput
	DiagnosticsProfile      DiagnosticsProfilePtrInput
	HardwareProfile         HardwareProfilePtrInput
	Identity                VirtualMachineIdentityPtrInput
	LicenseType             pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	NetworkProfile          NetworkProfilePtrInput
	OsProfile               OSProfilePtrInput
	Plan                    PlanPtrInput
	ProximityPlacementGroup SubResourcePtrInput
	ResourceGroupName       pulumi.StringInput
	StorageProfile          StorageProfilePtrInput
	Tags                    pulumi.StringMapInput
	VmName                  pulumi.StringPtrInput
	Zones                   pulumi.StringArrayInput
}

func (VirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineArgs)(nil)).Elem()
}

type VirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineOutput() VirtualMachineOutput
	ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput
}

func (*VirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachine)(nil)).Elem()
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) AvailabilitySet() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.AvailabilitySet }).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineOutput) DiagnosticsProfile() DiagnosticsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) DiagnosticsProfileResponsePtrOutput { return v.DiagnosticsProfile }).(DiagnosticsProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) Identity() VirtualMachineIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineIdentityResponsePtrOutput { return v.Identity }).(VirtualMachineIdentityResponsePtrOutput)
}

func (o VirtualMachineOutput) InstanceView() VirtualMachineInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineInstanceViewResponseOutput { return v.InstanceView }).(VirtualMachineInstanceViewResponseOutput)
}

func (o VirtualMachineOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) OsProfile() OSProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) OSProfileResponsePtrOutput { return v.OsProfile }).(OSProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) PlanResponsePtrOutput { return v.Plan }).(PlanResponsePtrOutput)
}

func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ProximityPlacementGroup() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) SubResourceResponsePtrOutput { return v.ProximityPlacementGroup }).(SubResourceResponsePtrOutput)
}

func (o VirtualMachineOutput) Resources() VirtualMachineExtensionResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachineExtensionResponseArrayOutput { return v.Resources }).(VirtualMachineExtensionResponseArrayOutput)
}

func (o VirtualMachineOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
