


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountPartnerContentCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountPartnerContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountPartnerContentCallbackUrlResult, error) {
	var rv ListIntegrationAccountPartnerContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20180701preview:listIntegrationAccountPartnerContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountPartnerContentCallbackUrlArgs struct {
	IntegrationAccountName string  `pulumi:"integrationAccountName"`
	KeyType                *string `pulumi:"keyType"`
	NotAfter               *string `pulumi:"notAfter"`
	PartnerName            string  `pulumi:"partnerName"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountPartnerContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListIntegrationAccountPartnerContentCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountPartnerContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountPartnerContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountPartnerContentCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountPartnerContentCallbackUrlArgs)
			r, err := ListIntegrationAccountPartnerContentCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountPartnerContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountPartnerContentCallbackUrlResultOutput)
}

type ListIntegrationAccountPartnerContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                pulumi.StringPtrInput `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	PartnerName            pulumi.StringInput    `pulumi:"partnerName"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountPartnerContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountPartnerContentCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountPartnerContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountPartnerContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountPartnerContentCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) ToListIntegrationAccountPartnerContentCallbackUrlResultOutput() ListIntegrationAccountPartnerContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) ToListIntegrationAccountPartnerContentCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountPartnerContentCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) []string {
		return v.RelativePathParameters
	}).(pulumi.StringArrayOutput)
}

func (o ListIntegrationAccountPartnerContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListIntegrationAccountPartnerContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountPartnerContentCallbackUrlResultOutput{})
}
