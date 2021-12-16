


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSetVM struct {
	pulumi.CustomResourceState

	AvailabilitySet    SubResourceResponsePtrOutput                       `pulumi:"availabilitySet"`
	DiagnosticsProfile DiagnosticsProfileResponsePtrOutput                `pulumi:"diagnosticsProfile"`
	HardwareProfile    HardwareProfileResponsePtrOutput                   `pulumi:"hardwareProfile"`
	InstanceId         pulumi.StringOutput                                `pulumi:"instanceId"`
	InstanceView       VirtualMachineScaleSetVMInstanceViewResponseOutput `pulumi:"instanceView"`
	LatestModelApplied pulumi.BoolOutput                                  `pulumi:"latestModelApplied"`
	LicenseType        pulumi.StringPtrOutput                             `pulumi:"licenseType"`
	Location           pulumi.StringOutput                                `pulumi:"location"`
	Name               pulumi.StringOutput                                `pulumi:"name"`
	NetworkProfile     NetworkProfileResponsePtrOutput                    `pulumi:"networkProfile"`
	OsProfile          OSProfileResponsePtrOutput                         `pulumi:"osProfile"`
	Plan               PlanResponsePtrOutput                              `pulumi:"plan"`
	ProvisioningState  pulumi.StringOutput                                `pulumi:"provisioningState"`
	Resources          VirtualMachineExtensionResponseArrayOutput         `pulumi:"resources"`
	Sku                SkuResponseOutput                                  `pulumi:"sku"`
	StorageProfile     StorageProfileResponsePtrOutput                    `pulumi:"storageProfile"`
	Tags               pulumi.StringMapOutput                             `pulumi:"tags"`
	Type               pulumi.StringOutput                                `pulumi:"type"`
	VmId               pulumi.StringOutput                                `pulumi:"vmId"`
	Zones              pulumi.StringArrayOutput                           `pulumi:"zones"`
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
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetVM
	err := ctx.RegisterResource("azure-native:compute/v20180401:VirtualMachineScaleSetVM", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSetVM(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVM, error) {
	var resource VirtualMachineScaleSetVM
	err := ctx.ReadResource("azure-native:compute/v20180401:VirtualMachineScaleSetVM", name, id, state, &resource, opts...)
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
	AvailabilitySet    *SubResource        `pulumi:"availabilitySet"`
	DiagnosticsProfile *DiagnosticsProfile `pulumi:"diagnosticsProfile"`
	HardwareProfile    *HardwareProfile    `pulumi:"hardwareProfile"`
	InstanceId         *string             `pulumi:"instanceId"`
	LicenseType        *string             `pulumi:"licenseType"`
	Location           *string             `pulumi:"location"`
	NetworkProfile     *NetworkProfile     `pulumi:"networkProfile"`
	OsProfile          *OSProfile          `pulumi:"osProfile"`
	Plan               *Plan               `pulumi:"plan"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	StorageProfile     *StorageProfile     `pulumi:"storageProfile"`
	Tags               map[string]string   `pulumi:"tags"`
	VmScaleSetName     string              `pulumi:"vmScaleSetName"`
}


type VirtualMachineScaleSetVMArgs struct {
	AvailabilitySet    SubResourcePtrInput
	DiagnosticsProfile DiagnosticsProfilePtrInput
	HardwareProfile    HardwareProfilePtrInput
	InstanceId         pulumi.StringPtrInput
	LicenseType        pulumi.StringPtrInput
	Location           pulumi.StringPtrInput
	NetworkProfile     NetworkProfilePtrInput
	OsProfile          OSProfilePtrInput
	Plan               PlanPtrInput
	ResourceGroupName  pulumi.StringInput
	StorageProfile     StorageProfilePtrInput
	Tags               pulumi.StringMapInput
	VmScaleSetName     pulumi.StringInput
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

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMOutput{})
}
