


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RedisAccessKeysResponse struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}





type RedisAccessKeysResponseInput interface {
	pulumi.Input

	ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput
	ToRedisAccessKeysResponseOutputWithContext(context.Context) RedisAccessKeysResponseOutput
}

type RedisAccessKeysResponseArgs struct {
	PrimaryKey   pulumi.StringInput `pulumi:"primaryKey"`
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
}

func (RedisAccessKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return i.ToRedisAccessKeysResponseOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput)
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput).ToRedisAccessKeysResponsePtrOutputWithContext(ctx)
}









type RedisAccessKeysResponsePtrInput interface {
	pulumi.Input

	ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput
	ToRedisAccessKeysResponsePtrOutputWithContext(context.Context) RedisAccessKeysResponsePtrOutput
}

type redisAccessKeysResponsePtrType RedisAccessKeysResponseArgs

func RedisAccessKeysResponsePtr(v *RedisAccessKeysResponseArgs) RedisAccessKeysResponsePtrInput {
	return (*redisAccessKeysResponsePtrType)(v)
}

func (*redisAccessKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponsePtrOutput)
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

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedisAccessKeysResponse) *RedisAccessKeysResponse {
		return &v
	}).(RedisAccessKeysResponsePtrOutput)
}

func (o RedisAccessKeysResponseOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o RedisAccessKeysResponseOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.SecondaryKey }).(pulumi.StringOutput)
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
		return &v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o RedisAccessKeysResponsePtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type RedisLinkedServerResponse struct {
	Id string `pulumi:"id"`
}





type RedisLinkedServerResponseInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput
	ToRedisLinkedServerResponseOutputWithContext(context.Context) RedisLinkedServerResponseOutput
}

type RedisLinkedServerResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (RedisLinkedServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return i.ToRedisLinkedServerResponseOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseOutput)
}





type RedisLinkedServerResponseArrayInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput
	ToRedisLinkedServerResponseArrayOutputWithContext(context.Context) RedisLinkedServerResponseArrayOutput
}

type RedisLinkedServerResponseArray []RedisLinkedServerResponseInput

func (RedisLinkedServerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return i.ToRedisLinkedServerResponseArrayOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseArrayOutput)
}

type RedisLinkedServerResponseOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RedisLinkedServerResponse) string { return v.Id }).(pulumi.StringOutput)
}

type RedisLinkedServerResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) Index(i pulumi.IntInput) RedisLinkedServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisLinkedServerResponse {
		return vs[0].([]RedisLinkedServerResponse)[vs[1].(int)]
	}).(RedisLinkedServerResponseOutput)
}

type ScheduleEntry struct {
	DayOfWeek         DayOfWeek `pulumi:"dayOfWeek"`
	MaintenanceWindow *string   `pulumi:"maintenanceWindow"`
	StartHourUtc      int       `pulumi:"startHourUtc"`
}





type ScheduleEntryInput interface {
	pulumi.Input

	ToScheduleEntryOutput() ScheduleEntryOutput
	ToScheduleEntryOutputWithContext(context.Context) ScheduleEntryOutput
}

type ScheduleEntryArgs struct {
	DayOfWeek         DayOfWeekInput        `pulumi:"dayOfWeek"`
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	StartHourUtc      pulumi.IntInput       `pulumi:"startHourUtc"`
}

func (ScheduleEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArgs) ToScheduleEntryOutput() ScheduleEntryOutput {
	return i.ToScheduleEntryOutputWithContext(context.Background())
}

func (i ScheduleEntryArgs) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryOutput)
}





type ScheduleEntryArrayInput interface {
	pulumi.Input

	ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput
	ToScheduleEntryArrayOutputWithContext(context.Context) ScheduleEntryArrayOutput
}

type ScheduleEntryArray []ScheduleEntryInput

func (ScheduleEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return i.ToScheduleEntryArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryArrayOutput)
}

type ScheduleEntryOutput struct{ *pulumi.OutputState }

func (ScheduleEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryOutput) ToScheduleEntryOutput() ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) DayOfWeek() DayOfWeekOutput {
	return o.ApplyT(func(v ScheduleEntry) DayOfWeek { return v.DayOfWeek }).(DayOfWeekOutput)
}

func (o ScheduleEntryOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntry) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntry) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) Index(i pulumi.IntInput) ScheduleEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntry {
		return vs[0].([]ScheduleEntry)[vs[1].(int)]
	}).(ScheduleEntryOutput)
}

type ScheduleEntryResponse struct {
	DayOfWeek         string  `pulumi:"dayOfWeek"`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	StartHourUtc      int     `pulumi:"startHourUtc"`
}





type ScheduleEntryResponseInput interface {
	pulumi.Input

	ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput
	ToScheduleEntryResponseOutputWithContext(context.Context) ScheduleEntryResponseOutput
}

type ScheduleEntryResponseArgs struct {
	DayOfWeek         pulumi.StringInput    `pulumi:"dayOfWeek"`
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	StartHourUtc      pulumi.IntInput       `pulumi:"startHourUtc"`
}

func (ScheduleEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return i.ToScheduleEntryResponseOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseOutput)
}





type ScheduleEntryResponseArrayInput interface {
	pulumi.Input

	ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput
	ToScheduleEntryResponseArrayOutputWithContext(context.Context) ScheduleEntryResponseArrayOutput
}

type ScheduleEntryResponseArray []ScheduleEntryResponseInput

func (ScheduleEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return i.ToScheduleEntryResponseArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseArrayOutput)
}

type ScheduleEntryResponseOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

func (o ScheduleEntryResponseOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryResponseOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) Index(i pulumi.IntInput) ScheduleEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntryResponse {
		return vs[0].([]ScheduleEntryResponse)[vs[1].(int)]
	}).(ScheduleEntryResponseOutput)
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

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Family   pulumi.StringInput `pulumi:"family"`
	Name     pulumi.StringInput `pulumi:"name"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedisAccessKeysResponseInput)(nil)).Elem(), RedisAccessKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisAccessKeysResponsePtrInput)(nil)).Elem(), RedisAccessKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisLinkedServerResponseInput)(nil)).Elem(), RedisLinkedServerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisLinkedServerResponseArrayInput)(nil)).Elem(), RedisLinkedServerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryInput)(nil)).Elem(), ScheduleEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryArrayInput)(nil)).Elem(), ScheduleEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryResponseInput)(nil)).Elem(), ScheduleEntryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryResponseArrayInput)(nil)).Elem(), ScheduleEntryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterOutputType(RedisAccessKeysResponseOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryOutput{})
	pulumi.RegisterOutputType(ScheduleEntryArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
