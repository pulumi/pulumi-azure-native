


package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	var rv LookupSolutionResult
	err := ctx.Invoke("azure-native:operationsmanagement/v20151101preview:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SolutionName      string `pulumi:"solutionName"`
}


type LookupSolutionResult struct {
	Id         string                     `pulumi:"id"`
	Location   *string                    `pulumi:"location"`
	Name       string                     `pulumi:"name"`
	Plan       *SolutionPlanResponse      `pulumi:"plan"`
	Properties SolutionPropertiesResponse `pulumi:"properties"`
	Tags       map[string]string          `pulumi:"tags"`
	Type       string                     `pulumi:"type"`
}

func LookupSolutionOutput(ctx *pulumi.Context, args LookupSolutionOutputArgs, opts ...pulumi.InvokeOption) LookupSolutionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSolutionResult, error) {
			args := v.(LookupSolutionArgs)
			r, err := LookupSolution(ctx, &args, opts...)
			var s LookupSolutionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSolutionResultOutput)
}

type LookupSolutionOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SolutionName      pulumi.StringInput `pulumi:"solutionName"`
}

func (LookupSolutionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSolutionArgs)(nil)).Elem()
}


type LookupSolutionResultOutput struct{ *pulumi.OutputState }

func (LookupSolutionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSolutionResult)(nil)).Elem()
}

func (o LookupSolutionResultOutput) ToLookupSolutionResultOutput() LookupSolutionResultOutput {
	return o
}

func (o LookupSolutionResultOutput) ToLookupSolutionResultOutputWithContext(ctx context.Context) LookupSolutionResultOutput {
	return o
}

func (o LookupSolutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSolutionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSolutionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSolutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSolutionResultOutput) Plan() SolutionPlanResponsePtrOutput {
	return o.ApplyT(func(v LookupSolutionResult) *SolutionPlanResponse { return v.Plan }).(SolutionPlanResponsePtrOutput)
}

func (o LookupSolutionResultOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSolutionResult) SolutionPropertiesResponse { return v.Properties }).(SolutionPropertiesResponseOutput)
}

func (o LookupSolutionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSolutionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSolutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSolutionResultOutput{})
}
