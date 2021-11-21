


package v20211201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScriptExecution(ctx *pulumi.Context, args *LookupScriptExecutionArgs, opts ...pulumi.InvokeOption) (*LookupScriptExecutionResult, error) {
	var rv LookupScriptExecutionResult
	err := ctx.Invoke("azure-native:avs/v20211201:getScriptExecution", args, &rv, opts...)
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
