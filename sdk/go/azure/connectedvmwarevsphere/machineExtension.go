


package connectedvmwarevsphere

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MachineExtension struct {
	pulumi.CustomResourceState

	AutoUpgradeMinorVersion pulumi.BoolPtrOutput                                    `pulumi:"autoUpgradeMinorVersion"`
	ForceUpdateTag          pulumi.StringPtrOutput                                  `pulumi:"forceUpdateTag"`
	InstanceView            MachineExtensionPropertiesResponseInstanceViewPtrOutput `pulumi:"instanceView"`
	Location                pulumi.StringPtrOutput                                  `pulumi:"location"`
	Name                    pulumi.StringOutput                                     `pulumi:"name"`
	ProtectedSettings       pulumi.AnyOutput                                        `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringOutput                                     `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrOutput                                  `pulumi:"publisher"`
	Settings                pulumi.AnyOutput                                        `pulumi:"settings"`
	SystemData              SystemDataResponseOutput                                `pulumi:"systemData"`
	Tags                    pulumi.StringMapOutput                                  `pulumi:"tags"`
	Type                    pulumi.StringOutput                                     `pulumi:"type"`
	TypeHandlerVersion      pulumi.StringPtrOutput                                  `pulumi:"typeHandlerVersion"`
}


func NewMachineExtension(ctx *pulumi.Context,
	name string, args *MachineExtensionArgs, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20201001preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:connectedvmwarevsphere/v20220110preview:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineExtension
	err := ctx.RegisterResource("azure-native:connectedvmwarevsphere:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azure-native:connectedvmwarevsphere:MachineExtension", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type machineExtensionState struct {
}

type MachineExtensionState struct {
}

func (MachineExtensionState) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionState)(nil)).Elem()
}

type machineExtensionArgs struct {
	AutoUpgradeMinorVersion *bool             `pulumi:"autoUpgradeMinorVersion"`
	ExtensionName           *string           `pulumi:"extensionName"`
	ForceUpdateTag          *string           `pulumi:"forceUpdateTag"`
	Location                *string           `pulumi:"location"`
	Name                    string            `pulumi:"name"`
	ProtectedSettings       interface{}       `pulumi:"protectedSettings"`
	Publisher               *string           `pulumi:"publisher"`
	ResourceGroupName       string            `pulumi:"resourceGroupName"`
	Settings                interface{}       `pulumi:"settings"`
	Tags                    map[string]string `pulumi:"tags"`
	Type                    *string           `pulumi:"type"`
	TypeHandlerVersion      *string           `pulumi:"typeHandlerVersion"`
}


type MachineExtensionArgs struct {
	AutoUpgradeMinorVersion pulumi.BoolPtrInput
	ExtensionName           pulumi.StringPtrInput
	ForceUpdateTag          pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	Name                    pulumi.StringInput
	ProtectedSettings       pulumi.Input
	Publisher               pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	Settings                pulumi.Input
	Tags                    pulumi.StringMapInput
	Type                    pulumi.StringPtrInput
	TypeHandlerVersion      pulumi.StringPtrInput
}

func (MachineExtensionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*machineExtensionArgs)(nil)).Elem()
}

type MachineExtensionInput interface {
	pulumi.Input

	ToMachineExtensionOutput() MachineExtensionOutput
	ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput
}

func (*MachineExtension) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (i *MachineExtension) ToMachineExtensionOutput() MachineExtensionOutput {
	return i.ToMachineExtensionOutputWithContext(context.Background())
}

func (i *MachineExtension) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionOutput)
}

type MachineExtensionOutput struct{ *pulumi.OutputState }

func (MachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtension)(nil)).Elem()
}

func (o MachineExtensionOutput) ToMachineExtensionOutput() MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) AutoUpgradeMinorVersion() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.BoolPtrOutput { return v.AutoUpgradeMinorVersion }).(pulumi.BoolPtrOutput)
}

func (o MachineExtensionOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionOutput) InstanceView() MachineExtensionPropertiesResponseInstanceViewPtrOutput {
	return o.ApplyT(func(v *MachineExtension) MachineExtensionPropertiesResponseInstanceViewPtrOutput {
		return v.InstanceView
	}).(MachineExtensionPropertiesResponseInstanceViewPtrOutput)
}

func (o MachineExtensionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MachineExtensionOutput) ProtectedSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.AnyOutput { return v.ProtectedSettings }).(pulumi.AnyOutput)
}

func (o MachineExtensionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o MachineExtensionOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o MachineExtensionOutput) Settings() pulumi.AnyOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.AnyOutput { return v.Settings }).(pulumi.AnyOutput)
}

func (o MachineExtensionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MachineExtension) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MachineExtensionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o MachineExtensionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o MachineExtensionOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtension) pulumi.StringPtrOutput { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MachineExtensionOutput{})
}
