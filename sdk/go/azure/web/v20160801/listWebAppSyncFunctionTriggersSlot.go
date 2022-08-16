


package v20160801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSyncFunctionTriggersSlot(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersSlotResult, error) {
	var rv ListWebAppSyncFunctionTriggersSlotResult
	err := ctx.Invoke("azure-native:web/v20160801:listWebAppSyncFunctionTriggersSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSyncFunctionTriggersSlotResult struct {
	Id         string  `pulumi:"id"`
	Key        *string `pulumi:"key"`
	Kind       *string `pulumi:"kind"`
	Name       string  `pulumi:"name"`
	TriggerUrl *string `pulumi:"triggerUrl"`
	Type       string  `pulumi:"type"`
}

func ListWebAppSyncFunctionTriggersSlotOutput(ctx *pulumi.Context, args ListWebAppSyncFunctionTriggersSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSyncFunctionTriggersSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSyncFunctionTriggersSlotResult, error) {
			args := v.(ListWebAppSyncFunctionTriggersSlotArgs)
			r, err := ListWebAppSyncFunctionTriggersSlot(ctx, &args, opts...)
			var s ListWebAppSyncFunctionTriggersSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSyncFunctionTriggersSlotResultOutput)
}

type ListWebAppSyncFunctionTriggersSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppSyncFunctionTriggersSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersSlotArgs)(nil)).Elem()
}


type ListWebAppSyncFunctionTriggersSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSyncFunctionTriggersSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersSlotResult)(nil)).Elem()
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) ToListWebAppSyncFunctionTriggersSlotResultOutput() ListWebAppSyncFunctionTriggersSlotResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) ToListWebAppSyncFunctionTriggersSlotResultOutputWithContext(ctx context.Context) ListWebAppSyncFunctionTriggersSlotResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSyncFunctionTriggersSlotResultOutput{})
}
