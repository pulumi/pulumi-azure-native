


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
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere/v20201001preview:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:VirtualMachine"),
		},
		{
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere:VirtualMachine"),
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
	return reflect.TypeOf((*VirtualMachine)(nil))
}

func (i *VirtualMachine) ToVirtualMachineOutput() VirtualMachineOutput {
	return i.ToVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachine) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineOutput)
}

type VirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachine)(nil))
}

func (o VirtualMachineOutput) ToVirtualMachineOutput() VirtualMachineOutput {
	return o
}

func (o VirtualMachineOutput) ToVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
