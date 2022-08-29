


package web

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSitePushSettingsSlot(ctx *pulumi.Context, args *ListWebAppSitePushSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsSlotResult, error) {
	var rv ListWebAppSitePushSettingsSlotResult
	err := ctx.Invoke("azure-native:web:listWebAppSitePushSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSitePushSettingsSlotResult struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	Id                string  `pulumi:"id"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	Type              string  `pulumi:"type"`
}

func ListWebAppSitePushSettingsSlotOutput(ctx *pulumi.Context, args ListWebAppSitePushSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSitePushSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSitePushSettingsSlotResult, error) {
			args := v.(ListWebAppSitePushSettingsSlotArgs)
			r, err := ListWebAppSitePushSettingsSlot(ctx, &args, opts...)
			var s ListWebAppSitePushSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSitePushSettingsSlotResultOutput)
}

type ListWebAppSitePushSettingsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppSitePushSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsSlotArgs)(nil)).Elem()
}


type ListWebAppSitePushSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSitePushSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsSlotResult)(nil)).Elem()
}

func (o ListWebAppSitePushSettingsSlotResultOutput) ToListWebAppSitePushSettingsSlotResultOutput() ListWebAppSitePushSettingsSlotResultOutput {
	return o
}

func (o ListWebAppSitePushSettingsSlotResultOutput) ToListWebAppSitePushSettingsSlotResultOutputWithContext(ctx context.Context) ListWebAppSitePushSettingsSlotResultOutput {
	return o
}

func (o ListWebAppSitePushSettingsSlotResultOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSitePushSettingsSlotResultOutput{})
}
