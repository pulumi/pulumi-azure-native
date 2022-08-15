


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowVersionTriggerCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionTriggerCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionTriggerCallbackUrlResult, error) {
	var rv ListWorkflowVersionTriggerCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:listWorkflowVersionTriggerCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionTriggerCallbackUrlArgs struct {
	KeyType           *string `pulumi:"keyType"`
	NotAfter          *string `pulumi:"notAfter"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	TriggerName       string  `pulumi:"triggerName"`
	VersionId         string  `pulumi:"versionId"`
	WorkflowName      string  `pulumi:"workflowName"`
}


type ListWorkflowVersionTriggerCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListWorkflowVersionTriggerCallbackUrlOutput(ctx *pulumi.Context, args ListWorkflowVersionTriggerCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowVersionTriggerCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowVersionTriggerCallbackUrlResult, error) {
			args := v.(ListWorkflowVersionTriggerCallbackUrlArgs)
			r, err := ListWorkflowVersionTriggerCallbackUrl(ctx, &args, opts...)
			var s ListWorkflowVersionTriggerCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowVersionTriggerCallbackUrlResultOutput)
}

type ListWorkflowVersionTriggerCallbackUrlOutputArgs struct {
	KeyType           pulumi.StringPtrInput `pulumi:"keyType"`
	NotAfter          pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput    `pulumi:"triggerName"`
	VersionId         pulumi.StringInput    `pulumi:"versionId"`
	WorkflowName      pulumi.StringInput    `pulumi:"workflowName"`
}

func (ListWorkflowVersionTriggerCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionTriggerCallbackUrlArgs)(nil)).Elem()
}


type ListWorkflowVersionTriggerCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowVersionTriggerCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionTriggerCallbackUrlResult)(nil)).Elem()
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) ToListWorkflowVersionTriggerCallbackUrlResultOutput() ListWorkflowVersionTriggerCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) ToListWorkflowVersionTriggerCallbackUrlResultOutputWithContext(ctx context.Context) ListWorkflowVersionTriggerCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListWorkflowVersionTriggerCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionTriggerCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowVersionTriggerCallbackUrlResultOutput{})
}
