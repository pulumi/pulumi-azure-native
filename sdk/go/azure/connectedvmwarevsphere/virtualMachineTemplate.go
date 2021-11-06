


package connectedvmwarevsphere

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
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:VirtualMachineTemplate"),
		},
		{
			Type: pulumi.String("azure-nextgen:connectedvmwarevsphere/v20201001preview:VirtualMachineTemplate"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineTemplate
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:VirtualMachineTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineTemplateState, opts ...pulumi.ResourceOption) (*VirtualMachineTemplate, error) {
	var resource VirtualMachineTemplate
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:VirtualMachineTemplate", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*VirtualMachineTemplate)(nil))
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return i.ToVirtualMachineTemplateOutputWithContext(context.Background())
}

func (i *VirtualMachineTemplate) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineTemplateOutput)
}

type VirtualMachineTemplateOutput struct{ *pulumi.OutputState }

func (VirtualMachineTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineTemplate)(nil))
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutput() VirtualMachineTemplateOutput {
	return o
}

func (o VirtualMachineTemplateOutput) ToVirtualMachineTemplateOutputWithContext(ctx context.Context) VirtualMachineTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineTemplateOutput{})
}
