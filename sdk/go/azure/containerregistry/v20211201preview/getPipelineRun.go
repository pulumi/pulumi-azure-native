


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipelineRun(ctx *pulumi.Context, args *LookupPipelineRunArgs, opts ...pulumi.InvokeOption) (*LookupPipelineRunResult, error) {
	var rv LookupPipelineRunResult
	err := ctx.Invoke("azure-native:containerregistry/v20211201preview:getPipelineRun", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupPipelineRunResult) Defaults() *LookupPipelineRunResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Request = tmp.Request.Defaults()

	tmp.Response = *tmp.Response.Defaults()

	return &tmp
}

func LookupPipelineRunOutput(ctx *pulumi.Context, args LookupPipelineRunOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineRunResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineRunResult, error) {
			args := v.(LookupPipelineRunArgs)
			r, err := LookupPipelineRun(ctx, &args, opts...)
			var s LookupPipelineRunResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineRunResultOutput)
}

type LookupPipelineRunOutputArgs struct {
	PipelineRunName   pulumi.StringInput `pulumi:"pipelineRunName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineRunOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineRunArgs)(nil)).Elem()
}


type LookupPipelineRunResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineRunResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineRunResult)(nil)).Elem()
}

func (o LookupPipelineRunResultOutput) ToLookupPipelineRunResultOutput() LookupPipelineRunResultOutput {
	return o
}

func (o LookupPipelineRunResultOutput) ToLookupPipelineRunResultOutputWithContext(ctx context.Context) LookupPipelineRunResultOutput {
	return o
}

func (o LookupPipelineRunResultOutput) ForceUpdateTag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) *string { return v.ForceUpdateTag }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineRunResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPipelineRunResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPipelineRunResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPipelineRunResultOutput) Request() PipelineRunRequestResponsePtrOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) *PipelineRunRequestResponse { return v.Request }).(PipelineRunRequestResponsePtrOutput)
}

func (o LookupPipelineRunResultOutput) Response() PipelineRunResponseResponseOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) PipelineRunResponseResponse { return v.Response }).(PipelineRunResponseResponseOutput)
}

func (o LookupPipelineRunResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPipelineRunResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineRunResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineRunResultOutput{})
}
