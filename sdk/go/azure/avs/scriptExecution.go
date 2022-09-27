


package avs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ScriptExecution struct {
	pulumi.CustomResourceState

	Errors            pulumi.StringArrayOutput `pulumi:"errors"`
	FailureReason     pulumi.StringPtrOutput   `pulumi:"failureReason"`
	FinishedAt        pulumi.StringOutput      `pulumi:"finishedAt"`
	HiddenParameters  pulumi.ArrayOutput       `pulumi:"hiddenParameters"`
	Information       pulumi.StringArrayOutput `pulumi:"information"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	NamedOutputs      pulumi.MapOutput         `pulumi:"namedOutputs"`
	Output            pulumi.StringArrayOutput `pulumi:"output"`
	Parameters        pulumi.ArrayOutput       `pulumi:"parameters"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	Retention         pulumi.StringPtrOutput   `pulumi:"retention"`
	ScriptCmdletId    pulumi.StringPtrOutput   `pulumi:"scriptCmdletId"`
	StartedAt         pulumi.StringOutput      `pulumi:"startedAt"`
	SubmittedAt       pulumi.StringOutput      `pulumi:"submittedAt"`
	Timeout           pulumi.StringOutput      `pulumi:"timeout"`
	Type              pulumi.StringOutput      `pulumi:"type"`
	Warnings          pulumi.StringArrayOutput `pulumi:"warnings"`
}


func NewScriptExecution(ctx *pulumi.Context,
	name string, args *ScriptExecutionArgs, opts ...pulumi.ResourceOption) (*ScriptExecution, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateCloudName == nil {
		return nil, errors.New("invalid value for required argument 'PrivateCloudName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:avs/v20210601:ScriptExecution"),
		},
		{
			Type: pulumi.String("azure-native:avs/v20211201:ScriptExecution"),
		},
	})
	opts = append(opts, aliases)
	var resource ScriptExecution
	err := ctx.RegisterResource("azure-native:avs:ScriptExecution", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetScriptExecution(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScriptExecutionState, opts ...pulumi.ResourceOption) (*ScriptExecution, error) {
	var resource ScriptExecution
	err := ctx.ReadResource("azure-native:avs:ScriptExecution", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scriptExecutionState struct {
}

type ScriptExecutionState struct {
}

func (ScriptExecutionState) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptExecutionState)(nil)).Elem()
}

type scriptExecutionArgs struct {
	FailureReason       *string                `pulumi:"failureReason"`
	HiddenParameters    []interface{}          `pulumi:"hiddenParameters"`
	NamedOutputs        map[string]interface{} `pulumi:"namedOutputs"`
	Output              []string               `pulumi:"output"`
	Parameters          []interface{}          `pulumi:"parameters"`
	PrivateCloudName    string                 `pulumi:"privateCloudName"`
	ResourceGroupName   string                 `pulumi:"resourceGroupName"`
	Retention           *string                `pulumi:"retention"`
	ScriptCmdletId      *string                `pulumi:"scriptCmdletId"`
	ScriptExecutionName *string                `pulumi:"scriptExecutionName"`
	Timeout             string                 `pulumi:"timeout"`
}


type ScriptExecutionArgs struct {
	FailureReason       pulumi.StringPtrInput
	HiddenParameters    pulumi.ArrayInput
	NamedOutputs        pulumi.MapInput
	Output              pulumi.StringArrayInput
	Parameters          pulumi.ArrayInput
	PrivateCloudName    pulumi.StringInput
	ResourceGroupName   pulumi.StringInput
	Retention           pulumi.StringPtrInput
	ScriptCmdletId      pulumi.StringPtrInput
	ScriptExecutionName pulumi.StringPtrInput
	Timeout             pulumi.StringInput
}

func (ScriptExecutionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scriptExecutionArgs)(nil)).Elem()
}

type ScriptExecutionInput interface {
	pulumi.Input

	ToScriptExecutionOutput() ScriptExecutionOutput
	ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput
}

func (*ScriptExecution) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptExecution)(nil)).Elem()
}

func (i *ScriptExecution) ToScriptExecutionOutput() ScriptExecutionOutput {
	return i.ToScriptExecutionOutputWithContext(context.Background())
}

func (i *ScriptExecution) ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScriptExecutionOutput)
}

type ScriptExecutionOutput struct{ *pulumi.OutputState }

func (ScriptExecutionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScriptExecution)(nil)).Elem()
}

func (o ScriptExecutionOutput) ToScriptExecutionOutput() ScriptExecutionOutput {
	return o
}

func (o ScriptExecutionOutput) ToScriptExecutionOutputWithContext(ctx context.Context) ScriptExecutionOutput {
	return o
}

func (o ScriptExecutionOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Errors }).(pulumi.StringArrayOutput)
}

func (o ScriptExecutionOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.FailureReason }).(pulumi.StringPtrOutput)
}

func (o ScriptExecutionOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.FinishedAt }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.ArrayOutput { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

func (o ScriptExecutionOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Information }).(pulumi.StringArrayOutput)
}

func (o ScriptExecutionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.MapOutput { return v.NamedOutputs }).(pulumi.MapOutput)
}

func (o ScriptExecutionOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Output }).(pulumi.StringArrayOutput)
}

func (o ScriptExecutionOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.ArrayOutput { return v.Parameters }).(pulumi.ArrayOutput)
}

func (o ScriptExecutionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.Retention }).(pulumi.StringPtrOutput)
}

func (o ScriptExecutionOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringPtrOutput { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

func (o ScriptExecutionOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.StartedAt }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.SubmittedAt }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Timeout }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScriptExecutionOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScriptExecution) pulumi.StringArrayOutput { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ScriptExecutionOutput{})
}
