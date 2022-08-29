


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSitePushSettings(ctx *pulumi.Context, args *ListWebAppSitePushSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSitePushSettingsResult, error) {
	var rv ListWebAppSitePushSettingsResult
	err := ctx.Invoke("azure-native:web/v20210301:listWebAppSitePushSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSitePushSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppSitePushSettingsResult struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	Id                string  `pulumi:"id"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	Type              string  `pulumi:"type"`
}

func ListWebAppSitePushSettingsOutput(ctx *pulumi.Context, args ListWebAppSitePushSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSitePushSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSitePushSettingsResult, error) {
			args := v.(ListWebAppSitePushSettingsArgs)
			r, err := ListWebAppSitePushSettings(ctx, &args, opts...)
			var s ListWebAppSitePushSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSitePushSettingsResultOutput)
}

type ListWebAppSitePushSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSitePushSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsArgs)(nil)).Elem()
}


type ListWebAppSitePushSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSitePushSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSitePushSettingsResult)(nil)).Elem()
}

func (o ListWebAppSitePushSettingsResultOutput) ToListWebAppSitePushSettingsResultOutput() ListWebAppSitePushSettingsResultOutput {
	return o
}

func (o ListWebAppSitePushSettingsResultOutput) ToListWebAppSitePushSettingsResultOutputWithContext(ctx context.Context) ListWebAppSitePushSettingsResultOutput {
	return o
}

func (o ListWebAppSitePushSettingsResultOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

func (o ListWebAppSitePushSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSitePushSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSitePushSettingsResultOutput{})
}
