


package v20180201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBuildStep(ctx *pulumi.Context, args *LookupBuildStepArgs, opts ...pulumi.InvokeOption) (*LookupBuildStepResult, error) {
	var rv LookupBuildStepResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:getBuildStep", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBuildStepArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StepName          string `pulumi:"stepName"`
}


type LookupBuildStepResult struct {
	Id         string                  `pulumi:"id"`
	Name       string                  `pulumi:"name"`
	Properties DockerBuildStepResponse `pulumi:"properties"`
	Type       string                  `pulumi:"type"`
}


func (val *LookupBuildStepResult) Defaults() *LookupBuildStepResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}
