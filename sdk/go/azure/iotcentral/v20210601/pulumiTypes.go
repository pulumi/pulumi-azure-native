


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSkuInfo struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoInput interface {
	pulumi.Input

	ToAppSkuInfoOutput() AppSkuInfoOutput
	ToAppSkuInfoOutputWithContext(context.Context) AppSkuInfoOutput
}

type AppSkuInfoArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return i.ToAppSkuInfoOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput)
}

func (i AppSkuInfoArgs) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return i.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput).ToAppSkuInfoPtrOutputWithContext(ctx)
}









type AppSkuInfoPtrInput interface {
	pulumi.Input

	ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput
	ToAppSkuInfoPtrOutputWithContext(context.Context) AppSkuInfoPtrOutput
}

type appSkuInfoPtrType AppSkuInfoArgs

func AppSkuInfoPtr(v *AppSkuInfoArgs) AppSkuInfoPtrInput {
	return (*appSkuInfoPtrType)(v)
}

func (*appSkuInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfo)(nil)).Elem()
}

func (i *appSkuInfoPtrType) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return i.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (i *appSkuInfoPtrType) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoPtrOutput)
}

type AppSkuInfoOutput struct{ *pulumi.OutputState }

func (AppSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return o.ToAppSkuInfoPtrOutputWithContext(context.Background())
}

func (o AppSkuInfoOutput) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSkuInfo) *AppSkuInfo {
		return &v
	}).(AppSkuInfoPtrOutput)
}

func (o AppSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoPtrOutput struct{ *pulumi.OutputState }

func (AppSkuInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoPtrOutput) ToAppSkuInfoPtrOutput() AppSkuInfoPtrOutput {
	return o
}

func (o AppSkuInfoPtrOutput) ToAppSkuInfoPtrOutputWithContext(ctx context.Context) AppSkuInfoPtrOutput {
	return o
}

func (o AppSkuInfoPtrOutput) Elem() AppSkuInfoOutput {
	return o.ApplyT(func(v *AppSkuInfo) AppSkuInfo {
		if v != nil {
			return *v
		}
		var ret AppSkuInfo
		return ret
	}).(AppSkuInfoOutput)
}

func (o AppSkuInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSkuInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type AppSkuInfoResponse struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoResponseInput interface {
	pulumi.Input

	ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput
	ToAppSkuInfoResponseOutputWithContext(context.Context) AppSkuInfoResponseOutput
}

type AppSkuInfoResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return i.ToAppSkuInfoResponseOutputWithContext(context.Background())
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponseOutput)
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return i.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i AppSkuInfoResponseArgs) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponseOutput).ToAppSkuInfoResponsePtrOutputWithContext(ctx)
}









type AppSkuInfoResponsePtrInput interface {
	pulumi.Input

	ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput
	ToAppSkuInfoResponsePtrOutputWithContext(context.Context) AppSkuInfoResponsePtrOutput
}

type appSkuInfoResponsePtrType AppSkuInfoResponseArgs

func AppSkuInfoResponsePtr(v *AppSkuInfoResponseArgs) AppSkuInfoResponsePtrInput {
	return (*appSkuInfoResponsePtrType)(v)
}

func (*appSkuInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfoResponse)(nil)).Elem()
}

func (i *appSkuInfoResponsePtrType) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return i.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i *appSkuInfoResponsePtrType) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoResponsePtrOutput)
}

type AppSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return o.ToAppSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSkuInfoResponse) *AppSkuInfoResponse {
		return &v
	}).(AppSkuInfoResponsePtrOutput)
}

func (o AppSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponsePtrOutput) ToAppSkuInfoResponsePtrOutput() AppSkuInfoResponsePtrOutput {
	return o
}

func (o AppSkuInfoResponsePtrOutput) ToAppSkuInfoResponsePtrOutputWithContext(ctx context.Context) AppSkuInfoResponsePtrOutput {
	return o
}

func (o AppSkuInfoResponsePtrOutput) Elem() AppSkuInfoResponseOutput {
	return o.ApplyT(func(v *AppSkuInfoResponse) AppSkuInfoResponse {
		if v != nil {
			return *v
		}
		var ret AppSkuInfoResponse
		return ret
	}).(AppSkuInfoResponseOutput)
}

func (o AppSkuInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSkuInfoOutput{})
	pulumi.RegisterOutputType(AppSkuInfoPtrOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponsePtrOutput{})
}
