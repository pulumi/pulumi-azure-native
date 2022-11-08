


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowRunActionExpressionTraces(ctx *pulumi.Context, args *ListWorkflowRunActionExpressionTracesArgs, opts ...pulumi.InvokeOption) (*ListWorkflowRunActionExpressionTracesResult, error) {
	var rv ListWorkflowRunActionExpressionTracesResult
	err := ctx.Invoke("azure-native:logic:listWorkflowRunActionExpressionTraces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowRunActionExpressionTracesArgs struct {
	ActionName        string `pulumi:"actionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RunName           string `pulumi:"runName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type ListWorkflowRunActionExpressionTracesResult struct {
	Inputs []ExpressionRootResponse `pulumi:"inputs"`
}

func ListWorkflowRunActionExpressionTracesOutput(ctx *pulumi.Context, args ListWorkflowRunActionExpressionTracesOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowRunActionExpressionTracesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowRunActionExpressionTracesResult, error) {
			args := v.(ListWorkflowRunActionExpressionTracesArgs)
			r, err := ListWorkflowRunActionExpressionTraces(ctx, &args, opts...)
			var s ListWorkflowRunActionExpressionTracesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowRunActionExpressionTracesResultOutput)
}

type ListWorkflowRunActionExpressionTracesOutputArgs struct {
	ActionName        pulumi.StringInput `pulumi:"actionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RunName           pulumi.StringInput `pulumi:"runName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (ListWorkflowRunActionExpressionTracesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionExpressionTracesArgs)(nil)).Elem()
}


type ListWorkflowRunActionExpressionTracesResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowRunActionExpressionTracesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowRunActionExpressionTracesResult)(nil)).Elem()
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) ToListWorkflowRunActionExpressionTracesResultOutput() ListWorkflowRunActionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) ToListWorkflowRunActionExpressionTracesResultOutputWithContext(ctx context.Context) ListWorkflowRunActionExpressionTracesResultOutput {
	return o
}

func (o ListWorkflowRunActionExpressionTracesResultOutput) Inputs() ExpressionRootResponseArrayOutput {
	return o.ApplyT(func(v ListWorkflowRunActionExpressionTracesResult) []ExpressionRootResponse { return v.Inputs }).(ExpressionRootResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowRunActionExpressionTracesResultOutput{})
}
