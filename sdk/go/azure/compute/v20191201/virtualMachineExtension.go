


package v20191201

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
			Type: pulumi.String("azure-native:compute/v20190701:VirtualMachineExtension"),
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
	})
	opts = append(opts, aliases)
	var resource VirtualMachineExtension
	err := ctx.RegisterResource("azure-native:compute/v20191201:VirtualMachineExtension", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineExtension(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineExtensionState, opts ...pulumi.ResourceOption) (*VirtualMachineExtension, error) {
	var resource VirtualMachineExtension
	err := ctx.ReadResource("azure-native:compute/v20191201:VirtualMachineExtension", name, id, state, &resource, opts...)
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

func init() {
	pulumi.RegisterOutputType(VirtualMachineExtensionOutput{})
}
