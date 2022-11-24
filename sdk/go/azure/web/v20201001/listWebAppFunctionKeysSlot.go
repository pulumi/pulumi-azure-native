


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppFunctionKeysSlot(ctx *pulumi.Context, args *ListWebAppFunctionKeysSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppFunctionKeysSlotResult, error) {
	var rv ListWebAppFunctionKeysSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppFunctionKeysSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppFunctionKeysSlotArgs struct {
	FunctionName      string `pulumi:"functionName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppFunctionKeysSlotResult struct {
	Id         string             `pulumi:"id"`
	Kind       *string            `pulumi:"kind"`
	Name       string             `pulumi:"name"`
	Properties map[string]string  `pulumi:"properties"`
	SystemData SystemDataResponse `pulumi:"systemData"`
	Type       string             `pulumi:"type"`
}

func ListWebAppFunctionKeysSlotOutput(ctx *pulumi.Context, args ListWebAppFunctionKeysSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppFunctionKeysSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppFunctionKeysSlotResult, error) {
			args := v.(ListWebAppFunctionKeysSlotArgs)
			r, err := ListWebAppFunctionKeysSlot(ctx, &args, opts...)
			var s ListWebAppFunctionKeysSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppFunctionKeysSlotResultOutput)
}

type ListWebAppFunctionKeysSlotOutputArgs struct {
	FunctionName      pulumi.StringInput `pulumi:"functionName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppFunctionKeysSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysSlotArgs)(nil)).Elem()
}


type ListWebAppFunctionKeysSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppFunctionKeysSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppFunctionKeysSlotResult)(nil)).Elem()
}

func (o ListWebAppFunctionKeysSlotResultOutput) ToListWebAppFunctionKeysSlotResultOutput() ListWebAppFunctionKeysSlotResultOutput {
	return o
}

func (o ListWebAppFunctionKeysSlotResultOutput) ToListWebAppFunctionKeysSlotResultOutputWithContext(ctx context.Context) ListWebAppFunctionKeysSlotResultOutput {
	return o
}

func (o ListWebAppFunctionKeysSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppFunctionKeysSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppFunctionKeysSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppFunctionKeysSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppFunctionKeysSlotResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ListWebAppFunctionKeysSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppFunctionKeysSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppFunctionKeysSlotResultOutput{})
}
