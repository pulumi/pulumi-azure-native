


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScriptExecution(ctx *pulumi.Context, args *LookupScriptExecutionArgs, opts ...pulumi.InvokeOption) (*LookupScriptExecutionResult, error) {
	var rv LookupScriptExecutionResult
	err := ctx.Invoke("azure-native:avs/v20210601:getScriptExecution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScriptExecutionArgs struct {
	PrivateCloudName    string `pulumi:"privateCloudName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ScriptExecutionName string `pulumi:"scriptExecutionName"`
}


type LookupScriptExecutionResult struct {
	Errors            []string               `pulumi:"errors"`
	FailureReason     *string                `pulumi:"failureReason"`
	FinishedAt        string                 `pulumi:"finishedAt"`
	HiddenParameters  []interface{}          `pulumi:"hiddenParameters"`
	Id                string                 `pulumi:"id"`
	Information       []string               `pulumi:"information"`
	Name              string                 `pulumi:"name"`
	NamedOutputs      map[string]interface{} `pulumi:"namedOutputs"`
	Output            []string               `pulumi:"output"`
	Parameters        []interface{}          `pulumi:"parameters"`
	ProvisioningState string                 `pulumi:"provisioningState"`
	Retention         *string                `pulumi:"retention"`
	ScriptCmdletId    *string                `pulumi:"scriptCmdletId"`
	StartedAt         string                 `pulumi:"startedAt"`
	SubmittedAt       string                 `pulumi:"submittedAt"`
	Timeout           string                 `pulumi:"timeout"`
	Type              string                 `pulumi:"type"`
	Warnings          []string               `pulumi:"warnings"`
}

func LookupScriptExecutionOutput(ctx *pulumi.Context, args LookupScriptExecutionOutputArgs, opts ...pulumi.InvokeOption) LookupScriptExecutionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScriptExecutionResult, error) {
			args := v.(LookupScriptExecutionArgs)
			r, err := LookupScriptExecution(ctx, &args, opts...)
			var s LookupScriptExecutionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScriptExecutionResultOutput)
}

type LookupScriptExecutionOutputArgs struct {
	PrivateCloudName    pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptExecutionName pulumi.StringInput `pulumi:"scriptExecutionName"`
}

func (LookupScriptExecutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptExecutionArgs)(nil)).Elem()
}


type LookupScriptExecutionResultOutput struct{ *pulumi.OutputState }

func (LookupScriptExecutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScriptExecutionResult)(nil)).Elem()
}

func (o LookupScriptExecutionResultOutput) ToLookupScriptExecutionResultOutput() LookupScriptExecutionResultOutput {
	return o
}

func (o LookupScriptExecutionResultOutput) ToLookupScriptExecutionResultOutputWithContext(ctx context.Context) LookupScriptExecutionResultOutput {
	return o
}

func (o LookupScriptExecutionResultOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Errors }).(pulumi.StringArrayOutput)
}

func (o LookupScriptExecutionResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

func (o LookupScriptExecutionResultOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.FinishedAt }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []interface{} { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

func (o LookupScriptExecutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Information }).(pulumi.StringArrayOutput)
}

func (o LookupScriptExecutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) map[string]interface{} { return v.NamedOutputs }).(pulumi.MapOutput)
}

func (o LookupScriptExecutionResultOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Output }).(pulumi.StringArrayOutput)
}

func (o LookupScriptExecutionResultOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []interface{} { return v.Parameters }).(pulumi.ArrayOutput)
}

func (o LookupScriptExecutionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.Retention }).(pulumi.StringPtrOutput)
}

func (o LookupScriptExecutionResultOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) *string { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

func (o LookupScriptExecutionResultOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.StartedAt }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.SubmittedAt }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Timeout }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupScriptExecutionResultOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScriptExecutionResult) []string { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScriptExecutionResultOutput{})
}
