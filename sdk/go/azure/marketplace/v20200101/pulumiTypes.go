


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Plan struct {
	Accessibility *string `pulumi:"accessibility"`
}





type PlanInput interface {
	pulumi.Input

	ToPlanOutput() PlanOutput
	ToPlanOutputWithContext(context.Context) PlanOutput
}

type PlanArgs struct {
	Accessibility pulumi.StringPtrInput `pulumi:"accessibility"`
}

func (PlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (i PlanArgs) ToPlanOutput() PlanOutput {
	return i.ToPlanOutputWithContext(context.Background())
}

func (i PlanArgs) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanOutput)
}





type PlanArrayInput interface {
	pulumi.Input

	ToPlanArrayOutput() PlanArrayOutput
	ToPlanArrayOutputWithContext(context.Context) PlanArrayOutput
}

type PlanArray []PlanInput

func (PlanArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (i PlanArray) ToPlanArrayOutput() PlanArrayOutput {
	return i.ToPlanArrayOutputWithContext(context.Background())
}

func (i PlanArray) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanArrayOutput)
}

type PlanOutput struct{ *pulumi.OutputState }

func (PlanOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Plan)(nil)).Elem()
}

func (o PlanOutput) ToPlanOutput() PlanOutput {
	return o
}

func (o PlanOutput) ToPlanOutputWithContext(ctx context.Context) PlanOutput {
	return o
}

func (o PlanOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Plan) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

type PlanArrayOutput struct{ *pulumi.OutputState }

func (PlanArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Plan)(nil)).Elem()
}

func (o PlanArrayOutput) ToPlanArrayOutput() PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) ToPlanArrayOutputWithContext(ctx context.Context) PlanArrayOutput {
	return o
}

func (o PlanArrayOutput) Index(i pulumi.IntInput) PlanOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Plan {
		return vs[0].([]Plan)[vs[1].(int)]
	}).(PlanOutput)
}

type PlanResponse struct {
	Accessibility     *string `pulumi:"accessibility"`
	AltStackReference string  `pulumi:"altStackReference"`
	PlanDisplayName   string  `pulumi:"planDisplayName"`
	PlanId            string  `pulumi:"planId"`
	SkuId             string  `pulumi:"skuId"`
	StackType         string  `pulumi:"stackType"`
}





type PlanResponseInput interface {
	pulumi.Input

	ToPlanResponseOutput() PlanResponseOutput
	ToPlanResponseOutputWithContext(context.Context) PlanResponseOutput
}

type PlanResponseArgs struct {
	Accessibility     pulumi.StringPtrInput `pulumi:"accessibility"`
	AltStackReference pulumi.StringInput    `pulumi:"altStackReference"`
	PlanDisplayName   pulumi.StringInput    `pulumi:"planDisplayName"`
	PlanId            pulumi.StringInput    `pulumi:"planId"`
	SkuId             pulumi.StringInput    `pulumi:"skuId"`
	StackType         pulumi.StringInput    `pulumi:"stackType"`
}

func (PlanResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (i PlanResponseArgs) ToPlanResponseOutput() PlanResponseOutput {
	return i.ToPlanResponseOutputWithContext(context.Background())
}

func (i PlanResponseArgs) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseOutput)
}





type PlanResponseArrayInput interface {
	pulumi.Input

	ToPlanResponseArrayOutput() PlanResponseArrayOutput
	ToPlanResponseArrayOutputWithContext(context.Context) PlanResponseArrayOutput
}

type PlanResponseArray []PlanResponseInput

func (PlanResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanResponse)(nil)).Elem()
}

func (i PlanResponseArray) ToPlanResponseArrayOutput() PlanResponseArrayOutput {
	return i.ToPlanResponseArrayOutputWithContext(context.Background())
}

func (i PlanResponseArray) ToPlanResponseArrayOutputWithContext(ctx context.Context) PlanResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PlanResponseArrayOutput)
}

type PlanResponseOutput struct{ *pulumi.OutputState }

func (PlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PlanResponse)(nil)).Elem()
}

func (o PlanResponseOutput) ToPlanResponseOutput() PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) ToPlanResponseOutputWithContext(ctx context.Context) PlanResponseOutput {
	return o
}

func (o PlanResponseOutput) Accessibility() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PlanResponse) *string { return v.Accessibility }).(pulumi.StringPtrOutput)
}

func (o PlanResponseOutput) AltStackReference() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.AltStackReference }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PlanDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanDisplayName }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.PlanId }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) SkuId() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.SkuId }).(pulumi.StringOutput)
}

func (o PlanResponseOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v PlanResponse) string { return v.StackType }).(pulumi.StringOutput)
}

type PlanResponseArrayOutput struct{ *pulumi.OutputState }

func (PlanResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PlanResponse)(nil)).Elem()
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutput() PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) ToPlanResponseArrayOutputWithContext(ctx context.Context) PlanResponseArrayOutput {
	return o
}

func (o PlanResponseArrayOutput) Index(i pulumi.IntInput) PlanResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PlanResponse {
		return vs[0].([]PlanResponse)[vs[1].(int)]
	}).(PlanResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PlanInput)(nil)).Elem(), PlanArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanArrayInput)(nil)).Elem(), PlanArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponseInput)(nil)).Elem(), PlanResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PlanResponseArrayInput)(nil)).Elem(), PlanResponseArray{})
	pulumi.RegisterOutputType(PlanOutput{})
	pulumi.RegisterOutputType(PlanArrayOutput{})
	pulumi.RegisterOutputType(PlanResponseOutput{})
	pulumi.RegisterOutputType(PlanResponseArrayOutput{})
}
