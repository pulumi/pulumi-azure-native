


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSchemaContentCallbackUrl(ctx *pulumi.Context, args *ListSchemaContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListSchemaContentCallbackUrlResult, error) {
	var rv ListSchemaContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listSchemaContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSchemaContentCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
	SchemaName             string   `pulumi:"schemaName"`
}


type ListSchemaContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListSchemaContentCallbackUrlOutput(ctx *pulumi.Context, args ListSchemaContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListSchemaContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSchemaContentCallbackUrlResult, error) {
			args := v.(ListSchemaContentCallbackUrlArgs)
			r, err := ListSchemaContentCallbackUrl(ctx, &args, opts...)
			var s ListSchemaContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSchemaContentCallbackUrlResultOutput)
}

type ListSchemaContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
	SchemaName             pulumi.StringInput    `pulumi:"schemaName"`
}

func (ListSchemaContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSchemaContentCallbackUrlArgs)(nil)).Elem()
}


type ListSchemaContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListSchemaContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSchemaContentCallbackUrlResult)(nil)).Elem()
}

func (o ListSchemaContentCallbackUrlResultOutput) ToListSchemaContentCallbackUrlResultOutput() ListSchemaContentCallbackUrlResultOutput {
	return o
}

func (o ListSchemaContentCallbackUrlResultOutput) ToListSchemaContentCallbackUrlResultOutputWithContext(ctx context.Context) ListSchemaContentCallbackUrlResultOutput {
	return o
}

func (o ListSchemaContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListSchemaContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListSchemaContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListSchemaContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListSchemaContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListSchemaContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListSchemaContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSchemaContentCallbackUrlResultOutput{})
}
