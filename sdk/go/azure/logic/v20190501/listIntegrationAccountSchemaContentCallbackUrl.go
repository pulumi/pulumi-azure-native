


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountSchemaContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountSchemaContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountSchemaContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountSchemaContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20190501:listIntegrationAccountSchemaContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountSchemaContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
	SchemaName             string  `pulumi:"schemaName"`
}


type ListIntegrationAccountSchemaContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListIntegrationAccountSchemaContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountSchemaContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountSchemaContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountSchemaContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountSchemaContentCallbackUrlArgs)
			r, err := ListIntegrationAccountSchemaContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountSchemaContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountSchemaContentCallbackUrlResultOutput)
}

type ListIntegrationAccountSchemaContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                pulumi.StringPtrInput `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput    `pulumi:"schemaName"`
}

func (ListIntegrationAccountSchemaContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountSchemaContentCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountSchemaContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountSchemaContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountSchemaContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) ToListIntegrationAccountSchemaContentCallbackUrlResultOutput() ListIntegrationAccountSchemaContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) ToListIntegrationAccountSchemaContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountSchemaContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListIntegrationAccountSchemaContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountSchemaContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountSchemaContentCallbackUrlResultOutput{})
}
