


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkflowVersionCallbackUrl(ctx *pulumi.Context, args *ListWorkflowVersionCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListWorkflowVersionCallbackUrlResult, error) {
	var rv ListWorkflowVersionCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listWorkflowVersionCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkflowVersionCallbackUrlArgs struct {
	KeyType           *KeyType `pulumi:"keyType"`
	NotAfter          *string  `pulumi:"notAfter"`
	ResourceGroupName string   `pulumi:"resourceGroupName"`
	TriggerName       string   `pulumi:"triggerName"`
	VersionId         string   `pulumi:"versionId"`
	WorkflowName      string   `pulumi:"workflowName"`
}


type ListWorkflowVersionCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListWorkflowVersionCallbackUrlOutput(ctx *pulumi.Context, args ListWorkflowVersionCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListWorkflowVersionCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkflowVersionCallbackUrlResult, error) {
			args := v.(ListWorkflowVersionCallbackUrlArgs)
			r, err := ListWorkflowVersionCallbackUrl(ctx, &args, opts...)
			var s ListWorkflowVersionCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkflowVersionCallbackUrlResultOutput)
}

type ListWorkflowVersionCallbackUrlOutputArgs struct {
	KeyType           KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter          pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
	TriggerName       pulumi.StringInput    `pulumi:"triggerName"`
	VersionId         pulumi.StringInput    `pulumi:"versionId"`
	WorkflowName      pulumi.StringInput    `pulumi:"workflowName"`
}

func (ListWorkflowVersionCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionCallbackUrlArgs)(nil)).Elem()
}


type ListWorkflowVersionCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListWorkflowVersionCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkflowVersionCallbackUrlResult)(nil)).Elem()
}

func (o ListWorkflowVersionCallbackUrlResultOutput) ToListWorkflowVersionCallbackUrlResultOutput() ListWorkflowVersionCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowVersionCallbackUrlResultOutput) ToListWorkflowVersionCallbackUrlResultOutputWithContext(ctx context.Context) ListWorkflowVersionCallbackUrlResultOutput {
	return o
}

func (o ListWorkflowVersionCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListWorkflowVersionCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListWorkflowVersionCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListWorkflowVersionCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkflowVersionCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkflowVersionCallbackUrlResultOutput{})
}
