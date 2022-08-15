


package v20180201preview

import (
	"context"
	"reflect"

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

func LookupBuildStepOutput(ctx *pulumi.Context, args LookupBuildStepOutputArgs, opts ...pulumi.InvokeOption) LookupBuildStepResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildStepResult, error) {
			args := v.(LookupBuildStepArgs)
			r, err := LookupBuildStep(ctx, &args, opts...)
			var s LookupBuildStepResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildStepResultOutput)
}

type LookupBuildStepOutputArgs struct {
	BuildTaskName     pulumi.StringInput `pulumi:"buildTaskName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StepName          pulumi.StringInput `pulumi:"stepName"`
}

func (LookupBuildStepOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildStepArgs)(nil)).Elem()
}


type LookupBuildStepResultOutput struct{ *pulumi.OutputState }

func (LookupBuildStepResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildStepResult)(nil)).Elem()
}

func (o LookupBuildStepResultOutput) ToLookupBuildStepResultOutput() LookupBuildStepResultOutput {
	return o
}

func (o LookupBuildStepResultOutput) ToLookupBuildStepResultOutputWithContext(ctx context.Context) LookupBuildStepResultOutput {
	return o
}

func (o LookupBuildStepResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildStepResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBuildStepResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildStepResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBuildStepResultOutput) Properties() DockerBuildStepResponseOutput {
	return o.ApplyT(func(v LookupBuildStepResult) DockerBuildStepResponse { return v.Properties }).(DockerBuildStepResponseOutput)
}

func (o LookupBuildStepResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildStepResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildStepResultOutput{})
}
