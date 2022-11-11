


package avs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetScriptExecutionLogs(ctx *pulumi.Context, args *GetScriptExecutionLogsArgs, opts ...pulumi.InvokeOption) (*GetScriptExecutionLogsResult, error) {
	var rv GetScriptExecutionLogsResult
	err := ctx.Invoke("azure-native:avs:getScriptExecutionLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetScriptExecutionLogsArgs struct {
	PrivateCloudName    string `pulumi:"privateCloudName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
	ScriptExecutionName string `pulumi:"scriptExecutionName"`
}


type GetScriptExecutionLogsResult struct {
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

func GetScriptExecutionLogsOutput(ctx *pulumi.Context, args GetScriptExecutionLogsOutputArgs, opts ...pulumi.InvokeOption) GetScriptExecutionLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetScriptExecutionLogsResult, error) {
			args := v.(GetScriptExecutionLogsArgs)
			r, err := GetScriptExecutionLogs(ctx, &args, opts...)
			var s GetScriptExecutionLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetScriptExecutionLogsResultOutput)
}

type GetScriptExecutionLogsOutputArgs struct {
	PrivateCloudName    pulumi.StringInput `pulumi:"privateCloudName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
	ScriptExecutionName pulumi.StringInput `pulumi:"scriptExecutionName"`
}

func (GetScriptExecutionLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScriptExecutionLogsArgs)(nil)).Elem()
}


type GetScriptExecutionLogsResultOutput struct{ *pulumi.OutputState }

func (GetScriptExecutionLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetScriptExecutionLogsResult)(nil)).Elem()
}

func (o GetScriptExecutionLogsResultOutput) ToGetScriptExecutionLogsResultOutput() GetScriptExecutionLogsResultOutput {
	return o
}

func (o GetScriptExecutionLogsResultOutput) ToGetScriptExecutionLogsResultOutputWithContext(ctx context.Context) GetScriptExecutionLogsResultOutput {
	return o
}

func (o GetScriptExecutionLogsResultOutput) Errors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Errors }).(pulumi.StringArrayOutput)
}

func (o GetScriptExecutionLogsResultOutput) FailureReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.FailureReason }).(pulumi.StringPtrOutput)
}

func (o GetScriptExecutionLogsResultOutput) FinishedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.FinishedAt }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) HiddenParameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []interface{} { return v.HiddenParameters }).(pulumi.ArrayOutput)
}

func (o GetScriptExecutionLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) Information() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Information }).(pulumi.StringArrayOutput)
}

func (o GetScriptExecutionLogsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) NamedOutputs() pulumi.MapOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) map[string]interface{} { return v.NamedOutputs }).(pulumi.MapOutput)
}

func (o GetScriptExecutionLogsResultOutput) Output() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Output }).(pulumi.StringArrayOutput)
}

func (o GetScriptExecutionLogsResultOutput) Parameters() pulumi.ArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []interface{} { return v.Parameters }).(pulumi.ArrayOutput)
}

func (o GetScriptExecutionLogsResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) Retention() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.Retention }).(pulumi.StringPtrOutput)
}

func (o GetScriptExecutionLogsResultOutput) ScriptCmdletId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) *string { return v.ScriptCmdletId }).(pulumi.StringPtrOutput)
}

func (o GetScriptExecutionLogsResultOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.StartedAt }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) SubmittedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.SubmittedAt }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Timeout }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o GetScriptExecutionLogsResultOutput) Warnings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetScriptExecutionLogsResult) []string { return v.Warnings }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetScriptExecutionLogsResultOutput{})
}
