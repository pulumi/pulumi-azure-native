


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStaticSiteBuildFunctionAppSettings(ctx *pulumi.Context, args *ListStaticSiteBuildFunctionAppSettingsArgs, opts ...pulumi.InvokeOption) (*ListStaticSiteBuildFunctionAppSettingsResult, error) {
	var rv ListStaticSiteBuildFunctionAppSettingsResult
	err := ctx.Invoke("azure-native:web/v20220301:listStaticSiteBuildFunctionAppSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStaticSiteBuildFunctionAppSettingsArgs struct {
	EnvironmentName   string `pulumi:"environmentName"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListStaticSiteBuildFunctionAppSettingsResult struct {
	Id         string            `pulumi:"id"`
	Kind       *string           `pulumi:"kind"`
	Name       string            `pulumi:"name"`
	Properties map[string]string `pulumi:"properties"`
	Type       string            `pulumi:"type"`
}

func ListStaticSiteBuildFunctionAppSettingsOutput(ctx *pulumi.Context, args ListStaticSiteBuildFunctionAppSettingsOutputArgs, opts ...pulumi.InvokeOption) ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStaticSiteBuildFunctionAppSettingsResult, error) {
			args := v.(ListStaticSiteBuildFunctionAppSettingsArgs)
			r, err := ListStaticSiteBuildFunctionAppSettings(ctx, &args, opts...)
			var s ListStaticSiteBuildFunctionAppSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStaticSiteBuildFunctionAppSettingsResultOutput)
}

type ListStaticSiteBuildFunctionAppSettingsOutputArgs struct {
	EnvironmentName   pulumi.StringInput `pulumi:"environmentName"`
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListStaticSiteBuildFunctionAppSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildFunctionAppSettingsArgs)(nil)).Elem()
}


type ListStaticSiteBuildFunctionAppSettingsResultOutput struct{ *pulumi.OutputState }

func (ListStaticSiteBuildFunctionAppSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStaticSiteBuildFunctionAppSettingsResult)(nil)).Elem()
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) ToListStaticSiteBuildFunctionAppSettingsResultOutput() ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) ToListStaticSiteBuildFunctionAppSettingsResultOutputWithContext(ctx context.Context) ListStaticSiteBuildFunctionAppSettingsResultOutput {
	return o
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o ListStaticSiteBuildFunctionAppSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListStaticSiteBuildFunctionAppSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStaticSiteBuildFunctionAppSettingsResultOutput{})
}
