


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MessageCountDetailsResponse struct {
	ActiveMessageCount             float64 `pulumi:"activeMessageCount"`
	DeadLetterMessageCount         float64 `pulumi:"deadLetterMessageCount"`
	ScheduledMessageCount          float64 `pulumi:"scheduledMessageCount"`
	TransferDeadLetterMessageCount float64 `pulumi:"transferDeadLetterMessageCount"`
	TransferMessageCount           float64 `pulumi:"transferMessageCount"`
}





type MessageCountDetailsResponseInput interface {
	pulumi.Input

	ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput
	ToMessageCountDetailsResponseOutputWithContext(context.Context) MessageCountDetailsResponseOutput
}

type MessageCountDetailsResponseArgs struct {
	ActiveMessageCount             pulumi.Float64Input `pulumi:"activeMessageCount"`
	DeadLetterMessageCount         pulumi.Float64Input `pulumi:"deadLetterMessageCount"`
	ScheduledMessageCount          pulumi.Float64Input `pulumi:"scheduledMessageCount"`
	TransferDeadLetterMessageCount pulumi.Float64Input `pulumi:"transferDeadLetterMessageCount"`
	TransferMessageCount           pulumi.Float64Input `pulumi:"transferMessageCount"`
}

func (MessageCountDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return i.ToMessageCountDetailsResponseOutputWithContext(context.Background())
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponseOutput)
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return i.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (i MessageCountDetailsResponseArgs) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponseOutput).ToMessageCountDetailsResponsePtrOutputWithContext(ctx)
}









type MessageCountDetailsResponsePtrInput interface {
	pulumi.Input

	ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput
	ToMessageCountDetailsResponsePtrOutputWithContext(context.Context) MessageCountDetailsResponsePtrOutput
}

type messageCountDetailsResponsePtrType MessageCountDetailsResponseArgs

func MessageCountDetailsResponsePtr(v *MessageCountDetailsResponseArgs) MessageCountDetailsResponsePtrInput {
	return (*messageCountDetailsResponsePtrType)(v)
}

func (*messageCountDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageCountDetailsResponse)(nil)).Elem()
}

func (i *messageCountDetailsResponsePtrType) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return i.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *messageCountDetailsResponsePtrType) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessageCountDetailsResponsePtrOutput)
}

type MessageCountDetailsResponseOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutput() MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponseOutputWithContext(ctx context.Context) MessageCountDetailsResponseOutput {
	return o
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o.ToMessageCountDetailsResponsePtrOutputWithContext(context.Background())
}

func (o MessageCountDetailsResponseOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MessageCountDetailsResponse) *MessageCountDetailsResponse {
		return &v
	}).(MessageCountDetailsResponsePtrOutput)
}

func (o MessageCountDetailsResponseOutput) ActiveMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ActiveMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) DeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.DeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) ScheduledMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.ScheduledMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferDeadLetterMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferDeadLetterMessageCount }).(pulumi.Float64Output)
}

func (o MessageCountDetailsResponseOutput) TransferMessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v MessageCountDetailsResponse) float64 { return v.TransferMessageCount }).(pulumi.Float64Output)
}

type MessageCountDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (MessageCountDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MessageCountDetailsResponse)(nil)).Elem()
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutput() MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) ToMessageCountDetailsResponsePtrOutputWithContext(ctx context.Context) MessageCountDetailsResponsePtrOutput {
	return o
}

func (o MessageCountDetailsResponsePtrOutput) Elem() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) MessageCountDetailsResponse {
		if v != nil {
			return *v
		}
		var ret MessageCountDetailsResponse
		return ret
	}).(MessageCountDetailsResponseOutput)
}

func (o MessageCountDetailsResponsePtrOutput) ActiveMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ActiveMessageCount
	}).(pulumi.Float64PtrOutput)
}

func (o MessageCountDetailsResponsePtrOutput) DeadLetterMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.DeadLetterMessageCount
	}).(pulumi.Float64PtrOutput)
}

func (o MessageCountDetailsResponsePtrOutput) ScheduledMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ScheduledMessageCount
	}).(pulumi.Float64PtrOutput)
}

func (o MessageCountDetailsResponsePtrOutput) TransferDeadLetterMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransferDeadLetterMessageCount
	}).(pulumi.Float64PtrOutput)
}

func (o MessageCountDetailsResponsePtrOutput) TransferMessageCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MessageCountDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.TransferMessageCount
	}).(pulumi.Float64PtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     string  `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringInput    `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Tier }).(pulumi.StringOutput)
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
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Name     *string `pulumi:"name"`
	Tier     string  `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	Tier     pulumi.StringInput    `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
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
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MessageCountDetailsResponseOutput{})
	pulumi.RegisterOutputType(MessageCountDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
