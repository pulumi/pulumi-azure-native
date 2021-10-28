


package v20150601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeepCreatedOrigin struct {
	HostName  string `pulumi:"hostName"`
	HttpPort  *int   `pulumi:"httpPort"`
	HttpsPort *int   `pulumi:"httpsPort"`
	Name      string `pulumi:"name"`
}





type DeepCreatedOriginInput interface {
	pulumi.Input

	ToDeepCreatedOriginOutput() DeepCreatedOriginOutput
	ToDeepCreatedOriginOutputWithContext(context.Context) DeepCreatedOriginOutput
}

type DeepCreatedOriginArgs struct {
	HostName  pulumi.StringInput `pulumi:"hostName"`
	HttpPort  pulumi.IntPtrInput `pulumi:"httpPort"`
	HttpsPort pulumi.IntPtrInput `pulumi:"httpsPort"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DeepCreatedOriginArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOrigin)(nil)).Elem()
}

func (i DeepCreatedOriginArgs) ToDeepCreatedOriginOutput() DeepCreatedOriginOutput {
	return i.ToDeepCreatedOriginOutputWithContext(context.Background())
}

func (i DeepCreatedOriginArgs) ToDeepCreatedOriginOutputWithContext(ctx context.Context) DeepCreatedOriginOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginOutput)
}





type DeepCreatedOriginArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput
	ToDeepCreatedOriginArrayOutputWithContext(context.Context) DeepCreatedOriginArrayOutput
}

type DeepCreatedOriginArray []DeepCreatedOriginInput

func (DeepCreatedOriginArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOrigin)(nil)).Elem()
}

func (i DeepCreatedOriginArray) ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput {
	return i.ToDeepCreatedOriginArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginArray) ToDeepCreatedOriginArrayOutputWithContext(ctx context.Context) DeepCreatedOriginArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginArrayOutput)
}

type DeepCreatedOriginOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOrigin)(nil)).Elem()
}

func (o DeepCreatedOriginOutput) ToDeepCreatedOriginOutput() DeepCreatedOriginOutput {
	return o
}

func (o DeepCreatedOriginOutput) ToDeepCreatedOriginOutputWithContext(ctx context.Context) DeepCreatedOriginOutput {
	return o
}

func (o DeepCreatedOriginOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) string { return v.HostName }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOrigin) string { return v.Name }).(pulumi.StringOutput)
}

type DeepCreatedOriginArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOrigin)(nil)).Elem()
}

func (o DeepCreatedOriginArrayOutput) ToDeepCreatedOriginArrayOutput() DeepCreatedOriginArrayOutput {
	return o
}

func (o DeepCreatedOriginArrayOutput) ToDeepCreatedOriginArrayOutputWithContext(ctx context.Context) DeepCreatedOriginArrayOutput {
	return o
}

func (o DeepCreatedOriginArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOrigin {
		return vs[0].([]DeepCreatedOrigin)[vs[1].(int)]
	}).(DeepCreatedOriginOutput)
}

type DeepCreatedOriginResponse struct {
	HostName  string `pulumi:"hostName"`
	HttpPort  *int   `pulumi:"httpPort"`
	HttpsPort *int   `pulumi:"httpsPort"`
	Name      string `pulumi:"name"`
}





type DeepCreatedOriginResponseInput interface {
	pulumi.Input

	ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput
	ToDeepCreatedOriginResponseOutputWithContext(context.Context) DeepCreatedOriginResponseOutput
}

type DeepCreatedOriginResponseArgs struct {
	HostName  pulumi.StringInput `pulumi:"hostName"`
	HttpPort  pulumi.IntPtrInput `pulumi:"httpPort"`
	HttpsPort pulumi.IntPtrInput `pulumi:"httpsPort"`
	Name      pulumi.StringInput `pulumi:"name"`
}

func (DeepCreatedOriginResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginResponse)(nil)).Elem()
}

func (i DeepCreatedOriginResponseArgs) ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput {
	return i.ToDeepCreatedOriginResponseOutputWithContext(context.Background())
}

func (i DeepCreatedOriginResponseArgs) ToDeepCreatedOriginResponseOutputWithContext(ctx context.Context) DeepCreatedOriginResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginResponseOutput)
}





type DeepCreatedOriginResponseArrayInput interface {
	pulumi.Input

	ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput
	ToDeepCreatedOriginResponseArrayOutputWithContext(context.Context) DeepCreatedOriginResponseArrayOutput
}

type DeepCreatedOriginResponseArray []DeepCreatedOriginResponseInput

func (DeepCreatedOriginResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginResponse)(nil)).Elem()
}

func (i DeepCreatedOriginResponseArray) ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput {
	return i.ToDeepCreatedOriginResponseArrayOutputWithContext(context.Background())
}

func (i DeepCreatedOriginResponseArray) ToDeepCreatedOriginResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeepCreatedOriginResponseArrayOutput)
}

type DeepCreatedOriginResponseOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeepCreatedOriginResponse)(nil)).Elem()
}

func (o DeepCreatedOriginResponseOutput) ToDeepCreatedOriginResponseOutput() DeepCreatedOriginResponseOutput {
	return o
}

func (o DeepCreatedOriginResponseOutput) ToDeepCreatedOriginResponseOutputWithContext(ctx context.Context) DeepCreatedOriginResponseOutput {
	return o
}

func (o DeepCreatedOriginResponseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) string { return v.HostName }).(pulumi.StringOutput)
}

func (o DeepCreatedOriginResponseOutput) HttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.HttpPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) HttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) *int { return v.HttpsPort }).(pulumi.IntPtrOutput)
}

func (o DeepCreatedOriginResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DeepCreatedOriginResponse) string { return v.Name }).(pulumi.StringOutput)
}

type DeepCreatedOriginResponseArrayOutput struct{ *pulumi.OutputState }

func (DeepCreatedOriginResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeepCreatedOriginResponse)(nil)).Elem()
}

func (o DeepCreatedOriginResponseArrayOutput) ToDeepCreatedOriginResponseArrayOutput() DeepCreatedOriginResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginResponseArrayOutput) ToDeepCreatedOriginResponseArrayOutputWithContext(ctx context.Context) DeepCreatedOriginResponseArrayOutput {
	return o
}

func (o DeepCreatedOriginResponseArrayOutput) Index(i pulumi.IntInput) DeepCreatedOriginResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeepCreatedOriginResponse {
		return vs[0].([]DeepCreatedOriginResponse)[vs[1].(int)]
	}).(DeepCreatedOriginResponseOutput)
}

type Sku struct {
	Name *SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name SkuNamePtrInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v Sku) *SkuName { return v.Name }).(SkuNamePtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return v.Name
	}).(SkuNamePtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeepCreatedOriginOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginArrayOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseOutput{})
	pulumi.RegisterOutputType(DeepCreatedOriginResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
