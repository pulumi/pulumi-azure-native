


package v20190701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("azure-native:devops/v20190701preview:getPipeline", args, &rv, opts...)
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
	Organization           OrganizationReferenceResponse  `pulumi:"organization"`
	PipelineId             int                            `pulumi:"pipelineId"`
	Project                ProjectReferenceResponse       `pulumi:"project"`
	Tags                   map[string]string              `pulumi:"tags"`
	Type                   string                         `pulumi:"type"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	PipelineName      pulumi.StringInput `pulumi:"pipelineName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}


type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) BootstrapConfiguration() BootstrapConfigurationResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) BootstrapConfigurationResponse { return v.BootstrapConfiguration }).(BootstrapConfigurationResponseOutput)
}

func (o LookupPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Organization() OrganizationReferenceResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) OrganizationReferenceResponse { return v.Organization }).(OrganizationReferenceResponseOutput)
}

func (o LookupPipelineResultOutput) PipelineId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPipelineResult) int { return v.PipelineId }).(pulumi.IntOutput)
}

func (o LookupPipelineResultOutput) Project() ProjectReferenceResponseOutput {
	return o.ApplyT(func(v LookupPipelineResult) ProjectReferenceResponse { return v.Project }).(ProjectReferenceResponseOutput)
}

func (o LookupPipelineResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
