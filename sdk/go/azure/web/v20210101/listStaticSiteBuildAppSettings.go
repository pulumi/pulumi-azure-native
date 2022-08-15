


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteBuildAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildAppSettingsResult, error) {
	var rv ListStaticSiteBuildAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20210101:listStaticSiteBuildAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildAppSettingsArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteBuildAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListStaticSiteBuildAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteBuildAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteBuildAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteBuildAppSettingsResult, error) {
			args := v.(ListStaticSiteBuildAppSettingsArgs)
			r, err := ListStaticSiteBuildAppSettings(ctx, &args, opts...)
			var s ListStaticSiteBuildAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteBuildAppSettingsResultOutput)
}

type ListStaticSiteBuildAppSettingsOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteBuildAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildAppSettingsArgs)(nil)).Elem()
}


type ListStaticSiteBuildAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteBuildAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteBuildAppSettingsResultOutput) ToListStaticSiteBuildAppSettingsResultOutput() ListStaticSiteBuildAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildAppSettingsResultOutput) ToListStaticSiteBuildAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteBuildAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteBuildAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteBuildAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteBuildAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListStaticSiteBuildAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteBuildAppSettingsResultOutput{})
}
