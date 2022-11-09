


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualmachineRetrieve struct {
	pulumi.CustomResourceState

	ExtendedLocation  ExtendedLocationResponsePtrOutput                         `pulumi:"extendedLocation"`
	HardwareProfile   VirtualmachinesPropertiesResponseHardwareProfilePtrOutput `pulumi:"hardwareProfile"`
	Location          pulumi.StringOutput                                       `pulumi:"location"`
	Name              pulumi.StringOutput                                       `pulumi:"name"`
	NetworkProfile    VirtualmachinesPropertiesResponseNetworkProfilePtrOutput  `pulumi:"networkProfile"`
	OsProfile         VirtualmachinesPropertiesResponseOsProfilePtrOutput       `pulumi:"osProfile"`
	ProvisioningState pulumi.StringOutput                                       `pulumi:"provisioningState"`
	ResourceName      pulumi.StringPtrOutput                                    `pulumi:"resourceName"`
	SecurityProfile   VirtualmachinesPropertiesResponseSecurityProfilePtrOutput `pulumi:"securityProfile"`
	Status            VirtualMachineStatusResponseOutput                        `pulumi:"status"`
	StorageProfile    VirtualmachinesPropertiesResponseStorageProfilePtrOutput  `pulumi:"storageProfile"`
	SystemData        SystemDataResponseOutput                                  `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput                                    `pulumi:"tags"`
	Type              pulumi.StringOutput                                       `pulumi:"type"`
}


func NewVirtualmachineRetrieve(ctx *pulumi.Context,
	name string, args *VirtualmachineRetrieveArgs, opts ...pulumi.ResourceOption) (*VirtualmachineRetrieve, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource VirtualmachineRetrieve
	err := ctx.RegisterResource("azure-native:azurestackhci/v20210701preview:virtualmachineRetrieve", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualmachineRetrieve(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualmachineRetrieveState, opts ...pulumi.ResourceOption) (*VirtualmachineRetrieve, error) {
	var resource VirtualmachineRetrieve
	err := ctx.ReadResource("azure-native:azurestackhci/v20210701preview:virtualmachineRetrieve", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualmachineRetrieveState struct {
}

type VirtualmachineRetrieveState struct {
}

func (VirtualmachineRetrieveState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualmachineRetrieveState)(nil)).Elem()
}

type virtualmachineRetrieveArgs struct {
	ExtendedLocation    *ExtendedLocation                         `pulumi:"extendedLocation"`
	HardwareProfile     *VirtualmachinesPropertiesHardwareProfile `pulumi:"hardwareProfile"`
	Location            *string                                   `pulumi:"location"`
	NetworkProfile      *VirtualmachinesPropertiesNetworkProfile  `pulumi:"networkProfile"`
	OsProfile           *VirtualmachinesPropertiesOsProfile       `pulumi:"osProfile"`
	ResourceGroupName   string                                    `pulumi:"resourceGroupName"`
	ResourceName        *string                                   `pulumi:"resourceName"`
	SecurityProfile     *VirtualmachinesPropertiesSecurityProfile `pulumi:"securityProfile"`
	StorageProfile      *VirtualmachinesPropertiesStorageProfile  `pulumi:"storageProfile"`
	Tags                map[string]string                         `pulumi:"tags"`
	VirtualmachinesName *string                                   `pulumi:"virtualmachinesName"`
}


type VirtualmachineRetrieveArgs struct {
	ExtendedLocation    ExtendedLocationPtrInput
	HardwareProfile     VirtualmachinesPropertiesHardwareProfilePtrInput
	Location            pulumi.StringPtrInput
	NetworkProfile      VirtualmachinesPropertiesNetworkProfilePtrInput
	OsProfile           VirtualmachinesPropertiesOsProfilePtrInput
	ResourceGroupName   pulumi.StringInput
	ResourceName        pulumi.StringPtrInput
	SecurityProfile     VirtualmachinesPropertiesSecurityProfilePtrInput
	StorageProfile      VirtualmachinesPropertiesStorageProfilePtrInput
	Tags                pulumi.StringMapInput
	VirtualmachinesName pulumi.StringPtrInput
}

func (VirtualmachineRetrieveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualmachineRetrieveArgs)(nil)).Elem()
}

type VirtualmachineRetrieveInput interface {
	pulumi.Input

	ToVirtualmachineRetrieveOutput() VirtualmachineRetrieveOutput
	ToVirtualmachineRetrieveOutputWithContext(ctx context.Context) VirtualmachineRetrieveOutput
}

func (*VirtualmachineRetrieve) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachineRetrieve)(nil)).Elem()
}

func (i *VirtualmachineRetrieve) ToVirtualmachineRetrieveOutput() VirtualmachineRetrieveOutput {
	return i.ToVirtualmachineRetrieveOutputWithContext(context.Background())
}

func (i *VirtualmachineRetrieve) ToVirtualmachineRetrieveOutputWithContext(ctx context.Context) VirtualmachineRetrieveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualmachineRetrieveOutput)
}

type VirtualmachineRetrieveOutput struct{ *pulumi.OutputState }

func (VirtualmachineRetrieveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualmachineRetrieve)(nil)).Elem()
}

func (o VirtualmachineRetrieveOutput) ToVirtualmachineRetrieveOutput() VirtualmachineRetrieveOutput {
	return o
}

func (o VirtualmachineRetrieveOutput) ToVirtualmachineRetrieveOutputWithContext(ctx context.Context) VirtualmachineRetrieveOutput {
	return o
}

func (o VirtualmachineRetrieveOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualmachineRetrieveOutput) HardwareProfile() VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualmachinesPropertiesResponseHardwareProfilePtrOutput {
		return v.HardwareProfile
	}).(VirtualmachinesPropertiesResponseHardwareProfilePtrOutput)
}

func (o VirtualmachineRetrieveOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualmachineRetrieveOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualmachineRetrieveOutput) NetworkProfile() VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualmachinesPropertiesResponseNetworkProfilePtrOutput {
		return v.NetworkProfile
	}).(VirtualmachinesPropertiesResponseNetworkProfilePtrOutput)
}

func (o VirtualmachineRetrieveOutput) OsProfile() VirtualmachinesPropertiesResponseOsProfilePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualmachinesPropertiesResponseOsProfilePtrOutput {
		return v.OsProfile
	}).(VirtualmachinesPropertiesResponseOsProfilePtrOutput)
}

func (o VirtualmachineRetrieveOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualmachineRetrieveOutput) ResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringPtrOutput { return v.ResourceName }).(pulumi.StringPtrOutput)
}

func (o VirtualmachineRetrieveOutput) SecurityProfile() VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualmachinesPropertiesResponseSecurityProfilePtrOutput {
		return v.SecurityProfile
	}).(VirtualmachinesPropertiesResponseSecurityProfilePtrOutput)
}

func (o VirtualmachineRetrieveOutput) Status() VirtualMachineStatusResponseOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualMachineStatusResponseOutput { return v.Status }).(VirtualMachineStatusResponseOutput)
}

func (o VirtualmachineRetrieveOutput) StorageProfile() VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) VirtualmachinesPropertiesResponseStorageProfilePtrOutput {
		return v.StorageProfile
	}).(VirtualmachinesPropertiesResponseStorageProfilePtrOutput)
}

func (o VirtualmachineRetrieveOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualmachineRetrieveOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualmachineRetrieveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualmachineRetrieve) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualmachineRetrieveOutput{})
}
