


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RedisAccessKeysResponse struct {
	PrimaryKey   *string `pulumi:"primaryKey"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}

type RedisAccessKeysResponseOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o RedisAccessKeysResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type RedisAccessKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) Elem() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) RedisAccessKeysResponse {
		if v != nil {
			return *v
		}
		var ret RedisAccessKeysResponse
		return ret
	}).(RedisAccessKeysResponseOutput)
}

func (o RedisAccessKeysResponsePtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o RedisAccessKeysResponsePtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Family   pulumi.StringInput `pulumi:"family"`
	Name     pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RedisAccessKeysResponseOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
