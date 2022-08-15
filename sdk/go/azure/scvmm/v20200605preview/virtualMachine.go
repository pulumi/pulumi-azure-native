


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	AvailabilitySets  VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput `pulumi:"availabilitySets"`
	CheckpointType    pulumi.StringPtrOutput                                      `pulumi:"checkpointType"`
	Checkpoints       CheckpointResponseArrayOutput                               `pulumi:"checkpoints"`
	CloudId           pulumi.StringPtrOutput                                      `pulumi:"cloudId"`
	ExtendedLocation  ExtendedLocationResponseOutput                              `pulumi:"extendedLocation"`
	Generation        pulumi.IntPtrOutput                                         `pulumi:"generation"`
	HardwareProfile   HardwareProfileResponsePtrOutput                            `pulumi:"hardwareProfile"`
	InventoryItemId   pulumi.StringPtrOutput                                      `pulumi:"inventoryItemId"`
	Location          pulumi.StringOutput                                         `pulumi:"location"`
	Name              pulumi.StringOutput                                         `pulumi:"name"`
	NetworkProfile    NetworkProfileResponsePtrOutput                             `pulumi:"networkProfile"`
	OsProfile         OsProfileResponsePtrOutput                                  `pulumi:"osProfile"`
	PowerState        pulumi.StringOutput                                         `pulumi:"powerState"`
	ProvisioningState pulumi.StringOutput                                         `pulumi:"provisioningState"`
	StorageProfile    StorageProfileResponsePtrOutput                             `pulumi:"storageProfile"`
	SystemData        SystemDataResponseOutput                                    `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                                      `pulumi:"tags"`
	TemplateId        pulumi.StringPtrOutput                                      `pulumi:"templateId"`
	Type              pulumi.StringOutput                                         `pulumi:"type"`
	Uuid              pulumi.StringPtrOutput                                      `pulumi:"uuid"`
	VmName            pulumi.StringPtrOutput                                      `pulumi:"vmName"`
	VmmServerId       pulumi.StringPtrOutput                                      `pulumi:"vmmServerId"`
}


func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExtendedLocation == nil {
		return nil, errors.New("invalid value for required argument 'ExtendedLocation'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:scvmm/v20200605preview:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:scvmm/v20200605preview:VirtualMachine", name, id, state, &resource, opts...)
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
	AvailabilitySets   []VirtualMachinePropertiesAvailabilitySets `pulumi:"availabilitySets"`
	CheckpointType     *string                                    `pulumi:"checkpointType"`
	Checkpoints        []Checkpoint                               `pulumi:"checkpoints"`
	CloudId            *string                                    `pulumi:"cloudId"`
	ExtendedLocation   ExtendedLocation                           `pulumi:"extendedLocation"`
	Generation         *int                                       `pulumi:"generation"`
	HardwareProfile    *HardwareProfile                           `pulumi:"hardwareProfile"`
	InventoryItemId    *string                                    `pulumi:"inventoryItemId"`
	Location           *string                                    `pulumi:"location"`
	NetworkProfile     *NetworkProfile                            `pulumi:"networkProfile"`
	OsProfile          *OsProfile                                 `pulumi:"osProfile"`
	ResourceGroupName  string                                     `pulumi:"resourceGroupName"`
	StorageProfile     *StorageProfile                            `pulumi:"storageProfile"`
	Tags               map[string]string                          `pulumi:"tags"`
	TemplateId         *string                                    `pulumi:"templateId"`
	Uuid               *string                                    `pulumi:"uuid"`
	VirtualMachineName *string                                    `pulumi:"virtualMachineName"`
	VmName             *string                                    `pulumi:"vmName"`
	VmmServerId        *string                                    `pulumi:"vmmServerId"`
}


type VirtualMachineArgs struct {
	AvailabilitySets   VirtualMachinePropertiesAvailabilitySetsArrayInput
	CheckpointType     pulumi.StringPtrInput
	Checkpoints        CheckpointArrayInput
	CloudId            pulumi.StringPtrInput
	ExtendedLocation   ExtendedLocationInput
	Generation         pulumi.IntPtrInput
	HardwareProfile    HardwareProfilePtrInput
	InventoryItemId    pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	NetworkProfile     NetworkProfilePtrInput
	OsProfile          OsProfilePtrInput
	ResourceGroupName  pulumi.StringInput
	StorageProfile     StorageProfilePtrInput
	Tags               pulumi.StringMapInput
	TemplateId         pulumi.StringPtrInput
	Uuid               pulumi.StringPtrInput
	VirtualMachineName pulumi.StringPtrInput
	VmName             pulumi.StringPtrInput
	VmmServerId        pulumi.StringPtrInput
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

func (o VirtualMachineOutput) AvailabilitySets() VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
		return v.AvailabilitySets
	}).(VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput)
}

func (o VirtualMachineOutput) CheckpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CheckpointType }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) Checkpoints() CheckpointResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) CheckpointResponseArrayOutput { return v.Checkpoints }).(CheckpointResponseArrayOutput)
}

func (o VirtualMachineOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.CloudId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualMachine) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o VirtualMachineOutput) Generation() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntPtrOutput { return v.Generation }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineOutput) HardwareProfile() HardwareProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) HardwareProfileResponsePtrOutput { return v.HardwareProfile }).(HardwareProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
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

func (o VirtualMachineOutput) OsProfile() OsProfileResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) OsProfileResponsePtrOutput { return v.OsProfile }).(OsProfileResponsePtrOutput)
}

func (o VirtualMachineOutput) PowerState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PowerState }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
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

func (o VirtualMachineOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) VmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VmName }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
