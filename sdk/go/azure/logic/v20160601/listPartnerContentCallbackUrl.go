


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListPartnerContentCallbackUrl(ctx *pulumi.Context, args *ListPartnerContentCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListPartnerContentCallbackUrlResult, error) {
	var rv ListPartnerContentCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listPartnerContentCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListPartnerContentCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	PartnerName            string   `pulumi:"partnerName"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListPartnerContentCallbackUrlResult struct {
	BasePath               string                                         `pulumi:"basePath"`
	Method                 string                                         `pulumi:"method"`
	Queries                *WorkflowTriggerListCallbackUrlQueriesResponse `pulumi:"queries"`
	RelativePath           string                                         `pulumi:"relativePath"`
	RelativePathParameters []string                                       `pulumi:"relativePathParameters"`
	Value                  string                                         `pulumi:"value"`
}

func ListPartnerContentCallbackUrlOutput(ctx *pulumi.Context, args ListPartnerContentCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListPartnerContentCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListPartnerContentCallbackUrlResult, error) {
			args := v.(ListPartnerContentCallbackUrlArgs)
			r, err := ListPartnerContentCallbackUrl(ctx, &args, opts...)
			var s ListPartnerContentCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListPartnerContentCallbackUrlResultOutput)
}

type ListPartnerContentCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	PartnerName            pulumi.StringInput    `pulumi:"partnerName"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListPartnerContentCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPartnerContentCallbackUrlArgs)(nil)).Elem()
}


type ListPartnerContentCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListPartnerContentCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListPartnerContentCallbackUrlResult)(nil)).Elem()
}

func (o ListPartnerContentCallbackUrlResultOutput) ToListPartnerContentCallbackUrlResultOutput() ListPartnerContentCallbackUrlResultOutput {
	return o
}

func (o ListPartnerContentCallbackUrlResultOutput) ToListPartnerContentCallbackUrlResultOutputWithContext(ctx context.Context) ListPartnerContentCallbackUrlResultOutput {
	return o
}

func (o ListPartnerContentCallbackUrlResultOutput) BasePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) string { return v.BasePath }).(pulumi.StringOutput)
}

func (o ListPartnerContentCallbackUrlResultOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) string { return v.Method }).(pulumi.StringOutput)
}

func (o ListPartnerContentCallbackUrlResultOutput) Queries() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) *WorkflowTriggerListCallbackUrlQueriesResponse {
		return v.Queries
	}).(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput)
}

func (o ListPartnerContentCallbackUrlResultOutput) RelativePath() pulumi.StringOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) string { return v.RelativePath }).(pulumi.StringOutput)
}

func (o ListPartnerContentCallbackUrlResultOutput) RelativePathParameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) []string { return v.RelativePathParameters }).(pulumi.StringArrayOutput)
}

func (o ListPartnerContentCallbackUrlResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ListPartnerContentCallbackUrlResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListPartnerContentCallbackUrlResultOutput{})
}
