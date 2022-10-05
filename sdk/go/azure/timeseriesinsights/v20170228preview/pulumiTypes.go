


package v20170228preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ReferenceDataSetKeyProperty struct {
	Name *string                       `pulumi:"name"`
	Type *ReferenceDataKeyPropertyType `pulumi:"type"`
}





type ReferenceDataSetKeyPropertyInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput
	ToReferenceDataSetKeyPropertyOutputWithContext(context.Context) ReferenceDataSetKeyPropertyOutput
}

type ReferenceDataSetKeyPropertyArgs struct {
	Name pulumi.StringPtrInput                `pulumi:"name"`
	Type ReferenceDataKeyPropertyTypePtrInput `pulumi:"type"`
}

func (ReferenceDataSetKeyPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return i.ToReferenceDataSetKeyPropertyOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyOutput)
}





type ReferenceDataSetKeyPropertyArrayInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput
	ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Context) ReferenceDataSetKeyPropertyArrayOutput
}

type ReferenceDataSetKeyPropertyArray []ReferenceDataSetKeyPropertyInput

func (ReferenceDataSetKeyPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return i.ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyArrayOutput)
}

type ReferenceDataSetKeyPropertyOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ReferenceDataSetKeyPropertyOutput) Type() ReferenceDataKeyPropertyTypePtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *ReferenceDataKeyPropertyType { return v.Type }).(ReferenceDataKeyPropertyTypePtrOutput)
}

type ReferenceDataSetKeyPropertyArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyProperty {
		return vs[0].([]ReferenceDataSetKeyProperty)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyOutput)
}

type ReferenceDataSetKeyPropertyResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ReferenceDataSetKeyPropertyResponseOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ReferenceDataSetKeyPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyPropertyResponse {
		return vs[0].([]ReferenceDataSetKeyPropertyResponse)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyResponseOutput)
}

type Sku struct {
	Capacity int     `pulumi:"capacity"`
	Name     SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput `pulumi:"capacity"`
	Name     SkuNameInput    `pulumi:"name"`
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

func (o SkuOutput) Name() SkuNameOutput {
	return o.ApplyT(func(v Sku) SkuName { return v.Name }).(SkuNameOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
