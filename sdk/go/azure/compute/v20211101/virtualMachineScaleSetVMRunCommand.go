


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineScaleSetVMRunCommand struct {
	pulumi.CustomResourceState

	AsyncExecution      pulumi.BoolPtrOutput                                  `pulumi:"asyncExecution"`
	ErrorBlobUri        pulumi.StringPtrOutput                                `pulumi:"errorBlobUri"`
	InstanceView        VirtualMachineRunCommandInstanceViewResponseOutput    `pulumi:"instanceView"`
	Location            pulumi.StringOutput                                   `pulumi:"location"`
	Name                pulumi.StringOutput                                   `pulumi:"name"`
	OutputBlobUri       pulumi.StringPtrOutput                                `pulumi:"outputBlobUri"`
	Parameters          RunCommandInputParameterResponseArrayOutput           `pulumi:"parameters"`
	ProtectedParameters RunCommandInputParameterResponseArrayOutput           `pulumi:"protectedParameters"`
	ProvisioningState   pulumi.StringOutput                                   `pulumi:"provisioningState"`
	RunAsPassword       pulumi.StringPtrOutput                                `pulumi:"runAsPassword"`
	RunAsUser           pulumi.StringPtrOutput                                `pulumi:"runAsUser"`
	Source              VirtualMachineRunCommandScriptSourceResponsePtrOutput `pulumi:"source"`
	Tags                pulumi.StringMapOutput                                `pulumi:"tags"`
	TimeoutInSeconds    pulumi.IntPtrOutput                                   `pulumi:"timeoutInSeconds"`
	Type                pulumi.StringOutput                                   `pulumi:"type"`
}


func NewVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetVMRunCommandArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMRunCommand, error) {
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
	if isZero(args.AsyncExecution) {
		args.AsyncExecution = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineScaleSetVMRunCommand"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineScaleSetVMRunCommand"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineScaleSetVMRunCommand
	err := ctx.RegisterResource("azure-native:compute/v20211101:VirtualMachineScaleSetVMRunCommand", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineScaleSetVMRunCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetVMRunCommandState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSetVMRunCommand, error) {
	var resource VirtualMachineScaleSetVMRunCommand
	err := ctx.ReadResource("azure-native:compute/v20211101:VirtualMachineScaleSetVMRunCommand", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScaleSetVMRunCommandState struct {
}

type VirtualMachineScaleSetVMRunCommandState struct {
}

func (VirtualMachineScaleSetVMRunCommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMRunCommandState)(nil)).Elem()
}

type virtualMachineScaleSetVMRunCommandArgs struct {
	AsyncExecution      *bool                                 `pulumi:"asyncExecution"`
	ErrorBlobUri        *string                               `pulumi:"errorBlobUri"`
	InstanceId          string                                `pulumi:"instanceId"`
	Location            *string                               `pulumi:"location"`
	OutputBlobUri       *string                               `pulumi:"outputBlobUri"`
	Parameters          []RunCommandInputParameter            `pulumi:"parameters"`
	ProtectedParameters []RunCommandInputParameter            `pulumi:"protectedParameters"`
	ResourceGroupName   string                                `pulumi:"resourceGroupName"`
	RunAsPassword       *string                               `pulumi:"runAsPassword"`
	RunAsUser           *string                               `pulumi:"runAsUser"`
	RunCommandName      *string                               `pulumi:"runCommandName"`
	Source              *VirtualMachineRunCommandScriptSource `pulumi:"source"`
	Tags                map[string]string                     `pulumi:"tags"`
	TimeoutInSeconds    *int                                  `pulumi:"timeoutInSeconds"`
	VmScaleSetName      string                                `pulumi:"vmScaleSetName"`
}


type VirtualMachineScaleSetVMRunCommandArgs struct {
	AsyncExecution      pulumi.BoolPtrInput
	ErrorBlobUri        pulumi.StringPtrInput
	InstanceId          pulumi.StringInput
	Location            pulumi.StringPtrInput
	OutputBlobUri       pulumi.StringPtrInput
	Parameters          RunCommandInputParameterArrayInput
	ProtectedParameters RunCommandInputParameterArrayInput
	ResourceGroupName   pulumi.StringInput
	RunAsPassword       pulumi.StringPtrInput
	RunAsUser           pulumi.StringPtrInput
	RunCommandName      pulumi.StringPtrInput
	Source              VirtualMachineRunCommandScriptSourcePtrInput
	Tags                pulumi.StringMapInput
	TimeoutInSeconds    pulumi.IntPtrInput
	VmScaleSetName      pulumi.StringInput
}

func (VirtualMachineScaleSetVMRunCommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetVMRunCommandArgs)(nil)).Elem()
}

type VirtualMachineScaleSetVMRunCommandInput interface {
	pulumi.Input

	ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput
	ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput
}

func (*VirtualMachineScaleSetVMRunCommand) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMRunCommand)(nil)).Elem()
}

func (i *VirtualMachineScaleSetVMRunCommand) ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput {
	return i.ToVirtualMachineScaleSetVMRunCommandOutputWithContext(context.Background())
}

func (i *VirtualMachineScaleSetVMRunCommand) ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScaleSetVMRunCommandOutput)
}

type VirtualMachineScaleSetVMRunCommandOutput struct{ *pulumi.OutputState }

func (VirtualMachineScaleSetVMRunCommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineScaleSetVMRunCommand)(nil)).Elem()
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ToVirtualMachineScaleSetVMRunCommandOutput() VirtualMachineScaleSetVMRunCommandOutput {
	return o
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ToVirtualMachineScaleSetVMRunCommandOutputWithContext(ctx context.Context) VirtualMachineScaleSetVMRunCommandOutput {
	return o
}

func (o VirtualMachineScaleSetVMRunCommandOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.BoolPtrOutput { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) VirtualMachineRunCommandInstanceViewResponseOutput {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandInputParameterResponseArrayOutput {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) RunCommandInputParameterResponseArrayOutput {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringPtrOutput { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) VirtualMachineRunCommandScriptSourceResponsePtrOutput {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.IntPtrOutput { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineScaleSetVMRunCommandOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineScaleSetVMRunCommand) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScaleSetVMRunCommandOutput{})
}
