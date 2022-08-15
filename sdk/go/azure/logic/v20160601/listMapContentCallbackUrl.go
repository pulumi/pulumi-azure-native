


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMapContentCallbackUrl(ctx *pulumi.Context, args *ListMapContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListMapContentCallbackUrlResult, error) {
	var rv ListMapContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listMapContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMapContentCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	MapName                string   `pulumi:"mapName"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListMapContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListMapContentCallbackUrlOutput(ctx *pulumi.Context, args ListMapContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListMapContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMapContentCallbackUrlResult, error) {
			args := v.(ListMapContentCallbackUrlArgs)
			r, err := ListMapContentCallbackUrl(ctx, &args, opts...)
			var s ListMapContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMapContentCallbackUrlResultOutput)
}

type ListMapContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                KeyTypePtrInput       `pulumi:"keyType"`
	MapName                pulumi.StringInput    `pulumi:"mapName"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListMapContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMapContentCallbackUrlArgs)(nil)).Elem()
}


type ListMapContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListMapContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMapContentCallbackUrlResult)(nil)).Elem()
}

func (o ListMapContentCallbackUrlResultOutput) ToListMapContentCallbackUrlResultOutput() ListMapContentCallbackUrlResultOutput {
	return o
}

func (o ListMapContentCallbackUrlResultOutput) ToListMapContentCallbackUrlResultOutputWithContext(ctx context.Context) ListMapContentCallbackUrlResultOutput {
	return o
}

func (o ListMapContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListMapContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListMapContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListMapContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListMapContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListMapContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListMapContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMapContentCallbackUrlResultOutput{})
}
