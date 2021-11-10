


package hybridcompute

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
	Location                pulumi.StringOutput                                     `pulumi:"location"`
	Name                    pulumi.StringOutput                                     `pulumi:"name"`
	ProtectedSettings       pulumi.AnyOutput                                        `pulumi:"protectedSettings"`
	ProvisioningState       pulumi.StringOutput                                     `pulumi:"provisioningState"`
	Publisher               pulumi.StringPtrOutput                                  `pulumi:"publisher"`
	Settings                pulumi.AnyOutput                                        `pulumi:"settings"`
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
			Type: pulumi.String("azure-native:hybridcompute/v20190802preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20191212:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200730preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200802:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:MachineExtension"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:MachineExtension"),
		},
	})
	opts = append(opts, aliases)
	var resource MachineExtension
	err := ctx.RegisterResource("azure-native:hybridcompute:MachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MachineExtensionState, opts ...pulumi.ResourceOption) (*MachineExtension, error) {
	var resource MachineExtension
	err := ctx.ReadResource("azure-native:hybridcompute:MachineExtension", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*MachineExtension)(nil))
}

func (i *MachineExtension) ToMachineExtensionOutput() MachineExtensionOutput {
	return i.ToMachineExtensionOutputWithContext(context.Background())
}

func (i *MachineExtension) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionOutput)
}

type MachineExtensionOutput struct{ *pulumi.OutputState }

func (MachineExtensionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtension)(nil))
}

func (o MachineExtensionOutput) ToMachineExtensionOutput() MachineExtensionOutput {
	return o
}

func (o MachineExtensionOutput) ToMachineExtensionOutputWithContext(ctx context.Context) MachineExtensionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MachineExtensionOutput{})
}
