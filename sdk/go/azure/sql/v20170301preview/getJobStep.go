


package v20170301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupJobStep(ctx *pulumi.Context, args *LookupJobStepArgs, opts ...pulumi.InvokeOption) (*LookupJobStepResult, error) {
	var rv LookupJobStepResult
	err := ctx.Invoke("azure-native:sql/v20170301preview:getJobStep", args, &rv, opts...)
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
