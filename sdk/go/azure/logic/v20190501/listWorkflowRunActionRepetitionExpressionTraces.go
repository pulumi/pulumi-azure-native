


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowRunActionRepetitionExpressionTraces(ctx *pulumi.Context, args *ListWorkflowRunActionRepetitionExpressionTracesArgs, opts ...pulumi.InvokeOption) (*ListWorkflowRunActionRepetitionExpressionTracesResult, error) {
	var rv ListWorkflowRunActionRepetitionExpressionTracesResult
	err := ctx.Invoke("azure-native:logic/v20190501:listWorkflowRunActionRepetitionExpressionTraces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowRunActionRepetitionExpressionTracesArgs struct {
	ActionName        string `pulumi:"actionName"`
	RepetitionName    string `pulumi:"repetitionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RunName           string `pulumi:"runName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type ListWorkflowRunActionRepetitionExpressionTracesResult struct {
	Inputs []ExpressionRootResponse `pulumi:"inputs"`
}

func ListWorkflowRunActionRepetitionExpressionTracesOutput(ctx *pulumi.Context, args ListWorkflowRunActionRepetitionExpressionTracesOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowRunActionRepetitionExpressionTracesResult, error) {
			args := v.(ListWorkflowRunActionRepetitionExpressionTracesArgs)
			r, err := ListWorkflowRunActionRepetitionExpressionTraces(ctx, &args, opts...)
			var s ListWorkflowRunActionRepetitionExpressionTracesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowRunActionRepetitionExpressionTracesResultOutput)
}

type ListWorkflowRunActionRepetitionExpressionTracesOutputArgs struct {
	ActionName        pulumi.StringInput `pulumi:"actionName"`
	RepetitionName    pulumi.StringInput `pulumi:"repetitionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RunName           pulumi.StringInput `pulumi:"runName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowRunActionRepetitionExpressionTracesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesArgs)(nil)).Elem()
}


type ListWorkflowRunActionRepetitionExpressionTracesResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionRepetitionExpressionTracesResult)(nil)).Elem()
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutput() ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) ToListWorkflowRunActionRepetitionExpressionTracesResultOutputWithContext(ctx context.Context) ListWorkflowRunActionRepetitionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionRepetitionExpressionTracesResultOutput) Inputs() ExpressionRootResponseArrayOutput {
	return o.ApplyT(func(v ListWorkflowRunActionRepetitionExpressionTracesResult) []ExpressionRootResponse {
		return v.Inputs
	}).(ExpressionRootResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowRunActionRepetitionExpressionTracesResultOutput{})
}
