


package v20180101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Permission struct {
	Actions        []string `pulumi:"actions"`
	DataActions    []string `pulumi:"dataActions"`
	NotActions     []string `pulumi:"notActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}





type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(context.Context) PermissionOutput
}

type PermissionArgs struct {
	Actions        pulumi.StringArrayInput `pulumi:"actions"`
	DataActions    pulumi.StringArrayInput `pulumi:"dataActions"`
	NotActions     pulumi.StringArrayInput `pulumi:"notActions"`
	NotDataActions pulumi.StringArrayInput `pulumi:"notDataActions"`
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (i PermissionArgs) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i PermissionArgs) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}





type PermissionArrayInput interface {
	pulumi.Input

	ToPermissionArrayOutput() PermissionArrayOutput
	ToPermissionArrayOutputWithContext(context.Context) PermissionArrayOutput
}

type PermissionArray []PermissionInput

func (PermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (i PermissionArray) ToPermissionArrayOutput() PermissionArrayOutput {
	return i.ToPermissionArrayOutputWithContext(context.Background())
}

func (i PermissionArray) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionArrayOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionArrayOutput struct{ *pulumi.OutputState }

func (PermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (o PermissionArrayOutput) ToPermissionArrayOutput() PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) Index(i pulumi.IntInput) PermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Permission {
		return vs[0].([]Permission)[vs[1].(int)]
	}).(PermissionOutput)
}

type PermissionResponse struct {
	Actions        []string `pulumi:"actions"`
	DataActions    []string `pulumi:"dataActions"`
	NotActions     []string `pulumi:"notActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}





type PermissionResponseInput interface {
	pulumi.Input

	ToPermissionResponseOutput() PermissionResponseOutput
	ToPermissionResponseOutputWithContext(context.Context) PermissionResponseOutput
}

type PermissionResponseArgs struct {
	Actions        pulumi.StringArrayInput `pulumi:"actions"`
	DataActions    pulumi.StringArrayInput `pulumi:"dataActions"`
	NotActions     pulumi.StringArrayInput `pulumi:"notActions"`
	NotDataActions pulumi.StringArrayInput `pulumi:"notDataActions"`
}

func (PermissionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (i PermissionResponseArgs) ToPermissionResponseOutput() PermissionResponseOutput {
	return i.ToPermissionResponseOutputWithContext(context.Background())
}

func (i PermissionResponseArgs) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionResponseOutput)
}





type PermissionResponseArrayInput interface {
	pulumi.Input

	ToPermissionResponseArrayOutput() PermissionResponseArrayOutput
	ToPermissionResponseArrayOutputWithContext(context.Context) PermissionResponseArrayOutput
}

type PermissionResponseArray []PermissionResponseInput

func (PermissionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (i PermissionResponseArray) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return i.ToPermissionResponseArrayOutputWithContext(context.Background())
}

func (i PermissionResponseArray) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionResponseArrayOutput)
}

type PermissionResponseOutput struct{ *pulumi.OutputState }

func (PermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseOutput) ToPermissionResponseOutput() PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) Actions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.Actions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotActions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (PermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) Index(i pulumi.IntInput) PermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionResponse {
		return vs[0].([]PermissionResponse)[vs[1].(int)]
	}).(PermissionResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(PermissionOutput{})
	pulumi.RegisterOutputType(PermissionArrayOutput{})
	pulumi.RegisterOutputType(PermissionResponseOutput{})
	pulumi.RegisterOutputType(PermissionResponseArrayOutput{})
}
