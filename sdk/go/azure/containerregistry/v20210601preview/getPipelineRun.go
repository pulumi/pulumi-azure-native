


package v20210601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipelineRun(ctx *pulumi.Context, args *LookupPipelineRunArgs, opts ...pulumi.InvokeOption) (*LookupPipelineRunResult, error) {
	var rv LookupPipelineRunResult
	err := ctx.Invoke("azure-native:containerregistry/v20210601preview:getPipelineRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineRunArgs struct {
	PipelineRunName   string `pulumi:"pipelineRunName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPipelineRunResult struct {
	ForceUpdateTag    *string                     `pulumi:"forceUpdateTag"`
	Id                string                      `pulumi:"id"`
	Name              string                      `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	Request           *PipelineRunRequestResponse `pulumi:"request"`
	Response          PipelineRunResponseResponse `pulumi:"response"`
	SystemData        SystemDataResponse          `pulumi:"systemData"`
	Type              string                      `pulumi:"type"`
}
