


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionSecretsSlot(ctx *pulumi.Context, args *ListWebAppFunctionSecretsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionSecretsSlotResult, error) {
	var rv ListWebAppFunctionSecretsSlotResult
	err := ctx.Invoke("azure-native:web/v20200601:listWebAppFunctionSecretsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionSecretsSlotArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppFunctionSecretsSlotResult struct {
	Key        *string `pulumi:"key"`
	TriggerUrl *string `pulumi:"triggerUrl"`
}

func ListWebAppFunctionSecretsSlotOutput(ctx *pulumi.Context, args ListWebAppFunctionSecretsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionSecretsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionSecretsSlotResult, error) {
			args := v.(ListWebAppFunctionSecretsSlotArgs)
			r, err := ListWebAppFunctionSecretsSlot(ctx, &args, opts...)
			var s ListWebAppFunctionSecretsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppFunctionSecretsSlotResultOutput)
}

type ListWebAppFunctionSecretsSlotOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppFunctionSecretsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsSlotArgs)(nil)).Elem()
}


type ListWebAppFunctionSecretsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionSecretsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionSecretsSlotResult)(nil)).Elem()
}

func (o ListWebAppFunctionSecretsSlotResultOutput) ToListWebAppFunctionSecretsSlotResultOutput() ListWebAppFunctionSecretsSlotResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsSlotResultOutput) ToListWebAppFunctionSecretsSlotResultOutputWithContext(ctx context.Context) ListWebAppFunctionSecretsSlotResultOutput {
	return o
}

func (o ListWebAppFunctionSecretsSlotResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsSlotResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ListWebAppFunctionSecretsSlotResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionSecretsSlotResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionSecretsSlotResultOutput{})
}
