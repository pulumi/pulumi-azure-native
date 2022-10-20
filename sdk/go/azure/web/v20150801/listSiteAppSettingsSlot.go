


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAppSettingsSlot(ctx *pulumi.Context, args *ListSiteAppSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsSlotResult, error) {
	var rv ListSiteAppSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteAppSettingsSlotResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}

func ListSiteAppSettingsSlotOutput(ctx *pulumi.Context, args ListSiteAppSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteAppSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteAppSettingsSlotResult, error) {
			args := v.(ListSiteAppSettingsSlotArgs)
			r, err := ListSiteAppSettingsSlot(ctx, &args, opts...)
			var s ListSiteAppSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteAppSettingsSlotResultOutput)
}

type ListSiteAppSettingsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteAppSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsSlotArgs)(nil)).Elem()
}


type ListSiteAppSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteAppSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsSlotResult)(nil)).Elem()
}

func (o ListSiteAppSettingsSlotResultOutput) ToListSiteAppSettingsSlotResultOutput() ListSiteAppSettingsSlotResultOutput {
	return o
}

func (o ListSiteAppSettingsSlotResultOutput) ToListSiteAppSettingsSlotResultOutputWithContext(ctx context.Context) ListSiteAppSettingsSlotResultOutput {
	return o
}

func (o ListSiteAppSettingsSlotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteAppSettingsSlotResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsSlotResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteAppSettingsSlotResultOutput{})
}
