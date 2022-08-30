


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineRunCommandByVirtualMachine struct {
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


func NewVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context,
	name string, args *VirtualMachineRunCommandByVirtualMachineArgs, opts ...pulumi.ResourceOption) (*VirtualMachineRunCommandByVirtualMachine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VmName == nil {
		return nil, errors.New("invalid value for required argument 'VmName'")
	}
	if isZero(args.AsyncExecution) {
		args.AsyncExecution = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:VirtualMachineRunCommandByVirtualMachine"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:VirtualMachineRunCommandByVirtualMachine"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineRunCommandByVirtualMachine
	err := ctx.RegisterResource("azure-native:compute/v20211101:VirtualMachineRunCommandByVirtualMachine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineRunCommandByVirtualMachine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineRunCommandByVirtualMachineState, opts ...pulumi.ResourceOption) (*VirtualMachineRunCommandByVirtualMachine, error) {
	var resource VirtualMachineRunCommandByVirtualMachine
	err := ctx.ReadResource("azure-native:compute/v20211101:VirtualMachineRunCommandByVirtualMachine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineRunCommandByVirtualMachineState struct {
}

type VirtualMachineRunCommandByVirtualMachineState struct {
}

func (VirtualMachineRunCommandByVirtualMachineState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineRunCommandByVirtualMachineState)(nil)).Elem()
}

type virtualMachineRunCommandByVirtualMachineArgs struct {
	AsyncExecution      *bool                                 `pulumi:"asyncExecution"`
	ErrorBlobUri        *string                               `pulumi:"errorBlobUri"`
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
	VmName              string                                `pulumi:"vmName"`
}


type VirtualMachineRunCommandByVirtualMachineArgs struct {
	AsyncExecution      pulumi.BoolPtrInput
	ErrorBlobUri        pulumi.StringPtrInput
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
	VmName              pulumi.StringInput
}

func (VirtualMachineRunCommandByVirtualMachineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineRunCommandByVirtualMachineArgs)(nil)).Elem()
}

type VirtualMachineRunCommandByVirtualMachineInput interface {
	pulumi.Input

	ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput
	ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput
}

func (*VirtualMachineRunCommandByVirtualMachine) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineRunCommandByVirtualMachine)(nil)).Elem()
}

func (i *VirtualMachineRunCommandByVirtualMachine) ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput {
	return i.ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(context.Background())
}

func (i *VirtualMachineRunCommandByVirtualMachine) ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineRunCommandByVirtualMachineOutput)
}

type VirtualMachineRunCommandByVirtualMachineOutput struct{ *pulumi.OutputState }

func (VirtualMachineRunCommandByVirtualMachineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualMachineRunCommandByVirtualMachine)(nil)).Elem()
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ToVirtualMachineRunCommandByVirtualMachineOutput() VirtualMachineRunCommandByVirtualMachineOutput {
	return o
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ToVirtualMachineRunCommandByVirtualMachineOutputWithContext(ctx context.Context) VirtualMachineRunCommandByVirtualMachineOutput {
	return o
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) AsyncExecution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.BoolPtrOutput { return v.AsyncExecution }).(pulumi.BoolPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ErrorBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringPtrOutput { return v.ErrorBlobUri }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) InstanceView() VirtualMachineRunCommandInstanceViewResponseOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) VirtualMachineRunCommandInstanceViewResponseOutput {
		return v.InstanceView
	}).(VirtualMachineRunCommandInstanceViewResponseOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) OutputBlobUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringPtrOutput { return v.OutputBlobUri }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Parameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) RunCommandInputParameterResponseArrayOutput {
		return v.Parameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ProtectedParameters() RunCommandInputParameterResponseArrayOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) RunCommandInputParameterResponseArrayOutput {
		return v.ProtectedParameters
	}).(RunCommandInputParameterResponseArrayOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) RunAsPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringPtrOutput { return v.RunAsPassword }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) RunAsUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringPtrOutput { return v.RunAsUser }).(pulumi.StringPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Source() VirtualMachineRunCommandScriptSourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) VirtualMachineRunCommandScriptSourceResponsePtrOutput {
		return v.Source
	}).(VirtualMachineRunCommandScriptSourceResponsePtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.IntPtrOutput { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

func (o VirtualMachineRunCommandByVirtualMachineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualMachineRunCommandByVirtualMachine) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineRunCommandByVirtualMachineOutput{})
}
