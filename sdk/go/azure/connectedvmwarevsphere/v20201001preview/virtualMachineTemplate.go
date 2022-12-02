


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineTemplate struct {
	pulumi.CustomResourceState

	CustomResourceName pulumi.StringOutput                 `pulumi:"customResourceName"`
	Disks              VirtualDiskResponseArrayOutput      `pulumi:"disks"`
	ExtendedLocation   ExtendedLocationResponsePtrOutput   `pulumi:"extendedLocation"`
	FirmwareType       pulumi.StringOutput                 `pulumi:"firmwareType"`
	FolderPath         pulumi.StringOutput                 `pulumi:"folderPath"`
	InventoryItemId    pulumi.StringPtrOutput              `pulumi:"inventoryItemId"`
	Kind               pulumi.StringPtrOutput              `pulumi:"kind"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	MemorySizeMB       pulumi.IntOutput                    `pulumi:"memorySizeMB"`
	MoName             pulumi.StringOutput                 `pulumi:"moName"`
	MoRefId            pulumi.StringPtrOutput              `pulumi:"moRefId"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	NetworkInterfaces  NetworkInterfaceResponseArrayOutput `pulumi:"networkInterfaces"`
	NumCPUs            pulumi.IntOutput                    `pulumi:"numCPUs"`
	NumCoresPerSocket  pulumi.IntOutput                    `pulumi:"numCoresPerSocket"`
	OsName             pulumi.StringOutput                 `pulumi:"osName"`
	OsType             pulumi.StringOutput                 `pulumi:"osType"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	Statuses           ResourceStatusResponseArrayOutput   `pulumi:"statuses"`
	SystemData         SystemDataResponseOutput            `pulumi:"systemData"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	ToolsVersion       pulumi.StringOutput                 `pulumi:"toolsVersion"`
	ToolsVersionStatus pulumi.StringOutput                 `pulumi:"toolsVersionStatus"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
	Uuid               pulumi.StringOutput                 `pulumi:"uuid"`
	VCenterId          pulumi.StringPtrOutput              `pulumi:"vCenterId"`
}


func NewVirtualMachineTemplate(ctx *pulumi.Context,
	name string, args *VirtualMachineTemplateArgs, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220715preview:VirtualMachineTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineTemplate
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachineTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	var resource VirtualMachineTemplate
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachineTemplate", name, id, state, &resource, opts...)
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
	ExtendedLocation           *ExtendedLocation `pulumi:"extendedLocation"`
	InventoryItemId            *string           `pulumi:"inventoryItemId"`
	Kind                       *string           `pulumi:"kind"`
	Location                   *string           `pulumi:"location"`
	MoRefId                    *string           `pulumi:"moRefId"`
	ResourceGroupName          string            `pulumi:"resourceGroupName"`
	Tags                       map[string]string `pulumi:"tags"`
	VCenterId                  *string           `pulumi:"vCenterId"`
	VirtualMachineTemplateName *string           `pulumi:"virtualMachineTemplateName"`
}


type VirtualMachineTemplateArgs struct {
	ExtendedLocation           ExtendedLocationPtrInput
	InventoryItemId            pulumi.StringPtrInput
	Kind                       pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	MoRefId                    pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	Tags                       pulumi.StringMapInput
	VCenterId                  pulumi.StringPtrInput
	VirtualMachineTemplateName pulumi.StringPtrInput
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

func (o VirtualMachineTemplateOutput) CustomResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.CustomResourceName }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) VirtualDiskResponseArrayOutput { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

func (o VirtualMachineTemplateOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o VirtualMachineTemplateOutput) FirmwareType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.FirmwareType }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) FolderPath() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.FolderPath }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) InventoryItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.InventoryItemId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineTemplateOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) MemorySizeMB() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.MemorySizeMB }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) MoName() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.MoName }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) MoRefId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.MoRefId }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) NetworkInterfaceResponseArrayOutput { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o VirtualMachineTemplateOutput) NumCPUs() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.NumCPUs }).(pulumi.IntOutput)
}

func (o VirtualMachineTemplateOutput) NumCoresPerSocket() pulumi.IntOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.IntOutput { return v.NumCoresPerSocket }).(pulumi.IntOutput)
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

func (o VirtualMachineTemplateOutput) Statuses() ResourceStatusResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) ResourceStatusResponseArrayOutput { return v.Statuses }).(ResourceStatusResponseArrayOutput)
}

func (o VirtualMachineTemplateOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VirtualMachineTemplateOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineTemplateOutput) ToolsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ToolsVersion }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) ToolsVersionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.ToolsVersionStatus }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

func (o VirtualMachineTemplateOutput) VCenterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineTemplate) pulumi.StringPtrOutput { return v.VCenterId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineTemplateOutput{})
}
