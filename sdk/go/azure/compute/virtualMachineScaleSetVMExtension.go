


package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSetVMExtension struct {
	pulumi.CustomResourceState

	AutoUpgradeMinorVersion pulumi.BoolPtrOutput                                 `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade  pulumi.BoolPtrOutput                                 `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag          pulumi.StringPtrOutput                               `pulumi:"forceUpdateTag"`
	InstanceView            VirtualMachineExtensionInstanceViewResponsePtrOutput `pulumi:"instanceView"`
	Name                    pulumi.StringOutput                                  `pulumi:"name"`
	ProtectedSettings       pulumi.AnyOutput                                     `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringOutput                                  `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrOutput                               `pulumi:"publisher"`
	Settings                pulumi.AnyOutput                                     `pulumi:"settings"`
	Type                    pulumi.StringOutput                                  `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrOutput                               `pulumi:"typeHandlerVersion"`
}


func NewVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmScaleSetName == nil {
		return nil, errors.New("invalid value for required argument 'VmScaleSetName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetVMExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetVMExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.RegisterResource("azure-native:compute:VirtualMachineScaleSetVMExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSetVMExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMExtension, error) {
	var resource VirtualMachineScaleSetVMExtension
	err := ctx.ReadResource("azure-native:compute:VirtualMachineScaleSetVMExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScaleSetVMExtensionState struct {
}

type VirtualMachineScaleSetVMExtensionState struct {
}

func (VirtualMachineScaleSetVMExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetVMExtensionArgs struct {
	AutoUpgradeMinorVersion *bool                                `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade  *bool                                `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag          *string                              `pulumi:"forceUpdateTag"`
	InstanceId              string                               `pulumi:"instanceId"`
	InstanceView            *VirtualMachineExtensionInstanceView `pulumi:"instanceView"`
	ProtectedSettings       interface{}                          `pulumi:"protectedSettings"`
	Publisher               *string                              `pulumi:"publisher"`
	ResourceGroupName       string                               `pulumi:"resourceGroupName"`
	Settings                interface{}                          `pulumi:"settings"`
	Type                    *string                              `pulumi:"type"`
	TypeHandlerVersion      *string                              `pulumi:"typeHandlerVersion"`
	VmExtensionName         *string                              `pulumi:"vmExtensionName"`
	VmScaleSetName          string                               `pulumi:"vmScaleSetName"`
}


type VirtualMachineScaleSetVMExtensionArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	EnableAutomaticUpgrade  pulumi.BoolPtrInput
	ForceUpdateTag          pulumi.StringPtrInput
	InstanceId              pulumi.StringInput
	InstanceView            VirtualMachineExtensionInstanceViewPtrInput
	ProtectedSettings       pulumi.Input
	Publisher               pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Settings                pulumi.Input
	Type                    pulumi.StringPtrInput
	TypeHandlerVersion      pulumi.StringPtrInput
	VmExtensionName         pulumi.StringPtrInput
	VmScaleSetName          pulumi.StringInput
}

func (VirtualMachineScaleSetVMExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMExtensionArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput
	ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput
}

func (*VirtualMachineScaleSetVMExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMExtension)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVMExtension) ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput {
	return i.ToVirtualMachineScaleSetVMExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVMExtension) ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMExtensionOutput)
}

type VirtualMachineScaleSetVMExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMExtensionOutput) ToVirtualMachineScaleSetVMExtensionOutput() VirtualMachineScaleSetVMExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetVMExtensionOutput) ToVirtualMachineScaleSetVMExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetVMExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) VirtualMachineExtensionInstanceViewResponsePtrOutput {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMExtensionOutput{})
}
