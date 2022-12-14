


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAgreementContentCallbackUrl(ctx *pulumi.Context, args *ListAgreementContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListAgreementContentCallbackUrlResult, error) {
	var rv ListAgreementContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listAgreementContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAgreementContentCallbackUrlArgs struct {
	AgreementName          string   `pulumi:"agreementName"`
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListAgreementContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListAgreementContentCallbackUrlOutput(ctx *pulumi.Context, args ListAgreementContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListAgreementContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListAgreementContentCallbackUrlResult, error) {
			args := v.(ListAgreementContentCallbackUrlArgs)
			r, err := ListAgreementContentCallbackUrl(ctx, &args, opts...)
			var s ListAgreementContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListAgreementContentCallbackUrlResultOutput)
}

type ListAgreementContentCallbackUrlOutputArgs struct {
	AgreementName          pulumi.StringInput    `pulumi:"agreementName"`
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListAgreementContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAgreementContentCallbackUrlArgs)(nil)).Elem()
}


type ListAgreementContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListAgreementContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListAgreementContentCallbackUrlResult)(nil)).Elem()
}

func (o ListAgreementContentCallbackUrlResultOutput) ToListAgreementContentCallbackUrlResultOutput() ListAgreementContentCallbackUrlResultOutput {
	return o
}

func (o ListAgreementContentCallbackUrlResultOutput) ToListAgreementContentCallbackUrlResultOutputWithContext(ctx context.Context) ListAgreementContentCallbackUrlResultOutput {
	return o
}

func (o ListAgreementContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListAgreementContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListAgreementContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListAgreementContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListAgreementContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListAgreementContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListAgreementContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListAgreementContentCallbackUrlResultOutput{})
}
