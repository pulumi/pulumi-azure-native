


package v20191101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStep(ctx *pulumi.Context, args *LookupStepArgs, opts ...pulumi.InvokeOption) (*LookupStepResult, error) {
	var rv LookupStepResult
	err := ctx.Invoke("azure-native:deploymentmanager/v20191101preview:getStep", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupStepArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StepName          string `pulumi:"stepName"`
}


type LookupStepResult struct {
	Id         string            `pulumi:"id"`
	Location   string            `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
