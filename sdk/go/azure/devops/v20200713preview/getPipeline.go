


package v20200713preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("azure-native:devops/v20200713preview:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	PipelineName      string `pulumi:"pipelineName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPipelineResult struct {
	BootstrapConfiguration BootstrapConfigurationResponse `pulumi:"bootstrapConfiguration"`
	Id                     string                         `pulumi:"id"`
	Location               *string                        `pulumi:"location"`
	Name                   string                         `pulumi:"name"`
	PipelineId             int                            `pulumi:"pipelineId"`
	PipelineType           string                         `pulumi:"pipelineType"`
	SystemData             SystemDataResponse             `pulumi:"systemData"`
	Tags                   map[string]string              `pulumi:"tags"`
	Type                   string                         `pulumi:"type"`
}
