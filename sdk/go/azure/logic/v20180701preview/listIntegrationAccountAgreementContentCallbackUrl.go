


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountAgreementContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountAgreementContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountAgreementContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountAgreementContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:listIntegrationAccountAgreementContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountAgreementContentCallbackUrlArgs struct {
	AgreementName          string  `pulumi:"agreementName"`
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountAgreementContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListIntegrationAccountAgreementContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountAgreementContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountAgreementContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountAgreementContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountAgreementContentCallbackUrlArgs)
			r, err := ListIntegrationAccountAgreementContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountAgreementContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountAgreementContentCallbackUrlResultOutput)
}

type ListIntegrationAccountAgreementContentCallbackUrlOutputArgs struct {
	AgreementName          pulumi.StringInput    `pulumi:"agreementName"`
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                pulumi.StringPtrInput `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountAgreementContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAgreementContentCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountAgreementContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountAgreementContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountAgreementContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) ToListIntegrationAccountAgreementContentCallbackUrlResultOutput() ListIntegrationAccountAgreementContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) ToListIntegrationAccountAgreementContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountAgreementContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) []string {
		return v.RelativePathParameters
	}).(pulumi.StringArrayOutput)
}

func (o ListIntegrationAccountAgreementContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountAgreementContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountAgreementContentCallbackUrlResultOutput{})
}
