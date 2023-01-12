


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput                `pulumi:"customResourceName"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput  `pulumi:"extendedLocation"`
	FirmwareType       pulumi.StringPtrOutput             `pulumi:"firmwareType"`
	FolderPath         pulumi.StringOutput                `pulumi:"folderPath"`
	GuestAgentProfile  GuestAgentProfileResponsePtrOutput `pulumi:"guestAgentProfile"`
	HardwareProfile    HardwareProfileResponsePtrOutput   `pulumi:"hardwareProfile"`
	Identity           IdentityResponsePtrOutput          `pulumi:"identity"`
	InstanceUuid       pulumi.StringOutput                `pulumi:"instanceUuid"`
	InventoryItemId    pulumi.StringPtrOutput             `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput             `pulumi:"kind"`
	Location           pulumi.StringOutput                `pulumi:"location"`
	MoName             pulumi.StringOutput                `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput             `pulumi:"moRefId"`
	Name               pulumi.StringOutput                `pulumi:"name"`
	NetworkProfile     NetworkProfileResponsePtrOutput    `pulumi:"networkProfile"`
	OsProfile          OsProfileResponsePtrOutput         `pulumi:"osProfile"`
	PlacementProfile   PlacementProfileResponsePtrOutput  `pulumi:"placementProfile"`
	PowerState         pulumi.StringOutput                `pulumi:"powerState"`
	ProvisioningState  pulumi.StringOutput                `pulumi:"provisioningState"`
	ResourcePoolId     pulumi.StringPtrOutput             `pulumi:"resourcePoolId"`
	SmbiosUuid         pulumi.StringPtrOutput             `pulumi:"smbiosUuid"`
	Statuses           ResourceStatusResponseArrayOutput  `pulumi:"statuses"`
	StorageProfile     StorageProfileResponsePtrOutput    `pulumi:"storageProfile"`
	SystemData         SystemDataResponseOutput           `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput             `pulumi:"tags"`
	TemplateId         pulumi.StringPtrOutput             `pulumi:"templateId"`
	Type               pulumi.StringOutput                `pulumi:"type"`
	Uuid               pulumi.StringOutput                `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput             `pulumi:"vCenterId"`
	VmId               pulumi.StringOutput                `pulumi:"vmId"`
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
			Type: pulumi.String("azure-native:connectedvmwarevsphere:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachine", name, id, state, &resource, opts...)
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
	ExtendedLocation   *ExtendedLocation `pulumi:"extendedLocation"`
	FirmwareType       *string           `pulumi:"firmwareType"`
	HardwareProfile    *HardwareProfile  `pulumi:"hardwareProfile"`
	Identity           *Identity         `pulumi:"identity"`
	InventoryItemId    *string           `pulumi:"inventoryItemId"`
	Kind               *string           `pulumi:"kind"`
	Location           *string           `pulumi:"location"`
	MoRefId            *string           `pulumi:"moRefId"`
	NetworkProfile     *NetworkProfile   `pulumi:"networkProfile"`
	OsProfile          *OsProfile        `pulumi:"osProfile"`
	PlacementProfile   *PlacementProfile `pulumi:"placementProfile"`
	ResourceGroupName  string            `pulumi:"resourceGroupName"`
	ResourcePoolId     *string           `pulumi:"resourcePoolId"`
	SmbiosUuid         *string           `pulumi:"smbiosUuid"`
	StorageProfile     *StorageProfile   `pulumi:"storageProfile"`
	Tags               map[string]string `pulumi:"tags"`
	TemplateId         *string           `pulumi:"templateId"`
	VCenterId          *string           `pulumi:"vCenterId"`
	VirtualMachineName *string           `pulumi:"virtualMachineName"`
}


type VirtualMachineArgs struct {
	ExtendedLocation   ExtendedLocationPtrInput
	FirmwareType       pulumi.StringPtrInput
	HardwareProfile    HardwareProfilePtrInput
	Identity           IdentityPtrInput
	InventoryItemId    pulumi.StringPtrInput
	Kind               pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	MoRefId            pulumi.StringPtrInput
	NetworkProfile     NetworkProfilePtrInput
	OsProfile          OsProfilePtrInput
	PlacementProfile   PlacementProfilePtrInput
	ResourceGroupName  pulumi.StringInput
	ResourcePoolId     pulumi.StringPtrInput
	SmbiosUuid         pulumi.StringPtrInput
	StorageProfile     StorageProfilePtrInput
	Tags               pulumi.StringMapInput
	TemplateId         pulumi.StringPtrInput
	VCenterId          pulumi.StringPtrInput
	VirtualMachineName pulumi.StringPtrInput
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

func (o VirtualMachineOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualMachineOutput) FirmwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.FirmwareType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) GuestAgentProfile() GuestAgentProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) GuestAgentProfileResponsePtrOutput { return v.GuestAgentProfile }).(GuestAgentProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) IdentityResponsePtrOutput { return v.Identity }).(IdentityResponsePtrOutput)
}

func (o VirtualMachineOutput) InstanceUuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.InstanceUuid }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) NetworkProfile() NetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) NetworkProfileResponsePtrOutput { return v.NetworkProfile }).(NetworkProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) OsProfileResponsePtrOutput { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) PlacementProfile() PlacementProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) PlacementProfileResponsePtrOutput { return v.PlacementProfile }).(PlacementProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ResourcePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.ResourcePoolId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) SmbiosUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.SmbiosUuid }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o VirtualMachineOutput) StorageProfile() StorageProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) StorageProfileResponsePtrOutput { return v.StorageProfile }).(StorageProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.TemplateId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
