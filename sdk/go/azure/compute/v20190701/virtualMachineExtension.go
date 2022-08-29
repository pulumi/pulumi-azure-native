


package v20190701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type VirtualMachineExtension struct {
	pulumi.CustomResourceState

	AutoUpgradeMinorVersion pulumi.BoolPtrOutput                                 `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          pulumi.StringPtrOutput                               `pulumi:"forceUpdateTag"`
	InstanceView            VirtualMachineExtensionInstanceViewResponsePtrOutput `pulumi:"instanceView"`
	Location                pulumi.StringOutput                                  `pulumi:"location"`
	Name                    pulumi.StringOutput                                  `pulumi:"name"`
	ProtectedSettings       pulumi.AnyOutput                                     `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringOutput                                  `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrOutput                               `pulumi:"publisher"`
	Settings                pulumi.AnyOutput                                     `pulumi:"settings"`
	Tags                    pulumi.StringMapOutput                               `pulumi:"tags"`
	Type                    pulumi.StringOutput                                  `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrOutput                               `pulumi:"typeHandlerVersion"`
}


func NewVirtualMachineExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineExtension
	err := ctx.RegisterResource("azure-native:compute/v20190701:VirtualMachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineExtension, error) {
	var resource VirtualMachineExtension
	err := ctx.ReadResource("azure-native:compute/v20190701:VirtualMachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineExtensionState struct {
}

type VirtualMachineExtensionState struct {
}

func (VirtualMachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineExtensionState)(nil)).Elem()
}

type virtualMachineExtensionArgs struct {
	AutoUpgradeMinorVersion *bool                                `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          *string                              `pulumi:"forceUpdateTag"`
	InstanceView            *VirtualMachineExtensionInstanceView `pulumi:"instanceView"`
	Location                *string                              `pulumi:"location"`
	ProtectedSettings       interface{}                          `pulumi:"protectedSettings"`
	Publisher               *string                              `pulumi:"publisher"`
	ResourceGroupName       string                               `pulumi:"resourceGroupName"`
	Settings                interface{}                          `pulumi:"settings"`
	Tags                    map[string]string                    `pulumi:"tags"`
	Type                    *string                              `pulumi:"type"`
	TypeHandlerVersion      *string                              `pulumi:"typeHandlerVersion"`
	VmExtensionName         *string                              `pulumi:"vmExtensionName"`
	VmName                  string                               `pulumi:"vmName"`
}


type VirtualMachineExtensionArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	ForceUpdateTag          pulumi.StringPtrInput
	InstanceView            VirtualMachineExtensionInstanceViewPtrInput
	Location                pulumi.StringPtrInput
	ProtectedSettings       pulumi.Input
	Publisher               pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Settings                pulumi.Input
	Tags                    pulumi.StringMapInput
	Type                    pulumi.StringPtrInput
	TypeHandlerVersion      pulumi.StringPtrInput
	VmExtensionName         pulumi.StringPtrInput
	VmName                  pulumi.StringInput
}

func (VirtualMachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineExtensionArgs)(nil)).Elem()
}

type VirtualMachineExtensionInput interface {
	pulumi.Input

	ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput
	ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput
}

func (*VirtualMachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtension)(nil)).Elem()
}

func (i *VirtualMachineExtension) ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput {
	return i.ToVirtualMachineExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineExtension) ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineExtensionOutput)
}

type VirtualMachineExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineExtension)(nil)).Elem()
}

func (o VirtualMachineExtensionOutput) ToVirtualMachineExtensionOutput() VirtualMachineExtensionOutput {
	return o
}

func (o VirtualMachineExtensionOutput) ToVirtualMachineExtensionOutputWithContext(ctx context.Context) VirtualMachineExtensionOutput {
	return o
}

func (o VirtualMachineExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionOutput) InstanceView() VirtualMachineExtensionInstanceViewResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) VirtualMachineExtensionInstanceViewResponsePtrOutput {
		return v.InstanceView
	}).(VirtualMachineExtensionInstanceViewResponsePtrOutput)
}

func (o VirtualMachineExtensionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineExtensionOutput{})
}
