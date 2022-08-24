


package v20220301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSetExtension struct {
	pulumi.CustomResourceState

	AutoUpgradeMinorVersion       pulumi.BoolPtrOutput     `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        pulumi.BoolPtrOutput     `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                pulumi.StringPtrOutput   `pulumi:"forceUpdateTag"`
	Name                          pulumi.StringPtrOutput   `pulumi:"name"`
	ProtectedSettings             pulumi.AnyOutput         `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault pulumi.AnyOutput         `pulumi:"protectedSettingsFromKeyVault"`
	ProvisionAfterExtensions      pulumi.StringArrayOutput `pulumi:"provisionAfterExtensions"`
	ProvisioningState             pulumi.StringOutput      `pulumi:"provisioningState"`
	Publisher                     pulumi.StringPtrOutput   `pulumi:"publisher"`
	Settings                      pulumi.AnyOutput         `pulumi:"settings"`
	SuppressFailures              pulumi.BoolPtrOutput     `pulumi:"suppressFailures"`
	Type                          pulumi.StringOutput      `pulumi:"type"`
	TypeHandlerVersion            pulumi.StringPtrOutput   `pulumi:"typeHandlerVersion"`
}


func NewVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetExtensionArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
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
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20171201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetExtension"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:VirtualMachineScaleSetExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetExtension
	err := ctx.RegisterResource("azure-native:compute/v20220301:VirtualMachineScaleSetExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSetExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetExtension, error) {
	var resource VirtualMachineScaleSetExtension
	err := ctx.ReadResource("azure-native:compute/v20220301:VirtualMachineScaleSetExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScaleSetExtensionState struct {
}

type VirtualMachineScaleSetExtensionState struct {
}

func (VirtualMachineScaleSetExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionState)(nil)).Elem()
}

type virtualMachineScaleSetExtensionArgs struct {
	AutoUpgradeMinorVersion       *bool       `pulumi:"autoUpgradeMinorVersion"`
	EnableAutomaticUpgrade        *bool       `pulumi:"enableAutomaticUpgrade"`
	ForceUpdateTag                *string     `pulumi:"forceUpdateTag"`
	Name                          *string     `pulumi:"name"`
	ProtectedSettings             interface{} `pulumi:"protectedSettings"`
	ProtectedSettingsFromKeyVault interface{} `pulumi:"protectedSettingsFromKeyVault"`
	ProvisionAfterExtensions      []string    `pulumi:"provisionAfterExtensions"`
	Publisher                     *string     `pulumi:"publisher"`
	ResourceGroupName             string      `pulumi:"resourceGroupName"`
	Settings                      interface{} `pulumi:"settings"`
	SuppressFailures              *bool       `pulumi:"suppressFailures"`
	Type                          *string     `pulumi:"type"`
	TypeHandlerVersion            *string     `pulumi:"typeHandlerVersion"`
	VmScaleSetName                string      `pulumi:"vmScaleSetName"`
	VmssExtensionName             *string     `pulumi:"vmssExtensionName"`
}


type VirtualMachineScaleSetExtensionArgs struct {
	AutoUpgradeMinorVersion       pulumi.BoolPtrInput
	EnableAutomaticUpgrade        pulumi.BoolPtrInput
	ForceUpdateTag                pulumi.StringPtrInput
	Name                          pulumi.StringPtrInput
	ProtectedSettings             pulumi.Input
	ProtectedSettingsFromKeyVault pulumi.Input
	ProvisionAfterExtensions      pulumi.StringArrayInput
	Publisher                     pulumi.StringPtrInput
	ResourceGroupName             pulumi.StringInput
	Settings                      pulumi.Input
	SuppressFailures              pulumi.BoolPtrInput
	Type                          pulumi.StringPtrInput
	TypeHandlerVersion            pulumi.StringPtrInput
	VmScaleSetName                pulumi.StringInput
	VmssExtensionName             pulumi.StringPtrInput
}

func (VirtualMachineScaleSetExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetExtensionArgs)(nil)).Elem()
}

type VirtualMachineScaleSetExtensionInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput
	ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput
}

func (*VirtualMachineScaleSetExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return i.ToVirtualMachineScaleSetExtensionOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetExtension) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetExtensionOutput)
}

type VirtualMachineScaleSetExtensionOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetExtension)(nil)).Elem()
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutput() VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) ToVirtualMachineScaleSetExtensionOutputWithContext(ctx context.Context) VirtualMachineScaleSetExtensionOutput {
	return o
}

func (o VirtualMachineScaleSetExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) EnableAutomaticUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.BoolPtrOutput { return v.EnableAutomaticUpgrade }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ProtectedSettingsFromKeyVault() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.AnyOutput { return v.ProtectedSettingsFromKeyVault }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ProvisionAfterExtensions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringArrayOutput { return v.ProvisionAfterExtensions }).(pulumi.StringArrayOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) SuppressFailures() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.BoolPtrOutput { return v.SuppressFailures }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetExtensionOutput{})
}
