


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSyncFunctionTriggers(ctx *pulumi.Context, args *ListWebAppSyncFunctionTriggersArgs, opts ...pulumi.InvokeOption) (*ListWebAppSyncFunctionTriggersResult, error) {
	var rv ListWebAppSyncFunctionTriggersResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppSyncFunctionTriggers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSyncFunctionTriggersArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppSyncFunctionTriggersResult struct {
	Id         string  `pulumi:"id"`
	Key        *string `pulumi:"key"`
	Kind       *string `pulumi:"kind"`
	Name       string  `pulumi:"name"`
	TriggerUrl *string `pulumi:"triggerUrl"`
	Type       string  `pulumi:"type"`
}

func ListWebAppSyncFunctionTriggersOutput(ctx *pulumi.Context, args ListWebAppSyncFunctionTriggersOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSyncFunctionTriggersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSyncFunctionTriggersResult, error) {
			args := v.(ListWebAppSyncFunctionTriggersArgs)
			r, err := ListWebAppSyncFunctionTriggers(ctx, &args, opts...)
			var s ListWebAppSyncFunctionTriggersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSyncFunctionTriggersResultOutput)
}

type ListWebAppSyncFunctionTriggersOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSyncFunctionTriggersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersArgs)(nil)).Elem()
}


type ListWebAppSyncFunctionTriggersResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSyncFunctionTriggersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSyncFunctionTriggersResult)(nil)).Elem()
}

func (o ListWebAppSyncFunctionTriggersResultOutput) ToListWebAppSyncFunctionTriggersResultOutput() ListWebAppSyncFunctionTriggersResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersResultOutput) ToListWebAppSyncFunctionTriggersResultOutputWithContext(ctx context.Context) ListWebAppSyncFunctionTriggersResultOutput {
	return o
}

func (o ListWebAppSyncFunctionTriggersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppSyncFunctionTriggersResultOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppSyncFunctionTriggersResultOutput) TriggerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) *string { return v.TriggerUrl }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSyncFunctionTriggersResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSyncFunctionTriggersResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSyncFunctionTriggersResultOutput{})
}
