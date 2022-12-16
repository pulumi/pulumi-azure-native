


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineTemplate struct {
	pulumi.CustomResourceState

	ComputerName         pulumi.StringOutput                  `pulumi:"computerName"`
	CpuCount             pulumi.IntOutput                     `pulumi:"cpuCount"`
	Disks                VirtualDiskResponseArrayOutput       `pulumi:"disks"`
	DynamicMemoryEnabled pulumi.StringOutput                  `pulumi:"dynamicMemoryEnabled"`
	DynamicMemoryMaxMB   pulumi.IntOutput                     `pulumi:"dynamicMemoryMaxMB"`
	DynamicMemoryMinMB   pulumi.IntOutput                     `pulumi:"dynamicMemoryMinMB"`
	ExtendedLocation     ExtendedLocationResponseOutput       `pulumi:"extendedLocation"`
	Generation           pulumi.IntOutput                     `pulumi:"generation"`
	InventoryItemId      pulumi.StringPtrOutput               `pulumi:"inventoryItemId"`
	IsCustomizable       pulumi.StringOutput                  `pulumi:"isCustomizable"`
	IsHighlyAvailable    pulumi.StringOutput                  `pulumi:"isHighlyAvailable"`
	LimitCpuForMigration pulumi.StringOutput                  `pulumi:"limitCpuForMigration"`
	Location             pulumi.StringOutput                  `pulumi:"location"`
	MemoryMB             pulumi.IntOutput                     `pulumi:"memoryMB"`
	Name                 pulumi.StringOutput                  `pulumi:"name"`
	NetworkInterfaces    NetworkInterfacesResponseArrayOutput `pulumi:"networkInterfaces"`
	OsName               pulumi.StringOutput                  `pulumi:"osName"`
	OsType               pulumi.StringOutput                  `pulumi:"osType"`
	ProvisioningState    pulumi.StringOutput                  `pulumi:"provisioningState"`
	SystemData           SystemDataResponseOutput             `pulumi:"systemData"`
	Tags                 pulumi.StringMapOutput               `pulumi:"tags"`
	Type                 pulumi.StringOutput                  `pulumi:"type"`
	Uuid                 pulumi.StringPtrOutput               `pulumi:"uuid"`
	VmmServerId          pulumi.StringPtrOutput               `pulumi:"vmmServerId"`
}


func NewVirtualMachineTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
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
			Type: pulumi.String("azure-native:scvmm:VirtualMachineTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineTemplate
	err := ctx.RegisterResource("azure-native:scvmm/v20200605preview:VirtualMachineTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	var resource VirtualMachineTemplate
	err := ctx.ReadResource("azure-native:scvmm/v20200605preview:VirtualMachineTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineTemplateState struct {
}

type VirtualMachineTemplateState struct {
}

func (VirtualMachineTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineTemplateState)(nil)).Elem()
}

type virtualMachineTemplateArgs struct {
	ExtendedLocation           ExtendedLocation  `pulumi:"extendedLocation"`
	InventoryItemId            *string           `pulumi:"inventoryItemId"`
	Location                   *string           `pulumi:"location"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	Tags                       map[string]string `pulumi:"tags"`
	Uuid                       *string           `pulumi:"uuid"`
	VirtualMachineTemplateName *string           `pulumi:"virtualMachineTemplateName"`
	VmmServerId                *string           `pulumi:"vmmServerId"`
}


type VirtualMachineTemplateArgs struct {
	ExtendedLocation           ExtendedLocationInput
	InventoryItemId            pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
	Uuid                       pulumi.StringPtrInput
	VirtualMachineTemplateName pulumi.StringPtrInput
	VmmServerId                pulumi.StringPtrInput
}

func (VirtualMachineTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineTemplateArgs)(nil)).Elem()
}

type VirtualMachineTemplateInput interface {
	pulumi.Input

	ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput
	ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput
}

func (*VirtualMachineTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineTemplate)(nil)).Elem()
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return i.ToVirtualMachineTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineTemplateOutput)
}

type VirtualMachineTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineTemplate)(nil)).Elem()
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return o
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return o
}

func (o VirtualMachineTemplateOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ComputerName }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) CpuCount() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.CpuCount }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) VirtualDiskResponseArrayOutput { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o VirtualMachineTemplateOutput) DynamicMemoryEnabled() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.DynamicMemoryEnabled }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) DynamicMemoryMaxMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.DynamicMemoryMaxMB }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) DynamicMemoryMinMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.DynamicMemoryMinMB }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) ExtendedLocation() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) ExtendedLocationResponseOutput { return v.ExtendedLocation }).(ExtendedLocationResponseOutput)
}

func (o VirtualMachineTemplateOutput) Generation() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.Generation }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineTemplateOutput) IsCustomizable() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.IsCustomizable }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) IsHighlyAvailable() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.IsHighlyAvailable }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) LimitCpuForMigration() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.LimitCpuForMigration }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) MemoryMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.MemoryMB }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) NetworkInterfaces() NetworkInterfacesResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) NetworkInterfacesResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfacesResponseArrayOutput)
}

func (o VirtualMachineTemplateOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.OsName }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.OsType }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualMachineTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.Uuid }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineTemplateOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineTemplateOutput{})
}
