


package v20210115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppApplicationSettings(ctx *pulumi.Context, args *ListWebAppApplicationSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppApplicationSettingsResult, error) {
	var rv ListWebAppApplicationSettingsResult
	err := ctx.Invoke("azure-native:web/v20210115:listWebAppApplicationSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppApplicationSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppApplicationSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListWebAppApplicationSettingsOutput(ctx *pulumi.Context, args ListWebAppApplicationSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppApplicationSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppApplicationSettingsResult, error) {
			args := v.(ListWebAppApplicationSettingsArgs)
			r, err := ListWebAppApplicationSettings(ctx, &args, opts...)
			var s ListWebAppApplicationSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppApplicationSettingsResultOutput)
}

type ListWebAppApplicationSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppApplicationSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsArgs)(nil)).Elem()
}


type ListWebAppApplicationSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppApplicationSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppApplicationSettingsResult)(nil)).Elem()
}

func (o ListWebAppApplicationSettingsResultOutput) ToListWebAppApplicationSettingsResultOutput() ListWebAppApplicationSettingsResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsResultOutput) ToListWebAppApplicationSettingsResultOutputWithContext(ctx context.Context) ListWebAppApplicationSettingsResultOutput {
	return o
}

func (o ListWebAppApplicationSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppApplicationSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppApplicationSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppApplicationSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListWebAppApplicationSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppApplicationSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppApplicationSettingsResultOutput{})
}
