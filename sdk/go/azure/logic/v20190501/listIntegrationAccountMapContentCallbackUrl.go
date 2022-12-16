


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountMapContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountMapContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountMapContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountMapContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listIntegrationAccountMapContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountMapContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	MapName                string  `pulumi:"mapName"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountMapContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListIntegrationAccountMapContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountMapContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountMapContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountMapContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountMapContentCallbackUrlArgs)
			r, err := ListIntegrationAccountMapContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountMapContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountMapContentCallbackUrlResultOutput)
}

type ListIntegrationAccountMapContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                pulumi.StringPtrInput `pulumi:"keyType"`
	MapName                pulumi.StringInput    `pulumi:"mapName"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountMapContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountMapContentCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountMapContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountMapContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountMapContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) ToListIntegrationAccountMapContentCallbackUrlResultOutput() ListIntegrationAccountMapContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) ToListIntegrationAccountMapContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountMapContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListIntegrationAccountMapContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountMapContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountMapContentCallbackUrlResultOutput{})
}
