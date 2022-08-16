


package v20180201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBuildStepBuildArguments(ctx *pulumi.Context, args *ListBuildStepBuildArgumentsArgs, opts ...pulumi.InvokeOption) (*ListBuildStepBuildArgumentsResult, error) {
	var rv ListBuildStepBuildArgumentsResult
	err := ctx.Invoke("azure-native:containerregistry/v20180201preview:listBuildStepBuildArguments", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBuildStepBuildArgumentsArgs struct {
	BuildTaskName     string `pulumi:"buildTaskName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	StepName          string `pulumi:"stepName"`
}


type ListBuildStepBuildArgumentsResult struct {
	NextLink *string                 `pulumi:"nextLink"`
	Value    []BuildArgumentResponse `pulumi:"value"`
}

func ListBuildStepBuildArgumentsOutput(ctx *pulumi.Context, args ListBuildStepBuildArgumentsOutputArgs, opts ...pulumi.InvokeOption) ListBuildStepBuildArgumentsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBuildStepBuildArgumentsResult, error) {
			args := v.(ListBuildStepBuildArgumentsArgs)
			r, err := ListBuildStepBuildArguments(ctx, &args, opts...)
			var s ListBuildStepBuildArgumentsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBuildStepBuildArgumentsResultOutput)
}

type ListBuildStepBuildArgumentsOutputArgs struct {
	BuildTaskName     pulumi.StringInput `pulumi:"buildTaskName"`
	RegistryName      pulumi.StringInput `pulumi:"registryName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	StepName          pulumi.StringInput `pulumi:"stepName"`
}

func (ListBuildStepBuildArgumentsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildStepBuildArgumentsArgs)(nil)).Elem()
}


type ListBuildStepBuildArgumentsResultOutput struct{ *pulumi.OutputState }

func (ListBuildStepBuildArgumentsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBuildStepBuildArgumentsResult)(nil)).Elem()
}

func (o ListBuildStepBuildArgumentsResultOutput) ToListBuildStepBuildArgumentsResultOutput() ListBuildStepBuildArgumentsResultOutput {
	return o
}

func (o ListBuildStepBuildArgumentsResultOutput) ToListBuildStepBuildArgumentsResultOutputWithContext(ctx context.Context) ListBuildStepBuildArgumentsResultOutput {
	return o
}

func (o ListBuildStepBuildArgumentsResultOutput) NextLink() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListBuildStepBuildArgumentsResult) *string { return v.NextLink }).(pulumi.StringPtrOutput)
}

func (o ListBuildStepBuildArgumentsResultOutput) Value() BuildArgumentResponseArrayOutput {
	return o.ApplyT(func(v ListBuildStepBuildArgumentsResult) []BuildArgumentResponse { return v.Value }).(BuildArgumentResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBuildStepBuildArgumentsResultOutput{})
}
