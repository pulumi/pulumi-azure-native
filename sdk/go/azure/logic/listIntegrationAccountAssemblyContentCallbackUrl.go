


package logic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountAssemblyContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountAssemblyContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountAssemblyContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountAssemblyContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic:listIntegrationAccountAssemblyContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountAssemblyContentCallbackUrlArgs struct {
	AssemblyArtifactName   string `pulumi:"assemblyArtifactName"`
	IntegrationAccountName string `pulumi:"integrationAccountName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountAssemblyContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListIntegrationAccountAssemblyContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountAssemblyContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountAssemblyContentCallbackUrlArgs)
			r, err := ListIntegrationAccountAssemblyContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountAssemblyContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountAssemblyContentCallbackUrlResultOutput)
}

type ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs struct {
	AssemblyArtifactName   pulumi.StringInput `pulumi:"assemblyArtifactName"`
	IntegrationAccountName pulumi.StringInput `pulumi:"integrationAccountName"`
	ResourceGroupName      pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountAssemblyContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAssemblyContentCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountAssemblyContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAssemblyContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ToListIntegrationAccountAssemblyContentCallbackUrlResultOutput() ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) ToListIntegrationAccountAssemblyContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountAssemblyContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) []string {
		return v.RelativePathParameters
	}).(pulumi.StringArrayOutput)
}

func (o ListIntegrationAccountAssemblyContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAssemblyContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountAssemblyContentCallbackUrlResultOutput{})
}
