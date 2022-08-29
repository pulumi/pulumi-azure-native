


package v20180201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppApplicationSettingsSlot(ctx *pulumi.Context, args *ListWebAppApplicationSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsSlotResult, error) {
	var rv ListWebAppApplicationSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20180201:listWebAppApplicationSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppApplicationSettingsSlotResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListWebAppApplicationSettingsSlotOutput(ctx *pulumi.Context, args ListWebAppApplicationSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppApplicationSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppApplicationSettingsSlotResult, error) {
			args := v.(ListWebAppApplicationSettingsSlotArgs)
			r, err := ListWebAppApplicationSettingsSlot(ctx, &args, opts...)
			var s ListWebAppApplicationSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppApplicationSettingsSlotResultOutput)
}

type ListWebAppApplicationSettingsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppApplicationSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsSlotArgs)(nil)).Elem()
}


type ListWebAppApplicationSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppApplicationSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsSlotResult)(nil)).Elem()
}

func (o ListWebAppApplicationSettingsSlotResultOutput) ToListWebAppApplicationSettingsSlotResultOutput() ListWebAppApplicationSettingsSlotResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsSlotResultOutput) ToListWebAppApplicationSettingsSlotResultOutputWithContext(ctx context.Context) ListWebAppApplicationSettingsSlotResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppApplicationSettingsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppApplicationSettingsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppApplicationSettingsSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppApplicationSettingsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppApplicationSettingsSlotResultOutput{})
}
