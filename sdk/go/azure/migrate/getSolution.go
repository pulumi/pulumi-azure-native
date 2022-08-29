


package migrate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSolution(ctx *pulumi.Context, args *LookupSolutionArgs, opts ...pulumi.InvokeOption) (*LookupSolutionResult, error) {
	var rv LookupSolutionResult
	err := ctx.Invoke("azure-native:migrate:getSolution", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSolutionArgs struct {
	MigrateProjectName string `pulumi:"migrateProjectName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SolutionName       string `pulumi:"solutionName"`
}


type LookupSolutionResult struct {
	Etag       *string                    `pulumi:"etag"`
	Id         string                     `pulumi:"id"`
	Name       string                     `pulumi:"name"`
	Properties SolutionPropertiesResponse `pulumi:"properties"`
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
	MigrateProjectName pulumi.StringInput `pulumi:"migrateProjectName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	SolutionName       pulumi.StringInput `pulumi:"solutionName"`
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

func (o LookupSolutionResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSolutionResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupSolutionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSolutionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSolutionResultOutput) Properties() SolutionPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSolutionResult) SolutionPropertiesResponse { return v.Properties }).(SolutionPropertiesResponseOutput)
}

func (o LookupSolutionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSolutionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSolutionResultOutput{})
}
