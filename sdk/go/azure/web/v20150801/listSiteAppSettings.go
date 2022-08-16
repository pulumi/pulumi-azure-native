


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAppSettings(ctx *pulumi.Context, args *ListSiteAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListSiteAppSettingsResult, error) {
	var rv ListSiteAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAppSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteAppSettingsResult struct {
	Id         *string           `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Location   string            `pulumi:"location"`
	Name       *string           `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       *string           `pulumi:"type"`
}

func ListSiteAppSettingsOutput(ctx *pulumi.Context, args ListSiteAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListSiteAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteAppSettingsResult, error) {
			args := v.(ListSiteAppSettingsArgs)
			r, err := ListSiteAppSettings(ctx, &args, opts...)
			var s ListSiteAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteAppSettingsResultOutput)
}

type ListSiteAppSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsArgs)(nil)).Elem()
}


type ListSiteAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListSiteAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAppSettingsResult)(nil)).Elem()
}

func (o ListSiteAppSettingsResultOutput) ToListSiteAppSettingsResultOutput() ListSiteAppSettingsResultOutput {
	return o
}

func (o ListSiteAppSettingsResultOutput) ToListSiteAppSettingsResultOutputWithContext(ctx context.Context) ListSiteAppSettingsResultOutput {
	return o
}

func (o ListSiteAppSettingsResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o ListSiteAppSettingsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ListSiteAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListSiteAppSettingsResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ListSiteAppSettingsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAppSettingsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteAppSettingsResultOutput{})
}
