


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowCallbackUrl(ctx *pulumi.Context, args *ListWorkflowCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowCallbackUrlResult, error) {
	var rv ListWorkflowCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listWorkflowCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowCallbackUrlArgs struct {
	KeyType           *KeyType `pulumi:"keyType"`
	NotAfter          *string  `pulumi:"notAfter"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	WorkflowName      string   `pulumi:"workflowName"`
}


type ListWorkflowCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListWorkflowCallbackUrlOutput(ctx *pulumi.Context, args ListWorkflowCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowCallbackUrlResult, error) {
			args := v.(ListWorkflowCallbackUrlArgs)
			r, err := ListWorkflowCallbackUrl(ctx, &args, opts...)
			var s ListWorkflowCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowCallbackUrlResultOutput)
}

type ListWorkflowCallbackUrlOutputArgs struct {
	KeyType           KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter          pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	WorkflowName      pulumi.StringInput    `pulumi:"workflowName"`
}

func (ListWorkflowCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowCallbackUrlArgs)(nil)).Elem()
}


type ListWorkflowCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowCallbackUrlResult)(nil)).Elem()
}

func (o ListWorkflowCallbackUrlResultOutput) ToListWorkflowCallbackUrlResultOutput() ListWorkflowCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowCallbackUrlResultOutput) ToListWorkflowCallbackUrlResultOutputWithContext(ctx context.Context) ListWorkflowCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListWorkflowCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListWorkflowCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse { return v.Queries }).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListWorkflowCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListWorkflowCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListWorkflowCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowCallbackUrlResultOutput{})
}
