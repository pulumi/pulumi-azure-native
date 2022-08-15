


package v20171115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnvironmentStateDetailsResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type EnvironmentStateDetailsResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentStateDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStateDetailsResponse)(nil)).Elem()
}

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponseOutput() EnvironmentStateDetailsResponseOutput {
	return o
}

func (o EnvironmentStateDetailsResponseOutput) ToEnvironmentStateDetailsResponseOutputWithContext(ctx context.Context) EnvironmentStateDetailsResponseOutput {
	return o
}

func (o EnvironmentStateDetailsResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o EnvironmentStateDetailsResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentStateDetailsResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type EnvironmentStatusResponse struct {
	Ingress IngressEnvironmentStatusResponse `pulumi:"ingress"`
}

type EnvironmentStatusResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentStatusResponse)(nil)).Elem()
}

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponseOutput() EnvironmentStatusResponseOutput {
	return o
}

func (o EnvironmentStatusResponseOutput) ToEnvironmentStatusResponseOutputWithContext(ctx context.Context) EnvironmentStatusResponseOutput {
	return o
}

func (o EnvironmentStatusResponseOutput) Ingress() IngressEnvironmentStatusResponseOutput {
	return o.ApplyT(func(v EnvironmentStatusResponse) IngressEnvironmentStatusResponse { return v.Ingress }).(IngressEnvironmentStatusResponseOutput)
}

type IngressEnvironmentStatusResponse struct {
	State        *string                         `pulumi:"state"`
	StateDetails EnvironmentStateDetailsResponse `pulumi:"stateDetails"`
}

type IngressEnvironmentStatusResponseOutput struct{ *pulumi.OutputState }

func (IngressEnvironmentStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressEnvironmentStatusResponse)(nil)).Elem()
}

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponseOutput() IngressEnvironmentStatusResponseOutput {
	return o
}

func (o IngressEnvironmentStatusResponseOutput) ToIngressEnvironmentStatusResponseOutputWithContext(ctx context.Context) IngressEnvironmentStatusResponseOutput {
	return o
}

func (o IngressEnvironmentStatusResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o IngressEnvironmentStatusResponseOutput) StateDetails() EnvironmentStateDetailsResponseOutput {
	return o.ApplyT(func(v IngressEnvironmentStatusResponse) EnvironmentStateDetailsResponse { return v.StateDetails }).(EnvironmentStateDetailsResponseOutput)
}

type PartitionKeyProperty struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PartitionKeyPropertyInput interface {
	pulumi.Input

	ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput
	ToPartitionKeyPropertyOutputWithContext(context.Context) PartitionKeyPropertyOutput
}

type PartitionKeyPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PartitionKeyPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyProperty)(nil)).Elem()
}

func (i PartitionKeyPropertyArgs) ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput {
	return i.ToPartitionKeyPropertyOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyArgs) ToPartitionKeyPropertyOutputWithContext(ctx context.Context) PartitionKeyPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyOutput)
}





type PartitionKeyPropertyArrayInput interface {
	pulumi.Input

	ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput
	ToPartitionKeyPropertyArrayOutputWithContext(context.Context) PartitionKeyPropertyArrayOutput
}

type PartitionKeyPropertyArray []PartitionKeyPropertyInput

func (PartitionKeyPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyProperty)(nil)).Elem()
}

func (i PartitionKeyPropertyArray) ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput {
	return i.ToPartitionKeyPropertyArrayOutputWithContext(context.Background())
}

func (i PartitionKeyPropertyArray) ToPartitionKeyPropertyArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartitionKeyPropertyArrayOutput)
}

type PartitionKeyPropertyOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyProperty)(nil)).Elem()
}

func (o PartitionKeyPropertyOutput) ToPartitionKeyPropertyOutput() PartitionKeyPropertyOutput {
	return o
}

func (o PartitionKeyPropertyOutput) ToPartitionKeyPropertyOutputWithContext(ctx context.Context) PartitionKeyPropertyOutput {
	return o
}

func (o PartitionKeyPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartitionKeyPropertyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyProperty) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PartitionKeyPropertyArrayOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyProperty)(nil)).Elem()
}

func (o PartitionKeyPropertyArrayOutput) ToPartitionKeyPropertyArrayOutput() PartitionKeyPropertyArrayOutput {
	return o
}

func (o PartitionKeyPropertyArrayOutput) ToPartitionKeyPropertyArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyArrayOutput {
	return o
}

func (o PartitionKeyPropertyArrayOutput) Index(i pulumi.IntInput) PartitionKeyPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartitionKeyProperty {
		return vs[0].([]PartitionKeyProperty)[vs[1].(int)]
	}).(PartitionKeyPropertyOutput)
}

type PartitionKeyPropertyResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type PartitionKeyPropertyResponseOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionKeyPropertyResponse)(nil)).Elem()
}

func (o PartitionKeyPropertyResponseOutput) ToPartitionKeyPropertyResponseOutput() PartitionKeyPropertyResponseOutput {
	return o
}

func (o PartitionKeyPropertyResponseOutput) ToPartitionKeyPropertyResponseOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseOutput {
	return o
}

func (o PartitionKeyPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PartitionKeyPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PartitionKeyPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PartitionKeyPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (PartitionKeyPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PartitionKeyPropertyResponse)(nil)).Elem()
}

func (o PartitionKeyPropertyResponseArrayOutput) ToPartitionKeyPropertyResponseArrayOutput() PartitionKeyPropertyResponseArrayOutput {
	return o
}

func (o PartitionKeyPropertyResponseArrayOutput) ToPartitionKeyPropertyResponseArrayOutputWithContext(ctx context.Context) PartitionKeyPropertyResponseArrayOutput {
	return o
}

func (o PartitionKeyPropertyResponseArrayOutput) Index(i pulumi.IntInput) PartitionKeyPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PartitionKeyPropertyResponse {
		return vs[0].([]PartitionKeyPropertyResponse)[vs[1].(int)]
	}).(PartitionKeyPropertyResponseOutput)
}

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
	pulumi.RegisterOutputType(EnvironmentStateDetailsResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(IngressEnvironmentStatusResponseOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(PartitionKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
