


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionSecrets(ctx *pulumi.Context, args *ListWebAppFunctionSecretsArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionSecretsResult, error) {
	var rv ListWebAppFunctionSecretsResult
	err := ctx.Invoke("azure-native:web/v20210301:listWebAppFunctionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionSecretsArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppFunctionSecretsResult struct {
	Key        *string `pulumi:"key"`
	TriggerUrl *string `pulumi:"triggerUrl"`
}

func ListWebAppFunctionSecretsOutput(ctx *pulumi.Context, args ListWebAppFunctionSecretsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionSecretsResult, error) {
			args := v.(ListWebAppFunctionSecretsArgs)
			r, err := ListWebAppFunctionSecrets(ctx, &args, opts...)
			var s ListWebAppFunctionSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppFunctionSecretsResultOutput)
}

type ListWebAppFunctionSecretsOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppFunctionSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsArgs)(nil)).Elem()
}


type ListWebAppFunctionSecretsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsResult)(nil)).Elem()
}

func (o ListWebAppFunctionSecretsResultOutput) ToListWebAppFunctionSecretsResultOutput() ListWebAppFunctionSecretsResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsResultOutput) ToListWebAppFunctionSecretsResultOutputWithContext(ctx context.Context) ListWebAppFunctionSecretsResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ListWebAppFunctionSecretsResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionSecretsResultOutput{})
}
