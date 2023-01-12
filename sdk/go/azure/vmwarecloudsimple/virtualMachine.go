


package vmwarecloudsimple

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachine struct {
	pulumi.CustomResourceState

	AmountOfRam       pulumi.IntOutput                         `pulumi:"amountOfRam"`
	Controllers       VirtualDiskControllerResponseArrayOutput `pulumi:"controllers"`
	Customization     GuestOSCustomizationResponsePtrOutput    `pulumi:"customization"`
	Disks             VirtualDiskResponseArrayOutput           `pulumi:"disks"`
	Dnsname           pulumi.StringOutput                      `pulumi:"dnsname"`
	ExposeToGuestVM   pulumi.BoolPtrOutput                     `pulumi:"exposeToGuestVM"`
	Folder            pulumi.StringOutput                      `pulumi:"folder"`
	GuestOS           pulumi.StringOutput                      `pulumi:"guestOS"`
	GuestOSType       pulumi.StringOutput                      `pulumi:"guestOSType"`
	Location          pulumi.StringOutput                      `pulumi:"location"`
	Name              pulumi.StringOutput                      `pulumi:"name"`
	Nics              VirtualNicResponseArrayOutput            `pulumi:"nics"`
	NumberOfCores     pulumi.IntOutput                         `pulumi:"numberOfCores"`
	Password          pulumi.StringPtrOutput                   `pulumi:"password"`
	PrivateCloudId    pulumi.StringOutput                      `pulumi:"privateCloudId"`
	ProvisioningState pulumi.StringOutput                      `pulumi:"provisioningState"`
	PublicIP          pulumi.StringOutput                      `pulumi:"publicIP"`
	ResourcePool      ResourcePoolResponsePtrOutput            `pulumi:"resourcePool"`
	Status            pulumi.StringOutput                      `pulumi:"status"`
	Tags              pulumi.StringMapOutput                   `pulumi:"tags"`
	TemplateId        pulumi.StringPtrOutput                   `pulumi:"templateId"`
	Type              pulumi.StringOutput                      `pulumi:"type"`
	Username          pulumi.StringPtrOutput                   `pulumi:"username"`
	VSphereNetworks   pulumi.StringArrayOutput                 `pulumi:"vSphereNetworks"`
	VmId              pulumi.StringOutput                      `pulumi:"vmId"`
	Vmwaretools       pulumi.StringOutput                      `pulumi:"vmwaretools"`
}


func NewVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AmountOfRam == nil {
		return nil, errors.New("invalid value for required argument 'AmountOfRam'")
	}
	if args.NumberOfCores == nil {
		return nil, errors.New("invalid value for required argument 'NumberOfCores'")
	}
	if args.PrivateCloudId == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:vmwarecloudsimple/v20190401:VirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachine
	err := ctx.RegisterResource("azure-native:vmwarecloudsimple:VirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachine, error) {
	var resource VirtualMachine
	err := ctx.ReadResource("azure-native:vmwarecloudsimple:VirtualMachine", name, id, state, &resource, opts...)
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
	AmountOfRam        int                   `pulumi:"amountOfRam"`
	Customization      *GuestOSCustomization `pulumi:"customization"`
	Disks              []VirtualDisk         `pulumi:"disks"`
	ExposeToGuestVM    *bool                 `pulumi:"exposeToGuestVM"`
	Location           *string               `pulumi:"location"`
	Nics               []VirtualNic          `pulumi:"nics"`
	NumberOfCores      int                   `pulumi:"numberOfCores"`
	Password           *string               `pulumi:"password"`
	PrivateCloudId     string                `pulumi:"privateCloudId"`
	ResourceGroupName  string                `pulumi:"resourceGroupName"`
	ResourcePool       *ResourcePool         `pulumi:"resourcePool"`
	Tags               map[string]string     `pulumi:"tags"`
	TemplateId         *string               `pulumi:"templateId"`
	Username           *string               `pulumi:"username"`
	VSphereNetworks    []string              `pulumi:"vSphereNetworks"`
	VirtualMachineName *string               `pulumi:"virtualMachineName"`
}


type VirtualMachineArgs struct {
	AmountOfRam        pulumi.IntInput
	Customization      GuestOSCustomizationPtrInput
	Disks              VirtualDiskArrayInput
	ExposeToGuestVM    pulumi.BoolPtrInput
	Location           pulumi.StringPtrInput
	Nics               VirtualNicArrayInput
	NumberOfCores      pulumi.IntInput
	Password           pulumi.StringPtrInput
	PrivateCloudId     pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	ResourcePool       ResourcePoolPtrInput
	Tags               pulumi.StringMapInput
	TemplateId         pulumi.StringPtrInput
	Username           pulumi.StringPtrInput
	VSphereNetworks    pulumi.StringArrayInput
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

func (o VirtualMachineOutput) AmountOfRam() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.AmountOfRam }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Controllers() VirtualDiskControllerResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualDiskControllerResponseArrayOutput { return v.Controllers }).(VirtualDiskControllerResponseArrayOutput)
}

func (o VirtualMachineOutput) Customization() GuestOSCustomizationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) GuestOSCustomizationResponsePtrOutput { return v.Customization }).(GuestOSCustomizationResponsePtrOutput)
}

func (o VirtualMachineOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualDiskResponseArrayOutput { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o VirtualMachineOutput) Dnsname() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Dnsname }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ExposeToGuestVM() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.BoolPtrOutput { return v.ExposeToGuestVM }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) GuestOS() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.GuestOS }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) GuestOSType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.GuestOSType }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Nics() VirtualNicResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) VirtualNicResponseArrayOutput { return v.Nics }).(VirtualNicResponseArrayOutput)
}

func (o VirtualMachineOutput) NumberOfCores() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.IntOutput { return v.NumberOfCores }).(pulumi.IntOutput)
}

func (o VirtualMachineOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) PrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PrivateCloudId }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) PublicIP() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.PublicIP }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) ResourcePool() ResourcePoolResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachine) ResourcePoolResponsePtrOutput { return v.ResourcePool }).(ResourcePoolResponsePtrOutput)
}

func (o VirtualMachineOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
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

func (o VirtualMachineOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringPtrOutput { return v.Username }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineOutput) VSphereNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringArrayOutput { return v.VSphereNetworks }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.VmId }).(pulumi.StringOutput)
}

func (o VirtualMachineOutput) Vmwaretools() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachine) pulumi.StringOutput { return v.Vmwaretools }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineOutput{})
}
