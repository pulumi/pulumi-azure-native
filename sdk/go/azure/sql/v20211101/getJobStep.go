


package v20211101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobStep(ctx *pulumi.Context, args *LookupJobStepArgs, opts ...pulumi.InvokeOption) (*LookupJobStepResult, error) {
	var rv LookupJobStepResult
	err := ctx.Invoke("azure-native:sql/v20211101:getJobStep", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupJobStepArgs struct {
	JobAgentName      string `pulumi:"jobAgentName"`
	JobName           string `pulumi:"jobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
	StepName          string `pulumi:"stepName"`
}


type LookupJobStepResult struct {
	Action           JobStepActionResponse            `pulumi:"action"`
	Credential       string                           `pulumi:"credential"`
	ExecutionOptions *JobStepExecutionOptionsResponse `pulumi:"executionOptions"`
	Id               string                           `pulumi:"id"`
	Name             string                           `pulumi:"name"`
	Output           *JobStepOutputResponse           `pulumi:"output"`
	StepId           *int                             `pulumi:"stepId"`
	TargetGroup      string                           `pulumi:"targetGroup"`
	Type             string                           `pulumi:"type"`
}


func (val *LookupJobStepResult) Defaults() *LookupJobStepResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Action = *tmp.Action.Defaults()

	tmp.ExecutionOptions = tmp.ExecutionOptions.Defaults()

	tmp.Output = tmp.Output.Defaults()

	return &tmp
}

func LookupJobStepOutput(ctx *pulumi.Context, args LookupJobStepOutputArgs, opts ...pulumi.InvokeOption) LookupJobStepResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupJobStepResult, error) {
			args := v.(LookupJobStepArgs)
			r, err := LookupJobStep(ctx, &args, opts...)
			var s LookupJobStepResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupJobStepResultOutput)
}

type LookupJobStepOutputArgs struct {
	JobAgentName      pulumi.StringInput `pulumi:"jobAgentName"`
	JobName           pulumi.StringInput `pulumi:"jobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
	StepName          pulumi.StringInput `pulumi:"stepName"`
}

func (LookupJobStepOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobStepArgs)(nil)).Elem()
}


type LookupJobStepResultOutput struct{ *pulumi.OutputState }

func (LookupJobStepResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupJobStepResult)(nil)).Elem()
}

func (o LookupJobStepResultOutput) ToLookupJobStepResultOutput() LookupJobStepResultOutput {
	return o
}

func (o LookupJobStepResultOutput) ToLookupJobStepResultOutputWithContext(ctx context.Context) LookupJobStepResultOutput {
	return o
}

func (o LookupJobStepResultOutput) Action() JobStepActionResponseOutput {
	return o.ApplyT(func(v LookupJobStepResult) JobStepActionResponse { return v.Action }).(JobStepActionResponseOutput)
}

func (o LookupJobStepResultOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobStepResult) string { return v.Credential }).(pulumi.StringOutput)
}

func (o LookupJobStepResultOutput) ExecutionOptions() JobStepExecutionOptionsResponsePtrOutput {
	return o.ApplyT(func(v LookupJobStepResult) *JobStepExecutionOptionsResponse { return v.ExecutionOptions }).(JobStepExecutionOptionsResponsePtrOutput)
}

func (o LookupJobStepResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobStepResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupJobStepResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobStepResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupJobStepResultOutput) Output() JobStepOutputResponsePtrOutput {
	return o.ApplyT(func(v LookupJobStepResult) *JobStepOutputResponse { return v.Output }).(JobStepOutputResponsePtrOutput)
}

func (o LookupJobStepResultOutput) StepId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupJobStepResult) *int { return v.StepId }).(pulumi.IntPtrOutput)
}

func (o LookupJobStepResultOutput) TargetGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobStepResult) string { return v.TargetGroup }).(pulumi.StringOutput)
}

func (o LookupJobStepResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupJobStepResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupJobStepResultOutput{})
}
