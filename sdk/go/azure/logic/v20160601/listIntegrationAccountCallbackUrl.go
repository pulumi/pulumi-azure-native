


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListIntegrationAccountCallbackUrl(ctx *pulumi.Context, args *ListIntegrationAccountCallbackUrlArgs, opts ...pulumi.InvokeOption) (*ListIntegrationAccountCallbackUrlResult, error) {
	var rv ListIntegrationAccountCallbackUrlResult
	err := ctx.Invoke("azure-native:logic/v20160601:listIntegrationAccountCallbackUrl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListIntegrationAccountCallbackUrlArgs struct {
	IntegrationAccountName string   `pulumi:"integrationAccountName"`
	KeyType                *KeyType `pulumi:"keyType"`
	NotAfter               *string  `pulumi:"notAfter"`
	ResourceGroupName      string   `pulumi:"resourceGroupName"`
}


type ListIntegrationAccountCallbackUrlResult struct {
	Value *string `pulumi:"value"`
}

func ListIntegrationAccountCallbackUrlOutput(ctx *pulumi.Context, args ListIntegrationAccountCallbackUrlOutputArgs, opts ...pulumi.InvokeOption) ListIntegrationAccountCallbackUrlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListIntegrationAccountCallbackUrlResult, error) {
			args := v.(ListIntegrationAccountCallbackUrlArgs)
			r, err := ListIntegrationAccountCallbackUrl(ctx, &args, opts...)
			var s ListIntegrationAccountCallbackUrlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListIntegrationAccountCallbackUrlResultOutput)
}

type ListIntegrationAccountCallbackUrlOutputArgs struct {
	IntegrationAccountName pulumi.StringInput    `pulumi:"integrationAccountName"`
	KeyType                KeyTypePtrInput       `pulumi:"keyType"`
	NotAfter               pulumi.StringPtrInput `pulumi:"notAfter"`
	ResourceGroupName      pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListIntegrationAccountCallbackUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountCallbackUrlArgs)(nil)).Elem()
}


type ListIntegrationAccountCallbackUrlResultOutput struct{ *pulumi.OutputState }

func (ListIntegrationAccountCallbackUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListIntegrationAccountCallbackUrlResult)(nil)).Elem()
}

func (o ListIntegrationAccountCallbackUrlResultOutput) ToListIntegrationAccountCallbackUrlResultOutput() ListIntegrationAccountCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountCallbackUrlResultOutput) ToListIntegrationAccountCallbackUrlResultOutputWithContext(ctx context.Context) ListIntegrationAccountCallbackUrlResultOutput {
	return o
}

func (o ListIntegrationAccountCallbackUrlResultOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListIntegrationAccountCallbackUrlResult) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListIntegrationAccountCallbackUrlResultOutput{})
}
