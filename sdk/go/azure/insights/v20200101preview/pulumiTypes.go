


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupLogSettings struct {
	Category string `pulumi:"category"`
	Enabled  bool   `pulumi:"enabled"`
}





type ManagementGroupLogSettingsInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput
	ToManagementGroupLogSettingsOutputWithContext(context.Context) ManagementGroupLogSettingsOutput
}

type ManagementGroupLogSettingsArgs struct {
	Category pulumi.StringInput `pulumi:"category"`
	Enabled  pulumi.BoolInput   `pulumi:"enabled"`
}

func (ManagementGroupLogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettings)(nil)).Elem()
}

func (i ManagementGroupLogSettingsArgs) ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput {
	return i.ToManagementGroupLogSettingsOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsArgs) ToManagementGroupLogSettingsOutputWithContext(ctx context.Context) ManagementGroupLogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsOutput)
}





type ManagementGroupLogSettingsArrayInput interface {
	pulumi.Input

	ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput
	ToManagementGroupLogSettingsArrayOutputWithContext(context.Context) ManagementGroupLogSettingsArrayOutput
}

type ManagementGroupLogSettingsArray []ManagementGroupLogSettingsInput

func (ManagementGroupLogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettings)(nil)).Elem()
}

func (i ManagementGroupLogSettingsArray) ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput {
	return i.ToManagementGroupLogSettingsArrayOutputWithContext(context.Background())
}

func (i ManagementGroupLogSettingsArray) ToManagementGroupLogSettingsArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupLogSettingsArrayOutput)
}

type ManagementGroupLogSettingsOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettings)(nil)).Elem()
}

func (o ManagementGroupLogSettingsOutput) ToManagementGroupLogSettingsOutput() ManagementGroupLogSettingsOutput {
	return o
}

func (o ManagementGroupLogSettingsOutput) ToManagementGroupLogSettingsOutputWithContext(ctx context.Context) ManagementGroupLogSettingsOutput {
	return o
}

func (o ManagementGroupLogSettingsOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) string { return v.Category }).(pulumi.StringOutput)
}

func (o ManagementGroupLogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagementGroupLogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagementGroupLogSettingsArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettings)(nil)).Elem()
}

func (o ManagementGroupLogSettingsArrayOutput) ToManagementGroupLogSettingsArrayOutput() ManagementGroupLogSettingsArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsArrayOutput) ToManagementGroupLogSettingsArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsArrayOutput) Index(i pulumi.IntInput) ManagementGroupLogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupLogSettings {
		return vs[0].([]ManagementGroupLogSettings)[vs[1].(int)]
	}).(ManagementGroupLogSettingsOutput)
}

type ManagementGroupLogSettingsResponse struct {
	Category string `pulumi:"category"`
	Enabled  bool   `pulumi:"enabled"`
}

type ManagementGroupLogSettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (o ManagementGroupLogSettingsResponseOutput) ToManagementGroupLogSettingsResponseOutput() ManagementGroupLogSettingsResponseOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseOutput) ToManagementGroupLogSettingsResponseOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) string { return v.Category }).(pulumi.StringOutput)
}

func (o ManagementGroupLogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ManagementGroupLogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type ManagementGroupLogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupLogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupLogSettingsResponse)(nil)).Elem()
}

func (o ManagementGroupLogSettingsResponseArrayOutput) ToManagementGroupLogSettingsResponseArrayOutput() ManagementGroupLogSettingsResponseArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseArrayOutput) ToManagementGroupLogSettingsResponseArrayOutputWithContext(ctx context.Context) ManagementGroupLogSettingsResponseArrayOutput {
	return o
}

func (o ManagementGroupLogSettingsResponseArrayOutput) Index(i pulumi.IntInput) ManagementGroupLogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupLogSettingsResponse {
		return vs[0].([]ManagementGroupLogSettingsResponse)[vs[1].(int)]
	}).(ManagementGroupLogSettingsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupLogSettingsOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupLogSettingsResponseArrayOutput{})
}
